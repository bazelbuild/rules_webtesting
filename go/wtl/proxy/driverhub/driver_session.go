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

package driverhub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/bazelbuild/rules_webtesting/go/errors"
	"github.com/bazelbuild/rules_webtesting/go/httphelper"
	"github.com/bazelbuild/rules_webtesting/go/metadata"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/webdriver"
	"github.com/bazelbuild/rules_webtesting/go/wtl/diagnostics"
	"github.com/gorilla/mux"
)

// WebDriverSession is an http.Handler for forwarding requests to a WebDriver session.
type WebDriverSession struct {
	*mux.Router
	diagnostics.Diagnostics
	WebDriverHub *WebDriverHub
	webdriver.WebDriver
	ID            int
	handler       HandlerFunc
	sessionPath   string
	RequestedCaps *capabilities.Capabilities
	Metadata      *metadata.Metadata

	mu      sync.RWMutex
	stopped bool
}

// HandlerProvider wraps another HandlerFunc to create a new HandlerFunc.
// If the second return value is false, then the provider did not construct a new HandlerFunc.
type HandlerProvider func(session *WebDriverSession, caps *capabilities.Capabilities, base HandlerFunc) (HandlerFunc, bool)

// HandlerFunc is a func for handling a request to a WebDriver session.
type HandlerFunc func(context.Context, Request) (Response, error)

// Request wraps a request to a WebDriver session.
type Request struct {
	// HTTP Method for this request (e.g. http.MethodGet, ...).
	Method string
	// Path of the request after the session id.
	Path []string
	// Any HTTP headers sent with the request.
	Header http.Header
	// The body of the request.
	Body []byte
}

// Response describes what response should be returned for a request to WebDriver session.
type Response struct {
	// HTTP status code to return (e.g. http.StatusOK, ...).
	Status int
	// Any HTTP Headers that should be included in the response.
	Header http.Header
	// The body of the response.
	Body []byte
}

var providers = []HandlerProvider{}

// HandlerProviderFunc adds additional handlers that will wrap any previously defined handlers.
//
// It is important to note that later added handlers will wrap earlier added handlers.
// E.g. if you call as follows:
//
//	HandlerProviderFunc(hp1)
//	HandlerProviderFunc(hp2)
//	HandlerProviderFunc(hp3)
//
// The generated handler will be constructed as follows:
//
//	hp3(session, caps, hp2(session, caps, hp1(session, caps, base)))
//
// where base is the a default function that forwards commands to WebDriver unchanged.
func HandlerProviderFunc(provider HandlerProvider) {
	providers = append(providers, provider)
}

func createHandler(session *WebDriverSession, caps *capabilities.Capabilities) HandlerFunc {
	handler := createBaseHandler(session.WebDriver)

	for _, provider := range providers {
		if h, ok := provider(session, caps, handler); ok {
			handler = h
		}
	}
	return handler
}

// CreateSession creates a WebDriverSession object.
func CreateSession(id int, hub *WebDriverHub, driver webdriver.WebDriver, caps *capabilities.Capabilities) (*WebDriverSession, error) {
	sessionPath := fmt.Sprintf("/wd/hub/session/%s", driver.SessionID())
	session := &WebDriverSession{
		ID:            id,
		Diagnostics:   hub.Diagnostics,
		WebDriverHub:  hub,
		WebDriver:     driver,
		sessionPath:   sessionPath,
		Router:        mux.NewRouter(),
		RequestedCaps: caps,
		Metadata:      hub.Metadata,
	}

	session.handler = createHandler(session, caps)
	// Route for commands for this session.
	session.PathPrefix(sessionPath).HandlerFunc(session.defaultHandler)
	// Route for commands for some other session. If this happens, the hub has
	// done something wrong.
	session.PathPrefix("/wd/hub/session/{sessionID}").HandlerFunc(session.wrongSession)
	// Route for all other paths that aren't WebDriver commands. This also implies
	// that the hub has done something wrong.
	session.PathPrefix("/").HandlerFunc(session.unknownCommand)

	return session, nil
}

// Name is the name of the component used in error messages.
func (s *WebDriverSession) Name() string {
	return "WebDriver Session Handler"
}

func (s *WebDriverSession) wrongSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.Severe(errors.New(s.Name(), "request routed to wrong session handler"))
	unknownError(w, fmt.Errorf("request for session %q was routed to handler for %q", vars["sessionID"], s.SessionID()))
}

func (s *WebDriverSession) unknownCommand(w http.ResponseWriter, r *http.Request) {
	s.Severe(errors.New(s.Name(), "unknown command routed to session handler"))
	unknownCommand(w, r)
}

// Quit can be called by handlers to quit this session.
func (s *WebDriverSession) Quit(ctx context.Context, _ Request) (Response, error) {
	if err := s.quit(ctx, capabilities.CanReuseSession(s.RequestedCaps)); err != nil {
		return ResponseFromError(err)
	}

	return Response{
		Status: http.StatusOK,
		Body:   []byte(`{"status": 0}`),
	}, nil
}

// Quit can be called by handlers to quit this session.
func (s *WebDriverSession) quit(ctx context.Context, reusable bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stopped = true

	var wdErr error

	if !reusable {
		wdErr = s.WebDriver.Quit(ctx)
		if wdErr != nil {
			s.Warning(wdErr)
		}
	}

	envErr := s.WebDriverHub.Env.StopSession(ctx, s.ID)
	if envErr != nil {
		s.Warning(envErr)
	}

	s.WebDriverHub.RemoveSession(s.SessionID())

	if wdErr != nil {
		return wdErr
	}
	if envErr != nil {
		return envErr
	}

	if reusable {
		s.WebDriverHub.AddReusableSession(s)
	}

	return nil
}

func (s *WebDriverSession) commandPathTokens(path string) []string {
	commandPath := strings.Trim(strings.TrimPrefix(path, s.sessionPath), "/")
	if commandPath == "" {
		return []string{}
	}
	return strings.Split(commandPath, "/")
}

// Unpause makes the session usable again and associates it with the given session id.
func (s *WebDriverSession) Unpause(id int) {
	s.mu.Lock()
	s.stopped = false
	s.ID = id
	s.mu.Unlock()
}

func (s *WebDriverSession) defaultHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	pathTokens := s.commandPathTokens(r.URL.Path)

	s.mu.Lock()
	stopped := s.stopped
	s.mu.Unlock()

	if stopped {
		invalidSessionID(w, vars["sessionID"])
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		unknownError(w, err)
		return
	}

	req := Request{
		Method: r.Method,
		Path:   pathTokens,
		Header: r.Header,
		Body:   body,
	}
	resp, err := s.handler(ctx, req)
	if err != nil {
		if ctx.Err() == context.Canceled {
			log.Printf("[%s] request %+v was canceled.", s.Name(), req)
			return
		}
		if ctx.Err() == context.DeadlineExceeded {
			s.Warning(errors.New(s.Name(), fmt.Errorf("request %+v exceeded deadline", req)))
			timeout(w, r.URL.Path)
			return
		}
		s.Severe(errors.New(s.Name(), err))
		unknownError(w, err)
		return
	}

	if len(resp.Body) != 0 {
		w.Header().Set("Content-Type", contentType)
	}
	if resp.Header != nil {
		// Copy response headers from resp to w
		for k, vs := range resp.Header {
			w.Header().Del(k)
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}
	}

	// TODO(fisherii): needed to play nice with Dart Sync WebDriver. Delete when Dart Sync WebDriver is deleted.
	w.Header().Set("Transfer-Encoding", "identity")
	w.Header().Set("Content-Length", strconv.Itoa(len(resp.Body)))

	httphelper.SetDefaultResponseHeaders(w.Header())

	// Copy status code from resp to w
	w.WriteHeader(resp.Status)

	// Write body from resp to w
	w.Write(resp.Body)
}

func createBaseHandler(driver webdriver.WebDriver) HandlerFunc {
	client := &http.Client{}

	return func(ctx context.Context, rq Request) (Response, error) {
		url, err := driver.CommandURL(rq.Path...)
		if err != nil {
			return Response{}, err
		}

		req, err := http.NewRequest(rq.Method, url.String(), bytes.NewReader(rq.Body))
		if err != nil {
			return Response{}, err
		}
		req = req.WithContext(ctx)
		for k, v := range rq.Header {
			if !strings.HasPrefix(k, "x-google-") {
				req.Header[k] = v
			}
		}

		resp, err := client.Do(req)
		if err != nil {
			return Response{}, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Response{}, err
		}
		return Response{resp.StatusCode, resp.Header, body}, nil
	}
}

// ResponseFromError generates a Response object for err.
func ResponseFromError(err error) (Response, error) {
	body, e := webdriver.MarshalError(err)
	return Response{
		Status: webdriver.ErrorHTTPStatus(err),
		Body:   body,
	}, e
}

// SuccessfulResponse generate a response object indicating success.
func SuccessfulResponse(value interface{}) (Response, error) {
	body := map[string]interface{}{
		"status": 0,
	}

	if value != nil {
		body["value"] = value
	}

	bytes, err := json.Marshal(body)
	return Response{
		Status: http.StatusOK,
		Body:   bytes,
	}, err
}
