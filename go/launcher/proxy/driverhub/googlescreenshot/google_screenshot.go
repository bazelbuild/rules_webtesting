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
	"context"
	"encoding/json"
	"image"
	"log"
	"net/http"

	"github.com/bazelbuild/rules_webtesting/go/cropper"
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/driverhub"
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/driverhub/mobileemulation"
	"github.com/bazelbuild/rules_webtesting/go/launcher/webdriver"
)

type request struct {
	Element map[string]interface{}
	Exclude []map[string]interface{}
}

// ProviderFunc provides a handler for an advanced screenshot endpoint at POST google/screenshot.
func ProviderFunc(session *driverhub.WebDriverSession, desired map[string]interface{}, base driverhub.HandlerFunc) (driverhub.HandlerFunc, bool) {
	useMobile := mobileemulation.Enabled(desired)

	return func(ctx context.Context, rq driverhub.Request) (driverhub.Response, error) {
		if rq.Method != http.MethodPost || len(rq.Path) != 2 || rq.Path[0] != "google" || rq.Path[1] != "screenshot" {
			return base(ctx, rq)
		}

		r := request{}

		if err := json.Unmarshal(rq.Body, &r); err != nil {
			// TODO(DrMarcII): better error message?
			return driverhub.ResponseFromError(err)
		}

		img, err := captureScreenshot(ctx, session.WebDriver, r, useMobile)
		if err != nil {
			// TODO(DrMarcII): better error message?
			return driverhub.ResponseFromError(err)
		}

		return mobileemulation.CreateResponse(img)
	}, true
}

func captureScreenshot(ctx context.Context, driver webdriver.WebDriver, r request, useMobile bool) (image.Image, error) {
	el, _ := driver.ElementFromMap(r.Element)

	if el != nil {
		if err := el.ScrollIntoView(ctx); err != nil {
			return nil, err
		}
	}

	// TODO(DrMarcII): stabilization to ensure that we have finished scrolling?
	var img image.Image
	if useMobile {
		i, err := mobileemulation.GetMobileScreenshot(ctx, driver)
		if err != nil {
			return nil, err
		}
		img = i
	} else {
		i, err := driver.Screenshot(ctx)
		if err != nil {
			return nil, err
		}
		img = i
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
