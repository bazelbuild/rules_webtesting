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
	"context"
	"errors"
	"net/http"

	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/webdriver"
)

const sizeScript = `return {"width": screen.width, "height": screen.height};`

func createChromeEmulatedDeviceHandler(_ webdriver.WebDriver, desired map[string]interface{}, base HandlerFunc) (HandlerFunc, error) {
	chromeOptions, ok := desired["chromeOptions"].(map[string]interface{})
	if !ok {
		return base, nil
	}

	if _, ok := chromeOptions["mobileEmulation"]; !ok {
		return base, nil
	}

	return func(ctx context.Context, rq Request) (Response, error) {
		if rq.method != http.MethodGet || len(rq.path) != 1 || rq.path[0] != "screenshot" {
			return base(ctx, rq)
		}
		return Response{}, errors.New("screenshot unsupported")
	}, nil
}
