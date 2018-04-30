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

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// See https://w3c.github.io/webdriver/webdriver-spec.html#capabilities
var w3cSupportedCapabilities = []string{
	"acceptInsecureCerts",
	"browserName",
	"browserVersion",
	"pageLoadStrategy",
	"platformName",
	"proxy",
	"setWindowRect",
	"timeouts",
	"unhandledPromptBehavior",
}

// Capabilities is a WebDriver capabilities object. It is modeled after W3C capabilities, but supports
// use as W3C, JWP, or mixed-mode.
type Capabilities struct {
	AlwaysMatch  map[string]interface{}
	FirstMatch   []map[string]interface{}
	W3CSupported bool
}

// FromNewSessionArgs creates a Capabilities object from the arguments to new session.
// AlwaysMatch will be the result of merging alwaysMatch, requiredCapabilities, and desiredCapabilities.
// Unlike Metadata capabilities merging and MergeOver, this is a shallow merge, and any conflicts will
// result in an error.
// Additionally if any capability in firstMatch conflicts with a capability in alwaysMatch, requiredCapabilities,
// or desiredCapabilities, an error will be returned.
func FromNewSessionArgs(args map[string]interface{}) (*Capabilities, error) {
	always := map[string]interface{}{}
	var first []map[string]interface{}

	w3c, _ := args["capabilities"].(map[string]interface{})

	if w3c != nil {
		if am, ok := w3c["alwaysMatch"].(map[string]interface{}); ok {
			for k, v := range am {
				always[k] = normalize(k, v)
			}
		}
	}

	if required, ok := args["requiredCapabilities"].(map[string]interface{}); ok {
		for k, v := range required {
			nv := normalize(k, v)
			if a, ok := always[k]; ok {
				if !reflect.DeepEqual(a, nv) {
					return nil, fmt.Errorf("alwaysMatch[%q] == %+v, required[%q] == %+v, they must be equal", k, a, k, v)
				}
				continue
			}
			always[k] = nv

		}
	}

	if desired, ok := args["desiredCapabilities"].(map[string]interface{}); ok {
		for k, v := range desired {
			nv := normalize(k, v)
			if a, ok := always[k]; ok {
				if !reflect.DeepEqual(a, nv) {
					return nil, fmt.Errorf("alwaysMatch|required[%q] == %+v, desired[%q] == %+v, they must be equal", k, a, k, v)
				}
				continue
			}
			always[k] = nv
		}
	}

	if w3c != nil {
		fm, _ := w3c["firstMatch"].([]interface{})

		for _, e := range fm {
			fme, ok := e.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("firstMatch entries must be JSON Objects, found %#v", e)
			}
			newFM := map[string]interface{}{}
			for k, v := range fme {
				nv := normalize(k, v)
				if a, ok := always[k]; ok {
					if !reflect.DeepEqual(a, nv) {
						return nil, fmt.Errorf("alwaysMatch|required|desired[%q] == %+v, firstMatch[%q] == %+v, they must be equal", k, a, k, v)
					}
					continue
				}
				newFM[k] = nv

			}
			first = append(first, newFM)
		}
	}

	return &Capabilities{
		AlwaysMatch:  always,
		FirstMatch:   first,
		W3CSupported: w3c != nil,
	}, nil
}

func normalize(key string, value interface{}) interface{} {
	if key != "proxy" {
		return value
	}

	proxy, ok := value.(map[string]interface{})
	if !ok {
		return value
	}

	// If the value if a proxy config, normalize by removing nulls and ensuring proxyType is lower case.
	out := map[string]interface{}{}

	for k, v := range proxy {
		if v == nil {
			continue
		}
		if k == "proxyType" {
			out[k] = strings.ToLower(v.(string))
			continue
		}
		out[k] = v
	}

	return out
}

// MergeOver creates a new Capabilities with AlwaysMatch == (c.AlwaysMatch deeply merged over other),
// FirstMatch == c.FirstMatch, and W3Supported == c.W3CSupported.
func (c *Capabilities) MergeOver(other map[string]interface{}) *Capabilities {
	if c == nil {
		return &Capabilities{
			AlwaysMatch: other,
		}
	}

	if len(other) == 0 {
		return c
	}
	always := map[string]interface{}{}
	first := map[string]interface{}{}

	for k, v := range other {
		if anyContains(c.FirstMatch, k) {
			first[k] = v
		} else {
			always[k] = v
		}
	}

	firstMatch := c.FirstMatch
	if len(first) != 0 {
		firstMatch = nil
		for _, fm := range c.FirstMatch {
			firstMatch = append(firstMatch, Merge(first, fm))
		}
	}

	alwaysMatch := Merge(always, c.AlwaysMatch)

	return &Capabilities{
		AlwaysMatch:  alwaysMatch,
		FirstMatch:   firstMatch,
		W3CSupported: c.W3CSupported,
	}
}

func anyContains(maps []map[string]interface{}, key string) bool {
	for _, m := range maps {
		_, ok := m[key]
		if ok {
			return true
		}
	}

	return false
}

// ToJWP creates a map suitable for use as arguments to a New Session request for JSON Wire Protocol remote ends.
// Since JWP does not support an equivalent to FirstMatch, if FirstMatch contains more than 1 entry
// then this returns an error (if it contains exactly 1 entry, it will be merged over AlwaysMatch).
func (c *Capabilities) ToJWP() (map[string]interface{}, error) {
	if c == nil {
		return map[string]interface{}{
			"desiredCapabilities": map[string]interface{}{},
		}, nil
	}

	if len(c.FirstMatch) > 1 {
		return nil, errors.New("can not convert Capabilities with multiple FirstMatch entries to JWP")
	}

	desired := c.AlwaysMatch
	if len(c.FirstMatch) == 1 {
		desired = Merge(desired, c.FirstMatch[0])
	}

	return map[string]interface{}{
		"desiredCapabilities": desired,
	}, nil
}

// ToW3C creates a map suitable for use as arguments to a New Session request for W3C remote ends.
func (c *Capabilities) ToW3C() map[string]interface{} {
	if c == nil {
		return map[string]interface{}{
			"capabilities": map[string]interface{}{},
		}
	}

	caps := map[string]interface{}{}

	alwaysMatch := w3cCapabilities(c.AlwaysMatch)
	var firstMatch []map[string]interface{}

	for _, fm := range c.FirstMatch {
		firstMatch = append(firstMatch, w3cCapabilities(fm))
	}

	if len(alwaysMatch) != 0 {
		caps["alwaysMatch"] = alwaysMatch
	}

	if len(firstMatch) != 0 {
		caps["firstMatch"] = firstMatch
	}

	return map[string]interface{}{
		"capabilities": caps,
	}
}

// w3cCapabilities remove non-W3C capabilities.
func w3cCapabilities(in map[string]interface{}) map[string]interface{} {
	out := map[string]interface{}{}

	for k, v := range in {
		// extension capabilities
		if strings.Contains(k, ":") {
			out[k] = v
			continue
		}
		for _, a := range w3cSupportedCapabilities {
			if k == a {
				out[k] = v
				break
			}
		}
	}

	return out
}

// ToMixedMode creates a map suitable for use as arguments to a New Session request for arbitrary remote ends.
// If FirstMatch contains more than 1 entry then this returns W3C-only capabilities.
// If W3CSupported is false then this will return JWP-only capabilities.
func (c *Capabilities) ToMixedMode() map[string]interface{} {
	if c == nil {
		return map[string]interface{}{
			"capabilities":        map[string]interface{}{},
			"desiredCapabilities": map[string]interface{}{},
		}
	}

	jwp, err := c.ToJWP()
	if err != nil {
		return c.ToW3C()
	}
	if !c.W3CSupported {
		return jwp
	}

	w3c := c.ToW3C()

	return map[string]interface{}{
		"capabilities":        w3c["capabilities"],
		"desiredCapabilities": jwp["desiredCapabilities"],
	}
}

func (c *Capabilities) Strip(capsToStrip ...string) *Capabilities {
	caps := &Capabilities{
		AlwaysMatch:  make(map[string]interface{}, len(c.AlwaysMatch)),
		FirstMatch:   make([]map[string]interface{}, 0, len(c.FirstMatch)),
		W3CSupported: c.W3CSupported,
	}

	for k, v := range c.AlwaysMatch {
		caps.AlwaysMatch[k] = v
	}

	for _, fm := range c.FirstMatch {
		newFM := make(map[string]interface{}, len(fm))
		for k, v := range fm {
			newFM[k] = v
		}
		caps.FirstMatch = append(caps.FirstMatch, newFM)
	}

	for _, c := range capsToStrip {
		delete(caps.AlwaysMatch, c)
		for _, fm := range caps.FirstMatch {
			delete(fm, c)
		}
	}

	return caps
}

// Merge takes two JSON objects, and merges them.
//
// The resulting object will have all of the keys in the two input objects.
// For each key that is in both objects:
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
		nm[k] = mergeValues(nm[k], v2, k)
	}
	return nm
}

func mergeValues(j1, j2 interface{}, name string) interface{} {
	switch t1 := j1.(type) {
	case map[string]interface{}:
		if t2, ok := j2.(map[string]interface{}); ok {
			return Merge(t1, t2)
		}
	case []interface{}:
		if t2, ok := j2.([]interface{}); ok {
			if name == "args" {
				return mergeArgs(t1, t2)
			}
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

func mergeArgs(m1, m2 []interface{}) []interface{} {
	if len(m1) == 0 {
		return m2
	}
	if len(m2) == 0 {
		return m1
	}

	m2Opts := map[string]bool{}

	for _, a := range m2 {
		if arg, ok := a.(string); ok {
			if strings.HasPrefix(arg, "--") {
				tokens := strings.Split(arg, "=")
				m2Opts[tokens[0]] = true
			}
		}
	}

	nl := make([]interface{}, 0, len(m1)+len(m2))

	for _, a := range m1 {
		if arg, ok := a.(string); ok {
			if strings.HasPrefix(arg, "--") {
				tokens := strings.Split(arg, "=")
				// Skip options from m1 that are redefined in m2
				if m2Opts[tokens[0]] {
					continue
				}
			}
		}
		nl = append(nl, a)
	}

	nl = append(nl, m2...)
	return nl
}

// CanReuseSession returns true if the "google:canReuseSession" is set.
func CanReuseSession(caps *Capabilities) bool {
	reuse, _ := caps.AlwaysMatch["google:canReuseSession"].(bool)
	return reuse
}
