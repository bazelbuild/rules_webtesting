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

// Package capabilities performs operations on maps representing WebDriver capabilities.
package capabilities

// Merge takes two JSON objects, and merges them.
//
// The resulting object will have all of the keys in the two input objects.
// For each key that is in both obejcts:
//   - if both objects have objects for values, then the result object will have
//     a value resulting from recursively calling Merge.
//   - if both objects have lists for values, then the result object will have
//     a value resulting from concatenating the two lists.
//   - Otherwise the result object will have the value from the second object.
func Merge(m1, m2 map[string]interface{}) map[string]interface{} {
	if len(m1) == 0 {
		return m2
	}
	if len(m2) == 0 {
		return m1
	}
	nm := make(map[string]interface{})
	for k, v := range m1 {
		nm[k] = v
	}
	for k, v2 := range m2 {
		nm[k] = mergeValues(nm[k], v2)
	}
	return nm
}

func mergeValues(j1, j2 interface{}) interface{} {
	switch t1 := j1.(type) {
	case map[string]interface{}:
		if t2, ok := j2.(map[string]interface{}); ok {
			return Merge(t1, t2)
		}
	case []interface{}:
		if t2, ok := j2.([]interface{}); ok {
			return mergeLists(t1, t2)
		}
	}
	return j2
}

func mergeLists(m1, m2 []interface{}) []interface{} {
	if len(m1) == 0 {
		return m2
	}
	if len(m2) == 0 {
		return m1
	}
	nl := make([]interface{}, 0, len(m1)+len(m2))
	nl = append(nl, m1...)
	nl = append(nl, m2...)
	return nl
}

// Equals compares two JSON objects, and returns true iff they are the same.
func Equals(e, v map[string]interface{}) bool {
	if len(e) != len(v) {
		return false
	}
	for ek, ev := range e {
		if vv, ok := v[ek]; !ok || !valueEquals(ev, vv) {
			return false
		}
	}
	return true
}

func valueEquals(e, v interface{}) bool {
	switch te := e.(type) {
	case []interface{}:
		tv, ok := v.([]interface{})
		return ok && sliceEquals(te, tv)
	case map[string]interface{}:
		tv, ok := v.(map[string]interface{})
		return ok && Equals(te, tv)
	default:
		return e == v
	}
}

func sliceEquals(e, v []interface{}) bool {
	if len(e) != len(v) {
		return false
	}
	for i := 0; i < len(e); i++ {
		if !valueEquals(e[i], v[i]) {
			return false
		}
	}
	return true
}
