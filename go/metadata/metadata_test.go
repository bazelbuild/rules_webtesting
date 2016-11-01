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

	if merged, err := Merge(cl, ab); err != nil {
		t.Error(err)
	} else if !Equals(merged, fb) {
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
			a, err := Merge(tc.input1, tc.input2)
			if err != nil {
				t.Error(err)
			} else if !Equals(a, tc.result) {
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

func TestMergeNamedFiles(t *testing.T) {
	testCases := []struct {
		name   string
		input1 map[string]string
		input2 map[string]string
		result map[string]string // nil indicates should return an error
	}{
		{
			"empty",
			map[string]string{},
			map[string]string{},
			map[string]string{},
		},
		{
			"duplicate names, different paths",
			map[string]string{"a": "b"},
			map[string]string{"a": "c"},
			nil,
		},
		{
			"duplicate names, same paths",
			map[string]string{"a": "b"},
			map[string]string{"a": "b"},
			map[string]string{"a": "b"},
		},
		{
			"multiple names, successful",
			map[string]string{"a": "A", "b": "B", "c": "C"},
			map[string]string{"a": "A", "d": "D", "e": "E"},
			map[string]string{"a": "A", "b": "B", "c": "C", "d": "D", "e": "E"},
		},
		{
			"multiple names, unsuccessful",
			map[string]string{"a": "A", "b": "B", "c": "C"},
			map[string]string{"a": "A", "d": "D", "e": "E", "c": "X"},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := mergeNamedFiles(tc.input1, tc.input2)
			if err != nil {
				if tc.result != nil {
					t.Error(err)
				}
				return
			}
			if tc.result == nil {
				t.Errorf("Got mergeNamedFiles(%+v, %+v) == %+v, expected error", tc.input1, tc.input2, result)
				return
			}
			if !mapEquals(result, tc.result) {
				t.Errorf("Got mergeNamedFiles(%+v, %+v) == %v, expected %v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}
}

func TestMergeArchive(t *testing.T) {
	testCases := []struct {
		name   string
		input1 *Archive
		input2 *Archive
		result *Archive // nil indicates should return an error
	}{
		{
			"empty",
			&Archive{NamedFiles: map[string]string{}},
			&Archive{NamedFiles: map[string]string{}},
			&Archive{NamedFiles: map[string]string{}},
		},
		{
			"different archive paths",
			&Archive{Path: "a", NamedFiles: map[string]string{}},
			&Archive{Path: "b", NamedFiles: map[string]string{}},
			nil,
		},
		{
			"different named file paths",
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}},
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "X"}},
			nil,
		},
		{
			"same named file paths",
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}},
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}},
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}},
		},
		{
			"multiple names, successful",
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "A", "b": "B", "c": "C"}},
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "A", "d": "D", "e": "E"}},
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "A", "b": "B", "c": "C", "d": "D", "e": "E"}},
		},
		{
			"multiple names, unsuccessful",
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "A", "b": "B", "c": "C"}},
			&Archive{Path: "a", NamedFiles: map[string]string{"a": "A", "d": "D", "e": "E", "c": "X"}},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := mergeArchive(tc.input1, tc.input2)
			if err != nil {
				if tc.result != nil {
					t.Error(err)
				}
				return
			}
			if tc.result == nil {
				t.Errorf("Got mergeArchive(%+v, %+v) == %+v, expected error", tc.input1, tc.input2, result)
				return
			}
			if result.Path != tc.result.Path || !mapEquals(result.NamedFiles, tc.result.NamedFiles) {
				t.Errorf("Got mergeArchive(%+v, %+v) == %+v, expected %+v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}

}

func TestMergeArchives(t *testing.T) {
	testCases := []struct {
		name   string
		input1 []*Archive
		input2 []*Archive
		// map of archive paths to NamedFiles maps
		// nil indicates should return an error
		result map[string]map[string]string
	}{
		{
			"empty",
			[]*Archive{},
			[]*Archive{},
			map[string]map[string]string{},
		},
		{
			"unmergeable archives",
			[]*Archive{&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}}},
			[]*Archive{&Archive{Path: "a", NamedFiles: map[string]string{"a": "X"}}},
			nil,
		},
		{
			"mergeable archives",
			[]*Archive{&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}}},
			[]*Archive{&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}}},
			map[string]map[string]string{"a": map[string]string{"a": "A"}},
		},
		{
			"multiple archives, success",
			[]*Archive{
				&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}},
				&Archive{Path: "b", NamedFiles: map[string]string{"b": "B"}},
			},
			[]*Archive{
				&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}},
				&Archive{Path: "c", NamedFiles: map[string]string{"c": "C"}},
			},
			map[string]map[string]string{
				"a": map[string]string{"a": "A"},
				"b": map[string]string{"b": "B"},
				"c": map[string]string{"c": "C"},
			},
		},
		{
			"multiple archives, failure",
			[]*Archive{
				&Archive{Path: "a", NamedFiles: map[string]string{"a": "A"}},
				&Archive{Path: "b", NamedFiles: map[string]string{"b": "B"}},
			},
			[]*Archive{
				&Archive{Path: "a", NamedFiles: map[string]string{"a": "X"}},
				&Archive{Path: "c", NamedFiles: map[string]string{"c": "C"}},
			},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := mergeArchives(tc.input1, tc.input2)
			if err != nil {
				if tc.result != nil {
					t.Error(err)
				}
				return
			}
			if tc.result == nil {
				t.Errorf("Got mergeArchives(%+v, %+v) == %+v, expected error", tc.input1, tc.input2, result)
				return
			}

			for _, a := range result {
				nf, ok := tc.result[a.Path]
				delete(tc.result, a.Path)
				if !ok {
					t.Errorf("Result included unexpected archive %+v", a)
					continue
				}
				if !mapEquals(a.NamedFiles, nf) {
					t.Errorf("Got archive %+v, expected NamedFiles to be %+v", a, nf)
				}
			}

			if len(tc.result) != 0 {
				t.Errorf("Missing archives %+v from result", tc.result)
			}
		})
	}
}

func TestValidateNoDuplicateNamedFiles(t *testing.T) {
	testCases := []struct {
		name       string
		namedFiles map[string]string
		archives   []*Archive
		err        bool
	}{
		{
			"empty",
			map[string]string{},
			[]*Archive{},
			false,
		},
		{
			"duplicate between NamedFiles and Archive",
			map[string]string{"a": "A"},
			[]*Archive{&Archive{NamedFiles: map[string]string{"a": "A"}}},
			true,
		},
		{
			"duplicate between two Archives",
			map[string]string{},
			[]*Archive{
				&Archive{NamedFiles: map[string]string{"a": "A"}},
				&Archive{NamedFiles: map[string]string{"a": "A"}},
			},
			true,
		},
		{
			"no duplicates",
			map[string]string{"a": "A", "b": "B"},
			[]*Archive{
				&Archive{NamedFiles: map[string]string{"c": "C", "d": "D"}},
				&Archive{NamedFiles: map[string]string{"e": "E"}},
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateNoDuplicateNamedFiles(tc.namedFiles, tc.archives)
			if err == nil && tc.err {
				t.Error("Got nil, expected err")
			}
			if err != nil && !tc.err {
				t.Errorf("Exoected nil, got %v", err)
			}
		})
	}
}
