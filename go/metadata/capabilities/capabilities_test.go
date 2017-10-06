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

package capabilities

import "testing"

func TestMerge(t *testing.T) {
	testCases := []struct {
		name   string
		input1 map[string]interface{}
		input2 map[string]interface{}
		result map[string]interface{}
	}{
		{
			name:   "int,int",
			input1: map[string]interface{}{"v": 1},
			input2: map[string]interface{}{"v": 2},
			result: map[string]interface{}{"v": 2},
		},
		{
			name:   "string,bool",
			input1: map[string]interface{}{"v": "a string"},
			input2: map[string]interface{}{"v": true},
			result: map[string]interface{}{"v": true},
		},
		{
			name:   "int,string",
			input1: map[string]interface{}{"v": 1},
			input2: map[string]interface{}{"v": "a string"},
			result: map[string]interface{}{"v": "a string"},
		},
		{
			name:   "int,slice",
			input1: map[string]interface{}{"v": 1},
			input2: map[string]interface{}{"v": []interface{}{"string 1", 2, true, nil}},
			result: map[string]interface{}{"v": []interface{}{"string 1", 2, true, nil}},
		},
		{
			name:   "slice,slice",
			input1: map[string]interface{}{"v": []interface{}{"string 1", "string 2"}},
			input2: map[string]interface{}{"v": []interface{}{1, 2}},
			result: map[string]interface{}{"v": []interface{}{"string 1", "string 2", 1, 2}},
		},
		{
			name:   "int,map",
			input1: map[string]interface{}{"v": 1},
			input2: map[string]interface{}{"v": map[string]interface{}{
				"1": 1,
				"2": 2,
			}},
			result: map[string]interface{}{"v": map[string]interface{}{
				"1": 1,
				"2": 2,
			}},
		},
		{
			name: "map,map",
			input1: map[string]interface{}{"v": map[string]interface{}{
				"1": 1,
				"2": 2,
			}},
			input2: map[string]interface{}{"v": map[string]interface{}{
				"2": 3,
				"3": 4,
			}},
			result: map[string]interface{}{"v": map[string]interface{}{"1": 1, "2": 3, "3": 4}},
		},
		{
			name: "mixed",
			input1: map[string]interface{}{
				"1": []interface{}{1, 2},
				"2": map[string]interface{}{"a": "an a", "b": "a b", "c": "a c"},
				"3": 3,
			},
			input2: map[string]interface{}{
				"1": []interface{}{"a"},
				"2": map[string]interface{}{"a": "a c", "b": "a d"},
			},
			result: map[string]interface{}{
				"1": []interface{}{1, 2, "a"},
				"2": map[string]interface{}{"a": "a c", "b": "a d", "c": "a c"},
				"3": 3,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := Merge(tc.input1, tc.input2); !JSONEquals(tc.result, result) {
				t.Errorf("Got Merge(%+v, %+v) == %+v, expected %+v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}
}

func TestMergeSpecOntoCaps(t *testing.T) {
	testCases := []struct {
		name   string
		caps   map[string]interface{}
		spec   Spec
		result Spec
	}{
		{
			name:   "all nil",
			caps:   map[string]interface{}{"v": 1},
			spec:   Spec{},
			result: Spec{},
		},
		{
			name:   "empty, not nil",
			caps:   map[string]interface{}{"v": 1},
			spec:   Spec{OSSCaps: map[string]interface{}{}, Always: map[string]interface{}{}},
			result: Spec{OSSCaps: map[string]interface{}{"v": 1}, Always: map[string]interface{}{"v": 1}},
		},
		{
			name:   "nil OSS",
			caps:   map[string]interface{}{"v": 1},
			spec:   Spec{Always: map[string]interface{}{"x": 2}},
			result: Spec{Always: map[string]interface{}{"v": 1, "x": 2}},
		},
		{
			name:   "nil W3C",
			caps:   map[string]interface{}{"v": 1},
			spec:   Spec{OSSCaps: map[string]interface{}{"x": 2}},
			result: Spec{OSSCaps: map[string]interface{}{"v": 1, "x": 2}},
		},
		{
			name: "both present",
			caps: map[string]interface{}{"v": 1},
			spec: Spec{
				OSSCaps: map[string]interface{}{"type": "oss"},
				Always:  map[string]interface{}{"type": "w3c"},
			},
			result: Spec{
				OSSCaps: map[string]interface{}{"v": 1, "type": "oss"},
				Always:  map[string]interface{}{"v": 1, "type": "w3c"},
			},
		},
		{
			name: "no overlaps with firstMatch",
			caps: map[string]interface{}{"three": 3, "four": 999},
			spec: Spec{
				Always: map[string]interface{}{"type": "w3c", "four": 4},
				First: []map[string]interface{}{
					map[string]interface{}{"one": 1},
					map[string]interface{}{"two": 2},
				},
			},
			result: Spec{
				Always: map[string]interface{}{"type": "w3c", "three": 3, "four": 4},
				First: []map[string]interface{}{
					map[string]interface{}{"one": 1},
					map[string]interface{}{"two": 2},
				},
			},
		},
		{
			name: "firstMatch key collision",
			caps: map[string]interface{}{"zero": 0, "one": 999},
			spec: Spec{
				Always: map[string]interface{}{"type": "w3c"},
				First: []map[string]interface{}{
					map[string]interface{}{"one": 1},
					map[string]interface{}{"two": 2},
					map[string]interface{}{"three": 3},
				},
			},
			result: Spec{
				Always: map[string]interface{}{"type": "w3c", "zero": 0},
				First: []map[string]interface{}{
					map[string]interface{}{"one": 1},
					map[string]interface{}{"two": 2, "one": 999},
					map[string]interface{}{"three": 3, "one": 999},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := MergeSpecOntoCaps(tc.caps, tc.spec); !SpecEquals(tc.result, result) {
				t.Errorf("Got MergeSpecOntoCaps(%+v, %+v) == %+v, expected %+v", tc.caps, tc.spec, result, tc.result)
			}
		})
	}
}

func TestSpecEquals(t *testing.T) {
	testCases := []struct {
		name   string
		input1 Spec
		input2 Spec
		result bool
	}{
		{
			"all nil",
			Spec{},
			Spec{},
			true,
		},
		{
			"nil is not the same as empty",
			Spec{},
			Spec{Always: map[string]interface{}{}},
			false,
		},
		{
			"nil vs non-empty",
			Spec{},
			Spec{Always: map[string]interface{}{"v": 1}},
			false,
		},
		{
			"equal with First absent",
			Spec{OSSCaps: map[string]interface{}{"v": 1}, Always: map[string]interface{}{"v": 2}},
			Spec{OSSCaps: map[string]interface{}{"v": 1}, Always: map[string]interface{}{"v": 2}},
			true,
		},
		{
			"equal, all fields present",
			Spec{
				OSSCaps: map[string]interface{}{"v": 1},
				Always:  map[string]interface{}{"v": 2},
				First: []map[string]interface{}{
					map[string]interface{}{"first": 1},
					map[string]interface{}{"second": 2},
				},
			},
			Spec{
				OSSCaps: map[string]interface{}{"v": 1},
				Always:  map[string]interface{}{"v": 2},
				First: []map[string]interface{}{
					map[string]interface{}{"first": 1},
					map[string]interface{}{"second": 2},
				},
			},
			true,
		},
		{
			"one dialect unequal",
			Spec{OSSCaps: map[string]interface{}{"v": 1}, Always: map[string]interface{}{"v": 2}},
			Spec{OSSCaps: map[string]interface{}{"v": 1}, Always: map[string]interface{}{"v": 999}},
			false,
		},
		{
			"dialects swapped",
			Spec{OSSCaps: map[string]interface{}{"v": 1}, Always: map[string]interface{}{"v": 2}},
			Spec{OSSCaps: map[string]interface{}{"v": 2}, Always: map[string]interface{}{"v": 1}},
			false,
		},
		{
			"First order matters",
			Spec{Always: map[string]interface{}{"v": 1}, First: []map[string]interface{}{
				map[string]interface{}{"first": 1}, map[string]interface{}{"second": 2},
			}},
			Spec{Always: map[string]interface{}{"v": 1}, First: []map[string]interface{}{
				map[string]interface{}{"second": 2}, map[string]interface{}{"first": 1},
			}},
			false,
		},
		{
			"First uneven length",
			Spec{Always: map[string]interface{}{"v": 1}, First: []map[string]interface{}{
				map[string]interface{}{"first": 1}, map[string]interface{}{"second": 2},
			}},
			Spec{Always: map[string]interface{}{"v": 1}, First: []map[string]interface{}{
				map[string]interface{}{"first": 1},
			}},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := SpecEquals(tc.input1, tc.input2); result != tc.result {
				t.Errorf("Got SpecEquals(%+v, %+v) == %v, expected %v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}
}

func TestJSONEquals(t *testing.T) {
	testCases := []struct {
		name   string
		input1 map[string]interface{}
		input2 map[string]interface{}
		result bool
	}{
		{
			"empty",
			map[string]interface{}{},
			map[string]interface{}{},
			true,
		},
		{
			"nil map,nil map",
			nil,
			nil,
			true,
		},
		{
			"nil map,empty map",
			nil,
			map[string]interface{}{},
			false,
		},
		{
			"nil map,non-empty",
			nil,
			map[string]interface{}{"v": nil},
			false,
		},
		{
			"int,nil",
			map[string]interface{}{"v": 1},
			map[string]interface{}{"v": nil},
			false,
		},
		{
			"string,nil",
			map[string]interface{}{"v": "hello"},
			map[string]interface{}{"v": nil},
			false,
		},
		{
			"bool,nil",
			map[string]interface{}{"v": true},
			map[string]interface{}{"v": nil},
			false,
		},
		{
			"slice,nil",
			map[string]interface{}{"v": []interface{}{}},
			map[string]interface{}{"v": nil},
			false,
		},
		{
			"map,nil",
			map[string]interface{}{"v": map[string]interface{}{}},
			map[string]interface{}{"v": nil},
			false,
		},
		{
			"map,map equals",
			map[string]interface{}{"v": map[string]interface{}{"a": 1}},
			map[string]interface{}{"v": map[string]interface{}{"a": 1}},
			true,
		},
		{
			"map,map different values",
			map[string]interface{}{"v": map[string]interface{}{"a": 1}},
			map[string]interface{}{"v": map[string]interface{}{"a": "hello"}},
			false,
		},
		{
			"map,map different keys",
			map[string]interface{}{"v": map[string]interface{}{"a": 1}},
			map[string]interface{}{"v": map[string]interface{}{"b": 1}},
			false,
		},
		{
			"map,map different lengths",
			map[string]interface{}{"v": map[string]interface{}{"a": 1}},
			map[string]interface{}{"v": map[string]interface{}{"a": 1, "b": 1}},
			false,
		},
		{
			"slice,slice equals",
			map[string]interface{}{"v": []interface{}{1}},
			map[string]interface{}{"v": []interface{}{1}},
			true,
		},
		{
			"slice,slice different values",
			map[string]interface{}{"v": []interface{}{1}},
			map[string]interface{}{"v": []interface{}{"hello"}},
			false,
		},
		{
			"slice,slice different lengths",
			map[string]interface{}{"v": []interface{}{1}},
			map[string]interface{}{"v": []interface{}{1, 1}},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := JSONEquals(tc.input1, tc.input2); result != tc.result {
				t.Errorf("Got JSONEquals(%+v, %+v) == %v, expected %v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}
}

type kv struct {
	k string
	v interface{}
}

func TestGoogleCap(t *testing.T) {
	testCases := []struct {
		name  string
		caps  Spec
		wants []kv
	}{
		{
			"not found",
			Spec{},
			[]kv{{"capName", nil}, {"otherCapName", nil}},
		},
		{
			"found only in oss caps",
			Spec{OSSCaps: map[string]interface{}{"google.capName": "vvvvvv", "google:otherCap": "xxx"}},
			[]kv{{"capName", "vvvvvv"}, {"otherCap", "xxx"}},
		},
		{
			"found only in w3c caps",
			Spec{Always: map[string]interface{}{"google:capName": "vvvvvv"}},
			[]kv{{"capName", "vvvvvv"}},
		},
		{
			"w3c caps value takes precedence",
			Spec{
				Always:  map[string]interface{}{"google:capName": "vvvvvv"},
				OSSCaps: map[string]interface{}{"google:capName": "xxx"},
			},
			[]kv{{"capName", "vvvvvv"}},
		},
		{
			"requires google prefix",
			Spec{OSSCaps: map[string]interface{}{"capName": "vvvvvv"}},
			[]kv{{"capName", nil}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, want := range tc.wants {
				if got := GoogleCap(tc.caps, want.k); got != want.v {
					t.Errorf("GoogleCap(%v, %q) is %v, want %v", tc.caps, want.k, got, want.v)
				}
				has := want.v != nil
				if got := HasGoogleCap(tc.caps, want.k); got != has {
					t.Errorf("HasGoogleCap(%v, %q) is %v, want %v", tc.caps, want.k, got, has)
				}
			}
		})
	}
}

func TestSetGoogleCap(t *testing.T) {
	testCases := []struct {
		name string
		caps Spec
		k    string
		v    interface{}
		want Spec
	}{
		{
			"set cap",
			Spec{
				OSSCaps: map[string]interface{}{},
				Always:  map[string]interface{}{},
			},
			"capName", "vvvvvv",
			Spec{
				OSSCaps: map[string]interface{}{"google.capName": "vvvvvv", "google:capName": "vvvvvv"},
				Always:  map[string]interface{}{"google:capName": "vvvvvv"},
			},
		},
		{
			"nil OSS caps are ignored",
			Spec{
				Always: map[string]interface{}{},
			},
			"capName", "vvvvvv",
			Spec{
				Always: map[string]interface{}{"google:capName": "vvvvvv"},
			},
		},
		{
			"nil W3C caps are ignored",
			Spec{
				OSSCaps: map[string]interface{}{},
			},
			"capName", "vvvvvv",
			Spec{
				OSSCaps: map[string]interface{}{"google.capName": "vvvvvv", "google:capName": "vvvvvv"},
			},
		},
		{
			"overwrite cap",
			Spec{
				OSSCaps: map[string]interface{}{"google:capName": "xyz"},
				Always:  map[string]interface{}{"google:capName": "xyz"},
			},
			"capName", "vvvvvv",
			Spec{
				OSSCaps: map[string]interface{}{"google:capName": "vvvvvv", "google.capName": "vvvvvv"},
				Always:  map[string]interface{}{"google:capName": "vvvvvv"},
			},
		},
		{
			"overwrite google-prefixed cap only",
			Spec{
				OSSCaps: map[string]interface{}{"capName": "xyz"},
				Always:  map[string]interface{}{"capName": "xyz"},
			},
			"capName", "vvvvvv",
			Spec{
				OSSCaps: map[string]interface{}{"google:capName": "vvvvvv", "google.capName": "vvvvvv", "capName": "xyz"},
				Always:  map[string]interface{}{"google:capName": "vvvvvv", "capName": "xyz"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SetGoogleCap(tc.caps, tc.k, tc.v)
			if !SpecEquals(tc.caps, tc.want) {
				t.Errorf("got %v, want %v", tc.caps, tc.want)
			}
		})
	}
}
