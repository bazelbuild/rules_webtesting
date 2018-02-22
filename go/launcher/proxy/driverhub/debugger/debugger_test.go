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

package debugger

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestBreakpointMatches(t *testing.T) {
	testCases := []struct {
		name       string
		breakpoint *breakpoint
		request    *request
		matches    bool
	}{
		{
			"empty breakpoint matches everything",
			&breakpoint{},
			&request{},
			true,
		},
		{
			"matching method",
			&breakpoint{
				Methods: []string{"POST", "GET"},
			},
			&request{
				Method: "POST",
			},
			true,
		},
		{
			"non-matching method",
			&breakpoint{
				Methods: []string{"POST", "GET"},
			},
			&request{
				Method: "DELETE",
			},
			false,
		},
		{
			"matching path",
			&breakpoint{
				Path: "url",
			},
			&request{
				Path: "/wd/hub/session/abc/url",
			},
			true,
		},
		{
			"non-matching path",
			&breakpoint{
				Path: "url",
			},
			&request{
				Path: "/wd/hub/session/abc/elements",
			},
			false,
		},
		{
			"non-matching method and matching path",
			&breakpoint{
				Methods: []string{"POST", "GET"},
				Path:    "url",
			},
			&request{
				Method: "DELETE",
				Path:   "/wd/hub/session/abc/url",
			},
			false,
		},
		{
			"matching method and non-matching path",
			&breakpoint{
				Methods: []string{"POST", "GET"},
				Path:    "url",
			},
			&request{
				Method: "POST",
				Path:   "/wd/hub/session/abc/elements",
			},
			false,
		},
		{
			"non-matching method and matching path",
			&breakpoint{
				Methods: []string{"POST", "GET"},
				Path:    "url",
			},
			&request{
				Method: "DELETE",
				Path:   "/wd/hub/session/abc/url",
			},
			false,
		},
		{
			"non-matching method and path",
			&breakpoint{
				Methods: []string{"POST", "GET"},
				Path:    "url",
			},
			&request{
				Method: "DELETE",
				Path:   "/wd/hub/session/abc/elements",
			},
			false,
		},
		{
			"matching body",
			&breakpoint{
				Body: `http://www\.google\.com`,
			},
			&request{
				Body: `{"url": "http://www.google.com"}`,
			},
			true,
		},
		{
			"non-matching body",
			&breakpoint{
				Body: `http://www\.google\.com`,
			},
			&request{
				Body: `{"url": "http://wwwagoogle.com"}`,
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.breakpoint.initialize(); err != nil {
				t.Fatal(err)
			}
			if m := tc.breakpoint.matches(tc.request); m != tc.matches {
				t.Fatalf("got %+v.matches(%+v) == %v, expected %v", tc.breakpoint, tc.request, m, tc.matches)
			}
		})
	}
}

type fakeReadCloser struct {
	io.Reader
	closed bool
}

func (f *fakeReadCloser) Close() error {
	if f.closed {
		return errors.New("Close called multiple times")
	}
	f.closed = true
	return nil
}

func TestCapture(t *testing.T) {
	f := &fakeReadCloser{
		Reader: strings.NewReader("a test string"),
	}

	c, err := capture(f)
	if err != nil {
		t.Fatal(err)
	}

	if c.captured != "a test string" {
		t.Errorf(`Got c.captured == %q, expected "a test string"`, c.captured)
	}

	b, err := ioutil.ReadAll(c)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "a test string" {
		t.Errorf(`Got ioutil.ReadAll(c) == %q, expected "a test string"`, string(b))
	}

	if err := c.Close(); err != nil {
		t.Fatal(err)
	}

	if !f.closed {
		t.Errorf("Got f.closed == false, expected true")
	}
}
