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
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/httphelper"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/wsl/driver"
	"github.com/bazelbuild/rules_webtesting/go/wsl/resolver"
)

// A Hub is an HTTP handler that manages incoming WebDriver requests.
type Hub struct {
	// Mutex to protext access to sessions.
	mu       sync.RWMutex
	sessions map[string]*driver.Driver

	localHost string
	uploader  http.Handler
}

// New creates a new Hub.
func New(localHost string, uploader http.Handler) *Hub {
	return &Hub{
		sessions:  map[string]*driver.Driver{},
		localHost: localHost,
		uploader:  uploader,
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

	if len(path) == 3 && path[2] == "file" {
		h.uploader.ServeHTTP(w, r)
		return
	}

	driver.Forward(r.Context(), w, r)
}

func (h *Hub) driver(session string) *driver.Driver {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.sessions[session]
}

func (h *Hub) newSession(w http.ResponseWriter, r *http.Request) {
	reqJSON := map[string]interface{}{}

	if err := json.NewDecoder(r.Body).Decode(&reqJSON); err != nil {
		errorResponse(w, http.StatusBadRequest, 13, "invalid argument", err.Error())
		return
	}

	caps, err := capabilities.FromNewSessionArgs(reqJSON)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, 13, "invalid argument", err.Error())
		return
	}

	session, driver, err := h.newSessionFromCaps(r.Context(), caps, w)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, 33, "session not created", fmt.Sprintf("unable to create session: %v", err))
		log.Printf("Error creating webdriver session: %v", err)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	h.sessions[session] = driver
}

func (h *Hub) newSessionFromCaps(ctx context.Context, caps *capabilities.Capabilities, w http.ResponseWriter) (string, *driver.Driver, error) {
	sessionID := "last"
	if i, ok := caps.AlwaysMatch["google:sessionId"]; ok {
		switch ii := i.(type) {
		case string:
			sessionID = ii
		case float64:
			sessionID = strconv.Itoa(int(ii))
		default:
			return "", nil, fmt.Errorf("google:sessionId %#v is not a string or number", i)
		}
	}
	resolver := resolver.New(sessionID)

	caps, err := caps.Resolve(resolver.Resolve)
	if err != nil {
		return "", nil, err
	}

	wslConfig, ok := caps.AlwaysMatch["google:wslConfig"].(map[string]interface{})

	if !ok {
		if err := resolver.RecyclePorts(); err != nil {
			log.Printf("Error recycling ports: %v", err)
		}
		return "", nil, errors.New("alwaysMatch capabilites must include google:wslConfig")
	}

	d, err := driver.New(ctx, h.localHost, sessionID, wslConfig, resolver)
	if err != nil {
		if err := resolver.RecyclePorts(); err != nil {
			log.Printf("Error recycling ports: %v", err)
		}
		return "", nil, err
	}

	s, err := d.NewSession(ctx, caps, w)
	if err != nil {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		if err := d.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down driver: %v", err)
		}
		return "", nil, err
	}

	return s, d, nil
}

func (h *Hub) quitSession(session string, driver *driver.Driver, w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	driver.Quit(w, r)

	delete(h.sessions, session)
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
