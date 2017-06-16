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

// Spec is a specification of capabilities, such as that included in the New
// Session request. Specs may include slightly different capabilities for
// different WebDriver protocol dialects.
//
// A Spec contains an optional set of OSS capabilities, and a list of zero or
// more sets of W3C capabilities. The W3C capabilities are expressed in two
// parts: the "always match" object that contains key-value pairs shared by
// every set of capabilites, and a list of "first match" objects that contain
// just the keys that differ.
//
// Any field may be nil if the request does not contain any capabilities for
// a dialect, i.e., the requestor does not support that dialect.
type Spec struct {
	OSSCaps map[string]interface{}
	// W3C spec capabilities
	Always map[string]interface{}
	First  []map[string]interface{}
}

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

// MergeSpecOntoCaps merges a set of capabilities and a capabilities spec. For
// each capabilities set C in spec, C is merged on top of caps. That is,
// capability values in the spec take precedence.
func MergeSpecOntoCaps(caps map[string]interface{}, spec Spec) Spec {
	newSpec := Spec{}
	if spec.OSSCaps != nil {
		newSpec.OSSCaps = Merge(caps, spec.OSSCaps)
	}
	if spec.Always == nil {
		return newSpec
	}

	// No key is allowed to appear in both alwaysMatch and firstMatch. Any key in
	// caps that is present in ANY of the firstMatch maps must be merged into each
	// firstMatch map, rather than into the alwaysMatch map.
	firstMatchKeys := map[string]bool{}
	for _, f := range spec.First {
		for fk := range f {
			firstMatchKeys[fk] = true
		}
	}
	alwaysMerge := map[string]interface{}{}
	firstMerge := map[string]interface{}{}
	for k, v := range caps {
		if firstMatchKeys[k] {
			firstMerge[k] = v
		} else {
			alwaysMerge[k] = v
		}
	}
	newSpec.Always = Merge(alwaysMerge, spec.Always)
	for _, f := range spec.First {
		newSpec.First = append(newSpec.First, Merge(firstMerge, f))
	}
	return newSpec
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

// SpecEquals compares two Specs, and returns true iff all the component JSON
// objects are the same.
func SpecEquals(e, v Spec) bool {
	return JSONEquals(e.OSSCaps, v.OSSCaps) && JSONEquals(e.Always, v.Always) && firstsEqual(e.First, v.First)
}

func firstsEqual(e, v []map[string]interface{}) bool {
	if len(e) != len(v) {
		return false
	}
	for i, em := range e {
		vm := v[i]
		if !JSONEquals(em, vm) {
			return false
		}
	}
	return true
}

// JSONEquals compares two JSON objects, and returns true iff they are the same.
func JSONEquals(e, v map[string]interface{}) bool {
	if e == nil || v == nil {
		return e == nil && v == nil
	}
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
		return ok && JSONEquals(te, tv)
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

// GoogleCap returns the value of a Google capability from the given
// Spec. Google capabilities are currently only extracted from the OSS specs,
// but this will eventually change to recognize vendor-prefixed capabilities in
// the W3C spec as well.
func GoogleCap(caps Spec, name string) interface{} {
	return caps.OSSCaps["google."+name]
}

// HasGoogleCap returns whether the named Google capability is present.
func HasGoogleCap(caps Spec, name string) bool {
	_, ok := caps.OSSCaps["google."+name]
	return ok
}

// SetGoogleCap mutates the given Spec by setting a Google capability.
func SetGoogleCap(caps Spec, name string, value interface{}) {
	caps.OSSCaps["google."+name] = value
}

// CanReuseSession returns whether the Google capability "canReuseSession" is
// set.
func CanReuseSession(caps Spec) bool {
	// default value is false
	v, _ := GoogleCap(caps, "canReuseSession").(bool)
	return v
}
