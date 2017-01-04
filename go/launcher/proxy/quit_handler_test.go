// Copyright 2017 Google Inc.
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

package quithandler

import (
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/util/bazel"
	"github.com/bazelbuild/rules_webtesting/go/webtest"
	"github.com/tebeka/selenium/selenium"
)

func TestCloseOneWindow(t *testing.T) {
	testpage, err := bazel.Runfile("go/launcher/proxy/testdata/testpage.html")
	if err != nil {
		t.Fatal(err)
	}

	driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	// Avoid Fatal so that this defer gets called
	defer driver.Quit()

	if err := driver.Get("file://" + testpage); err != nil {
		t.Error(err)
		return
	}

	if err := driver.CloseWindow(""); err != nil {
		t.Error(err)
		return
	}

	if _, err := driver.WindowHandles(); err == nil {
		t.Error("got nil, expected invalid session error")
		return
	}
}

func TestQuit(t *testing.T) {
	testpage, err := bazel.Runfile("go/launcher/proxy/testdata/testpage.html")
	if err != nil {
		t.Fatal(err)
	}

	driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	// Avoid Fatal so that this defer gets called
	defer driver.Quit()

	if err := driver.Get("file://" + testpage); err != nil {
		t.Error(err)
		return
	}

	button, err := driver.FindElement("tag name", "input")
	if err != nil {
		t.Error(err)
		return
	}

	if err := button.Click(); err != nil {
		t.Error(err)
		return
	}

	if err := driver.Quit(); err != nil {
		t.Error(err)
		return
	}

	if _, err := driver.WindowHandles(); err == nil {
		t.Error("got nil, expected invalid session error")
		return
	}
}

func TestCloseTwoWindows(t *testing.T) {
	testpage, err := bazel.Runfile("go/launcher/proxy/testdata/testpage.html")
	if err != nil {
		t.Fatal(err)
	}

	driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	// Avoid Fatal so that this defer gets called
	defer driver.Quit()

	if err := driver.Get("file://" + testpage); err != nil {
		t.Error(err)
		return
	}

	button, err := driver.FindElement("tag name", "input")
	if err != nil {
		t.Error(err)
		return
	}

	if err := button.Click(); err != nil {
		t.Error(err)
		return
	}

	if err := driver.CloseWindow(""); err != nil {
		t.Error(err)
		return
	}

	handles, err := driver.WindowHandles()
	if err != nil {
		t.Error(err)
		return
	}

	if len(handles) != 1 {
		t.Errorf("Got %d handles, expected 1", len(handles))
		return
	}

	if err := driver.SwitchWindow(handles[0]); err != nil {
		t.Error(err)
	}

	if err := driver.CloseWindow(""); err != nil {
		t.Error(err)
		return
	}

	if _, err := driver.WindowHandles(); err == nil {
		t.Error("got nil, expected invalid session error")
		return
	}
}
