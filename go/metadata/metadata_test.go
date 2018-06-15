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

package metadata

import (
	"reflect"
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/bazel"
)

const (
	allFields      = "testdata/all-fields.json"
	chromeLinux    = "testdata/chrome-linux.json"
	androidBrowser = "testdata/android-browser-gingerbread-nexus-s.json"
	fakeBrowser    = "testdata/merge-from-file-result.json"
	badNamedFiles  = "testdata/bad-named-files.json"
)

func TestFromFile(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		f, err := bazel.Runfile(allFields)
		if err != nil {
			t.Fatal(err)
		}
		file, err := FromFile(f, nil)
		if err != nil {
			t.Fatal(err)
		}

		expected := &Metadata{
			Capabilities: map[string]interface{}{},
			Environment:  "chromeos",
			BrowserLabel: "//browsers:figaro",
			TestLabel:    "//go/wtl:tests",
			Extension:    &extension{},
		}

		if !reflect.DeepEqual(expected, file) {
			t.Errorf("Got %#v, expected %#v", file, expected)
		}
	})

	t.Run("bad named files", func(t *testing.T) {
		f, err := bazel.Runfile(badNamedFiles)
		if err != nil {
			t.Fatal(err)
		}
		d, err := FromFile(f, nil)
		if err == nil {
			t.Errorf("Got %#v, expected err", d)
		}
	})
}

func TestMergeFromFile(t *testing.T) {
	f1, err := bazel.Runfile(chromeLinux)
	if err != nil {
		t.Fatal(err)
	}
	cl, err := FromFile(f1, nil)
	if err != nil {
		t.Fatal(err)
	}

	f2, err := bazel.Runfile(androidBrowser)
	if err != nil {
		t.Fatal(err)
	}
	ab, err := FromFile(f2, nil)
	if err != nil {
		t.Fatal(err)
	}

	f3, err := bazel.Runfile(fakeBrowser)
	if err != nil {
		t.Fatal(err)
	}
	fb, err := FromFile(f3, nil)
	if err != nil {
		t.Fatal(err)
	}

	merged, err := Merge(cl, ab)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(merged, fb) {
		t.Errorf("Got %#v, expected %#v", merged, fb)
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
		{
			"EnableDebugger, no override",
			&Metadata{DebuggerPort: 1},
			&Metadata{DebuggerPort: 0},
			&Metadata{DebuggerPort: 1},
		},
		{
			"EnableDebugger, override",
			&Metadata{DebuggerPort: 1},
			&Metadata{DebuggerPort: 2},
			&Metadata{DebuggerPort: 2},
		},
		{
			"EnableDebugger, not set",
			&Metadata{DebuggerPort: 0},
			&Metadata{DebuggerPort: 0},
			&Metadata{DebuggerPort: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, err := Merge(tc.input1, tc.input2)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(a, tc.result) {
				t.Errorf("Got %#v, expected %#v", a, tc.result)
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
				t.Errorf("Got %s, expected error", result)
				return
			}
			if !reflect.DeepEqual(result, tc.result) {
				t.Errorf("Got %s, expected %s", result, tc.result)
			}
		})
	}
}

func TestMergeWebTestFiles(t *testing.T) {
	testCases := []struct {
		name   string
		input1 *WebTestFiles
		input2 *WebTestFiles
		result *WebTestFiles // nil indicates should return an error
	}{
		{
			"empty",
			&WebTestFiles{NamedFiles: map[string]string{}},
			&WebTestFiles{NamedFiles: map[string]string{}},
			&WebTestFiles{NamedFiles: map[string]string{}},
		},
		{
			"different archive paths",
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{}},
			&WebTestFiles{ArchiveFile: "b", NamedFiles: map[string]string{}},
			nil,
		},
		{
			"different named file paths",
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "X"}},
			nil,
		},
		{
			"same named file paths",
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
		},
		{
			"multiple names, successful",
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A", "b": "B", "c": "C"}},
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A", "d": "D", "e": "E"}},
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A", "b": "B", "c": "C", "d": "D", "e": "E"}},
		},
		{
			"multiple names, unsuccessful",
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A", "b": "B", "c": "C"}},
			&WebTestFiles{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A", "d": "D", "e": "E", "c": "X"}},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := mergeWebTestFiles(tc.input1, tc.input2)
			if err != nil {
				if tc.result != nil {
					t.Fatal(err)
				}
				return
			}
			if tc.result == nil {
				t.Fatalf("Got %v, expected error", result)
			}
			if !reflect.DeepEqual(result, tc.result) {
				t.Errorf("Got %v, expected %v", result, tc.result)
			}
		})
	}

}

func TestNormalizeWebTestFiles(t *testing.T) {
	testCases := []struct {
		name  string
		input []*WebTestFiles
		// map of archive paths to NamedFiles maps
		// nil indicates should return an error
		result []*WebTestFiles
		err    bool
	}{
		{
			"empty",
			nil,
			nil,
			false,
		},
		{
			"unnormalizable WebTestFiles",
			[]*WebTestFiles{
				{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
				{ArchiveFile: "a", NamedFiles: map[string]string{"a": "X"}},
			},
			nil,
			true,
		},
		{
			"normalizable WebTestFiles",
			[]*WebTestFiles{
				{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
				{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
			},
			[]*WebTestFiles{
				{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
			},
			false,
		},
		{
			"multiple WebTestFiles, success",
			[]*WebTestFiles{
				{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
				{ArchiveFile: "b", NamedFiles: map[string]string{"b": "B"}},
				{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A", "d": "D"}},
				{ArchiveFile: "c", NamedFiles: map[string]string{"c": "C"}},
			},
			[]*WebTestFiles{
				{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A", "d": "D"}},
				{ArchiveFile: "b", NamedFiles: map[string]string{"b": "B"}},
				{ArchiveFile: "c", NamedFiles: map[string]string{"c": "C"}},
			},
			false,
		},
		{
			"multiple WebTestFiles, failure",
			[]*WebTestFiles{
				{ArchiveFile: "a", NamedFiles: map[string]string{"a": "A"}},
				{ArchiveFile: "b", NamedFiles: map[string]string{"b": "B"}},
				{ArchiveFile: "a", NamedFiles: map[string]string{"d": "D"}},
				{ArchiveFile: "c", NamedFiles: map[string]string{"a": "A", "c": "C"}},
			},
			nil,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := normalizeWebTestFiles(tc.input)
			if err != nil {
				if !tc.err {
					t.Fatal(err)
				}
				return
			}
			if tc.err {
				t.Fatalf("Got %v, expected error", result)
			}

			if !reflect.DeepEqual(result, tc.result) {
				t.Fatalf("Got  %v, expected %v", result, tc.result)
			}
		})
	}
}
