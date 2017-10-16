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

// Package quithandler checks if a window close command is closing the last window and treats it
// as a quit if it is.
package quithandler

import (
	"context"
	"net/http"

	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/driverhub"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
)

// ProviderFunc provides a handler for quit and close commands that end the session within the environment when the browser exits.
func ProviderFunc(session *driverhub.WebDriverSession, caps capabilities.Spec, base driverhub.HandlerFunc) (driverhub.HandlerFunc, bool) {
	return func(ctx context.Context, rq driverhub.Request) (driverhub.Response, error) {
		// If quit command, then quit.
		if rq.Method == http.MethodDelete && len(rq.Path) == 0 {
			return session.Quit(ctx, rq)
		}

		// If not window close command, then forward as normal
		if rq.Method != http.MethodDelete || len(rq.Path) != 1 || rq.Path[0] != "window" {
			return base(ctx, rq)
		}

		// If closing last window, then quit.
		if windows, err := session.WindowHandles(ctx); err != nil {
			return driverhub.ResponseFromError(err)
		} else if len(windows) == 1 {
			return session.Quit(ctx, rq)
		}

		// Otherwise forward the close window
		return base(ctx, rq)
	}, true
}
