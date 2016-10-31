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
//
////////////////////////////////////////////////////////////////////////////////
// Package webtest provides WebDriver provisioning and information APIs.
//
// Provision a browser:
//   import "github.com/bazelbuild/rules_webtesting/go/webtest"
//
//   driver, err := webtest.NewWebDriverSession(nil)
//
// Provision a browser with capabilities (as an example profiling):
//   capabilities := map[string]interface{}{
//     "webdriver.logging.profiler.enabled": true,
//   }
//   driver, err := webtest.NewWebDriverSession(capabilities)
//
// Get basic information about the browser defined by the web test environment:
//   info, err := webtest.GetBrowserInfo()
package webtest

import (
	"errors"
	"os"
	"strings"

	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
	"github.com/bazelbuild/rules_webtesting/go/util/bazel"
	"github.com/tebeka/selenium/selenium"
)

var info *BrowserInfo

// BrowserInfo represents basic information about the browser defined by the web test environment.
type BrowserInfo struct {
	// The Bazel label for the browser.
	BrowserLabel string
	// The Environment that Web Test Launcher is using for the specified configuration.
	Environment string
}

// GetBrowserInfo returns basic information about the browser defined by the web test environment.
func GetBrowserInfo() (*BrowserInfo, error) {
	if info == nil {
		i, err := newInfo(os.Getenv("WEB_TEST_BROWSER_METADATA"))
		if err != nil {
			return nil, err
		}
		info = i
	}
	return info, nil
}

func newInfo(mf string) (*BrowserInfo, error) {
	f, err := bazel.Runfile(mf)
	if err != nil {
		return nil, err
	}
	m, err := metadata.FromFile(f)
	if err != nil {
		return nil, err
	}
	return &BrowserInfo{
		BrowserLabel: m.BrowserLabel,
		Environment:  m.Environment,
	}, nil
}

// NewWebDriverSession provisions and returns a new WebDriver session.
func NewWebDriverSession(capabilities selenium.Capabilities) (selenium.WebDriver, error) {
	address, ok := os.LookupEnv("REMOTE_WEBDRIVER_SERVER")
	if !ok {
		return nil, errors.New(`environment variable "REMOTE_WEBDRIVER_SERVER" not set.`)
	}

	return selenium.NewRemote(capabilities, strings.TrimSuffix(address, "/"))
}
