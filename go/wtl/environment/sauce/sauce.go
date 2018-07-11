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
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/bazelbuild/rules_webtesting/go/metadata"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/wtl/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/wtl/environment"
	sc "github.com/bazelbuild/rules_webtesting/go/wtl/service/sauce"
)

const (
	name = "Sauce WebDriver Environment"
)

type sauce struct {
	*environment.Base
	address string
	connect *sc.Connect
	opts    sauceOptions
}

// NewEnv creates a new environment that uses an externally started Selenium Server.
func NewEnv(m *metadata.Metadata, d diagnostics.Diagnostics) (environment.Env, error) {
	opts := extractOptions(m)

	address := fmt.Sprintf("http://%s:%s@ondemand.saucelabs.com/wd/hub/", opts.username, opts.accessKey)

	var connect *sc.Connect

	if opts.startConnect {
		c, err := sc.New(m, opts.username, opts.accessKey, opts.tunnelID)
		if err != nil {
			return nil, err
		}

		connect = c
		address = c.Address
	}

	base, err := environment.NewBase(name, m, d)
	if err != nil {
		return nil, err
	}

	return &sauce{
		Base:    base,
		address: address,
		connect: connect,
		opts:    opts,
	}, nil
}

func (s *sauce) SetUp(ctx context.Context) error {
	if err := s.Base.SetUp(ctx); err != nil {
		return err
	}

	if s.connect != nil {
		return s.connect.Start(ctx)
	}
	return nil
}

func (s *sauce) TearDown(ctx context.Context) error {
	if err := s.Base.TearDown(ctx); err != nil {
		return err
	}

	if s.connect != nil {
		return s.connect.Stop(ctx)
	}
	return nil
}

func (s *sauce) Healthy(ctx context.Context) error {
	if err := s.Base.Healthy(ctx); err != nil {
		return err
	}
	if s.connect != nil {
		return s.connect.Healthy(ctx)
	}
	return nil
}

func (s *sauce) StartSession(ctx context.Context, id int, caps *capabilities.Capabilities) (*capabilities.Capabilities, error) {
	updated, err := s.Base.StartSession(ctx, id, caps)
	if err != nil {
		return nil, err
	}

	return updated.Resolve(capabilities.MapResolver("SAUCE", map[string]string{
		"TUNNEL_ID": s.opts.tunnelID,
	}))
}

// WDAddress returns the user-provided selenium address.
func (s *sauce) WDAddress(context.Context) string {
	return s.address
}

// sauceOptions is the set of options that can be defined in the sauceOptions section of
// a Metadata.Extension field.
type sauceOptions struct {
	// The Sauce username. If not defined in Extension, uses the env variable SAUCE_USERNAME.
	username string
	// The Sauce access key. If not defined in Extension, uses the env variable SAUCE_ACCESS_KEY.
	accessKey string
	// The Sauce tunnel id. This will be used to replace %SAUCE:TUNNEL_ID% in capabilities. If startConnect
	// then will be used to define the tunnel id for the started sauce connect. Can contain a single '%d', which will
	// be replaced with a random number in the range [0, 10000). If not defined, and startConnect is true, will
	// default to "tunnel-%d". Otherwise, if not defined, sues the env variable TUNNEL_IDENTIFIER.
	tunnelID string
	// Whether to start sauce connect. If true then will start Sauce Connect binary referenced by Web Test File SAUCE_CONNECT.
	// If not defined in Extension defaults to false.
	startConnect bool
}

func extractOptions(m *metadata.Metadata) sauceOptions {
	opts := sauceOptions{
		username:  os.Getenv("SAUCE_USERNAME"),
		accessKey: os.Getenv("SAUCE_ACCESS_KEY"),
		tunnelID:  os.Getenv("TUNNEL_IDENTIFIER"),
	}

	extMap, ok := m.ExtensionMap()
	if !ok {
		return opts
	}

	soMap, ok := extMap["sauceOptions"].(map[string]interface{})

	if !ok {
		return opts
	}

	if u, ok := soMap["username"].(string); ok {
		opts.username = u
	}

	if ak, ok := soMap["accessKey"].(string); ok {
		opts.accessKey = ak
	}

	if sc, ok := soMap["startConnect"].(bool); ok {
		opts.startConnect = sc
	}

	if tid, ok := soMap["tunnelId"].(string); ok {
		opts.tunnelID = tid
	} else if opts.startConnect {
		opts.tunnelID = "tunnel-%d"
	}

	if strings.Contains(opts.tunnelID, "%d") {
		r, err := rand.Int(rand.Reader, big.NewInt(10000))
		if err != nil {
			return opts
		}

		opts.tunnelID = fmt.Sprintf(opts.tunnelID, r)
	}

	return opts
}
