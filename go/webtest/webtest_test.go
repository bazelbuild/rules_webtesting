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

package webtest

import (
	"strings"
	"testing"

	"github.com/tebeka/selenium"
)

func TestProvisionBrowser_NoCaps(t *testing.T) {
	wd, err := NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	if err := wd.Get("about:"); err != nil {
		t.Error(err)
	}

	url, err := wd.CurrentURL()
	if err != nil {
		t.Error(err)
	}
	if url == "" {
		t.Error("Got empty url")
	}

	if err := wd.Quit(); err != nil {
		t.Error(err)
	}
}

func TestProvisionBrowser_WithCaps(t *testing.T) {
	wd, err := NewWebDriverSession(selenium.Capabilities{
		"acceptInsecureCerts": false,
		"pageLoadStrategy":    "normal",
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := wd.Get("about:"); err != nil {
		t.Error(err)
	}

	url, err := wd.CurrentURL()
	if err != nil {
		t.Error(err)
	}
	if url == "" {
		t.Error("Got empty url")
	}

	if err := wd.Quit(); err != nil {
		t.Error(err)
	}
}

func TestGetInfo(t *testing.T) {
	i, err := GetBrowserInfo()

	if err != nil {
		t.Fatal(err)
	}

	switch {
	case strings.Contains(i.BrowserLabel, "sauce"):
		if i.Environment != "sauce" {
			t.Errorf(`got Environment = %q, expected "sauce"`, i.Environment)
		}
	case strings.Contains(i.BrowserLabel, "chrome"):
		if i.Environment != "chrome" {
			t.Errorf(`got Environment = %q, expected "chrome"`, i.Environment)
		}
	case strings.Contains(i.BrowserLabel, "firefox"):
		if i.Environment != "firefox" {
			t.Errorf(`got Environment = %q, expected "firefox"`, i.Environment)
		}
	}
}

func TestReusableBrowser(t *testing.T) {
	wd1, err := NewWebDriverSession(selenium.Capabilities{
		"google.canReuseSession": true,
	})
	if err != nil {
		t.Fatal(err)
	}

	id1 := wd1.SessionID()
	if err := wd1.Quit(); err != nil {
		t.Error(err)
	}

	wd2, err := NewWebDriverSession(selenium.Capabilities{
		"google.canReuseSession": true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if id1 != wd2.SessionID() {
		t.Errorf("got different ids %q, %q, expected them to be the same.", id1, wd2.SessionID())
	}

	if err := wd2.Get("about:"); err != nil {
		t.Error(err)
	}

	url, err := wd2.CurrentURL()
	if err != nil {
		t.Error(err)
	}
	if url == "" {
		t.Error("Got empty url")
	}

	if err := wd2.Quit(); err != nil {
		t.Error(err)
	}
}
