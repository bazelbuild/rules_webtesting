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
package metadata

import (
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/util/bazel"
)

const (
	allFields      = "io_bazel_rules_webtesting/go/metadata/testdata/all-fields.json"
	chromeLinux    = "io_bazel_rules_webtesting/go/metadata/testdata/chrome-linux.json"
	androidBrowser = "io_bazel_rules_webtesting/go/metadata/testdata/android-browser-gingerbread-nexus-s.json"
	fakeBrowser    = "io_bazel_rules_webtesting/go/metadata/testdata/merge-from-file-result.json"
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

	expected := &Metadata{
		Environment:  "chromeos",
		BrowserLabel: "//browsers:figaro",
		TestLabel:    "//go/launcher:tests",
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
		input1 *Metadata
		input2 *Metadata
		result *Metadata
	}{
		{
			"Environment override",
			&Metadata{Environment: "linux"},
			&Metadata{Environment: "android"},
			&Metadata{Environment: "android"},
		},
		{
			"Environment no override",
			&Metadata{Environment: "linux"},
			&Metadata{Environment: ""},
			&Metadata{Environment: "linux"},
		},
		{
			"BrowserLabel override",
			&Metadata{BrowserLabel: "//browsers:figaro"},
			&Metadata{BrowserLabel: "//browsers:murphy"},
			&Metadata{BrowserLabel: "//browsers:murphy"},
		},
		{
			"BrowserLabel no override",
			&Metadata{BrowserLabel: "//browsers:figaro"},
			&Metadata{BrowserLabel: ""},
			&Metadata{BrowserLabel: "//browsers:figaro"},
		},
		{
			"TestLabel override",
			&Metadata{TestLabel: "//browsers:figaro"},
			&Metadata{TestLabel: "//browsers:murphy"},
			&Metadata{TestLabel: "//browsers:murphy"},
		},
		{
			"TestLabel no override",
			&Metadata{TestLabel: "//browsers:figaro"},
			&Metadata{TestLabel: ""},
			&Metadata{TestLabel: "//browsers:figaro"},
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
		input1 *Metadata
		input2 *Metadata
		result bool
	}{
		{
			"empty",
			&Metadata{},
			&Metadata{},
			true,
		},
		{
			"Environment same",
			&Metadata{Environment: "local"},
			&Metadata{Environment: "local"},
			true,
		},
		{
			"Environment different",
			&Metadata{Environment: "local"},
			&Metadata{Environment: "running"},
			false,
		},
		{
			"BrowserLabel same",
			&Metadata{BrowserLabel: "//browsers:firefox"},
			&Metadata{BrowserLabel: "//browsers:firefox"},
			true,
		},
		{
			"BrowserLabel different",
			&Metadata{BrowserLabel: "//browsers:chrome"},
			&Metadata{BrowserLabel: "//browsers:firefox"},
			false,
		},
		{
			"TestLabel same",
			&Metadata{TestLabel: "//test:test1"},
			&Metadata{TestLabel: "//test:test1"},
			true,
		},
		{
			"TestLabel different",
			&Metadata{TestLabel: "//test:test1"},
			&Metadata{TestLabel: "//test:test2"},
			false,
		},
		{
			"Capabilities same",
			&Metadata{Capabilities: map[string]interface{}{"browser": "chrome"}},
			&Metadata{Capabilities: map[string]interface{}{"browser": "chrome"}},
			true,
		},
		{
			"Capabilities different",
			&Metadata{Capabilities: map[string]interface{}{"browser": "chrome"}},
			&Metadata{Capabilities: map[string]interface{}{"browser": "firefox"}},
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
