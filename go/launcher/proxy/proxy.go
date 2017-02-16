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
	"io"
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
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/driverhub"
	"github.com/bazelbuild/rules_webtesting/go/metadata"
	"github.com/bazelbuild/rules_webtesting/go/portpicker"
)

const (
	compName   = "WebDriver proxy"
	timeout    = 1 * time.Second
	envTimeout = 1 * time.Second // some environments such as Android take a long time to start up.
)

// Proxy starts a WebDriver protocol proxy.
type Proxy struct {
	diagnostics.Diagnostics
	env      environment.Env
	metadata *metadata.Metadata
	Address  string
	port     int
}

// New creates a new Proxy object.
func New(env environment.Env, m *metadata.Metadata, d diagnostics.Diagnostics) (*Proxy, error) {
	port, err := portpicker.PickUnusedPort()
	if err != nil {
		return nil, errors.New(compName, err)
	}
	return &Proxy{
		Diagnostics: d,
		env:         env,
		metadata:    m,
		Address:     net.JoinHostPort("localhost", strconv.Itoa(port)),
		port:        port,
	}, nil
}

// Component returns the name used in error messages.
func (*Proxy) Name() string {
	return compName
}

// Start configures the proxy with handlers, starts its listen loop, and waits for it to respond to a health check.
func (p *Proxy) Start(ctx context.Context) error {
	log.Printf("launching server at: %v", p.Address)

	http.Handle("/wd/hub/", driverhub.NewHandler(p.env, p.metadata, p.Diagnostics))

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "ok")
	})

	go func() {
		log.Printf("Proxy has exited: %v", http.ListenAndServe(":"+strconv.Itoa(p.port), nil))
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
