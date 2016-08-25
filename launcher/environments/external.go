/* Copyright 2016 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package external

import (
	"context"
	"fmt"
	"os"

	"github.com/bazelbuild/rules_web/launcher/environments/environment"
	"github.com/bazelbuild/rules_web/launcher/errors"
	"github.com/bazelbuild/rules_web/metadata/metadata"
)

const (
	name            = "External WebDriver Environment"
	address_env_var = "EXTERNAL_WEBDRIVER_SERVER_ADDRESS"
)

type external struct {
	*environment.Base
	address string
}

// NewEnv creates a new environment that uses an externally started Selenium Server.
func NewEnv(m metadata.Metadata) (environment.Env, error) {
	address, ok := os.LookupEnv(address_env_var)
	if !ok {
		return nil, errors.New(name, fmt.Errorf("environment variable %q not set", address_env_var))
	}

	base, err := environment.NewBase(name, m)
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
