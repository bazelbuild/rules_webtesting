/* Copyright 2016 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metadata

import (
	"testing"

	"github.com/web_test_launcher/util/bazel"
)

const (
	allFields      = "__main__/metadata/testdata/all-fields.json"
	chromeLinux    = "__main__/metadata/testdata/chrome-linux.json"
	androidBrowser = "__main__/metadata/testdata/android-browser-gingerbread-nexus-s.json"
	fakeBrowser    = "__main__/metadata/testdata/merge-from-file-result.json"
)

func TestFromFile(t *testing.T) {
	f, err := bazel.Runfile(allFields)
	if err != nil {
		t.Fatal(err)
	}
	file, err := FromFile(f)
	if err != nil {
		t.Fatal(err)
	}

	expected := Metadata{
		FormFactor:      "PHONE",
		BrowserName:     "chrome",
		Environment:     "chromeos",
		BrowserLabel:    "//testing/web/browsers:figaro",
		TestLabel:       "//testing/web/launcher:tests",
		CropScreenshots: true,
		RecordVideo:     RecordAlways,
	}

	if !Equals(expected, file) {
		t.Errorf("Got %+v, expected %+v", file, expected)
	}
}

func TestMergeFromFile(t *testing.T) {
	f1, err := bazel.Runfile(chromeLinux)
	if err != nil {
		t.Fatal(err)
	}
	cl, err := FromFile(f1)
	if err != nil {
		t.Fatal(err)
	}

	f2, err := bazel.Runfile(androidBrowser)
	if err != nil {
		t.Fatal(err)
	}
	ab, err := FromFile(f2)
	if err != nil {
		t.Fatal(err)
	}

	f3, err := bazel.Runfile(fakeBrowser)
	if err != nil {
		t.Fatal(err)
	}
	fb, err := FromFile(f3)
	if err != nil {
		t.Fatal(err)
	}

	if merged := Merge(cl, ab); !Equals(merged, fb) {
		t.Errorf("Got Merge(%+v, %+v) == %+v, expected %+v", cl, ab, merged, fb)
	}
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		name   string
		input1 Metadata
		input2 Metadata
		result Metadata
	}{
		{
			"FormFactor override",
			Metadata{FormFactor: "PHONE"},
			Metadata{FormFactor: "TABLET"},
			Metadata{FormFactor: "TABLET"},
		},
		{
			"FormFactor no override",
			Metadata{FormFactor: "PHONE"},
			Metadata{FormFactor: ""},
			Metadata{FormFactor: "PHONE"},
		},
		{
			"BrowserName override",
			Metadata{BrowserName: "chrome"},
			Metadata{BrowserName: "firefox"},
			Metadata{BrowserName: "firefox"},
		},
		{
			"BrowserName no override",
			Metadata{BrowserName: "chrome"},
			Metadata{BrowserName: ""},
			Metadata{BrowserName: "chrome"},
		},
		{
			"Environment override",
			Metadata{Environment: "linux"},
			Metadata{Environment: "android"},
			Metadata{Environment: "android"},
		},
		{
			"Environment no override",
			Metadata{Environment: "linux"},
			Metadata{Environment: ""},
			Metadata{Environment: "linux"},
		},
		{
			"BrowserLabel override",
			Metadata{BrowserLabel: "//testing/web/browsers:figaro"},
			Metadata{BrowserLabel: "//testing/web/browsers:murphy"},
			Metadata{BrowserLabel: "//testing/web/browsers:murphy"},
		},
		{
			"BrowserLabel no override",
			Metadata{BrowserLabel: "//testing/web/browsers:figaro"},
			Metadata{BrowserLabel: ""},
			Metadata{BrowserLabel: "//testing/web/browsers:figaro"},
		},
		{
			"TestLabel override",
			Metadata{TestLabel: "//testing/web/browsers:figaro"},
			Metadata{TestLabel: "//testing/web/browsers:murphy"},
			Metadata{TestLabel: "//testing/web/browsers:murphy"},
		},
		{
			"TestLabel no override",
			Metadata{TestLabel: "//testing/web/browsers:figaro"},
			Metadata{TestLabel: ""},
			Metadata{TestLabel: "//testing/web/browsers:figaro"},
		},
		{
			"CropScreenshots override with false",
			Metadata{CropScreenshots: true},
			Metadata{CropScreenshots: false},
			Metadata{CropScreenshots: false},
		},
		{
			"CropScreenshots override with true",
			Metadata{CropScreenshots: false},
			Metadata{CropScreenshots: true},
			Metadata{CropScreenshots: true},
		},
		{
			"CropScreenshots no override false",
			Metadata{CropScreenshots: false},
			Metadata{CropScreenshots: nil},
			Metadata{CropScreenshots: false},
		},
		{
			"CropScreenshots no override true",
			Metadata{CropScreenshots: true},
			Metadata{CropScreenshots: nil},
			Metadata{CropScreenshots: true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := Merge(tc.input1, tc.input2)
			if !Equals(a, tc.result) {
				t.Errorf("Got Merge(%+v, %+v) == %+v, expected %+v", tc.input1, tc.input2, a, tc.result)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	testCases := []struct {
		name   string
		input1 Metadata
		input2 Metadata
		result bool
	}{
		{
			"empty",
			Metadata{},
			Metadata{},
			true,
		},
		{
			"FormFactor same",
			Metadata{FormFactor: "PHONE"},
			Metadata{FormFactor: "PHONE"},
			true,
		},
		{
			"FormFactor different",
			Metadata{FormFactor: "PHONE"},
			Metadata{FormFactor: "TABLET"},
			false,
		},
		{
			"BrowserName same",
			Metadata{BrowserName: "chrome"},
			Metadata{BrowserName: "chrome"},
			true,
		},
		{
			"BrowserName different",
			Metadata{BrowserName: "chrome"},
			Metadata{BrowserName: "firefox"},
			false,
		},
		{
			"Environment same",
			Metadata{Environment: "local"},
			Metadata{Environment: "local"},
			true,
		},
		{
			"Environment different",
			Metadata{Environment: "local"},
			Metadata{Environment: "running"},
			false,
		},
		{
			"BrowserLabel same",
			Metadata{BrowserLabel: "//testing/web/browsers:firefox"},
			Metadata{BrowserLabel: "//testing/web/browsers:firefox"},
			true,
		},
		{
			"BrowserLabel different",
			Metadata{BrowserLabel: "//testing/web/browsers:chrome"},
			Metadata{BrowserLabel: "//testing/web/browsers:firefox"},
			false,
		},
		{
			"TestLabel same",
			Metadata{TestLabel: "//test:test1"},
			Metadata{TestLabel: "//test:test1"},
			true,
		},
		{
			"TestLabel different",
			Metadata{TestLabel: "//test:test1"},
			Metadata{TestLabel: "//test:test2"},
			false,
		},
		{
			"CropScreenshots same",
			Metadata{CropScreenshots: true},
			Metadata{CropScreenshots: true},
			true,
		},
		{
			"CropScreenshots different",
			Metadata{CropScreenshots: nil},
			Metadata{CropScreenshots: false},
			false,
		},
		{
			"RecordVideo same",
			Metadata{RecordVideo: RecordAlways},
			Metadata{RecordVideo: RecordAlways},
			true,
		},
		{
			"RecordVideo different",
			Metadata{RecordVideo: RecordNever},
			Metadata{RecordVideo: RecordAlways},
			false,
		},
		{
			"HealthyBeforeTest same",
			Metadata{HealthyBeforeTest: true},
			Metadata{HealthyBeforeTest: true},
			true,
		},
		{
			"HealthyBeforeTest different",
			Metadata{HealthyBeforeTest: nil},
			Metadata{HealthyBeforeTest: false},
			false,
		},
		{
			"Capabilities same",
			Metadata{Capabilities: map[string]interface{}{"browser": "chrome"}},
			Metadata{Capabilities: map[string]interface{}{"browser": "chrome"}},
			true,
		},
		{
			"Capabilities different",
			Metadata{Capabilities: map[string]interface{}{"browser": "chrome"}},
			Metadata{Capabilities: map[string]interface{}{"browser": "firefox"}},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := Equals(tc.input1, tc.input2); result != tc.result {
				t.Errorf("Got Equals(%+v, %+v) == %v, expected %v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}
}
