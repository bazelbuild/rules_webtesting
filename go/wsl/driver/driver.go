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
	"io"
	"net/http"
	"os"
	"os/exec"
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
	*exec.Cmd
	Address string

	// Mutex to prevent overlapping commands to remote end.
	mu sync.Mutex
}

// New creates starts a WebDriver endpoint binary based on caps. Argument caps should just be
// the google:wslConfig capability extracted from the capabilities passed into a new session request.
func New(ctx context.Context, caps map[string]interface{}) (*Driver, error) {
	wslCaps, err := extractWSLCaps(caps)
	if err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(context.Background(), wslCaps.binary, wslCaps.args...)

	cmd.Env = cmdhelper.BulkUpdateEnv(os.Environ(), wslCaps.env)

	// TODO(DrMarcII): figure out how to deal with log files.
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	go io.Copy(os.Stderr, stderr)

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	deadline, cancel := context.WithTimeout(ctx, wslCaps.timeout)
	defer cancel()

	statusURL := fmt.Sprintf("http://localhost:%d/status", wslCaps.port)

	for {
		if response, err := httphelper.Get(deadline, statusURL); err == nil && response.StatusCode == http.StatusOK {
			respJSON := map[string]interface{}{}
			if err := json.NewDecoder(response.Body).Decode(&respJSON); err == nil {
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

		if deadline.Err() != nil {
			return nil, deadline.Err()
		}

		time.Sleep(10 * time.Millisecond)
	}

	return &Driver{
		Cmd:     cmd,
		Address: fmt.Sprintf("http://localhost:%d", wslCaps.port),
	}, nil
}

type wslCaps struct {
	binary  string
	args    []string
	port    int
	timeout time.Duration
	env     map[string]string
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

	if e, ok := caps["timeout"].(map[string]interface{}); ok {
		for k, v := range e {
			if vs, ok := v.(string); ok {
				env[k] = strings.Replace(vs, "%WSL:PORT%", portStr, -1)
			}
		}
	}

	return &wslCaps{
		binary:  binary,
		args:    args,
		port:    port,
		timeout: timeout,
		env:     env,
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
func (d *Driver) NewSessionW3C(ctx context.Context, alwaysMatch map[string]interface{}, firstMatch []map[string]interface{}, w http.ResponseWriter) (string, error) {
	delete(alwaysMatch, "google:wslConfig")

	for _, fm := range firstMatch {
		delete(fm, "google:wslConfig")
	}

	wd, err := webdriver.CreateSession(ctx, d.Address, 1, capabilities.Spec{
		Always: alwaysMatch,
		First:  firstMatch,
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

// NewSessionJWP creates a new session using the Selenium JSON wire protocol.
func (d *Driver) NewSessionJWP(ctx context.Context, desired, required map[string]interface{}, w http.ResponseWriter) (string, error) {
	caps := map[string]interface{}{}

	for k, v := range desired {
		if k != "google:wslConfig" {
			caps[k] = v
		}
	}

	for k, v := range required {
		if k != "google:wslConfig" {
			caps[k] = v
		}
	}

	wd, err := webdriver.CreateSession(ctx, d.Address, 1, capabilities.Spec{
		OSSCaps: caps,
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
	respJSON := map[string]interface{}{
		"value": map[string]interface{}{
			"capabilities": wd.Capabilities(),
			"sessionId":    wd.SessionID(),
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respJSON)
}

func writeJWPNewSessionResponse(wd webdriver.WebDriver, w http.ResponseWriter) {
	respJSON := map[string]interface{}{
		"value":     wd.Capabilities(),
		"sessionId": wd.SessionID(),
		"status":    0,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respJSON)
}

// Kill kills a running WebDriver server.
func (d *Driver) Kill() error {
	return d.Process.Kill()
}

func errorResponse(w http.ResponseWriter, httpStatus, status int, err, message string) {
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
