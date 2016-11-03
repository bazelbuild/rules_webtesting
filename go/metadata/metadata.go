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
	"regexp"

	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
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
	// A list of archive files with named files in them.
	// Note: An archive will only be extracted if GetExecutablePath is called
	// with one of the named files.
	WebTestFiles []*WebTestFiles `json:"webTestFiles,omitempty"`
}

// Merge takes two Metadata objects and merges them into a new Metadata object.
func Merge(m1, m2 *Metadata) (*Metadata, error) {
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

	webTestFiles := []*WebTestFiles{}
	webTestFiles = append(webTestFiles, m1.WebTestFiles...)
	webTestFiles = append(webTestFiles, m2.WebTestFiles...)

	webTestFiles, err := normalizeWebTestFiles(webTestFiles)
	if err != nil {
		return nil, err
	}

	return &Metadata{
		Capabilities: capabilities,
		Environment:  environment,
		BrowserLabel: browserLabel,
		TestLabel:    testLabel,
		WebTestFiles: webTestFiles,
	}, nil
}

// FromFile reads a Metadata object from a json file.
func FromFile(filename string) (*Metadata, error) {
	metadata := &Metadata{}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, metadata); err != nil {
		return nil, err
	}
	webTestFiles, err := normalizeWebTestFiles(metadata.WebTestFiles)
	if err != nil {
		return nil, err
	}
	metadata.WebTestFiles = webTestFiles

	return metadata, nil
}

// ToFile writes m to filename as json.
func (m *Metadata) ToFile(filename string) error {
	bytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

// Equals compares two Metadata object and return true iff they are the same.
func Equals(e, a *Metadata) bool {
	// TODO(DrMarcII): should consider equality of WebTestFiles.
	return capabilities.Equals(e.Capabilities, a.Capabilities) &&
		e.Environment == a.Environment &&
		e.BrowserLabel == a.BrowserLabel &&
		e.TestLabel == a.TestLabel &&
		webTestFilesSliceEquals(e.WebTestFiles, a.WebTestFiles)
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

// GetFilePath returns the path to a file specified by web_test_archive,
// web_test_named_executable, or web_test_named_file.
func (m *Metadata) GetFilePath(name string) (string, error) {
	for _, a := range m.WebTestFiles {
		filename, err := a.getFilePath(name)
		if err != nil {
			return "", err
		}
		if filename != "" {
			return filename, nil
		}
	}
	return "", fmt.Errorf("no named file %q", name)
}

var varRegExp = regexp.MustCompile(`%\w+%`)

// ResolvedCapabilities returns Capabilities with any strings/substrings
// of the form %NAME% resolved to a file path retrieved with GetFilePath.
func (m *Metadata) ResolvedCapabilities() (map[string]interface{}, error) {
	var resolve func(v interface{}) (interface{}, error)

	resolveMap := func(m map[string]interface{}) (map[string]interface{}, error) {
		caps := map[string]interface{}{}
		for k, v := range m {
			u, err := resolve(v)
			if err != nil {
				return nil, err
			}
			caps[k] = u
		}
		return caps, nil
	}
	resolveSlice := func(l []interface{}) ([]interface{}, error) {
		caps := []interface{}{}
		for _, v := range l {
			u, err := resolve(v)
			if err != nil {
				return nil, err
			}
			caps = append(caps, u)
		}
		return caps, nil
	}
	resolveString := func(s string) (string, error) {
		result := ""
		previous := 0
		for _, match := range varRegExp.FindAllStringIndex(s, -1) {
			result += s[previous:match[0]]
			name := s[match[0]+1 : match[1]-1]
			path, err := m.GetFilePath(name)
			if err != nil {
				return "", err
			}
			result += path
			previous = match[1]
		}
		result += s[previous:]
		return result, nil
	}
	resolve = func(v interface{}) (interface{}, error) {
		switch tv := v.(type) {
		case string:
			return resolveString(tv)
		case []interface{}:
			return resolveSlice(tv)
		case map[string]interface{}:
			return resolveMap(tv)
		default:
			return v, nil
		}
	}
	return resolveMap(m.Capabilities)
}
