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

// Package hub launches WebDriver servers and correctly dispatches requests to the correct server
// based on session id.
package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/bazelbuild/rules_webtesting/go/wsl/driver"
)

// A Hub is an HTTP handler that manages incoming WebDriver requests.
type Hub struct {
	// Mutex to protext access to sessions.
	mu       sync.RWMutex
	sessions map[string]*driver.Driver
}

// New creates a new Hub.
func New() *Hub {
	return &Hub{
		sessions: map[string]*driver.Driver{},
	}
}

func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1:]

	if len(path) < 1 || path[0] != "session" {
		errorResponse(w, http.StatusNotFound, 9, "unknown command", fmt.Sprintf("%q is not a known command", r.URL.Path))
		return
	}

	if r.Method == http.MethodPost && len(path) == 1 {
		h.newSession(w, r)
		return
	}

	if len(path) < 2 {
		errorResponse(w, http.StatusMethodNotAllowed, 9, "unknown method", fmt.Sprintf("%s is not a supported method for /session", r.Method))
		return
	}

	driver := h.driver(path[1])
	if driver == nil {
		errorResponse(w, http.StatusNotFound, 6, "invalid session id", fmt.Sprintf("%q is not an active session", path[1]))
		return
	}

	if r.Method == http.MethodDelete && len(path) == 2 {
		h.quitSession(path[1], driver, w, r)
		return
	}

	driver.Forward(w, r)
}

func (h *Hub) driver(session string) *driver.Driver {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.sessions[session]
}

type capabilities struct {
	DesiredCapabilities  map[string]interface{} `json:"desiredCapabilities,omitempty"`
	RequiredCapabilities map[string]interface{} `json:"requiredCapabilities,omitempty"`
	Capabilities         *w3cCaps               `json:"capabilities,omitempty"`
}

type w3cCaps struct {
	AlwaysMatch map[string]interface{}   `json:"alwaysMatch,omitempty"`
	FirstMatch  []map[string]interface{} `json:"firstMatch,omitempty"`
}

func (h *Hub) newSession(w http.ResponseWriter, r *http.Request) {
	reqJSON := capabilities{}

	if err := json.NewDecoder(r.Body).Decode(&reqJSON); err != nil {
		errorResponse(w, http.StatusBadRequest, 13, "invalid argument", err.Error())
		return
	}

	var driver *driver.Driver
	var session string

	if reqJSON.Capabilities != nil {
		session, driver = h.newSessionW3CCaps(r.Context(), *reqJSON.Capabilities, w)
	}

	if driver == nil && reqJSON.DesiredCapabilities != nil {
		session, driver = h.newSessionJWPCaps(r.Context(), reqJSON.DesiredCapabilities, reqJSON.RequiredCapabilities, w)
	}

	if driver == nil {
		errorResponse(w, http.StatusInternalServerError, 33, "session not created", "unable to create session")
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	h.sessions[session] = driver
}

func (h *Hub) newSessionW3CCaps(ctx context.Context, caps w3cCaps, w http.ResponseWriter) (string, *driver.Driver) {
	if caps.AlwaysMatch != nil {
		wslConfig, ok := caps.AlwaysMatch["google:wslConfig"].(map[string]interface{})
		if ok {
			d, err := driver.New(ctx, wslConfig)
			if err != nil {
				return "", nil
			}

			s, err := d.NewSessionW3C(ctx, caps.AlwaysMatch, caps.FirstMatch, w)
			if err != nil {
				d.Kill()
				return "", nil
			}

			return s, d
		}
	}

	for _, fm := range caps.FirstMatch {
		wslConfig, ok := fm["google:wslConfig"].(map[string]interface{})

		if ok {
			d, err := driver.New(ctx, wslConfig)
			if err != nil {
				continue
			}

			s, err := d.NewSessionW3C(ctx, caps.AlwaysMatch, []map[string]interface{}{fm}, w)
			if err != nil {
				d.Kill()
				continue
			}

			return s, d
		}
	}

	return "", nil
}

func (h *Hub) newSessionJWPCaps(ctx context.Context, desired, required map[string]interface{}, w http.ResponseWriter) (string, *driver.Driver) {
	if required != nil {
		wslConfig, ok := required["google:wslConfig"].(map[string]interface{})
		if ok {
			d, err := driver.New(ctx, wslConfig)
			if err != nil {
				return "", nil
			}

			s, err := d.NewSessionJWP(ctx, desired, required, w)
			if err != nil {
				d.Kill()
				return "", nil
			}

			return s, d
		}
	}

	if desired != nil {
		wslConfig, ok := desired["google:wslConfig"].(map[string]interface{})

		if ok {
			d, err := driver.New(ctx, wslConfig)
			if err != nil {
				return "", nil
			}

			s, err := d.NewSessionJWP(ctx, desired, required, w)
			if err != nil {
				d.Kill()
				return "", nil
			}

			return s, d
		}
	}

	return "", nil
}

func (h *Hub) quitSession(session string, driver *driver.Driver, w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	driver.Forward(w, r)

	if err := driver.Kill(); err != nil {
		log.Printf("Error killing driver: %v", err)
	}

	driver.Wait()

	delete(h.sessions, session)
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
