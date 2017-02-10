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

// Package screenshotter provides a client-side API for accessing the google/screenshot Web Test Launcher endpoint.
package screenshotter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tebeka/selenium"
)

// Screenshotter allows configurable screenshotting.
type Screenshotter interface {
	// Of returns a copy of Screenshotter configured to crop the screenshot using the given WebElement.
	Of(selenium.WebElement) Screenshotter
	// Excluding returns a copy of Screenshotter configured to exclude the given WebElement by covering it with a black box.
	Excluding(selenium.WebElement) Screenshotter
	// TakeScreenshot returns the screenshot.
	TakeScreenshot() (Screenshot, error)
}

type screenshotter struct {
	url     string
	Element selenium.WebElement
	Exclude []selenium.WebElement
}

// Screenshot is a PNG-encoded screenshot.
type Screenshot []byte

// New creates a new Screenshotter. It looks up the address of the WebDriver remote end in
// the WEB_TEST_WEBDRIVER_SERVER environment variable, and use driver to get the session id
// that will be used for screenshotting.
func New(driver selenium.WebDriver) (Screenshotter, error) {
	address, ok := os.LookupEnv("WEB_TEST_WEBDRIVER_SERVER")
	if !ok {
		return nil, errors.New(`environment variable "WEB_TEST_WEBDRIVER_SERVER" not set`)
	}
	return &screenshotter{
		url: fmt.Sprintf("%s/session/%s/google/screenshot", strings.TrimSuffix(address, "/"), driver.SessionID()),
	}, nil
}

// Of returns a copy of Screenshotter configured to crop the screenshot using the given WebElement.
func (s *screenshotter) Of(e selenium.WebElement) Screenshotter {
	log.Printf("e: %+v", e)
	return &screenshotter{
		url:     s.url,
		Element: e,
		Exclude: s.Exclude,
	}
}

// Excluding returns a copy of Screenshotter configured to exclude the given WebElement by covering it with a black box.
func (s *screenshotter) Excluding(e selenium.WebElement) Screenshotter {
	return &screenshotter{
		url:     s.url,
		Element: s.Element,
		Exclude: append([]selenium.WebElement{e}, s.Exclude...),
	}
}

// TakeScreenshot returns the screenshot.
func (s *screenshotter) TakeScreenshot() (Screenshot, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(s.url, "application/json; charset=utf-8", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	result := struct {
		Status  int         `json:"status"`
		Value   interface{} `json:"value"`
		Error   string      `json:"error"`
		Message string      `json:"message"`
	}{}

	if err := json.Unmarshal(b, &result); err != nil {
		return nil, fmt.Errorf("error %v decoding json %q", err, b)
	}

	if result.Status != 0 || result.Error != "" {
		msg := result.Message
		if msg == "" {
			msg, _ = result.Value.(string)

		}
		if msg == "" {
			msg = fmt.Sprintf("server returned status %d error %q", result.Status, result.Error)
		}
		return nil, errors.New(msg)
	}

	v, ok := result.Value.(string)
	if !ok {
		return nil, fmt.Errorf("expected value to be string, but was %T", result.Value)
	}

	return base64.StdEncoding.DecodeString(v)
}

// AsImage converts s into an Image object.
func (s Screenshot) AsImage() (image.Image, error) {
	return png.Decode(bytes.NewReader(s))
}
