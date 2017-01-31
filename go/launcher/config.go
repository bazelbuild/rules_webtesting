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

package main

import (
	"github.com/bazelbuild/rules_webtesting/go/launcher/environments/chrome"
	"github.com/bazelbuild/rules_webtesting/go/launcher/environments/external"
	"github.com/bazelbuild/rules_webtesting/go/launcher/environments/firefox"
	"github.com/bazelbuild/rules_webtesting/go/launcher/environments/sauce"
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/driverhub"
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/googlescreenshot"
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/quithandler"
	"github.com/bazelbuild/rules_webtesting/go/launcher/proxy/screenshot"
)

func init() {
	// Configure WebDriver handlers.
	driverhub.HandlerProviderFunc(screenshot.ProviderFunc)
	driverhub.HandlerProviderFunc(quithandler.ProviderFunc)
	driverhub.HandlerProviderFunc(googlescreenshot.ProviderFunc)

	// Configure Environments.
	RegisterEnvProviderFunc("external", external.NewEnv)
	RegisterEnvProviderFunc("chrome", chrome.NewEnv)
	RegisterEnvProviderFunc("firefox", firefox.NewEnv)
	RegisterEnvProviderFunc("sauce", sauce.NewEnv)
}
