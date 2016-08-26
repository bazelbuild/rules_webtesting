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
// Package selenium provides a service for launching GoogleSeleniumServer on the local
// host.
package selenium

import (
	"time"

	"github.com/bazelbuild/rules_web/launcher/services/service"
)

// Selenium is a service that starts GoogleSeleniumServer.
type Selenium struct {
	*service.Server
}

// New creates a new service for starting GoogleSeleniumServer on the host machine.
func New(xvfb bool, env map[string]string) (*Selenium, error) {
	server, err := service.NewServer(
		"SeleniumServer",
		"io_bazel_rules_web/java/SeleniumServer",
		"http://%s/wd/hub/status",
		xvfb,
		60*time.Second,
		env, "-port", "{port}")
	if err != nil {
		return nil, err
	}
	return &Selenium{server}, nil
}
