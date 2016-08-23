package driverhub

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux/mux"
	"github.com/web_test_launcher/launcher/proxy/webdriver"
)

type handlerFunc func(context.Context, request) (response, error)

type session struct {
	id          int
	hub         *hub
	driver      webdriver.WebDriver
	handler     handlerFunc
	sessionPath string
	router      *mux.Router

	mu      sync.Mutex
	stopped bool
}

type request struct {
	method string
	path   []string
	header http.Header
	body   []byte
}

type response struct {
	status int
	header http.Header
	body   []byte
}

func createSession(id int, hub *hub, driver webdriver.WebDriver, desired map[string]interface{}) (http.Handler, error) {
	// create base handler function
	handler := createBaseHandler(driver)

	sessionPath := fmt.Sprintf("/wd/hub/session/%s", driver.SessionID())
	sess := &session{id: id, hub: hub, driver: driver, handler: handler, sessionPath: sessionPath}

	r := mux.NewRouter()
	r.Path(sessionPath).Methods("DELETE").HandlerFunc(sess.quit)
	// Route for commands for this session.
	r.PathPrefix(sessionPath).HandlerFunc(sess.defaultHandler)
	// Route for commands for some other session. If this happens, the hub has
	// done something wrong.
	r.PathPrefix("/wd/hub/session/{sessionID}").HandlerFunc(sess.wrongSession)
	// Route for all other paths that aren't WebDriver commands. This also implies
	// that the hub has done something wrong.
	r.PathPrefix("/").HandlerFunc(sess.unknownCommand)
	sess.router = r

	return sess, nil
}

func (s *session) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *session) wrongSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	unknownError(w, fmt.Errorf("request for session %q was routed to handler for %q", vars["sessionID"], s.driver.SessionID()))
}

func (s *session) unknownCommand(w http.ResponseWriter, r *http.Request) {
	unknownCommand(w, r)
}

func (s *session) quit(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ctx := r.Context()
	vars := mux.Vars(r)

	if s.stopped {
		invalidSessionID(w, vars["sessionID"])
		return
	}

	s.stopped = true

	wdErr := s.driver.Quit(ctx)
	if wdErr != nil {
		log.Printf("Error quitting wendrover: %v", wdErr)
	}

	envErr := s.hub.env.StopSession(ctx, s.id)
	if envErr != nil {
		log.Printf("Error stopping session: %v", envErr)
	}

	s.hub.mu.Lock()
	delete(s.hub.sessions, s.driver.SessionID())
	s.hub.mu.Unlock()

	if wdErr != nil {
		unknownError(w, wdErr)
		return
	}
	if envErr != nil {
		unknownError(w, envErr)
		return
	}

	success(w, nil)
	return
}

func (s *session) commandPathTokens(path string) []string {
	commandPath := strings.TrimPrefix(path, s.sessionPath)
	return strings.Split(strings.Trim(commandPath, "/"), "/")
}

func (s *session) defaultHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	pathTokens := s.commandPathTokens(r.URL.Path)

	if s.stopped {
		invalidSessionID(w, vars["sessionID"])
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		unknownError(w, err)
		return
	}

	req := request{
		method: r.Method,
		path:   pathTokens,
		header: r.Header,
		body:   body,
	}
	resp, err := s.handler(ctx, req)
	if err != nil {
		unknownError(w, err)
		return
	}

	if resp.header != nil {
		// Copy response headers from resp to w
		for k, vs := range resp.header {
			w.Header().Del(k)
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}
	}

	// TODO(fisherii): needed to play nice with Dart Sync WebDriver. Delete when Dart Sync WebDriver is deleted.
	w.Header().Set("Transfer-Encoding", "identity")
	w.Header().Set("Content-Length", strconv.Itoa(len(resp.body)))

	// Copy status code from resp to w
	w.WriteHeader(resp.status)

	// Write body from resp to w
	w.Write(resp.body)
}

func createBaseHandler(driver webdriver.WebDriver) handlerFunc {
	client := &http.Client{}

	return func(ctx context.Context, rq request) (response, error) {
		url, err := driver.CommandURL(rq.path...)
		if err != nil {
			return response{}, err
		}

		req, err := http.NewRequest(rq.method, url.String(), bytes.NewReader(rq.body))
		if err != nil {
			return response{}, err
		}
		req = req.WithContext(ctx)
		if rq.header != nil {
			req.Header = rq.header
		}

		resp, err := client.Do(req)
		if err != nil {
			return response{}, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return response{}, err
		}
		return response{resp.StatusCode, resp.Header, body}, nil
	}
}
