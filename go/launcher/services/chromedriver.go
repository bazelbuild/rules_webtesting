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
// Package chromedriver provides a service.Server for managing an instance of chromedriver.
package chromedriver

import (
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/services/service"
	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
)

// New creates a new service.Server instance that manages chromedriver.
func New(m *metadata.Metadata, xvfb bool) (*service.Server, error) {
	chromedriverPath, err := m.GetFilePath("CHROMEDRIVER")
	if err != nil {
		return nil, errors.New("ChromeDriver", err)
	}

	return service.NewServer(
		"ChromeDriver",
		chromedriverPath,
		"http://%s/wd/hub/status",
		xvfb,
		60*time.Second,
		nil,
		"--port={port}",
		"--url-base=wd/hub")
}
