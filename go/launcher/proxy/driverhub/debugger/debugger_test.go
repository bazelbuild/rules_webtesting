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
  "net/http"
  "net/url"
  "testing"
)

func TestBreakpointMatches(t *testing.T) {
  testCases := []struct {
    name       string
    breakpoint *breakpoint
    request    *http.Request
    matches    bool
  }{
    {
      "empty breakpoint matches everything",
      &breakpoint{},
      &http.Request{},
      true,
    },
    {
      "matching method",
      &breakpoint{
        Methods: []string{"POST", "GET"},
      },
      &http.Request{
        Method: "POST",
      },
      true,
    },
    {
      "non-matching method",
      &breakpoint{
        Methods: []string{"POST", "GET"},
      },
      &http.Request{
        Method: "DELETE",
      },
      false,
    },
    {
      "matching path",
      &breakpoint{
        Path: "url",
      },
      &http.Request{
        URL: &url.URL{
          Path: "/wd/hub/session/abc/url",
        },
      },
      true,
    },
    {
      "non-matching path",
      &breakpoint{
        Path: "url",
      },
      &http.Request{
        URL: &url.URL{
          Path: "/wd/hub/session/abc/elements",
        },
      },
      false,
    },
    {
      "non-matching method and matching path",
      &breakpoint{
        Methods: []string{"POST", "GET"},
        Path:    "url",
      },
      &http.Request{
        Method: "DELETE",
        URL: &url.URL{
          Path: "/wd/hub/session/abc/url",
        },
      },
      false,
    },
    {
      "matching method and non-matching path",
      &breakpoint{
        Methods: []string{"POST", "GET"},
        Path:    "url",
      },
      &http.Request{
        Method: "POST",
        URL: &url.URL{
          Path: "/wd/hub/session/abc/elements",
        },
      },
      false,
    },
    {
      "non-matching method and matching path",
      &breakpoint{
        Methods: []string{"POST", "GET"},
        Path:    "url",
      },
      &http.Request{
        Method: "DELETE",
        URL: &url.URL{
          Path: "/wd/hub/session/abc/url",
        },
      },
      false,
    },
    {
      "non-matching method and path",
      &breakpoint{
        Methods: []string{"POST", "GET"},
        Path:    "url",
      },
      &http.Request{
        Method: "DELETE",
        URL: &url.URL{
          Path: "/wd/hub/session/abc/elements",
        },
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
