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
//
////////////////////////////////////////////////////////////////////////////////
//
// Package metadata provides a struct for storing browser metadata.
package metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bazelbuild/rules_web/go/metadata/capabilities"
	"github.com/bazelbuild/rules_web/go/util/bazel"
)

// Values for Metadata.RecordVideo.
const (
	RecordNever  = "never"
	RecordFailed = "failed"
	RecordAlways = "always"
)

// Metadata provides necessary metadata for launching a browser.
type Metadata struct {
	// The Capabilities that should be used for this browser.
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
	// The Environment that web test launcher should use to to launch the browser.
	Environment string `json:"environment,omitempty"`
	// Browser label set in the web_test rule.
	BrowserLabel string `json:"browserLabel,omitempty"`
	// Test label set in the web_test rule.
	TestLabel string `json:"testLabel,omitempty"`
	// A map of names to TEST_SRCDIR-relative file paths.
	// Prefer using GetExecutablePath instead of accessing this map directly.
	NamedExecutables map[string]string `json:"NamedExecutables,omitempty"`
}

// Merge takes two Metadata objects and merges them into a new Metadata object.
func Merge(m1, m2 Metadata) Metadata {
	capabilities := capabilities.Merge(m1.Capabilities, m2.Capabilities)

	environment := m1.Environment
	if m2.Environment != "" {
		environment = m2.Environment
	}

	browserLabel := m1.BrowserLabel
	if m2.BrowserLabel != "" {
		browserLabel = m2.BrowserLabel
	}

	testLabel := m1.TestLabel
	if m2.TestLabel != "" {
		testLabel = m2.TestLabel
	}

	// TODO(fisherii): propagate merge error
	namedExecutables, _ := mergeNamedExecutables(m1.NamedExecutables, m2.NamedExecutables)

	return Metadata{
		Capabilities:     capabilities,
		Environment:      environment,
		BrowserLabel:     browserLabel,
		TestLabel:        testLabel,
		NamedExecutables: namedExecutables,
	}
}

// FromFile reads a Metadata object from a json file.
func FromFile(filename string) (Metadata, error) {
	metadata := Metadata{}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return metadata, err
	}

	if err := json.Unmarshal(bytes, &metadata); err != nil {
		return metadata, err
	}
	return metadata, nil
}

// ToFile writes m to filename as json.
func (m Metadata) ToFile(filename string) error {
	bytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

// Equals compares two Metadata object and return true iff they are the same.
func Equals(e, a Metadata) bool {
	return capabilities.Equals(e.Capabilities, a.Capabilities) &&
		e.Environment == a.Environment &&
		e.BrowserLabel == a.BrowserLabel &&
		e.TestLabel == a.TestLabel &&
		mapEquals(e.NamedExecutables, a.NamedExecutables)
}

func mergeNamedExecutables(n1, n2 map[string]string) (map[string]string, error) {
	result := map[string]string{}

	for k, v := range n1 {
		result[k] = v
	}

	for k, v2 := range n2 {
		if v1, ok := result[k]; ok && v1 != v2 {
			return nil, fmt.Errorf("key %q exists in both NamedFiles with different values", k)
		}
		result[k] = v2
	}
	return result, nil
}

func mapEquals(e, a map[string]string) bool {
	if len(e) != len(a) {
		return false
	}
	for k, ev := range e {
		if av, ok := a[k]; !ok || ev != av {
			return false
		}
	}
	return true
}

func (m Metadata) GetExecutablePath(name string) (string, error) {
	filename, ok := m.NamedExecutables[name]
	if !ok {
		return "", fmt.Errorf("no named executable %q", name)
	}
	return bazel.Runfile(filename)
}
