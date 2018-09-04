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
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/bazel"
	"github.com/bazelbuild/rules_webtesting/go/portpicker"
	"github.com/bazelbuild/rules_webtesting/go/webtest"
	"github.com/tebeka/selenium"
)

var testpage = ""

func TestMain(m *testing.M) {
	port, err := portpicker.PickUnusedPort()
	if err != nil {
		log.Fatal(err)
	}

	dir, err := bazel.Runfile("testdata/")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir))))
	}()

	testpage = fmt.Sprintf("http://localhost:%d/testpage.html", port)

	os.Exit(m.Run())
}

func TestCloseOneWindow(t *testing.T) {
	driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	defer driver.Quit()

	if err := driver.Get(testpage); err != nil {
		t.Fatal(err)
	}

	if err := driver.CloseWindow(""); err != nil {
		t.Fatal(err)
	}

	if _, err := driver.WindowHandles(); err == nil {
		t.Fatal("got nil, expected invalid session error")
	}
}

func TestQuit(t *testing.T) {
	driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	defer driver.Quit()

	if err := driver.Get(testpage); err != nil {
		t.Fatal(err)
	}

	button, err := driver.FindElement("tag name", "input")
	if err != nil {
		t.Fatal(err)
	}

	if err := button.Click(); err != nil {
		t.Fatal(err)
	}

	if err := driver.Quit(); err != nil {
		t.Fatal(err)
	}

	if _, err := driver.WindowHandles(); err == nil {
		t.Fatal("got nil, expected invalid session error")
	}
}

func TestCloseTwoWindows(t *testing.T) {
	driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	defer driver.Quit()

	if err := driver.Get(testpage); err != nil {
		t.Fatal(err)
	}

	button, err := driver.FindElement("tag name", "input")
	if err != nil {
		t.Fatal(err)
	}

	if err := button.Click(); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		handles, err := driver.WindowHandles()
		if err != nil {
			t.Fatal(err)
		}
		if len(handles) == 2 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	handles, err := driver.WindowHandles()
	if err != nil {
		t.Fatal(err)
	}
	if len(handles) != 2 {
		t.Fatalf("got %v, expected two window handles")
	}

	if err := driver.CloseWindow(""); err != nil {
		t.Fatal(err)
	}

	handles, err = driver.WindowHandles()
	if err != nil {
		t.Fatal(err)
	}

	if len(handles) != 1 {
		t.Fatalf("Got %d handles, expected 1", len(handles))
	}

	if err := driver.SwitchWindow(handles[0]); err != nil {
		t.Fatal(err)
	}

	if err := driver.CloseWindow(""); err != nil {
		t.Fatal(err)
	}

	if _, err := driver.WindowHandles(); err == nil {
		t.Fatal("got nil, expected invalid session error")
	}
}
