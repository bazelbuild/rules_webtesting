// Copyright 2018 Google Inc.
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

// Package driver launches a WebDriver driver endpoint binary (e.g. ChromeDriver, SafariDriver, etc)
// based on a google:wslConfig capability.
package driver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/cmdhelper"
	"github.com/bazelbuild/rules_webtesting/go/httphelper"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/webdriver"
)

const compName = "WSL Driver"

// Driver is wrapper around a running WebDriver endpoint binary.
type Driver struct {
	Address      string
	caps         *wslCaps
	stopped      chan error
	cmd          *exec.Cmd
	portRecycler PortRecycler

	// Mutex to prevent overlapping commands to remote end.
	mu sync.Mutex
}

type wslCaps struct {
	binary      string
	args        []string
	port        int
	timeout     time.Duration
	env         map[string]string
	shutdown    bool
	status      bool
	stdout      string
	stderr      string
	quitTimeout time.Duration
}

// PortRecycler is an interface for an object that includes ports that are owned by this driver.
type PortRecycler interface {
	RecyclePorts() error
}

// New creates starts a WebDriver endpoint binary based on caps. Argument caps should just be
// the google:wslConfig capability extracted from the capabilities passed into a new session request.
func New(ctx context.Context, localHost, sessionID string, caps map[string]interface{}, portRecycler PortRecycler) (*Driver, error) {
	wslCaps, err := extractWSLCaps(sessionID, caps)
	if err != nil {
		return nil, err
	}
	hostPort := net.JoinHostPort(localHost, strconv.Itoa(wslCaps.port))

	d := &Driver{
		Address:      fmt.Sprintf("http://%s", hostPort),
		caps:         wslCaps,
		stopped:      make(chan error),
		portRecycler: portRecycler,
	}

	errChan, err := d.startDriver()
	if err != nil {
		return nil, err
	}

	deadline, cancel := context.WithTimeout(ctx, d.caps.timeout)
	defer cancel()

	statusURL := fmt.Sprintf("http://%s/status", hostPort)

	for {
		select {
		case err := <-errChan:
			return nil, err
		default:
		}

		if response, err := httphelper.Get(deadline, statusURL); err == nil {
			if !d.caps.status {
				// just wait for successful connection because status endpoint doesn't work.
				break
			}
			if response.StatusCode == http.StatusOK {
				respJSON := map[string]interface{}{}
				if err := json.NewDecoder(response.Body).Decode(&respJSON); err == nil {
					log.Printf("Response: %+v", respJSON)
					if status, ok := respJSON["status"].(float64); ok {
						if int(status) == 0 {
							break
						}
					} else if value, ok := respJSON["value"].(map[string]interface{}); ok {
						if ready, _ := value["ready"].(bool); ready {
							break
						}
					}
				}
			}
		}

		if deadline.Err() != nil {
			if d.cmd != nil {
				go d.cmd.Process.Kill()
			}
			return nil, deadline.Err()
		}

		time.Sleep(10 * time.Millisecond)
	}

	return d, nil
}

func extractWSLCaps(sessionID string, caps map[string]interface{}) (*wslCaps, error) {
	binary := ""
	if b, ok := caps["binary"]; ok {
		bs, ok := b.(string)
		if !ok {
			return nil, fmt.Errorf("binary %#v is not a string", b)
		}
		binary = bs
	}

	port := 0
	if p, ok := caps["port"]; ok {
		switch pt := p.(type) {
		case float64:
			port = int(pt)
		case string:
			pi, err := strconv.Atoi(pt)
			if err != nil {
				return nil, err
			}
			port = pi
		default:
			return nil, fmt.Errorf("port %#v is not a number or string", p)
		}
	}

	if port == 0 {
		return nil, errors.New(`port must be set (use "%WSLPORT:DRIVER%" if you don't care what port is used)`)
	}

	var args []string
	if a, ok := caps["args"]; ok {
		if binary == "" {
			return nil, fmt.Errorf("args set to %#v when binary is not set", a)
		}

		argsInterface, ok := a.([]interface{})
		if !ok {
			return nil, fmt.Errorf("args %#v is not a list", a)
		}

		for _, argInterface := range argsInterface {
			arg, ok := argInterface.(string)
			if !ok {
				return nil, fmt.Errorf("element %#v in args is not a string", argInterface)
			}
			args = append(args, arg)
		}
	}

	timeout := 1 * time.Second
	if t, ok := caps["timeout"]; ok {
		switch tt := t.(type) {
		case float64:
			// Incoming value is in seconds.
			to, err := time.ParseDuration(fmt.Sprintf("%fs", tt))
			if err != nil {
				return nil, err
			}
			timeout = to
		case string:
			to, err := time.ParseDuration(tt)
			if err != nil {
				return nil, err
			}
			timeout = to
		default:
			return nil, fmt.Errorf("timeout %#v is not a number or string", t)
		}
	}

	env := map[string]string{}
	if e, ok := caps["env"]; ok {
		if binary == "" {
			return nil, fmt.Errorf("env set to %#v when binary is not set", e)
		}
		em, ok := e.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("env %#v is not a map", e)
		}
		for k, v := range em {
			vs, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("value %#v for key %q in env is not a string", v, k)
			}
			env[k] = vs
		}
	}

	shutdown := true
	if s, ok := caps["shutdown"]; ok {
		sb, ok := s.(bool)
		if !ok {
			return nil, fmt.Errorf("shutdown %#v is not a boolean", s)
		}
		shutdown = sb
	}

	status := true
	if s, ok := caps["status"]; ok {
		sb, ok := s.(bool)
		if !ok {
			return nil, fmt.Errorf("status %#v is not a boolean", s)
		}
		status = sb
	}

	stdout := ""
	if s, ok := caps["stdout"]; ok {
		if binary == "" {
			return nil, fmt.Errorf("stdout set to %#v when binary is not set", s)
		}
		sb, ok := s.(string)
		if !ok {
			return nil, fmt.Errorf("stdout %#v is not a string", s)
		}
		stdout = sb
	}

	stderr := ""
	if s, ok := caps["stderr"]; ok {
		if binary == "" {
			return nil, fmt.Errorf("stderr set to %#v when binary is not set", s)
		}
		sb, ok := s.(string)
		if !ok {
			return nil, fmt.Errorf("stderr %#v is not a string", s)
		}
		stderr = sb
	}

	quitTimeout := 0 * time.Second
	if t, ok := caps["quitTimeout"]; ok {
		switch tt := t.(type) {
		case float64:
			// Incoming value is in seconds.
			to, err := time.ParseDuration(fmt.Sprintf("%fs", tt))
			if err != nil {
				return nil, err
			}
			quitTimeout = to
		case string:
			to, err := time.ParseDuration(tt)
			if err != nil {
				return nil, err
			}
			quitTimeout = to
		default:
			return nil, fmt.Errorf("quitTimeout %#v is not a number or string", t)
		}
	}

	return &wslCaps{
		binary:      binary,
		args:        args,
		port:        port,
		timeout:     timeout,
		env:         env,
		shutdown:    shutdown,
		status:      status,
		stdout:      stdout,
		stderr:      stderr,
		quitTimeout: quitTimeout,
	}, nil
}

func (d *Driver) startDriver() (chan error, error) {
	errChan := make(chan error)
	if d.caps.binary == "" {
		return errChan, nil
	}

	cmd := exec.CommandContext(context.Background(), d.caps.binary, d.caps.args...)

	cmd.Env = cmdhelper.BulkUpdateEnv(os.Environ(), d.caps.env)

	stdout := os.Stdout

	if d.caps.stdout != "" {
		s, err := os.Create(d.caps.stdout)
		if err != nil {
			return nil, err
		}
		stdout = s
	}
	cmd.Stdout = stdout

	stderr := os.Stderr

	if d.caps.stderr != "" {
		if d.caps.stderr == d.caps.stdout {
			stderr = stdout
		} else {
			s, err := os.Create(d.caps.stderr)
			if err != nil {
				return nil, err
			}
			stderr = s
		}
	}
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	go func() {
		err := cmd.Wait()
		log.Printf("%s has exited: %v", d.caps.binary, err)

		if err := d.portRecycler.RecyclePorts(); err != nil {
			log.Printf("Error cleaning up used ports: %v", err)
		}

		errChan <- err
		d.stopped <- err
		if stdout != os.Stdout {
			stdout.Close()
		}
		if stderr != os.Stderr {
			stdout.Close()
		}
	}()

	d.cmd = cmd

	return errChan, nil
}

// Forward forwards a request to the WebDriver endpoint managed by this server.
func (d *Driver) Forward(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if err := httphelper.Forward(ctx, d.Address, "", w, r); err != nil {
		errorResponse(w, http.StatusInternalServerError, 13, "unknown error", err.Error())
	}
}

// NewSessionW3C creates a new session using the W3C protocol.
func (d *Driver) NewSession(ctx context.Context, caps *capabilities.Capabilities, w http.ResponseWriter) (string, error) {
	// IEDriver does not handle capabilities with unknown prefixes, so they must be stripped.
	if bn, ok := caps.AlwaysMatch["browserName"].(string); ok && bn == "internet explorer" {
		caps = caps.Strip("goog:chromeOptions", "moz:firefoxOptions")
	}

	wd, err := webdriver.CreateSession(ctx, d.Address, 1, caps.Strip("google:wslConfig", "google:sessionId"))

	if err != nil {
		return "", err
	}

	if wd.W3C() {
		writeW3CNewSessionResponse(wd, w)
	} else {
		writeJWPNewSessionResponse(wd, w)
	}

	return wd.SessionID(), nil
}

func writeW3CNewSessionResponse(wd webdriver.WebDriver, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	httphelper.SetDefaultResponseHeaders(w.Header())
	w.WriteHeader(http.StatusOK)

	respJSON := map[string]interface{}{
		"value": map[string]interface{}{
			"capabilities": wd.Capabilities(),
			"sessionId":    wd.SessionID(),
		},
	}

	json.NewEncoder(w).Encode(respJSON)
}

func writeJWPNewSessionResponse(wd webdriver.WebDriver, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	httphelper.SetDefaultResponseHeaders(w.Header())
	w.WriteHeader(http.StatusOK)

	respJSON := map[string]interface{}{
		"value":     wd.Capabilities(),
		"sessionId": wd.SessionID(),
		"status":    0,
	}

	json.NewEncoder(w).Encode(respJSON)
}

// Wait waits for the driver binary to exit, and returns an error if the binary exited with an error.
func (d *Driver) Wait(ctx context.Context) error {
	select {
	case err := <-d.stopped:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Quit forwards a Quit command and shuts down a driver.
func (d *Driver) Quit(w http.ResponseWriter, r *http.Request) {
	defer func() {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		if err := d.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down: %v", err)
		}
	}()

	if d.caps.quitTimeout <= 0 {
		// If no quitTimeout is set, then forward normally
		d.Forward(r.Context(), w, r)
		return
	}

	// Otherwise forward with timeout, and swallow any error response.
	ctx, cancel := context.WithTimeout(r.Context(), d.caps.quitTimeout)
	defer cancel()

	d.Forward(ctx, &fakeResponseWriter{
		prefix: "quit session",
		header: http.Header{},
	}, r)

	respJSON := map[string]interface{}{
		"status": 0,
		"value":  nil,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respJSON)
}

// Shutdown shuts down a running WebDriver server.
func (d *Driver) Shutdown(ctx context.Context) error {
	if d.cmd == nil {
		close(d.stopped)
		return nil
	}
	if d.caps.shutdown {
		httphelper.Get(ctx, d.Address+"/shutdown")
	} else if err := d.cmd.Process.Signal(syscall.SIGTERM); err != nil {
		if err := d.cmd.Process.Signal(os.Interrupt); err != nil {
			return d.cmd.Process.Kill()
		}
	}

	if err := d.Wait(ctx); err != nil {
		if err == context.DeadlineExceeded {
			return d.cmd.Process.Kill()
		}
		return err
	}
	return nil
}

func errorResponse(w http.ResponseWriter, httpStatus, status int, err, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	httphelper.SetDefaultResponseHeaders(w.Header())
	w.WriteHeader(httpStatus)

	respJSON := map[string]interface{}{
		"status": status,
		"value": map[string]interface{}{
			"error":   err,
			"message": message,
		},
	}

	json.NewEncoder(w).Encode(respJSON)
}

// fakeResponseWriter is used when we don't want to send responses from forwarded commands back to caller.
type fakeResponseWriter struct {
	prefix string
	header http.Header
}

func (f *fakeResponseWriter) Header() http.Header {
	return f.header
}

func (f *fakeResponseWriter) Write(b []byte) (int, error) {
	log.Printf("%s body: %s", f.prefix, string(b))
	return len(b), nil
}

func (f *fakeResponseWriter) WriteHeader(statusCode int) {
	log.Printf("%s status code: %d", f.prefix, statusCode)
}
