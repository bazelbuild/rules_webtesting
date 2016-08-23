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

	"github.com/web_test_launcher/launcher/environments/environment"
	"github.com/web_test_launcher/launcher/errors"
	"github.com/web_test_launcher/launcher/healthreporter"
	"github.com/web_test_launcher/launcher/proxy/driverhub"
	"github.com/web_test_launcher/util/httphelper"
)

const (
	compName   = "WebDriver proxy"
	timeout    = 1 * time.Second
	envTimeout = 1 * time.Second // some environments such as Android take a long time to start up.
)

// Proxy starts a WebDriver protocol proxy.
type Proxy struct {
	env     environment.Env
	Address string
	port    int
}

// New creates a new Proxy object.
func New(env environment.Env, port int) (*Proxy, error) {
	return &Proxy{
		env:     env,
		Address: net.JoinHostPort("localhost", strconv.Itoa(port)),
		port:    port,
	}, nil
}

// Component returns the name used in error messages.
func (*Proxy) Name() string {
	return compName
}

// Start configures the proxy with handlers, starts its listen loop, and waits for it to respond to a health check.
func (p *Proxy) Start(ctx context.Context) error {
	log.Printf("launching server at: %v", p.Address)

	http.Handle("/wd/hub/", driverhub.NewHandler(p.env))

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
