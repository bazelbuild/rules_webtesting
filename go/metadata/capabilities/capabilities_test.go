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

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDenormalizeW3C(t *testing.T) {
	testCases := []struct {
		name string
		args map[string]interface{}
		want map[string]interface{}
	}{
		{
			name: "nil args",
			args: nil,
			want: map[string]interface{}{},
		},
		{
			name: "W3C args preserved",
			args: map[string]interface{}{
				"browserName":    "chrome",
				"browserVersion": "1",
			},
			want: map[string]interface{}{
				"browserName":    "chrome",
				"browserVersion": "1",
			},
		},
		{
			name: "non-W3C args removed",
			args: map[string]interface{}{
				"notW3C": "foo",
			},
			want: map[string]interface{}{},
		},
		{
			name: "extension capabilities preserved",
			args: map[string]interface{}{
				"goog:notW3C": "foo",
			},
			want: map[string]interface{}{
				"goog:notW3C": "foo",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := denormalizeW3C(tc.args)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}

func TestDenormalizeJWP(t *testing.T) {
	testCases := []struct {
		name string
		args map[string]interface{}
		want map[string]interface{}
	}{
		{
			name: "nil args",
			args: nil,
			want: map[string]interface{}{},
		},
		{
			name: "duplicates goog:chromeOptions into chromeOptions",
			args: map[string]interface{}{
				"goog:chromeOptions": map[string]interface{}{
					"key": "value",
				},
			},
			want: map[string]interface{}{
				"goog:chromeOptions": map[string]interface{}{
					"key": "value",
				},
				"chromeOptions": map[string]interface{}{
					"key": "value",
				},
			},
		},
		{
			name: "duplicates goog:loggingPrefs into loggingPrefs",
			args: map[string]interface{}{
				"goog:loggingPrefs": map[string]interface{}{
					"key": "value",
				},
			},
			want: map[string]interface{}{
				"goog:loggingPrefs": map[string]interface{}{
					"key": "value",
				},
				"loggingPrefs": map[string]interface{}{
					"key": "value",
				},
			},
		},
		{
			name: "converts noProxy from []interface{} into string",
			args: map[string]interface{}{
				"proxy": map[string]interface{}{
					"noProxy": []interface{}{
						"127.0.0.1",
						"localhost",
					},
				},
			},
			want: map[string]interface{}{
				"proxy": map[string]interface{}{
					"noProxy": "127.0.0.1,localhost",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := denormalizeJWP(tc.args)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	testCases := []struct {
		name    string
		args    map[string]interface{}
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "empty args",
			args:    map[string]interface{}{},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "normalizes chromeOptions",
			args: map[string]interface{}{
				"goog:chromeOptions": map[string]interface{}{
					"key1": "value1",
				},
				"chromeOptions": map[string]interface{}{
					"key2": "value2",
				},
			},
			want: map[string]interface{}{
				"goog:chromeOptions": map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			wantErr: false,
		},
		{
			name: "err if chromeOptions not map[string]{}",
			args: map[string]interface{}{
				"chromeOptions": "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "err if goog:chromeOptions not map[string]{}",
			args: map[string]interface{}{
				"goog:chromeOptions": "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "err if goog:chromeOptions and chromeOptions same key different value",
			args: map[string]interface{}{
				"goog:chromeOptions": map[string]interface{}{
					"key": "value1",
				},
				"chromeOptions": map[string]interface{}{
					"key": "value2",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "normalizes loggingPrefs",
			args: map[string]interface{}{
				"goog:loggingPrefs": map[string]interface{}{
					"key1": "value1",
				},
				"loggingPrefs": map[string]interface{}{
					"key2": "value2",
				},
			},
			want: map[string]interface{}{
				"goog:loggingPrefs": map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			wantErr: false,
		},
		{
			name: "normalizes proxy",
			args: map[string]interface{}{
				"proxy": map[string]interface{}{
					"proxyType": "DIRECT",
					"noProxy":   "[::1]:8080,localhost",
				},
			},
			want: map[string]interface{}{
				"proxy": map[string]interface{}{
					"proxyType": "direct",
					"noProxy": []interface{}{
						"[::1]:8080",
						"localhost",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "err if proxy not map[string]{}",
			args: map[string]interface{}{
				"proxy": "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "err if proxy[\"proxyType\"] not string",
			args: map[string]interface{}{
				"proxy": map[string]interface{}{
					"proxyType": map[string]interface{}{},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "err if proxy[\"noProxy\"] not string or []interface{}",
			args: map[string]interface{}{
				"proxy": map[string]interface{}{
					"noProxy": 0,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "strips underscore-prefixed caps",
			args: map[string]interface{}{
				"_comment": "stripped",
				"comment":  "not stripped",
				"nestedInMap": map[string]interface{}{
					"_another": "stripped",
					"another":  "not stripped",
				},
				"nestedInSlice": []interface{}{
					map[string]interface{}{
						"_onemore": "stripped",
						"onemore":  "not stripped",
					},
					[]interface{}{
						map[string]interface{}{
							"_lastone": "stripped",
							"lastone":  "not stripped",
						},
					},
				},
			},
			want: map[string]interface{}{
				"comment": "not stripped",
				"nestedInMap": map[string]interface{}{
					"another": "not stripped",
				},
				"nestedInSlice": []interface{}{
					map[string]interface{}{
						"onemore": "not stripped",
					},
					[]interface{}{
						map[string]interface{}{
							"lastone": "not stripped",
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := normalize(tc.args)

			if err != nil || tc.wantErr {
				if (err != nil) != tc.wantErr {
					t.Fatalf("got err %v, wantErr==%t", err, tc.wantErr)
				}
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}

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
		{
			name: "args -- no redefines",
			input1: map[string]interface{}{
				"args": []interface{}{
					"an option",
					"--anOption",
					"--anOption=true",
					"-anotherOption",
					map[string]interface{}{
						"some": "map",
					},
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"an option",
					"anOption",
					"-anOption=true",
					"-anotherOption",
					map[string]interface{}{
						"some": "map",
					},
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"an option",
					"--anOption",
					"--anOption=true",
					map[string]interface{}{
						"some": "map",
					},
					"an option",
					"anOption",
					"-anOption=true",
					"-anotherOption",
					map[string]interface{}{
						"some": "map",
					},
				},
			},
		},
		{
			name: "overrides '--' prefixed args",
			input1: map[string]interface{}{
				"args": []interface{}{
					"an option",
					"--anOption",
					"--anOption=true",
					"--optionToLeave=this",
					map[string]interface{}{
						"some": "map",
					},
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"an option",
					"--anOption=false",
					"--anotherOption",
					"-optionToLeave=that",
					map[string]interface{}{
						"some": "map",
					},
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"an option",
					"--optionToLeave=this",
					map[string]interface{}{
						"some": "map",
					},
					"an option",
					"--anOption=false",
					"--anotherOption",
					"-optionToLeave=that",
					map[string]interface{}{
						"some": "map",
					},
				},
			},
		},
		{
			name: "args -- removals",
			input1: map[string]interface{}{
				"args": []interface{}{
					"an option",
					"--anOption",
					"--anOption=true",
					"--optionToLeave=this",
					map[string]interface{}{
						"some": "map",
					},
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"an option",
					"REMOVE:--anOption",
					"REMOVE:--noSuchOption",
					"REMOVE:an option",
					"-optionToLeave=that",
					map[string]interface{}{
						"some": "map",
					},
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"an option",
					"--optionToLeave=this",
					map[string]interface{}{
						"some": "map",
					},
					"an option",
					"REMOVE:an option",
					"-optionToLeave=that",
					map[string]interface{}{
						"some": "map",
					},
				},
			},
		},
		{
			name: "overrides '-' prefixed args",
			input1: map[string]interface{}{
				"args": []interface{}{
					"option",
					"--anOption",
					"-width=1024",
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"option",
					"--anOption",
					"-width=2048",
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"option",
					"option",
					"--anOption",
					"-width=2048",
				},
			},
		},
		{
			name: "chromeOptions to goog:chromeOptions",
			input1: map[string]interface{}{
				"chromeOptions": []interface{}{
					"foo",
				},
			},
			input2: map[string]interface{}{
				"chromeOptions": []interface{}{
					"bar",
				},
			},
			result: map[string]interface{}{
				"goog:chromeOptions": []interface{}{
					"foo",
					"bar",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := Merge(tc.input1, tc.input2); !reflect.DeepEqual(tc.result, result) {
				t.Errorf("Got Merge(%+v, %+v) == %+v, expected %+v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}
}

func TestMergeFeatures(t *testing.T) {
	testCases := []struct {
		name   string
		input1 map[string]interface{}
		input2 map[string]interface{}
		result map[string]interface{}
	}{
		{
			name: "merge features that only exist on one side",
			input1: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f2",
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"--disable-blink-features=f3,f5",
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f2",
					"--disable-blink-features=f3,f5",
				},
			},
		},
		{
			name: "merge only --enable-feature and disable-blink-feature",
			input1: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f2",
					"--disable-blink-features=f3,f5",
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f4",
					"--disable-blink-features=f3,f5",
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f2,f4",
					"--disable-blink-features=f3,f5",
				},
			},
		},
		{
			name: "merge --enable/disable-feature with a left hand side feature stated both in enable and disable, result should keep both if none of them is covered by right side args",
			input1: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f3",
					"--disable-features=f3,f4",
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f4",
					"--disable-features=f1,f5",
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f3,f4",
					"--disable-features=f1,f3,f5",
				},
			},
		},
		{
			name: "merge --enable/disable-feature with a right hand side feature stated both in enable and disable, result should keep both",
			input1: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f2,f4",
					"--disable-features=f3,f4",
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f4",
					"--disable-features=f4,f5",
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f2,f4",
					"--disable-features=f3,f4,f5",
				},
			},
		},
		{
			name: "merge both --enable/disable-feature and --enable/disable-blink-feature",
			input1: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f2",
					"--disable-features=f3,f4",
					"--enable-blink-features=f1",
					"--disable-blink-features=f3",
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f4",
					"--disable-features=f1,f5",
					"--enable-blink-features=f3",
					"--disable-blink-features=f1",
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f2,f4",
					"--disable-features=f1,f3,f5",
					"--enable-blink-features=f3",
					"--disable-blink-features=f1",
				},
			},
		},
		{
			name: "merge only --enable/disable-feature with a feature stated twice in one place",
			input1: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f1,f2",
					"--disable-features=f3",
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f4",
					"--disable-features=f3",
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f2,f4",
					"--disable-features=f3",
				},
			},
		},
		{
			name: "merge only --enable-feature with a feature has empty value",
			input1: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=",
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f4,f1",
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f4",
				},
			},
		},
		{
			name: "merge only --enable-feature with a feature has empty value on the right",
			input1: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=f1,f4",
					"--disable-features=f1,f3",
				},
			},
			input2: map[string]interface{}{
				"args": []interface{}{
					"--enable-features=",
					"--disable-features=f2,f5",
				},
			},
			result: map[string]interface{}{
				"args": []interface{}{
					"--disable-features=f1,f2,f3,f5",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := Merge(tc.input1, tc.input2); !reflect.DeepEqual(tc.result, result) {
				t.Errorf("Got Merge(%+v, %+v) == %+v, expected %+v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}
}

func TestFromNewSessionArgs(t *testing.T) {
	testCases := []struct {
		name    string
		args    map[string]interface{}
		want    *Capabilities
		wantErr bool
	}{
		{
			name: "empty args",
			args: map[string]interface{}{},
			want: &Capabilities{
				AlwaysMatch:  map[string]interface{}{},
				W3CSupported: false,
			},
			wantErr: false,
		},
		{
			name: "alwaysMatch",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value1",
					},
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				W3CSupported: true,
			},
			wantErr: false,
		},
		{
			name: "requiredCapabilities",
			args: map[string]interface{}{
				"requiredCapabilities": map[string]interface{}{
					"key1": "value1",
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				W3CSupported: false,
			},
			wantErr: false,
		},
		{
			name: "desiredCapabilities",
			args: map[string]interface{}{
				"desiredCapabilities": map[string]interface{}{
					"key1": "value1",
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				W3CSupported: false,
			},
			wantErr: false,
		},
		{
			name: "all three",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value1",
					},
				},
				"desiredCapabilities": map[string]interface{}{
					"key2": "value2",
				},
				"requiredCapabilities": map[string]interface{}{
					"key3": "value3",
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				W3CSupported: true,
			},
			wantErr: false,
		},
		{
			name: "all three normalized",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"chromeOptions": map[string]interface{}{
							"key": "value",
						},
					},
				},
				"desiredCapabilities": map[string]interface{}{
					"chromeOptions": map[string]interface{}{
						"key": "value",
					},
				},
				"requiredCapabilities": map[string]interface{}{
					"chromeOptions": map[string]interface{}{
						"key": "value",
					},
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"goog:chromeOptions": map[string]interface{}{
						"key": "value",
					},
				},
				W3CSupported: true,
			},
			wantErr: false,
		},
		{
			name: "all three, same value ok",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value1",
					},
				},
				"desiredCapabilities": map[string]interface{}{
					"key1": "value1",
				},
				"requiredCapabilities": map[string]interface{}{
					"key1": "value1",
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				W3CSupported: true,
			},
			wantErr: false,
		},
		{
			name: "always, required != desired",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value1",
					},
				},
				"desiredCapabilities": map[string]interface{}{
					"key1": "value12",
				},
				"requiredCapabilities": map[string]interface{}{
					"key1": "value1",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "always, desired != required",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value1",
					},
				},
				"desiredCapabilities": map[string]interface{}{
					"key1": "value1",
				},
				"requiredCapabilities": map[string]interface{}{
					"key1": "value12",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "always != desired. required",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value12",
					},
				},
				"desiredCapabilities": map[string]interface{}{
					"key1": "value1",
				},
				"requiredCapabilities": map[string]interface{}{
					"key1": "value1",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "firstMatch, no conflicts",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value1",
					},
					"firstMatch": []interface{}{
						map[string]interface{}{
							"key2": "value2",
						},
						map[string]interface{}{
							"key2": "value3",
						},
					},
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key2": "value3",
					},
				},
				W3CSupported: true,
			},
			wantErr: false,
		},
		{
			name: "firstMatch, same value as alwaysMatch",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value1",
					},
					"firstMatch": []interface{}{
						map[string]interface{}{
							"key1": "value1",
							"key2": "value2",
						},
						map[string]interface{}{
							"key2": "value3",
						},
					},
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key2": "value3",
					},
				},
				W3CSupported: true,
			},
			wantErr: false,
		},
		{
			name: "firstMatch, different value than alwaysMatch",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value1",
					},
					"firstMatch": []interface{}{
						map[string]interface{}{
							"key1": "value12",
							"key2": "value2",
						},
						map[string]interface{}{
							"key2": "value3",
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "firstMatch, deduped",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"firstMatch": []interface{}{
						map[string]interface{}{
							"key1": "value1",
						},
						map[string]interface{}{
							"key2": "value2",
						},
						map[string]interface{}{
							"key2": "value2",
						},
					},
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
				FirstMatch: []map[string]interface{}{
					{
						"key1": "value1",
					},
					{
						"key2": "value2",
					},
				},
				W3CSupported: true,
			},
			wantErr: false,
		},
		{
			name: "single deduped firstMatch entry becomes alwaysMatch",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"firstMatch": []interface{}{
						map[string]interface{}{
							"key1": "value1",
						},
						map[string]interface{}{
							"key1": "value1",
						},
					},
				},
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch:   []map[string]interface{}(nil),
				W3CSupported: true,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FromNewSessionArgs(tc.args)

			if err != nil || tc.wantErr {
				if (err != nil) != tc.wantErr {
					t.Fatalf("got err %v, wantErr==%t", err, tc.wantErr)
				}
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}

func TestMergeOver(t *testing.T) {
	testCases := []struct {
		name  string
		this  *Capabilities
		other map[string]interface{}
		want  *Capabilities
	}{
		{
			name: "empty",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
			other: map[string]interface{}{},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
		},
		{
			name: "no overlap",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
			other: map[string]interface{}{
				"key4": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
					"key4": "value4",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
		},
		{
			name: "overlaps always",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
			other: map[string]interface{}{
				"key1": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
		},
		{
			name: "overlaps first[0]",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
			other: map[string]interface{}{
				"key2": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key2": "value4",
						"key3": "value3",
					},
				},
			},
		},
		{
			name: "overlaps first[1]",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
			other: map[string]interface{}{
				"key3": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
						"key3": "value4",
					},
					{
						"key3": "value3",
					},
				},
			},
		},
		{
			name: "overlap and non-overlap",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
			other: map[string]interface{}{
				"key1": "value11",
				"key2": "value22",
				"key3": "value33",
				"key4": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
					"key4": "value4",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
						"key3": "value33",
					},
					{
						"key2": "value22",
						"key3": "value3",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.this.MergeOver(tc.other)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}

func TestStripAllPrefixedExcept(t *testing.T) {
	testCases := []struct {
		name   string
		in     *Capabilities
		exempt []string
		out    *Capabilities
	}{
		{
			name: "strips non-exempt AlwaysMatch prefixed capabilities",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"nonexempt:cap": "val",
				},
			},
			exempt: []string{""},
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
		},
		{
			name: "does not strip exempt AlwaysMatch prefixed capabilities",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"foo:cap": "val",
				},
			},
			exempt: []string{"foo"},
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"foo:cap": "val",
				},
			},
		},
		{
			name: "strips non-exempt FirstMatch prefixed capabilities",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
				FirstMatch: []map[string]interface{}{
					{
						"foo:cap": "val",
					},
				},
			},
			exempt: []string{""},
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
				FirstMatch:  []map[string]interface{}{{}},
			},
		},
		{
			name: "does not strip exempt FirstMatch prefixed capabilities",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
				FirstMatch: []map[string]interface{}{
					{
						"foo:cap": "val",
					},
				},
			},
			exempt: []string{"foo"},
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
				FirstMatch: []map[string]interface{}{
					{
						"foo:cap": "val",
					},
				},
			},
		},
		{
			name: "does not strip FirstMatch un-prefixed capabilities",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
				FirstMatch: []map[string]interface{}{
					{
						"cap": "val",
					},
				},
			},
			exempt: []string{""},
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
				FirstMatch: []map[string]interface{}{
					{
						"cap": "val",
					},
				},
			},
		},
		{
			name: "does not strip AlwaysMatch un-prefixed capabilities",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"cap": "val",
				},
			},
			exempt: []string{""},
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"cap": "val",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := tc.in.StripAllPrefixedExcept(tc.exempt...)

			if !reflect.DeepEqual(out, tc.out) {
				t.Fatalf("got %#v, want %#v", out, tc.out)
			}
		})
	}
}

func TestResolve(t *testing.T) {
	testCases := []struct {
		name     string
		in       *Capabilities
		resolver Resolver
		out      *Capabilities
		err      bool
	}{
		{
			name: "empty capabilities",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
			resolver: func(prefix, name string) (string, error) {
				return "", fmt.Errorf("resolver called with %s:%s", prefix, name)
			},
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
			err: false,
		},
		{
			name: "NoOP resolver",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"abc": "%p:n%",
				},
				FirstMatch: []map[string]interface{}{
					{
						"xyz": "%n:p%",
					},
				},
			},
			resolver: NoOPResolver,
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"abc": "%p:n%",
				},
				FirstMatch: []map[string]interface{}{
					{
						"xyz": "%n:p%",
					},
				},
			},
			err: false,
		},
		{
			name: "MapResolver",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"abc": "%p:n%",
				},
				FirstMatch: []map[string]interface{}{
					{
						"xyz": "%n:p%",
					},
				},
			},
			resolver: MapResolver("p", map[string]string{"n": "some value"}),
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"abc": "some value",
				},
				FirstMatch: []map[string]interface{}{
					{
						"xyz": "%n:p%",
					},
				},
			},
			err: false,
		},
		{
			name: "complex input",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"abc": []interface{}{"%p:n%"},
				},
				FirstMatch: []map[string]interface{}{
					{
						"xyz": map[string]interface{}{"zyx": "%n:p%=%p:n%"},
					},
				},
			},
			resolver: func(prefix, name string) (string, error) {
				if prefix == "p" && name == "n" {
					return "some-value", nil
				}
				if prefix == "n" && name == "p" {
					return "value-some", nil
				}
				return "", fmt.Errorf("unknown %s:%s", prefix, name)
			},
			out: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"abc": []interface{}{"some-value"},
				},
				FirstMatch: []map[string]interface{}{
					{
						"xyz": map[string]interface{}{"zyx": "value-some=some-value"},
					},
				},
			},
			err: false,
		},
		{
			name: "resolver returns error",
			in: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"abc": []interface{}{"%p:n%"},
				},
				FirstMatch: []map[string]interface{}{
					{
						"xyz": map[string]interface{}{"zyx": "%n:p%=%p:n%"},
					},
					{
						"bad": "%x:y%",
					},
				},
			},
			resolver: func(prefix, name string) (string, error) {
				if prefix == "p" && name == "n" {
					return "some-value", nil
				}
				if prefix == "n" && name == "p" {
					return "value-some", nil
				}
				return "", fmt.Errorf("unknown %s:%s", prefix, name)
			},
			out: nil,
			err: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := tc.in.Resolve(tc.resolver)

			if err != nil {
				if !tc.err {
					t.Fatal(err)
				}
				return
			}
			if tc.err {
				t.Fatalf("got nil err, want err")
			}

			if !reflect.DeepEqual(out, tc.out) {
				t.Fatalf("got %#v, want %#v", out, tc.out)
			}
		})
	}
}

func TestMergeUnder(t *testing.T) {
	testCases := []struct {
		name  string
		this  *Capabilities
		other map[string]interface{}
		want  *Capabilities
	}{
		{
			name:  "nil, nil",
			this:  nil,
			other: nil,
			want:  nil,
		},
		{
			name: "empty, nil",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
			other: nil,
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
		},
		{
			name:  "nil, empty",
			this:  nil,
			other: map[string]interface{}{},
			want:  nil,
		},
		{
			name: "empty, empty",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
			other: map[string]interface{}{},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
		},
		{
			name: "no overlap",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
			other: map[string]interface{}{
				"key4": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
					"key4": "value4",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
		},
		{
			name: "overlaps always",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
			other: map[string]interface{}{
				"key1": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value4",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
		},
		{
			name: "overlaps first[0]",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
			other: map[string]interface{}{
				"key2": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
					"key2": "value4",
				},
				FirstMatch: []map[string]interface{}{
					{},
					{
						"key3": "value3",
					},
				},
			},
		},
		{
			name: "overlaps first[1]",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{
						"key3": "value3",
					},
				},
			},
			other: map[string]interface{}{
				"key3": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
					"key3": "value4",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
					},
					{},
				},
			},
		},
		{
			name: "overlap and non-overlap",
			this: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key1": "value1",
					"key5": "value5",
				},
				FirstMatch: []map[string]interface{}{
					{
						"key2": "value2",
						"key6": "value6",
					},
					{
						"key3": "value3",
						"key6": "value6",
					},
				},
			},
			other: map[string]interface{}{
				"key1": "value11",
				"key2": "value22",
				"key3": "value33",
				"key4": "value4",
			},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{
					"key4": "value4",
					"key1": "value11",
					"key2": "value22",
					"key3": "value33",
					"key5": "value5",
					"key6": "value6",
				},
				FirstMatch: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.this.MergeUnder(tc.other)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}

