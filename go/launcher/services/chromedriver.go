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
	"fmt"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/services/service"
	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
)

type ChromeDriver struct {
	*service.Server
}

// New creates a new service.Server instance that manages chromedriver.
func New(m *metadata.Metadata, xvfb bool) (*ChromeDriver, error) {
	chromeDriverPath, err := m.GetFilePath("CHROMEDRIVER")
	if err != nil {
		return nil, errors.New("ChromeDriver", err)
	}

	server, err := service.NewServer(
		"ChromeDriver",
		chromeDriverPath,
		"http://%s/status",
		xvfb,
		60*time.Second,
		nil,
		"--port={port}")
	if err != nil {
		return nil, err
	}
	return &ChromeDriver{server}, nil
}

func (c *ChromeDriver) Address() string {
	return fmt.Sprintf("http://%s/", c.Server.Address())
}
