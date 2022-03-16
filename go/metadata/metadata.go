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

	"github.com/bazelbuild/rules_webtesting/go/httphelper"
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

	var webTestFiles []*WebTestFiles
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

// GetFilePath returns the path to a file specified by platform_archive,
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

// Resolver returns a Resolver that processes ENV, FILE, and METADATA prefixed
// capabilities variables.
func (m *Metadata) Resolver() capabilities.Resolver {
	metadataResolver := capabilities.MapResolver("METADATA", map[string]string{
		"LABEL":         m.Label,
		"TEST_LABEL":    m.TestLabel,
		"BROWSER_LABEL": m.BrowserLabel,
		"CONFIG_LABEL":  m.ConfigLabel,
		"ENVIRONMENT":   m.Environment,
	})

	return func(prefix, name string) (string, error) {
		switch prefix {
		case "ENV":
			v, ok := os.LookupEnv(name)
			if !ok {
				return "", fmt.Errorf("environment variable %q is not defined", name)
			}
			return v, nil
		case "FILE":
			return m.GetFilePath(name)
		case "WTL":
			switch name {
			case "FQDN":
				return httphelper.FQDN()
			default:
				return "", fmt.Errorf("WTL:%q is not defined", name)
			}
		default:
			return metadataResolver(prefix, name)
		}
	}
}

type extension map[string]interface{}

func (e extension) Merge(other Extension) (Extension, error) {
	if other == nil {
		return e, nil
	}
	if e == nil || len(e) == 0 {
		return other, nil
	}
	o, ok := other.(extension)
	if !ok || len(o) == 0 {
		return e, nil
	}
	ext := extension{}
	for k, v := range e {
		ext[k] = v
	}
	for k, v := range o {
		ext[k] = v
	}
	return ext, nil
}

func (e extension) Normalize() error {
	return nil
}

// ExtensionMap returns the Extension field as a map if, when the metadata was read,
// no extension type was specified.
func (m *Metadata) ExtensionMap() (map[string]interface{}, bool) {
	if m.Extension == nil {
		return nil, false
	}

	ext, ok := m.Extension.(*extension)
	if !ok {
		return nil, false
	}

	return map[string]interface{}(*ext), true
}
