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
	"reflect"
	"testing"
)

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
					"-anotherOption",
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
			name: "args -- redefines",
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
	testCases := []struct{
		name string
		args map[string]interface{}
		want *Capabilities
		wantErr bool
	}{
		{
			name: "empty args",
			args: map[string]interface{}{},
			want: &Capabilities{
				AlwaysMatch: map[string]interface{}{},
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
			want: nil,
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
			want: nil,
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
			want: nil,
			wantErr: true,
		},
		{
			name: "firstMatch, no conflicts",
			args: map[string]interface{}{
				"capabilities": map[string]interface{}{
					"alwaysMatch": map[string]interface{}{
						"key1": "value1",
					},
					"firstMatch": []map[string]interface{}{
						{
							"key2": "value2",
						},
						{
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
					"firstMatch": []map[string]interface{}{
						{
							"key1": "value1",							
							"key2": "value2",
						},
						{
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
					"firstMatch": []map[string]interface{}{
						{
							"key1": "value12",							
							"key2": "value2",
						},
						{
							"key2": "value3",
						},					
					},
				},						
			},
			want: nil,
			wantErr: true,
		},	
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
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
	testCases := []struct{
		name string
		this *Capabilities
		other map[string]interface{}
		want *Capabilities
	}{
		{
			name: "empty",
			this:  &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
			other: map[string]interface{}{},
			want:  &Capabilities{
				AlwaysMatch: map[string]interface{}{},
			},
		},
		{
			name: "no overlap",
			this:  &Capabilities{
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
			want:  &Capabilities{
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
			this:  &Capabilities{
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
			want:  &Capabilities{
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
			this:  &Capabilities{
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
			want:  &Capabilities{
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
			this:  &Capabilities{
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
			want:  &Capabilities{
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
			this:  &Capabilities{
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
			want:  &Capabilities{
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
