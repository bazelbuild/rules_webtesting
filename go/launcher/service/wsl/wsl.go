// Copyright 2018 Google Inc.
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

// Package wsl provides a Service for launching WebDriver Server Light (WSL).
package wsl

import (
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/service"
	"github.com/bazelbuild/rules_webtesting/go/metadata"
)

const (
	compName     = "WSL Service"
	wslNamedFile = "WEBDRIVER_SERVER_LIGHT"
)

// New creates a new service.Server instance that manages chromedriver.
func New(d diagnostics.Diagnostics, m *metadata.Metadata) (*service.Server, error) {
	wslPath, err := m.GetFilePath(wslNamedFile)
	if err != nil {
		return nil, errors.New(compName, err)
	}

	server, err := service.NewServer(
		compName,
		d,
		wslPath,
		"http://%s/healthz",
		false,
		1*time.Second,
		nil,
		"--port={port}")
	if err != nil {
		return nil, err
	}
	return server, nil
}
