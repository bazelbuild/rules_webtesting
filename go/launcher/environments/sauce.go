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

// Package sauce provides a simple environment for accessing a SauceLabs browser.
// It depends on environment variables SAUCE_USERNAME and SAUCE_ACCESS_KEY being defined.
// TODO(DrMarcII): Add SauceConnect support.
package sauce

import (
	"context"
	"os"

	"github.com/bazelbuild/rules_webtesting/go/launcher/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/launcher/environments/environment"
	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
)

const (
	name = "Sauce WebDriver Environment"
)

type sauce struct {
	*environment.Base
	address string
}

// NewEnv creates a new environment that uses an externally started Selenium Server.
func NewEnv(m *metadata.Metadata, d diagnostics.Diagnostics) (environment.Env, error) {
	address := os.ExpandEnv("http://${SAUCE_USERNAME}:${SAUCE_ACCESS_KEY}@ondemand.saucelabs.com/wd/hub/")

	base, err := environment.NewBase(name, m, d)
	if err != nil {
		return nil, err
	}

	return &sauce{
		Base:    base,
		address: address,
	}, nil
}

// WDAddress returns the user-provided selenium address.
func (s *sauce) WDAddress(context.Context) string {
	return s.address
}
