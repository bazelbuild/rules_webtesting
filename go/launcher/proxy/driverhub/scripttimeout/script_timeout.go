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

// Package scripttimeout translates calls to set script timeout into calls
// on the WebDriver object so it can record the last set script timeout.
package scripttimeout

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/driverhub"
)

// ProviderFunc provides a handler for set script timeout commands.
func ProviderFunc(session *driverhub.WebDriverSession, desired map[string]interface{}, base driverhub.HandlerFunc) (driverhub.HandlerFunc, bool) {
	return func(ctx context.Context, rq driverhub.Request) (driverhub.Response, error) {

		if rq.Method == http.MethodPost && len(rq.Path) == 1 && rq.Path[0] == "timeouts" {
			var request map[string]interface{}

			if err := json.Unmarshal(rq.Body, &request); err != nil {
				return base(ctx, rq)
			}

			if timeout, ok := request["script"].(int); ok {
				if err := session.WebDriver.SetScriptTimeout(ctx, time.Duration(timeout)*time.Millisecond); err == nil {
					delete(request, "script")
				}
			}

			if timeout, ok := request["script"].(float64); ok {
				if err := session.WebDriver.SetScriptTimeout(ctx, time.Duration(timeout)*time.Millisecond); err == nil {
					delete(request, "script")
				}
			}

			if t, ok := request["type"].(string); ok && t == "script" {
				if timeout, ok := request["ms"].(int); ok {
					if err := session.WebDriver.SetScriptTimeout(ctx, time.Duration(timeout)*time.Millisecond); err == nil {
						delete(request, "ms")
						delete(request, "type")
					}
				}

				if timeout, ok := request["ms"].(float64); ok {
					if err := session.WebDriver.SetScriptTimeout(ctx, time.Duration(timeout)*time.Millisecond); err == nil {
						delete(request, "ms")
						delete(request, "type")
					}
				}
			}

			if len(request) == 0 {
				return driverhub.Response{
					Status: http.StatusOK,
					Body:   []byte(`{"status": 0}`),
				}, nil
			}

			body, err := json.Marshal(request)
			if err == nil {
				rq.Body = body
			}
		}

		if rq.Method == http.MethodPost && len(rq.Path) == 2 && rq.Path[0] == "timeouts" && rq.Path[1] == "async_script" {
			var request map[string]interface{}

			if err := json.Unmarshal(rq.Body, &request); err != nil {
				return base(ctx, rq)
			}

			if timeout, ok := request["ms"].(int); ok {
				if err := session.WebDriver.SetScriptTimeout(ctx, time.Duration(timeout)*time.Millisecond); err == nil {
					return driverhub.Response{
						Status: http.StatusOK,
						Body:   []byte(`{"status": 0}`),
					}, nil
				}
			}

			if timeout, ok := request["ms"].(float64); ok {
				if err := session.WebDriver.SetScriptTimeout(ctx, time.Duration(timeout)*time.Millisecond); err == nil {
					return driverhub.Response{
						Status: http.StatusOK,
						Body:   []byte(`{"status": 0}`),
					}, nil
				}
			}
		}

		return base(ctx, rq)
	}, true
}
