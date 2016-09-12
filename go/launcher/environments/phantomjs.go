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
// Package phantomjs provides an Env for launching phantomjs locally using
// GoogleSeleniumServer.
package phantomjs

import (
	"context"

	"github.com/bazelbuild/rules_web/go/launcher/environments/environment"
	"github.com/bazelbuild/rules_web/go/launcher/services/selenium"
	"github.com/bazelbuild/rules_web/go/launcher/services/service"
	"github.com/bazelbuild/rules_web/go/metadata/capabilities"
	"github.com/bazelbuild/rules_web/go/metadata/metadata"
)

const (
	compName  = "PhantomJS environment"
	binaryCap = "phantomjs.binary.path"
)

type phantomJS struct {
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
	s, err := selenium.NewSelenium(m, false)
	if err != nil {
		return nil, err
	}

	return &phantomJS{
		Base:     base,
		selenium: s,
	}, nil
}

func (p *phantomJS) SetUp(ctx context.Context) error {
	if err := p.Base.SetUp(ctx); err != nil {
		return err
	}
	return p.selenium.Start(ctx)
}

func (p *phantomJS) TearDown(ctx context.Context) error {
	if err := p.Base.TearDown(ctx); err != nil {
		return err
	}
	return p.selenium.Stop(ctx)
}

func (p *phantomJS) WDAddress(context.Context) string {
	return p.selenium.Address
}

func (p *phantomJS) Healthy(ctx context.Context) error {
	if err := p.Base.Healthy(ctx); err != nil {
		return err
	}
	return p.selenium.Healthy(ctx)
}

// StartSession merges the passed in caps with b.metadata.caps and returns the merged
// capabilities that should be used when calling new session on the WebDriver
// server.
func (p *phantomJS) StartSession(ctx context.Context, id int, caps map[string]interface{}) (map[string]interface{}, error) {
	updated, err := p.Base.StartSession(ctx, id, caps)
	if err != nil {
		return nil, err
	}
	// TODO: Figure out a general mechanism for this.
	if phantomJS, err := p.Metadata.GetExecutablePath("PHANTOMJS"); err == nil {
		updated = capabilities.Merge(updated, map[string]interface{}{
			binaryCap: phantomJS,
		})
	}
	return updated, nil
}
