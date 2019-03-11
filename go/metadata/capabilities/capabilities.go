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
			n, err := normalize(am)
			if err != nil {
				return nil, err
			}
			if err := mergeIntoNoReplace(always, n); err != nil {
				return nil, err
			}
		}
	}

	if required, ok := args["requiredCapabilities"].(map[string]interface{}); ok {
		n, err := normalize(required)
		if err != nil {
			return nil, err
		}
		if err := mergeIntoNoReplace(always, n); err != nil {
			return nil, err
		}
	}

	if desired, ok := args["desiredCapabilities"].(map[string]interface{}); ok {
		n, err := normalize(desired)
		if err != nil {
			return nil, err
		}
		if err := mergeIntoNoReplace(always, n); err != nil {
			return nil, err
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
			nfme, err := normalize(fme)
			if err != nil {
				return nil, err
			}
			for k, v := range nfme {
				if a, ok := always[k]; ok {
					if !reflect.DeepEqual(a, v) {
						return nil, fmt.Errorf("alwaysMatch|required|desired[%q] == %+v, firstMatch[%q] == %+v, they must be equal", k, a, k, v)
					}
					continue
				}
				newFM[k] = v
			}
			first = append(first, newFM)
		}
	}

	var reduced []map[string]interface{}

	for i, v := range first {

		duped := false
		for _, v2 := range first[i+1:] {
			if reflect.DeepEqual(v, v2) {
				duped = true
				break
			}
		}

		if !duped {
			reduced = append(reduced, v)
		}
	}

	first = reduced

	if len(first) == 1 {
		if err := mergeIntoNoReplace(always, first[0]); err != nil {
			return nil, err
		}
		first = nil
	}

	return &Capabilities{
		AlwaysMatch:  always,
		FirstMatch:   first,
		W3CSupported: w3c != nil,
	}, nil
}

func normalize(in map[string]interface{}) (map[string]interface{}, error) {
	// Normalize chromeOptions
	out, err := duplicateLegacyGoogCapability(in, map[string]interface{}{}, "chromeOptions")
	if err != nil {
		return nil, err
	}

	// Normalize loggingPrefs
	out, err = duplicateLegacyGoogCapability(in, out, "loggingPrefs")
	if err != nil {
		return nil, err
	}

	// Normalize proxy
	outProxy := map[string]interface{}{}

	if proxy, ok := in["proxy"]; ok {
		proxyMap, ok := proxy.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("proxy %v is %T, should be map[string]interface{}", proxy, proxy)
		}
		for k, v := range proxyMap {
			switch k {
			case "proxyType":
				pt, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("proxyType %v is %T, should be string", k, k)
				}
				outProxy["proxyType"] = strings.ToLower(pt)
				continue

			case "noProxy":
				var outNP []interface{}
				switch np := v.(type) {
				case []interface{}:
					outNP = append(outNP, np...)
				case string:
					for _, h := range strings.Split(np, ",") {
						outNP = append(outNP, h)
					}
				default:
					return nil, fmt.Errorf("noProxy %v is %T, should be string or []interface{}", k, k)
				}
				if len(outNP) != 0 {
					outProxy["noProxy"] = outNP
				}
			default:
				if v != nil {
					outProxy[k] = v
				}
			}
		}
	}

	if len(outProxy) != 0 {
		out["proxy"] = outProxy
	}

	for k, v := range in {
		if k != "chromeOptions" && k != "goog:chromeOptions" && k != "proxy" && k != "loggingPrefs" && k != "goog:loggingPrefs" {
			out[k] = v
		}
	}

	return out, nil
}

// duplicateLegacyGoogCapability duplicates and merges paramName from "in" to "out" with name "goog:"+paramName.
func duplicateLegacyGoogCapability(in map[string]interface{}, out map[string]interface{}, paramName string) (map[string]interface{}, error) {
	outParamVal := map[string]interface{}{}
	if param, ok := in[paramName]; ok {
		inParamVal, ok := param.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("%s %v is %T, should be map[string]interface{}", paramName, param, param)
		}
		if err := mergeIntoNoReplace(outParamVal, inParamVal); err != nil {
			return nil, err
		}
	}

	if param, ok := in["goog:"+paramName]; ok {
		inParamVal, ok := param.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("goog:%s %v is %T, should be map[string]interface{}", paramName, param, param)
		}
		if err := mergeIntoNoReplace(outParamVal, inParamVal); err != nil {
			return nil, err
		}
	}

	if len(outParamVal) != 0 {
		out["goog:"+paramName] = outParamVal
	}
	return out, nil
}

func mergeIntoNoReplace(dst, src map[string]interface{}) error {
	for k, sv := range src {
		dv, ok := dst[k]
		if ok && !reflect.DeepEqual(dv, sv) {
			return fmt.Errorf("dst[%q] == %v, src[%q] == %v, they must be equal", k, dv, k, sv)
		}
		dst[k] = sv
	}
	return nil
}

// denormalizeW3C removes non-W3C capabilities.
func denormalizeW3C(in map[string]interface{}) map[string]interface{} {
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

// denormalizeJWP duplicates goog:chromeOptions into chromeOptions and converts noProxy from []interface{} into string.
func denormalizeJWP(caps map[string]interface{}) map[string]interface{} {
	out := map[string]interface{}{}

	for k, v := range caps {
		if k == "goog:chromeOptions" {
			out["chromeOptions"] = v
			out["goog:chromeOptions"] = v
			continue
		} else if k == "goog:loggingPrefs" {
			out["loggingPrefs"] = v
			out["goog:loggingPrefs"] = v
			continue
		}

		if k != "proxy" {
			out[k] = v
			continue
		}

		proxy, ok := v.(map[string]interface{})
		if !ok {
			out[k] = v
			continue
		}

		outProxy := map[string]interface{}{}
		for pk, pv := range proxy {
			if pk != "noProxy" {
				outProxy[pk] = pv
				continue
			}

			noProxy, ok := pv.([]interface{})
			if !ok {
				outProxy["noProxy"] = pv
				continue
			}

			var noProxys []string

			for _, npv := range noProxy {
				nps := npv.(string)
				noProxys = append(noProxys, nps)
			}

			outProxy["noProxy"] = strings.Join(noProxys, ",")
		}

		out["proxy"] = outProxy
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

	// Partition other into those that should affect entries in FirstMatch, and those that should affect AlwaysMatch.
	for k, v := range other {
		if anyContains(c.FirstMatch, k) {
			first[k] = v
		} else {
			always[k] = v
		}
	}

	firstMatch := c.FirstMatch

	// If any of the entries in other should affect FirstMatch, then merge each FirstMatch of those entries.
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

// MergeUnder creates a new Capabilities with caps in other removed from all entries in c.FirstMatch and merged over c.AlwaysMatch.
func (c *Capabilities) MergeUnder(other map[string]interface{}) *Capabilities {
	if len(other) == 0 {
		return c
	}

	if c == nil {
		return &Capabilities{
			AlwaysMatch: other,
		}
	}

	var first []map[string]interface{}

	// Remove any keys that are in other from all FirstMatch capabilities.
	for _, oldF := range c.FirstMatch {
		newF := map[string]interface{}{}

		for k, v := range oldF {
			if _, ok := other[k]; !ok {
				newF[k] = v
			}
		}

		// Since we are removing keys from FirstMatch capabilities, it is possible for some of them to now
		// be identical, so only add newF if there isn't one in first that is the same.
		duped := false
		for _, v := range first {
			if reflect.DeepEqual(newF, v) {
				duped = true
				break
			}
		}

		if !duped {
			first = append(first, newF)
		}
	}

	always := c.AlwaysMatch

	// If first now only contains one entry, then merge it with always.
	if len(first) == 1 {
		always = Merge(first[0], always)
		first = nil
	}

	always = Merge(always, other)

	return &Capabilities{
		AlwaysMatch:  always,
		FirstMatch:   first,
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
		"desiredCapabilities": denormalizeJWP(desired),
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

	alwaysMatch := denormalizeW3C(c.AlwaysMatch)
	var firstMatch []map[string]interface{}

	for _, fm := range c.FirstMatch {
		firstMatch = append(firstMatch, denormalizeW3C(fm))
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
		if k == "chromeOptions" {
			k = "goog:chromeOptions"
		} else if k == "loggingPrefs" {
			k = "goog:loggingPrefs"
		}
		nm[k] = v
	}
	for k, v := range m2 {
		if k == "chromeOptions" {
			k = "goog:chromeOptions"
		} else if k == "loggingPrefs" {
			k = "goog:loggingPrefs"
		}
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

