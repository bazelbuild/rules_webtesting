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

package screenshotter

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"

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

	dir = filepath.Join(dir, bazel.TestWorkspace())

	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir))))
	}()

	testpage = fmt.Sprintf("http://localhost:%d/testpage.html", port)

	os.Exit(m.Run())
}

func TestOfElement(t *testing.T) {
	driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	defer driver.Quit()

	shotter, err := New(driver)

	if err := driver.Get(testpage); err != nil {
		t.Fatal(err)
	}

	el, err := driver.FindElement(selenium.ByTagName, "input")
	if err != nil {
		t.Fatal(err)
	}

	img, err := shotter.Of(el).TakeScreenshot()

	if err != nil {
		t.Fatal(err)
	}

	tmpDir, err := bazel.NewTmpDir("crop_test")
	if err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(tmpDir, "of-element.png")

	log.Printf("Output file: %q", outPath)

	if err := ioutil.WriteFile(outPath, img, 0660); err != nil {
		t.Fatal(err)
	}

	check, err := os.Open(outPath)
	if err != nil {
		t.Fatal(err)
	}
	defer check.Close()

	config, err := png.DecodeConfig(check)
	if err != nil {
		t.Fatal(err)
	}

	if config.Width != 200 || config.Height != 200 {
		t.Fatalf("got size == %d, %d, expected 200, 200", config.Width, config.Height)
	}
}

func TestExcluding(t *testing.T) {
	driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	defer driver.Quit()

	shotter, err := New(driver)

	if err := driver.Get(testpage); err != nil {
		t.Fatal(err)
	}

	outer, err := driver.FindElement(selenium.ByID, "outer-div")
	if err != nil {
		t.Fatal(err)
	}

	ex1, err := driver.FindElement(selenium.ByID, "inner-div1")
	if err != nil {
		t.Fatal(err)
	}

	ex2, err := driver.FindElement(selenium.ByID, "b")
	if err != nil {
		t.Fatal(err)
	}

	img, err := shotter.Of(outer).Excluding(ex1).Excluding(ex2).TakeScreenshot()

	if err != nil {
		t.Fatal(err)
	}

	tmpDir, err := bazel.NewTmpDir("exclude-element")
	if err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(tmpDir, "cropped.png")

	log.Printf("Output file: %q", outPath)

	if err := ioutil.WriteFile(outPath, img, 0660); err != nil {
		t.Fatal(err)
	}
}
