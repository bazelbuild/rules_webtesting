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

package screenshot

import (
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/util/bazel"
	"github.com/bazelbuild/rules_webtesting/go/webtest"
	"github.com/tebeka/selenium/selenium"
)

func TestScreenshot(t *testing.T) {
	driver, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	defer driver.Quit()

	if err := driver.Get("http://www.google.com"); err != nil {
		t.Error(err)
		return
	}

	img, err := driver.Screenshot()

	if err != nil {
		t.Error(err)
		return
	}

	tmpDir, err := bazel.NewTmpDir("crop_test")
	if err != nil {
		t.Error(err)
		return
	}

	outPath := filepath.Join(tmpDir, "cropped.png")

	if err := ioutil.WriteFile(outPath, img, 0660); err != nil {
		t.Error(err)
		return
	}

	check, err := os.Open(outPath)
	if err != nil {
		t.Error(err)
		return
	}

	config, err := png.DecodeConfig(check)
	check.Close()
	if err != nil {
		t.Error(err)
		return
	}

	if config.Width != 480 || config.Height != 800 {
		t.Errorf("got size == %d, %d, expected 800, 480", config.Width, config.Height)
	}
}
