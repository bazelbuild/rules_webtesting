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
//
// Package browser provides WebDriver provisioning and information APIs.
//
// Provision a browser:
//   driver, err := browser.NewSession(nil)
//
// Provision a browser with capabilities (as an example profiling):
//   capabilities := map[string]interface{}{
//     "webdriver.logging.profiler.enabled": true,
//   }
//   driver, err := browser.NewSession(capabilities)
//
// Get basic information about the browser defined by the web test environment:
//   info, err := browser.GetInfo()
package browser

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium/selenium"
)

const (
	metadataEnvVar        = "WEB_TEST_BROWSER_METADATA"
	webdriverServerEnvVar = "REMOTE_WEBDRIVER_SERVER"
)

// NewSession provisions and returns a new WebDriver session.
func NewSession(capabilities selenium.Capabilities) (selenium.WebDriver, error) {
	hostport, ok := os.LookupEnv(webdriverServerEnvVar)
	if !ok {
		return nil, fmt.Errorf("environment variable %q is not defined, are you using web_test", webdriverServerEnvVar)
	}

	address := fmt.Sprintf("http://%s/wd/hub", hostport)

	return selenium.NewRemote(capabilities, address)
}
