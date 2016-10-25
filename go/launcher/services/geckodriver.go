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
// Package chromedriver provides a service.Server for managing an instance of GeckoDriver.
package geckodriver

import (
	"fmt"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/services/service"
	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
)

type GeckoDriver struct {
	*service.Server
}

// New creates a new service.Server instance that manages GeckoDriver.
func New(m *metadata.Metadata, xvfb bool) (*GeckoDriver, error) {
	driverPath, err := m.GetFilePath("GECKODRIVER")
	if err != nil {
		return nil, errors.New("GeckoDriver", err)
	}

	firefoxPath, err := m.GetFilePath("FIREFOX")
	if err != nil {
		return nil, errors.New("GeckoDriver", err)
	}

	server, err := service.NewServer(
		"GeckoDriver",
		driverPath,
		"",
		xvfb,
		60*time.Second,
		nil,
		"--binary", firefoxPath,
		"--port", "{port}")
	if err != nil {
		return nil, err
	}
	return &GeckoDriver{server}, nil
}

func (c *GeckoDriver) Address() string {
	return fmt.Sprintf("http://%s/", c.Server.Address())
}
