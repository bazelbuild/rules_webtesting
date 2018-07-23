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

// Package driverhub provides a handler for proxying connections to a Selenium server.
package driverhub

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/errors"
	"github.com/bazelbuild/rules_webtesting/go/healthreporter"
	"github.com/bazelbuild/rules_webtesting/go/httphelper"
	"github.com/bazelbuild/rules_webtesting/go/metadata"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/webdriver"
	"github.com/bazelbuild/rules_webtesting/go/wtl/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/wtl/environment"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy/driverhub/debugger"
	"github.com/gorilla/mux"
)

const envTimeout = 5 * time.Minute // some environments such as Android take a long time to start up.

// WebDriverHub routes message to the various WebDriver sessions.
type WebDriverHub struct {
	*mux.Router
	environment.Env
	*metadata.Metadata
	*http.Client
	diagnostics.Diagnostics
	Proxy    *proxy.Proxy
	Debugger *debugger.Debugger

	healthyOnce sync.Once

	mu               sync.RWMutex
	sessions         map[string]*WebDriverSession
	reusableSessions []*WebDriverSession
	nextID           int
}

// NewHandler creates a handler for /wd/hub paths that delegates to a WebDriver server instance provided by env.
func HTTPHandlerProvider(p *proxy.Proxy) (proxy.HTTPHandler, error) {
	var d *debugger.Debugger
	if p.Metadata.DebuggerPort != 0 {
		d = debugger.New(p.Metadata.DebuggerPort)
	}
	h := &WebDriverHub{
		Router:      mux.NewRouter(),
		Env:         p.Env,
		sessions:    map[string]*WebDriverSession{},
		Client:      &http.Client{},
		Diagnostics: p.Diagnostics,
		Metadata:    p.Metadata,
		Proxy:       p,
		Debugger:    d,
	}

	h.Path("/wd/hub/session").Methods("POST").HandlerFunc(h.createSession)
	h.Path("/wd/hub/session").HandlerFunc(unknownMethod)
	h.PathPrefix("/wd/hub/session/{sessionID}").HandlerFunc(h.routeToSession)
	h.PathPrefix("/wd/hub/{command}").HandlerFunc(h.defaultForward)
	h.PathPrefix("/").HandlerFunc(unknownCommand)

	return h, nil
}

func (h *WebDriverHub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Debugger != nil {
		// allow debugger to pause for breakpoint, log to debugger front-end.
		h.Debugger.Request(r)
	}
	// TODO(DrMarcII) add support for breakpointing on responses.
	h.Router.ServeHTTP(w, r)
}

// Name is the name of the component used in error messages.
func (h *WebDriverHub) Name() string {
	return "WebDriver Hub"
}

// Healthy returns nil if the WebDriverHub is ready for use, and an error otherwise.
func (h *WebDriverHub) Healthy(ctx context.Context) error {
	if h.Debugger != nil {
		return h.Debugger.Healthy(ctx)
	}
	return nil
}

// AddSession adds a session to WebDriverHub.
func (h *WebDriverHub) AddSession(id string, session *WebDriverSession) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.sessions == nil {
		h.sessions = map[string]*WebDriverSession{}
	}
	h.sessions[id] = session
}

// RemoveSession removes a session from WebDriverHub.
func (h *WebDriverHub) RemoveSession(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.sessions == nil {
		return
	}
	delete(h.sessions, id)
}

// GetSession gets the session for a given WebDriver session id..
func (h *WebDriverHub) GetSession(id string) *WebDriverSession {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.sessions[id]
}

// NextID gets the next available internal id for a session.
func (h *WebDriverHub) NextID() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	id := h.nextID
	h.nextID++
	return id
}

// GetActiveSessions returns the ids for all currently active sessions.
func (h *WebDriverHub) GetActiveSessions() []string {
	result := []string{}
	h.mu.RLock()
	defer h.mu.RUnlock()
	for id := range h.sessions {
		result = append(result, id)
	}
	return result
}

// Shutdown  shuts down any running sessions.
func (h *WebDriverHub) Shutdown(ctx context.Context) error {
	for _, id := range h.GetActiveSessions() {
		session := h.GetSession(id)
		if session != nil {
			session.quit(ctx, false)
		}
	}
	for _, session := range h.reusableSessions {
		session.quit(ctx, false)
	}
	return nil
}

// GetReusableSession grabs a reusable session if one is available that matches caps.
func (h *WebDriverHub) GetReusableSession(ctx context.Context, caps *capabilities.Capabilities) (*WebDriverSession, bool) {
	if !capabilities.CanReuseSession(caps) {
		return nil, false
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	for i, session := range h.reusableSessions {
		if reflect.DeepEqual(caps, session.RequestedCaps) {
			h.reusableSessions = append(h.reusableSessions[:i], h.reusableSessions[i+1:]...)
			if err := session.WebDriver.Healthy(ctx); err == nil {
				return session, true
			}
			return session, true
		}
	}
	return nil, false
}

// AddReusableSession adds a session that can be reused.
func (h *WebDriverHub) AddReusableSession(session *WebDriverSession) error {
	if !capabilities.CanReuseSession(session.RequestedCaps) {
		return errors.New(h.Name(), "session is not reusable.")
	}
	h.reusableSessions = append(h.reusableSessions, session)
	return nil
}

func (h *WebDriverHub) routeToSession(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sessionID"]
	session := h.GetSession(sid)

	if session == nil {
		invalidSessionID(w, sid)
		return
	}
	session.ServeHTTP(w, r)
}

func (h *WebDriverHub) createSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Print("Creating session\n\n")

	if err := h.waitForHealthyEnv(ctx); err != nil {
		sessionNotCreated(w, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sessionNotCreated(w, err)
		return
	}

	j := map[string]interface{}{}

	if err := json.Unmarshal(body, &j); err != nil {
		sessionNotCreated(w, err)
		return
	}

	requestedCaps, err := capabilities.FromNewSessionArgs(j)
	if err != nil {
		sessionNotCreated(w, err)
		return
	}

	id := h.NextID()

	caps, err := h.Env.StartSession(ctx, id, requestedCaps)
	if err != nil {
		sessionNotCreated(w, err)
		return
	}

	log.Printf("Caps: %+v", caps)

	var session *WebDriverSession

	if reusable, ok := h.GetReusableSession(ctx, caps); ok {
		reusable.Unpause(id)
		session = reusable
	} else {
		// TODO(DrMarcII) parameterize attempts based on browser metadata
		driver, err := webdriver.CreateSession(ctx, h.Env.WDAddress(ctx), 3, caps.Strip("google:canReuseSession"))
		if err != nil {
			if err2 := h.Env.StopSession(ctx, id); err2 != nil {
				log.Printf("error stopping session after failing to launch webdriver: %v", err2)
			}
			sessionNotCreated(w, err)
			return
		}

		s, err := CreateSession(id, h, driver, caps)
		if err != nil {
			sessionNotCreated(w, err)
			return
		}
		session = s
	}

	h.AddSession(session.WebDriver.SessionID(), session)

	var respJSON map[string]interface{}

	if session.WebDriver.W3C() {
		respJSON = map[string]interface{}{
			"value": map[string]interface{}{
				"capabilities": session.WebDriver.Capabilities(),
				"sessionId":    session.WebDriver.SessionID(),
			},
		}
	} else {
		respJSON = map[string]interface{}{
			"value":     session.WebDriver.Capabilities(),
			"sessionId": session.WebDriver.SessionID(),
			"status":    0,
		}
	}

	bytes, err := json.Marshal(respJSON)
	if err != nil {
		unknownError(w, err)
		return
	}

	w.Header().Set("Content-Type", contentType)
	httphelper.SetDefaultResponseHeaders(w.Header())
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h *WebDriverHub) defaultForward(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := h.waitForHealthyEnv(ctx); err != nil {
		unknownError(w, err)
		return
	}

	if err := httphelper.Forward(r.Context(), h.Env.WDAddress(ctx), "/wd/hub/", w, r); err != nil {
		unknownError(w, err)
	}
}

func (h *WebDriverHub) waitForHealthyEnv(ctx context.Context) error {
	h.healthyOnce.Do(func() {
		healthyCtx, cancel := context.WithTimeout(ctx, envTimeout)
		defer cancel()
		// ignore error here as we will call and return Healthy below.
		healthreporter.WaitForHealthy(healthyCtx, h.Env)
	})
	err := h.Env.Healthy(ctx)
	if err != nil {
		err = errors.New(h.Name(), fmt.Sprintf("environment is unhealthy: %v", err))
	}
	return err
}
