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

// Package metadata provides a struct for storing browser metadata.
package metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
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
	// Label for the web_test rule.
	Label string `json:"label,omitempty"`
	// Browser label set in the web_test rule.
	BrowserLabel string `json:"browserLabel,omitempty"`
	// Test label set in the web_test rule.
	TestLabel string `json:"testLabel,omitempty"`
	// Config label set in the web_test rule.
	ConfigLabel string `json:"configLabel,omitempty"`
	// Port to connect debugger to. If 0, debugger will not be started.
	DebuggerPort int `json:"debuggerPort,omitempty"`
	// A list of WebTestFiles with named files in them.
	WebTestFiles []*WebTestFiles `json:"webTestFiles,omitempty"`
	// An object for any additional metadata fields on this object.
	Extension `json:"extension,omitempty"`
}

// Extension is an interface for adding additional fields that will be parsed as part of the metadata.
type Extension interface {
	// Merge merges this extension data with another set of Extension data. It should not mutate either
	// Extension object, but it is allowed to return one of the Extension objects unchanged if needed.
	// In general values in other should take precedence over values in this object.
	Merge(other Extension) (Extension, error)
	// Normalize normalizes and validate the extension data.
	Normalize() error
	// Equals returns true iff other should be treated as equal to this.
	Equals(other Extension) bool
}

// Merge takes two Metadata objects and merges them into a new Metadata object.
func Merge(m1, m2 *Metadata) (*Metadata, error) {
	capabilities := capabilities.Merge(m1.Capabilities, m2.Capabilities)

	environment := m1.Environment
	if m2.Environment != "" {
		environment = m2.Environment
	}

	label := m1.Label
	if m2.Label != "" {
		label = m2.Label
	}

	browserLabel := m1.BrowserLabel
	if m2.BrowserLabel != "" {
		browserLabel = m2.BrowserLabel
	}

	testLabel := m1.TestLabel
	if m2.TestLabel != "" {
		testLabel = m2.TestLabel
	}

	configLabel := m1.ConfigLabel
	if m2.ConfigLabel != "" {
		configLabel = m2.ConfigLabel
	}

	debuggerPort := m1.DebuggerPort
	if m2.DebuggerPort != 0 {
		debuggerPort = m2.DebuggerPort
	}

	webTestFiles := []*WebTestFiles{}
	webTestFiles = append(webTestFiles, m1.WebTestFiles...)
	webTestFiles = append(webTestFiles, m2.WebTestFiles...)

	webTestFiles, err := normalizeWebTestFiles(webTestFiles)
	if err != nil {
		return nil, err
	}

	extension := m1.Extension
	if extension == nil {
		extension = m2.Extension
	} else if m2.Extension != nil {
		e, err := extension.Merge(m2.Extension)
		if err != nil {
			return nil, err
		}
		extension = e
	}

	return &Metadata{
		Capabilities: capabilities,
		Environment:  environment,
		Label:        label,
		BrowserLabel: browserLabel,
		TestLabel:    testLabel,
		ConfigLabel:  configLabel,
		DebuggerPort: debuggerPort,
		WebTestFiles: webTestFiles,
		Extension:    extension,
	}, nil
}

// FromFile reads a Metadata object from a json file.
func FromFile(filename string, ext Extension) (*Metadata, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return FromBytes(bytes, ext)
}

// FromBytes reads a Metadata object from a byte array.
func FromBytes(bytes []byte, ext Extension) (*Metadata, error) {
	if ext == nil {
		ext = &extension{}
	}
	metadata := &Metadata{Extension: ext}

	if err := json.Unmarshal(bytes, metadata); err != nil {
		return nil, err
	}
	webTestFiles, err := normalizeWebTestFiles(metadata.WebTestFiles)
	if err != nil {
		return nil, err
	}
	metadata.WebTestFiles = webTestFiles

	if metadata.Extension != nil {
		if err := metadata.Extension.Normalize(); err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

// ToFile writes m to filename as json.
func (m *Metadata) ToFile(filename string) error {
	bytes, err := m.ToBytes()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

// ToBytes serializes metadata.
func (m *Metadata) ToBytes() ([]byte, error) {
	return json.MarshalIndent(m, "", "  ")
}

// Equals compares two Metadata object and return true iff they are the same.
func Equals(e, a *Metadata) bool {
	var extsEqual bool
	if e.Extension == nil {
		extsEqual = (a.Extension == nil) || a.Extension.Equals(e.Extension)
	} else {
		extsEqual = e.Extension.Equals(a.Extension)
	}
	return extsEqual &&
		e.Environment == a.Environment &&
		e.Label == a.Label &&
		e.BrowserLabel == a.BrowserLabel &&
		e.TestLabel == a.TestLabel &&
		e.ConfigLabel == a.ConfigLabel &&
		e.DebuggerPort == a.DebuggerPort &&
		reflect.DeepEqual(e.Capabilities, a.Capabilities) &&
		webTestFilesSliceEquals(e.WebTestFiles, a.WebTestFiles)
}

// GetFilePath returns the path to a file specified by web_test_archive,
// web_test_named_executable, or web_test_named_file.
func (m *Metadata) GetFilePath(name string) (string, error) {
	for _, a := range m.WebTestFiles {
		filename, err := a.getFilePath(name, m)
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

// ResolvedCapabilities returns Capabilities with any strings updated such that:
//   1. %NAME% will be replaced:
//      if NAME one of LABEL, TEST_LABEL, BROWSER_LABEL, CONFIG_LABEL, ENVIRONMENT then the corresponding field on m.
//      otherwise the value m.GetFilePath(NAME).
//   2. $ENV_VAR or ${ENV_VAR} will be replaced with the value of the environment variable ENV_VAR.
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
			value := ""
			switch name := s[match[0]+1 : match[1]-1]; name {
			case "LABEL":
				value = m.Label
			case "TEST_LABEL":
				value = m.TestLabel
			case "BROWSER_LABEL":
				value = m.BrowserLabel
			case "CONFIG_LABEL":
				value = m.ConfigLabel
			case "ENVIRONMENT":
				value = m.Environment
			default:
				path, err := m.GetFilePath(name)
				if err != nil {
					return "", err
				}
				value = path
			}
			result += value
			previous = match[1]
		}
		result += s[previous:]
		return os.ExpandEnv(result), nil
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

type extension struct {
	value map[string]interface{}
}

func (e *extension) Merge(other Extension) (Extension, error) {
	if other == nil {
		return e, nil
	}
	o, ok := other.(*extension)
	if !ok {
		return other, nil
	}
	return &extension{capabilities.Merge(e.value, o.value)}, nil
}

func (e *extension) Normalize() error {
	return nil
}

func (e *extension) Equals(other Extension) bool {
	return reflect.DeepEqual(e, other)
}

func (e *extension) UnmarshalJSON(b []byte) error {
	v := map[string]interface{}{}
	err := json.Unmarshal(b, &v)
	e.value = v
	return err
}
