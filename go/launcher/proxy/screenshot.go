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

// Package screenshot includes a handler for the WebDriver screenshot operation.
package screenshot

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"net/http"

	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/driverhub"
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/webdriver"
	"github.com/bazelbuild/rules_webtesting/go/util/cropper"
)

const sizeScript = `return {"width": screen.width, "height": screen.height};`

// ProviderFunc provides a handler for /screenshot command that crops the image if Chrome mobile emulation is enabled.
func ProviderFunc(session *driverhub.WebDriverSession, desired map[string]interface{}, base driverhub.HandlerFunc) (driverhub.HandlerFunc, bool) {
	chromeOptions, ok := desired["chromeOptions"].(map[string]interface{})
	if !ok {
		return base, false
	}

	if _, ok := chromeOptions["mobileEmulation"]; !ok {
		return base, false
	}

	return func(ctx context.Context, rq driverhub.Request) (driverhub.Response, error) {
		if rq.Method != http.MethodGet || len(rq.Path) != 1 || rq.Path[0] != "screenshot" {
			return base(ctx, rq)
		}
		img, err := session.Screenshot(ctx)
		if err != nil {
			body, _ := webdriver.MarshalError(err)
			return driverhub.Response{
				Status: webdriver.ErrorHTTPStatus(err),
				Body:   body,
			}, nil
		}

		val := struct {
			Width  int
			Height int
		}{}

		if err := session.ExecuteScript(ctx, sizeScript, nil, &val); err != nil {
			body, _ := webdriver.MarshalError(err)
			return driverhub.Response{
				Status: webdriver.ErrorHTTPStatus(err),
				Body:   body,
			}, nil
		}

		cropped, err := cropper.Crop(img, image.Rect(0, 0, val.Width, val.Height))
		if err != nil {
			body, _ := webdriver.MarshalError(err)
			return driverhub.Response{
				Status: webdriver.ErrorHTTPStatus(err),
				Body:   body,
			}, nil
		}

		buffer := &bytes.Buffer{}
		b64 := base64.NewEncoder(base64.StdEncoding, buffer)

		if err := png.Encode(b64, cropped); err != nil {
			body, _ := webdriver.MarshalError(err)
			return driverhub.Response{
				Status: webdriver.ErrorHTTPStatus(err),
				Body:   body,
			}, nil
		}

		b64.Close()

		return driverhub.Response{
			Status: http.StatusOK,
			Body:   []byte(fmt.Sprintf(`{"status": 0, "value": %q}`, buffer.String())),
		}, nil
	}, true
}
