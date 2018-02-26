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

// Package external works with an externally started WebDriver server
// located at EXTERNAL_WEBDRIVER_SERVER_ADDRESS.
package external

import (
	"context"
	"fmt"
	"os"

	"github.com/bazelbuild/rules_webtesting/go/metadata"
	"github.com/bazelbuild/rules_webtesting/go/wtl/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/wtl/environment"
	"github.com/bazelbuild/rules_webtesting/go/wtl/errors"
)

const (
	name          = "External WebDriver Environment"
	addressEnvVar = "EXTERNAL_WEBDRIVER_SERVER_ADDRESS"
)

type external struct {
	*environment.Base
	address string
}

// NewEnv creates a new environment that uses an externally started Selenium Server.
func NewEnv(m *metadata.Metadata, d diagnostics.Diagnostics) (environment.Env, error) {
	address, ok := os.LookupEnv(addressEnvVar)
	if !ok {
		return nil, errors.New(name, fmt.Errorf("environment variable %q not set", addressEnvVar))
	}

	base, err := environment.NewBase(name, m, d)
	if err != nil {
		return nil, err
	}

	return &external{
		Base:    base,
		address: address,
	}, nil
}

// WDAddress returns the user-provided selenium address.
func (e *external) WDAddress(context.Context) string {
	return e.address
}
