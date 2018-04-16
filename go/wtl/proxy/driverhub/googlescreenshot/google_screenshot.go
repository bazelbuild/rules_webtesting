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

// Package googlescreenshot includes a handler for an advanced screenshot endpoint at POST google/screenshot.
package googlescreenshot

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"

	"github.com/bazelbuild/rules_webtesting/go/cropper"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/webdriver"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy/driverhub"
)

type request struct {
	Element map[string]interface{}
	Exclude []map[string]interface{}
}

// ProviderFunc provides a handler for an advanced screenshot endpoint at POST google/screenshot.
func ProviderFunc(session *driverhub.WebDriverSession, _ *capabilities.Capabilities, base driverhub.HandlerFunc) (driverhub.HandlerFunc, bool) {
	return func(ctx context.Context, rq driverhub.Request) (driverhub.Response, error) {
		if rq.Method != http.MethodPost || len(rq.Path) != 2 || rq.Path[0] != "google" || rq.Path[1] != "screenshot" {
			return base(ctx, rq)
		}

		r := request{}

		if err := json.Unmarshal(rq.Body, &r); err != nil {
			// TODO(DrMarcII): better error message?
			return driverhub.ResponseFromError(err)
		}

		img, err := captureScreenshot(ctx, session.WebDriver, r)
		if err != nil {
			// TODO(DrMarcII): better error message?
			return driverhub.ResponseFromError(err)
		}

		return createResponse(img)
	}, true
}

func captureScreenshot(ctx context.Context, driver webdriver.WebDriver, r request) (image.Image, error) {
	el, _ := driver.ElementFromMap(r.Element)

	if el != nil {
		if err := el.ScrollIntoView(ctx); err != nil {
			return nil, err
		}
	}

	// TODO(DrMarcII): stabilization to ensure that we have finished scrolling?
	img, err := driver.Screenshot(ctx)
	if err != nil {
		return nil, err
	}

	for _, e := range r.Exclude {
		exclude, err := driver.ElementFromMap(e)
		if err != nil {
			log.Print(err)
			continue
		}

		bounds, err := exclude.Bounds(ctx)
		if err != nil {
			log.Print(err)
			continue
		}

		i, err := cropper.Blackout(img, bounds)
		if err != nil {
			return nil, err
		}
		img = i
	}

	if el != nil {
		bounds, err := el.Bounds(ctx)
		if err != nil {
			return nil, err
		}

		return cropper.Crop(img, bounds)
	}
	return img, nil
}

func createResponse(img image.Image) (driverhub.Response, error) {
	buffer := &bytes.Buffer{}
	b64 := base64.NewEncoder(base64.StdEncoding, buffer)
	defer b64.Close()

	if err := png.Encode(b64, img); err != nil {
		return driverhub.ResponseFromError(err)
	}

	return driverhub.Response{
		Status: http.StatusOK,
		Body:   []byte(fmt.Sprintf(`{"status": 0, "value": %q}`, buffer.String())),
	}, nil
}
