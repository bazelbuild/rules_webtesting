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
// Package native provides an Env for launching a browser locally using
// GoogleSeleniumServer.
package native

import (
	"context"
	"os"

	"github.com/bazelbuild/rules_web/launcher/cmdhelper"
	"github.com/bazelbuild/rules_web/launcher/environments/environment"
	"github.com/bazelbuild/rules_web/launcher/services/selenium"
	"github.com/bazelbuild/rules_web/launcher/services/service"
	"github.com/bazelbuild/rules_web/metadata/metadata"
)

const (
	compName     = "native environment"
	forceXvfbEnv = "FORCE_DEDICATED_X_DISPLAY"
)

type native struct {
	*environment.Base
	selenium *service.Server
}

// NewEnv creates a new environment for launching a browser locally using
// GoogleSeleniumServer.
func NewEnv(m metadata.Metadata) (environment.Env, error) {
	base, err := environment.NewBase(compName, m)
	if err != nil {
		return nil, err
	}
	s, err := selenium.NewSelenium(m, useXvfb())
	if err != nil {
		return nil, err
	}

	return &native{
		Base:     base,
		selenium: s,
	}, nil
}

func (n *native) SetUp(ctx context.Context) error {
	if err := n.Base.SetUp(ctx); err != nil {
		return err
	}
	return n.selenium.Start(ctx)
}

func (n *native) TearDown(ctx context.Context) error {
	if err := n.Base.TearDown(ctx); err != nil {
		return err
	}
	return n.selenium.Stop(ctx)
}

func (n *native) WDAddress(context.Context) string {
	return n.selenium.Address
}

func (n *native) Healthy(ctx context.Context) error {
	if err := n.Base.Healthy(ctx); err != nil {
		return err
	}
	return n.selenium.Healthy(ctx)
}

func useXvfb() bool {
	return os.Getenv("DISPLAY") == "" || cmdhelper.IsTruthyEnv(forceXvfbEnv)
}
