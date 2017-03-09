// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package proxy provides a WebDriver protocol that forwards requests
// to a WebDriver server provided by an environment.Env.
package proxy

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/httphelper"
	"github.com/bazelbuild/rules_webtesting/go/launcher/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/launcher/environment"
	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/healthreporter"
	"github.com/bazelbuild/rules_webtesting/go/metadata"
	"github.com/bazelbuild/rules_webtesting/go/portpicker"
)

const (
	compName = "WebDriver proxy"
	timeout  = 1 * time.Second
)

var handlerProviders = map[string]HTTPHandlerProvider{}

// A HTTPHandlerProvider is a function that provides a HTTPHandler.
type HTTPHandlerProvider func(*Proxy) (HTTPHandler, error)

// A HTTPHandler implements http.Handler plus a Shutdown method.
type HTTPHandler interface {
	http.Handler

	// Shutdown is called when the proxy is in the process of shutting down.
	Shutdown(context.Context) error
}

// AddHTTPHandlerProvider adds a HTTPHandlerProvider used to create handlers for
// specified routes in any Proxy structs creates by New.
func AddHTTPHandlerProvider(route string, provider HTTPHandlerProvider) {
	handlerProviders[route] = provider
}

// Proxy starts a WebDriver protocol proxy.
type Proxy struct {
	Diagnostics diagnostics.Diagnostics
	Env         environment.Env
	Metadata    *metadata.Metadata
	Address     string
	handlers    []HTTPHandler
	srv         *http.Server
	port        int
}

// New creates a new Proxy object.
func New(env environment.Env, m *metadata.Metadata, d diagnostics.Diagnostics) (*Proxy, error) {
	port, err := portpicker.PickUnusedPort()
	if err != nil {
		return nil, errors.New(compName, err)
	}
	p := &Proxy{
		Diagnostics: d,
		Env:         env,
		Metadata:    m,
		Address:     net.JoinHostPort("localhost", strconv.Itoa(port)),
		port:        port,
	}

	mux := http.NewServeMux()

	for route, provider := range handlerProviders {
		h, err := provider(p)
		if err != nil {
			return nil, err
		}
		p.handlers = append(p.handlers, h)
		mux.Handle(route, h)
	}

	p.srv = &http.Server{
		Addr:    ":" + strconv.Itoa(p.port),
		Handler: mux,
	}

	return p, nil
}

// Component returns the name used in error messages.
func (*Proxy) Name() string {
	return compName
}

// Start configures the proxy with handlers, starts its listen loop, and waits for it to respond to a health check.
func (p *Proxy) Start(ctx context.Context) error {
	log.Printf("launching server at: %v", p.Address)

	go func() {
		log.Printf("Proxy has exited: %v", p.srv.ListenAndServe())
	}()

	healthyCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return healthreporter.WaitForHealthy(healthyCtx, p)
}

// Healthy returns nil if the proxy is able to receive requests.
func (p *Proxy) Healthy(ctx context.Context) error {
	url := fmt.Sprintf("http://%s/healthz", p.Address)
	resp, err := httphelper.Get(ctx, url)
	if err != nil {
		return errors.New(p.Name(), fmt.Errorf("error getting %s: %v", url, err))
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(p.Name(), fmt.Errorf("request to %s returned status %v", url, resp.StatusCode))
	}
	return nil
}

// Shutdown calls Shutdown on all handlers, then shuts the HTTP server down.
func (p *Proxy) Shutdown(ctx context.Context) error {
	for _, handler := range p.handlers {
		if err := handler.Shutdown(ctx); err != nil {
			p.Diagnostics.Warning(err)
		}
	}
	// TODO(DrMarcII) figure out why Shutdown doesn't work.
	return nil
}
