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
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/launcher/environments/environment"
	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/healthreporter"
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/webdriver"
	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
	"github.com/bazelbuild/rules_webtesting/go/util/httphelper"
	"github.com/gorilla/mux/mux"
)

const envTimeout = 5 * time.Minute // some environments such as Android take a long time to start up.

// WebDriverHub routes message to the various WebDriver sessions.
type WebDriverHub struct {
	*mux.Router
	environment.Env
	*metadata.Metadata
	*http.Client
	diagnostics.Diagnostics

	healthyOnce sync.Once

	mu       sync.RWMutex
	sessions map[string]http.Handler
	nextID   int
}

// NewHandler creates a handler for /wd/hub paths that delegates to a WebDriver server instance provided by env.
func NewHandler(env environment.Env, m *metadata.Metadata, d diagnostics.Diagnostics) http.Handler {
	h := &WebDriverHub{
		Router:      mux.NewRouter(),
		Env:         env,
		sessions:    map[string]http.Handler{},
		Client:      &http.Client{},
		Diagnostics: d,
		Metadata:    m,
	}

	h.Path("/wd/hub/session").Methods("POST").HandlerFunc(h.createSession)
	h.Path("/wd/hub/session").HandlerFunc(unknownMethod)
	h.PathPrefix("/wd/hub/session/{sessionID}").HandlerFunc(h.routeToSession)
	h.PathPrefix("/wd/hub/{command}").HandlerFunc(h.defaultForward)
	h.PathPrefix("/").HandlerFunc(unknownCommand)

	return h
}

// Name is the name of the component used in error messages.
func (h *WebDriverHub) Name() string {
	return "WebDriver Hub"
}

func (h *WebDriverHub) AddSession(id string, session http.Handler) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.sessions[id] = session
}

func (h *WebDriverHub) RemoveSession(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.sessions, id)
}

func (h *WebDriverHub) GetSession(id string) http.Handler {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.sessions[id]
}

func (h *WebDriverHub) NextID() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	id := h.nextID
	h.nextID++
	return id
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
	var desired map[string]interface{}

	if err := h.waitForHealthyEnv(ctx); err != nil {
		sessionNotCreated(w, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sessionNotCreated(w, err)
		return
	}

	j := struct {
		Desired map[string]interface{} `json:"desiredCapabilities"`
	}{}

	if err := json.Unmarshal(body, &j); err != nil {
		sessionNotCreated(w, err)
		return
	}

	if j.Desired == nil {
		sessionNotCreated(w, errors.New(h.Name(), "new session request body missing desired capabilities"))
		return
	}

	id := h.NextID()

	desired, err = h.Env.StartSession(ctx, id, j.Desired)
	if err != nil {
		sessionNotCreated(w, err)
		return
	}

	log.Printf("Caps: %+v", desired)
	// TODO(fisherii) parameterize attempts based on browser metadata
	driver, err := webdriver.CreateSession(ctx, h.Env.WDAddress(ctx), 3, desired)
	if err != nil {
		if err2 := h.Env.StopSession(ctx, id); err2 != nil {
			log.Printf("error stopping session after failing to launch webdriver: %v", err2)
		}
		sessionNotCreated(w, err)
		return
	}

	session, err := CreateSession(id, h, driver, desired)
	if err != nil {
		sessionNotCreated(w, err)
		return
	}

	h.AddSession(driver.SessionID(), session)

	respJSON := map[string]interface{}{
		"status":    0,
		"sessionId": driver.SessionID(),
		"value":     driver.Capabilities(),
	}

	bytes, err := json.Marshal(respJSON)
	if err != nil {
		unknownError(w, err)
		return
	}

	w.Header().Set("Content-Type", contentType)
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
	return h.Env.Healthy(ctx)
}
