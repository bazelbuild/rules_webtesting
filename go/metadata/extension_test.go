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
	"fmt"
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/bazel"
)

var extensionFile = "testdata/extension.json"

func TestFromFileWithExtension(t *testing.T) {
	f, err := bazel.Runfile(extensionFile)
	if err != nil {
		t.Fatal(err)
	}
	extension := &SampleExtension{}
	_, err = FromFile(f, extension)
	if err != nil {
		t.Fatal(err)
	}

	expected := &SampleExtension{"hello", 512}

	if !extension.Equals(expected) {
		t.Errorf("Got %+v, expected %+v", extension, expected)
	}

}

func TestMergeWithExtension(t *testing.T) {
	testCases := []struct {
		name   string
		input1 *Metadata
		input2 *Metadata
		result *Metadata // nil means it should return an error
	}{
		{
			"Second empty",
			&Metadata{Extension: &SampleExtension{"hello", 1024}},
			&Metadata{},
			&Metadata{Extension: &SampleExtension{"hello", 1024}},
		},
		{
			"First empty",
			&Metadata{},
			&Metadata{Extension: &SampleExtension{"hello", 1024}},
			&Metadata{Extension: &SampleExtension{"hello", 1024}},
		},
		{
			"Different types",
			&Metadata{Extension: &SampleExtension{"hello", 1024}},
			&Metadata{Extension: &SampleExtension2{5}},
			nil,
		},
		{
			"Successful merge",
			&Metadata{Extension: &SampleExtension{"hello", 1024}},
			&Metadata{Extension: &SampleExtension{"goodbye", 2048}},
			&Metadata{Extension: &SampleExtension{"hello", 2048}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a, err := Merge(tc.input1, tc.input2)

			if err == nil {
				if tc.result == nil {
					t.Errorf("Got Merge(%+v, %+v) == %+v, expected error", tc.input1, tc.input2, a)
				} else if !Equals(a, tc.result) {
					t.Errorf("Got Merge(%+v, %+v) == %+v, expected %+v", tc.input1, tc.input2, a, tc.result)
				}
			} else {
				if tc.result != nil {
					t.Error(err)
				}
			}
		})
	}
}

type SampleExtension struct {
	ExtensionField1 string
	ExtensionField2 int64
}

func (s *SampleExtension) Merge(other Extension) (Extension, error) {
	if other == nil {
		return s, nil
	}
	o, ok := other.(*SampleExtension)
	if !ok {
		return nil, fmt.Errorf("cannot merge %+v with %+v", s, other)
	}
	return &SampleExtension{
		ExtensionField1: s.ExtensionField1,
		ExtensionField2: o.ExtensionField2,
	}, nil
}

func (s *SampleExtension) Normalize() error {
	s.ExtensionField2 = s.ExtensionField2 / 2
	return nil
}

func (s *SampleExtension) Equals(other Extension) bool {
	o, ok := other.(*SampleExtension)
	return ok && o.ExtensionField1 == s.ExtensionField1 && o.ExtensionField2 == s.ExtensionField2
}

type SampleExtension2 struct {
	X int
}

func (s *SampleExtension2) Merge(other Extension) (Extension, error) {
	return s, nil
}

func (s *SampleExtension2) Normalize() error {
	return nil
}

func (s *SampleExtension2) Equals(other Extension) bool {
	o, ok := other.(*SampleExtension2)
	return ok && o.X == s.X
}
