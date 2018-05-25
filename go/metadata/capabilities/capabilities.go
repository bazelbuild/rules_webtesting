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
	"regexp"
	"strings"
)

// See https://w3c.github.io/webdriver/webdriver-spec.html#capabilities
var w3cSupportedCapabilities = map[string]bool{
	"acceptInsecureCerts": true,
	"browserName": true,
	"browserVersion": true,
	"pageLoadStrategy": true,
	"platformName": true,
	"proxy": true,
	"setWindowRect": true,
	"timeouts": true,
	"unhandledPromptBehavior": true,
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

func isW3CCapName(name string) bool {
	return w3cSupportedCapabilities[name] || strings.Contains(name, ":")
}

// W3CCompatible returns true iff all capability names in these capabilities are
// W3C compatible.
func (c *Capabilities) W3CCompatible() bool {
	if c == nil {
		return true
	}
	for k := range c.AlwaysMatch {
		if !isW3CCapName(k) {
			return false
		}
	}

	for _, fm := range c.FirstMatch {
		for k := range fm {
			if !isW3CCapName(k) {
				return false
			}
		}
	}

	return true
}

// JWPCompatible returns true iff there is at most 1 entry in FirstMatch.
func (c *Capabilities) JWPCompatible() bool {
	if c == nil {
		return true
	}
	return len(c.FirstMatch) < 2
}

// ToJWP creates a map suitable for use as arguments to a New Session request for JSON Wire Protocol remote ends.
// Since JWP does not support an equivalent to FirstMatch, if FirstMatch contains more than 1 entry
// then this returns an error (if it contains exactly 1 entry, it will be merged over AlwaysMatch).
func (c *Capabilities) ToJWP() (map[string]interface{}, error) {
	if !c.JWPCompatible {
		return nil, errors.New("capabilities includes multiple FirstMatch entries")
	}

	if c == nil {
		return map[string]interface{}{
			"desiredCapabilities": map[string]interface{}{},
		}, nil
	}

	desired := c.AlwaysMatch
	for _, fm := range c.FirstMatch {
		desired = Merge(desired, fm)
	}

	return map[string]interface{}{
		"desiredCapabilities": desired,
	}, nil
}

// ToW3C creates a map suitable for use as arguments to a New Session request for W3C remote ends.
func (c *Capabilities) ToW3C() (map[string]interface{}, error) {
	if !c.W3CCompatible() {
		return nil, errors.New("capabilities includes non-W3C compatible keys")
	}

	if c == nil {
		return map[string]interface{}{
			"capabilities": map[string]interface{}{},
		}
	}

	caps := map[string]interface{}{}

	if len(c.AlwaysMatch) != 0 {
		caps["alwaysMatch"] = c.AlwaysMatch
	}

	if len(c.FirstMatch) != 0 {
		caps["firstMatch"] = c.FirstMatch
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

// Strip returns a copy of c with all top-level capabilities capsToStrip and with nil values removed.
func (c *Capabilities) Strip(capsToStrip ...string) *Capabilities {
	am := map[string]interface{}{}
	var fms []map[string]interface{}

	for k, v := range c.AlwaysMatch {
		if v != nil {
			am[k] = v
		}
	}

	for _, fm := range c.FirstMatch {
		newFM := map[string]interface{}{}
		for k, v := range fm {
			if v != nil {
				newFM[k] = v
			}
		}
		fms = append(fms, newFM)
	}

	for _, c := range capsToStrip {
		delete(am, c)
		for _, fm := range fms {
			delete(fm, c)
		}
	}

	return &Capabilities{
		AlwaysMatch:  am,
		FirstMatch:   fms,
		W3CSupported: c.W3CSupported,
	}
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
	if m1 == nil {
		return m2
	}
	if m2 == nil {
		return m1
	}
	nm := map[string]interface{}{}
	for k, v := range m1 {
		nm[k] = v
	}
	for k, v := range m2 {
		nm[k] = mergeValues(nm[k], v, k)
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
	if m1 == nil {
		return m2
	}
	if m2 == nil {
		return m1
	}
	nl := []interface{}{}
	nl = append(nl, m1...)
	nl = append(nl, m2...)
	return nl
}

func mergeArgs(m1, m2 []interface{}) []interface{} {
	m2Opts := map[string]bool{}

	for _, a := range m2 {
		if arg, ok := a.(string); ok {
			if strings.HasPrefix(arg, "--") {
				tokens := strings.Split(arg, "=")
				m2Opts[tokens[0]] = true
			}
		}
	}

	nl := []interface{}{}

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

// A Resolver resolves a prefix, name pair to a replacement value.
type Resolver func(prefix, name string) (string, error)

// NoOPResolver resolves to %prefix:name%.
func NoOPResolver(prefix, name string) (string, error) {
	return "%" + prefix + ":" + name + "%", nil
}

// MapResolver returns a new Resolver that uses key-value pairs in names to
// resolve names for prefix, and otherwise uses the NoOPResolver.
func MapResolver(prefix string, names map[string]string) Resolver {
	return func(p, n string) (string, error) {
		if p == prefix {
			v, ok := names[n]
			if !ok {
				return "", fmt.Errorf("unable to resolve %s:%s", p, n)
			}
			return v, nil
		}
		return NoOPResolver(p, n)
	}
}

// Resolve returns a new Capabilities object with all %PREFIX:NAME% substrings replaced using resolver.
func (c *Capabilities) Resolve(resolver Resolver) (*Capabilities, error) {
	am, err := resolveMap(c.AlwaysMatch, resolver)
	if err != nil {
		return nil, err
	}

	var fms []map[string]interface{}

	for _, fm := range c.FirstMatch {
		u, err := resolveMap(fm, resolver)
		if err != nil {
			return nil, err
		}
		fms = append(fms, u)
	}

	return &Capabilities{
		AlwaysMatch:  am,
		FirstMatch:   fms,
		W3CSupported: c.W3CSupported,
	}, nil
}

func resolve(v interface{}, resolver Resolver) (interface{}, error) {
	switch tv := v.(type) {
	case string:
		return resolveString(tv, resolver)
	case []interface{}:
		return resolveSlice(tv, resolver)
	case map[string]interface{}:
		return resolveMap(tv, resolver)
	default:
		return v, nil
	}
}

func resolveMap(m map[string]interface{}, resolver Resolver) (map[string]interface{}, error) {
	caps := map[string]interface{}{}

	for k, v := range m {
		u, err := resolve(v, resolver)
		if err != nil {
			return nil, err
		}

		caps[k] = u
	}

	return caps, nil
}

func resolveSlice(l []interface{}, resolver Resolver) ([]interface{}, error) {
	caps := []interface{}{}

	for _, v := range l {
		u, err := resolve(v, resolver)
		if err != nil {
			return nil, err
		}
		caps = append(caps, u)
	}

	return caps, nil
}

var varRegExp = regexp.MustCompile(`%(\w+):(\w+)%`)

func resolveString(s string, resolver Resolver) (string, error) {
	result := ""
	previous := 0
	for _, match := range varRegExp.FindAllStringSubmatchIndex(s, -1) {
		// Append everything after the previous match to the beginning of this match
		result += s[previous:match[0]]
		// Set previous to the first character after this match
		previous = match[1]

		prefix := s[match[2]:match[3]]
		varName := s[match[4]:match[5]]

		value, err := resolver(prefix, varName)
		if err != nil {
			return "", err
		}

		result += value
	}

	// Append everything after the last match
	result += s[previous:]

	return result, nil
}
