// Copyright 2017 Google Inc.
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

// Package debugger enables WTL Debugger.
package debugger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"sync"

	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
)

type breakpoint struct {
	ID      int      `json:"id"`
	Path    string   `json:"path,omitempty"`
	Methods []string `json:"methods,omitempty"`
	Body    string   `josn:"body,omitempty"`

	pathRegex *regexp.Regexp
	bodyRegex *regexp.Regexp
}

type command struct {
	ID         int         `json:"id"`
	Command    string      `json:"command"`
	Breakpoint *breakpoint `json:"breakpoint,omitempty"`
}

type request struct {
	Method string `json:"method,omitempty"`
	Path   string `json:"path,omitempty"`
	Body   string `json:"body,omitempty"`
}

type response struct {
	ID      int      `json:"id"`
	Status  string   `json:"status"`
	Request *request `json:"request,omitempty"`
}

// Debugger is an implementation of the WTL Debugger server.
type Debugger struct {
	conn net.Conn

	mu          sync.RWMutex
	connError   error
	healthy     bool
	step        bool
	breakpoints map[int]*breakpoint
	waiting     chan<- interface{}
}

// New returns a Debugger waiting for a connection on TCP port.
func New(port int) *Debugger {
	d := &Debugger{
		breakpoints: map[int]*breakpoint{},
	}

	go d.waitForConnection(port)
	return d
}

// Name is the name of this component used in errors and logging.
func (*Debugger) Name() string {
	return "WTL Debugger Server"
}

// Healthy returns nil iff a frontend is connected and has sent a step or continue command.
func (d *Debugger) Healthy(context.Context) error {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.connError != nil {
		return d.connError
	}

	if !d.healthy {
		return errors.New(d.Name(), "debugger frontend is not connected.")
	}
	return nil
}

// Request logs r to the debugger frontend. If r matches a breakpoint or the debugger is in step mode,
// Request will not return until a continue message from the front-end is received.
func (d *Debugger) Request(r *http.Request) {
	// Capture request body
	body, err := capture(r.Body)
	if err != nil {
		log.Fatalf("Error reading request body: %v", err)
	}
	r.Body = body

	resp := &response{
		Request: &request{
			Method: r.Method,
			Path:   r.URL.Path,
			Body:   body.captured,
		},
	}

	// Identify if we should be continuing or waiting
	d.mu.RLock()
	step := d.step

	if !step {
		for _, bp := range d.breakpoints {
			if bp.matches(resp.Request) {
				step = true
				break
			}
		}
	}

	d.mu.RUnlock()

	if step {
		resp.Status = "waiting"
	} else {
		resp.Status = "running"
	}

	// Send request info to client client
	bytes, err := json.Marshal(resp)
	if err != nil {
		log.Print(err)
		return
	}

	if _, err := d.conn.Write(bytes); err != nil {
		log.Print(err)
	}

	if _, err := d.conn.Write([]byte("\n")); err != nil {
		log.Print(err)
	}

	// Not stepping, so return.
	if !step {
		return
	}

	// Wait for step/continue command from front end.

	// TODO(DrMarcII): Race condition here, but it is racing a human, so not too worried about it.
	waiting := make(chan interface{})
	d.mu.Lock()
	d.waiting = waiting
	d.mu.Unlock()

	<-waiting
	d.mu.Lock()
	d.waiting = nil
	d.mu.Unlock()
}

func (d *Debugger) waitForConnection(port int) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		d.mu.Lock()
		defer d.mu.Unlock()
		d.connError = errors.NewPermanent(d.Name(), err)
		return
	}

	fmt.Printf(`
***********************************************

Waiting for debugger connection on port %d.

***********************************************
`, port)

	conn, err := l.Accept()
	d.mu.Lock()
	defer d.mu.Unlock()
	if err != nil {
		d.connError = errors.NewPermanent(d.Name(), err)
		return
	}
	d.conn = conn

	go d.readLoop()
}

func (d *Debugger) readLoop() {
	decoder := json.NewDecoder(d.conn)

	for {
		cmd := &command{}
		if err := decoder.Decode(cmd); err != nil {
			// do something here...
			log.Fatalf("Error reading from debugger: %v", err)
		}

		d.processCommand(cmd)
	}
}

func (d *Debugger) processCommand(cmd *command) {
	d.mu.Lock()
	defer d.mu.Unlock()

	response := &response{ID: cmd.ID, Status: "error"}

	switch cmd.Command {
	case "continue":
		d.healthy = true
		d.step = false
		if d.waiting != nil {
			close(d.waiting)
		}
		response.Status = "running"

	case "step":
		d.healthy = true
		d.step = true
		if d.waiting != nil {
			close(d.waiting)
		}
		response.Status = "running"

	case "stop":
		os.Exit(-1)
	case "set breakpoint":
		if cmd.Breakpoint == nil {
			break
		}
		bp := cmd.Breakpoint
		if err := bp.initialize(); err != nil {
			break
		}
		d.breakpoints[bp.ID] = bp
		response.Status = "waiting"

	case "delete breakpoint":
		if cmd.Breakpoint == nil {
			break
		}
		delete(d.breakpoints, cmd.Breakpoint.ID)
		response.Status = "waiting"
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		log.Print(err)
		return
	}

	if _, err := d.conn.Write(bytes); err != nil {
		log.Print(err)
	}
}

func (bp *breakpoint) initialize() error {
	if bp.Path != "" {
		r, err := regexp.Compile(bp.Path)
		if err != nil {
			return err
		}
		bp.pathRegex = r
	}

	if bp.Body != "" {
		r, err := regexp.Compile(bp.Body)
		if err != nil {
			return err
		}
		bp.bodyRegex = r
	}
	return nil
}

func (bp *breakpoint) matches(r *request) bool {
	if bp.pathRegex != nil && bp.pathRegex.FindString(r.Path) == "" {
		return false
	}

	if bp.bodyRegex != nil && bp.bodyRegex.FindString(r.Body) == "" {
		return false
	}

	if len(bp.Methods) != 0 {
		found := false
		for _, method := range bp.Methods {
			if r.Method == method {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

type capturedReader struct {
	io.Reader
	io.Closer
	captured string
}

func capture(r io.ReadCloser) (*capturedReader, error) {
	c, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &capturedReader{
		bytes.NewReader(c),
		r,
		string(c),
	}, nil
}
