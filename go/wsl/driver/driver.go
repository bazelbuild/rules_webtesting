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
	"strings"
	"sync"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/cmdhelper"
	"github.com/bazelbuild/rules_webtesting/go/httphelper"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/portpicker"
	"github.com/bazelbuild/rules_webtesting/go/webdriver"
)

const compName = "WSL Driver"

// Driver is wrapper around a running WebDriver endpoint binary.
type Driver struct {
	Address string

	cmd      *exec.Cmd
	waitChan chan error

	shutdownEndpoint bool

	// Mutex to prevent overlapping commands to remote end.
	mu sync.Mutex
}

// New creates starts a WebDriver endpoint binary based on caps. Argument caps should just be
// the google:wslConfig capability extracted from the capabilities passed into a new session request.
func New(ctx context.Context, localHost string, caps map[string]interface{}) (*Driver, error) {
	wslCaps, err := extractWSLCaps(caps)
	if err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(context.Background(), wslCaps.binary, wslCaps.args...)

	cmd.Env = cmdhelper.BulkUpdateEnv(os.Environ(), wslCaps.env)

	stdout := os.Stdout

	if wslCaps.stdout != "" {
		s, err := os.Create(wslCaps.stdout)
		if err != nil {
			return nil, err
		}
		stdout = s
	}
	cmd.Stdout = stdout

	stderr := os.Stderr

	if wslCaps.stderr != "" {
		if wslCaps.stderr == wslCaps.stdout {
			stderr = stdout
		} else {
			s, err := os.Create(wslCaps.stderr)
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

	deadline, cancel := context.WithTimeout(ctx, wslCaps.timeout)
	defer cancel()

	hostPort := net.JoinHostPort(localHost, strconv.Itoa(wslCaps.port))

	statusURL := fmt.Sprintf("http://%s/status", hostPort)

	errChan := make(chan error, 1)
	driverChan := make(chan error, 1)
	go func() {
		err := cmd.Wait()
		log.Printf("%s has exited: %v", wslCaps.binary, err)
		errChan <- err
		driverChan <- err
		if stdout != os.Stdout {
			stdout.Close()
		}
		if stderr != os.Stderr {
			stdout.Close()
		}
	}()

	for {
		select {
		case err := <-errChan:
			return nil, err
		default:
		}

		if response, err := httphelper.Get(deadline, statusURL); err == nil {
			if wslCaps.statusBroken {
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
			go cmd.Process.Kill()
			return nil, deadline.Err()
		}

		time.Sleep(10 * time.Millisecond)
	}

	return &Driver{
		Address:  fmt.Sprintf("http://%s", hostPort),
		cmd:      cmd,
		shutdownEndpoint: wslCaps.shutdownEndpoint,
		waitChan: driverChan,
	}, nil
}

type wslCaps struct {
	binary       string
	args         []string
	port         int
	timeout      time.Duration
	env          map[string]string
	shutdownEndpoint	bool	
	statusBroken bool
	stdout       string
	stderr       string
}

func extractWSLCaps(caps map[string]interface{}) (*wslCaps, error) {
	binary, ok := caps["binary"].(string)

	if !ok {
		return nil, errors.New("binary not set or wrong type")
	}

	portNum, _ := caps["port"].(float64)
	port := int(portNum)

	if port == 0 {
		p, err := portpicker.PickUnusedPort()
		if err != nil {
			return nil, err
		}
		port = p
	}

	argsInterface, ok := caps["args"].([]interface{})

	if !ok {
		return nil, errors.New("args not set or wrong type")
	}

	var args []string

	portStr := fmt.Sprintf("%d", port)
	for _, argInterface := range argsInterface {
		arg, ok := argInterface.(string)
		if !ok {
			return nil, errors.New("arg had wrong type")
		}

		arg = strings.Replace(arg, "%WSL:PORT%", portStr, -1)
		args = append(args, arg)
	}

	timeout := 1 * time.Second

	if t, ok := caps["timeout"].(float64); ok {
		timeout = time.Duration(t*1000) * time.Millisecond
	}

	env := map[string]string{}

	if e, ok := caps["env"].(map[string]interface{}); ok {
		for k, v := range e {
			if vs, ok := v.(string); ok {
				env[k] = strings.Replace(vs, "%WSL:PORT%", portStr, -1)
			}
		}
	}

	shutdownEndpoint := false
	if s, ok := caps["shutdownEndpoint"].(bool); ok {
		shutdownEndpoint = s
	}

	statusBroken := false
	if s, ok := caps["statusBroken"].(bool); ok {
		statusBroken = s
	}

	stdout := ""
	if s, ok := caps["stdout"].(string); ok {
		stdout = s
	}

	stderr := ""
	if s, ok := caps["stderr"].(string); ok {
		stderr = s
	}

	return &wslCaps{
		binary:       binary,
		args:         args,
		port:         port,
		timeout:      timeout,
		env:          env,
		shutdownEndpoint: shutdownEndpoint,
		statusBroken: statusBroken,
		stdout:       stdout,
		stderr:       stderr,

	}, nil
}

// Forward forwards a request to the WebDriver endpoint managed by this server.
func (d *Driver) Forward(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if err := httphelper.Forward(r.Context(), d.Address, "", w, r); err != nil {
		errorResponse(w, http.StatusInternalServerError, 13, "unknown error", err.Error())
	}
}

// NewSessionW3C creates a new session using the W3C protocol.
func (d *Driver) NewSession(ctx context.Context, caps *capabilities.Capabilities, w http.ResponseWriter) (string, error) {
	always := map[string]interface{}{}
	for k, v := range caps.AlwaysMatch {
		if k != "google:wslConfig" {
			always[k] = v
		}
	}

	var first []map[string]interface{}

	for _, fm := range caps.FirstMatch {
		newFM := map[string]interface{}{}
		for k, v := range fm {
			if k != "google:wslConfig" {
				newFM[k] = v
			}
		}
		first = append(first, newFM)
	}

	wd, err := webdriver.CreateSession(ctx, d.Address, 1, &capabilities.Capabilities{
		AlwaysMatch:  always,
		FirstMatch:   first,
		W3CSupported: caps.W3CSupported,
	})

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
	w.Header().Set("Cache-Control", "no-cache")
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
	w.Header().Set("Cache-Control", "no-cache")
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
	case err := <-d.waitChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Kill kills a running WebDriver server.
func (d *Driver) Shutdown(ctx context.Context) error {
	if !d.shutdownEndpoint {
		return d.cmd.Process.Kill()
	}

	httphelper.Get(ctx, d.Address + "/shutdown")

	if err := d.Wait(ctx); err != nil {
		return d.cmd.Process.Kill()
	}
	return nil
}

func errorResponse(w http.ResponseWriter, httpStatus, status int, err, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
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
