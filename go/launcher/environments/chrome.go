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

// Package native provides an Env for launching a chrome browser locally using
// ChromeDriver.
package chrome

import (
	"context"
	"os"

	"github.com/bazelbuild/rules_webtesting/go/launcher/cmdhelper"
	"github.com/bazelbuild/rules_webtesting/go/launcher/environments/environment"
	"github.com/bazelbuild/rules_webtesting/go/launcher/services/chromedriver"
	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
)

const (
	compName       = "chrome environment"
	forceXvfbEnv   = "FORCE_DEDICATED_X_DISPLAY"
	disableXvfbEnv = "DISABLE_X_DISPLAY"
)

type chrome struct {
	*environment.Base
	chromedriver *chromedriver.ChromeDriver
}

// NewEnv creates a new environment for launching a chrome browser locally using
// ChromeDriver.
func NewEnv(m *metadata.Metadata) (environment.Env, error) {
	base, err := environment.NewBase(compName, m)
	if err != nil {
		return nil, err
	}
	cd, err := chromedriver.New(m, useXvfb())
	if err != nil {
		return nil, err
	}

	return &chrome{
		Base:         base,
		chromedriver: cd,
	}, nil
}

func (n *chrome) SetUp(ctx context.Context) error {
	if err := n.Base.SetUp(ctx); err != nil {
		return err
	}
	return n.chromedriver.Start(ctx)
}

func (n *chrome) TearDown(ctx context.Context) error {
	if err := n.Base.TearDown(ctx); err != nil {
		return err
	}
	return n.chromedriver.Stop(ctx)
}

func (n *chrome) WDAddress(context.Context) string {
	return n.chromedriver.Address()
}

func (n *chrome) Healthy(ctx context.Context) error {
	if err := n.Base.Healthy(ctx); err != nil {
		return err
	}
	return n.chromedriver.Healthy(ctx)
}

func useXvfb() bool {
	return !cmdhelper.IsTruthyEnv(disableXvfbEnv) && (os.Getenv("DISPLAY") == "" || cmdhelper.IsTruthyEnv(forceXvfbEnv))
}
