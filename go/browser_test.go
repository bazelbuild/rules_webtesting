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
package browser

import (
	"strings"
	"testing"

	"github.com/tebeka/selenium/selenium"
)

func TestProvisionBrowser_NoCaps(t *testing.T) {
	wd, err := NewSession(selenium.Capabilities{})
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
	wd, err := NewSession(selenium.Capabilities{
		"unexpectedAlertBehaviour": "dismiss",
		"elementScrollBehavior":    1,
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
	i, err := GetInfo()

	if err != nil {
		t.Fatal(err)
	}

	if strings.HasSuffix(i.TargetBrowserName, "linux") {
		if i.Environment != "native" {
			t.Errorf(`got Environment = %q, expected "native"`, i.Environment)
		}
		if i.FormFactor != "DESKTOP" {
			t.Errorf(`got FormFactor = %q, expected "DESKTOP"`, i.FormFactor)
		}
	}
}
