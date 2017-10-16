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

// Package mobileemulation includes a handler that modifies some WebDriver
// commands when Chrome Mobile Emulation is enabled.
package mobileemulation

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"

	"github.com/bazelbuild/rules_webtesting/go/cropper"
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/driverhub"
	"github.com/bazelbuild/rules_webtesting/go/launcher/webdriver"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
)

const sizeScript = `return {"width": screen.width, "height": screen.height};`

// ProviderFunc provides a handler for /screenshot command that crops the image if Chrome mobile emulation is enabled.
func ProviderFunc(session *driverhub.WebDriverSession, caps capabilities.Spec, base driverhub.HandlerFunc) (driverhub.HandlerFunc, bool) {
	if !Enabled(caps) {
		return base, false
	}

	return func(ctx context.Context, rq driverhub.Request) (driverhub.Response, error) {
		switch {
		case rq.Method == http.MethodGet && len(rq.Path) == 1 && rq.Path[0] == "screenshot":
			return screenshot(ctx, session.WebDriver)

		case rq.Method == http.MethodPost && len(rq.Path) == 2 && rq.Path[0] == "window" && rq.Path[1] == "rect":
			return noOp("set rect")
		case rq.Method == http.MethodPost && len(rq.Path) == 2 && rq.Path[0] == "window" && rq.Path[1] == "maximize":
			return noOp("maximize")
		case rq.Method == http.MethodPost && len(rq.Path) == 2 && rq.Path[0] == "window" && rq.Path[1] == "minimize":
			return noOp("minimize")
		case rq.Method == http.MethodPost && len(rq.Path) == 2 && rq.Path[0] == "window" && rq.Path[1] == "fullscreen":
			return noOp("fullscreen")
		case rq.Method == http.MethodPost && len(rq.Path) == 3 && rq.Path[0] == "window" && rq.Path[2] == "size":
			return noOp("set size")
		case rq.Method == http.MethodPost && len(rq.Path) == 3 && rq.Path[0] == "window" && rq.Path[2] == "position":
			return noOp("set position")
		case rq.Method == http.MethodPost && len(rq.Path) == 3 && rq.Path[0] == "window" && rq.Path[2] == "maximize":
			return noOp("maximize")

		default:
			return base(ctx, rq)
		}
	}, true
}

func screenshot(ctx context.Context, driver webdriver.WebDriver) (driverhub.Response, error) {
	img, err := GetMobileScreenshot(ctx, driver)
	if err != nil {
		return driverhub.ResponseFromError(err)
	}

	return CreateResponse(img)
}

func noOp(c string) (driverhub.Response, error) {
	log.Printf("Window %s is unsupported when mobile emulation is enabled", c)
	return driverhub.Response{
		Status: 200,
		Body:   []byte(fmt.Sprintf(`{"status": 0, "value": "window %s is unsupported."}`, c)),
	}, nil
}

func getCap(caps capabilities.Spec, key string) interface{} {
	if v, ok := caps.Always[key]; ok {
		return v
	}
	v, _ := caps.OSSCaps[key]
	return v
}

// MobileEmulationEnabled determines if the capabilities define a mobile emulate config.
func Enabled(caps capabilities.Spec) bool {
	browserName, _ := getCap(caps, "browserName").(string)
	if browserName != "chrome" {
		return false
	}

	chromeOptions, ok := getCap(caps, "chromeOptions").(map[string]interface{})
	if !ok {
		return false
	}

	if _, ok := chromeOptions["mobileEmulation"]; !ok {
		return false
	}

	return true
}

// GetMobileScreenshot crops a screenshot based on the emulated viewport size.
func GetMobileScreenshot(ctx context.Context, driver webdriver.WebDriver) (image.Image, error) {
	img, err := driver.Screenshot(ctx)
	if err != nil {
		return nil, err
	}

	val := struct {
		Width  int
		Height int
	}{}

	if err := driver.ExecuteScript(ctx, sizeScript, nil, &val); err != nil {
		return nil, err
	}

	return cropper.Crop(img, image.Rect(0, 0, val.Width, val.Height))
}

// CreateResponse creates a /screenshot endpoint response for a given image.
func CreateResponse(img image.Image) (driverhub.Response, error) {
	buffer := &bytes.Buffer{}
	b64 := base64.NewEncoder(base64.StdEncoding, buffer)

	if err := png.Encode(b64, img); err != nil {
		return driverhub.ResponseFromError(err)
	}

	b64.Close()

	return driverhub.Response{
		Status: http.StatusOK,
		Body:   []byte(fmt.Sprintf(`{"status": 0, "value": %q}`, buffer.String())),
	}, nil
}
