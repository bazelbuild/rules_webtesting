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
			if result := Merge(tc.input1, tc.input2); !Equals(tc.result, result) {
				t.Errorf("Got Merge(%+v, %+v) == %+v, expected %+v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}
}

func TestEquals(t *testing.T) {
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
			if result := Equals(tc.input1, tc.input2); result != tc.result {
				t.Errorf("Got Equals(%+v, %+v) == %v, expected %v", tc.input1, tc.input2, result, tc.result)
			}
		})
	}
}
