/* Copyright 2016 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package metadata provides a struct for storing browser metadata.
package metadata

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bazelbuild/rules_web/metadata/capabilities"
)

// Values for Metadata.RecordVideo.
const (
	RecordNever  = "never"
	RecordFailed = "failed"
	RecordAlways = "always"
)

// Metadata provide necessary metadata launching a browser.
type Metadata struct {
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
	// FormFactor that this browser is/pretends to be. One of "DESKTOP", "TABLET", "PHONE".
	FormFactor string `json:"formFactor,omitempty"`
	// BrowserName that this browser pretends to be.
	BrowserName string `json:"browserName,omitempty"`
	// The Environment that web test launcher should use to to launch the browser.
	Environment string `json:"environment,omitempty"`
	// Browser label set in the web_test rule.
	BrowserLabel string `json:"browserLabel,omitempty"`
	// Test label set in the web_test rule.
	TestLabel string `json:"testLabel,omitempty"`
	// Whether to crop screenshots. Cannot be a bool as absent is treated differently from false
	// for merging.
	CropScreenshots interface{} `json:"cropScreenshots,omitempty"`
	// Whether to record and keep videos. One of "always", "failed", "never".
	RecordVideo string `json:"recordVideo,omitempty"`
	// Whether to wait for the environment to be healthy before the test starts.
	HealthyBeforeTest interface{} `json:"healthyBeforeTest,omitempty"`
}

// Merge takes two Metadata objects and merges them into a new Metadata object.
func Merge(m1, m2 Metadata) Metadata {
	capabilities := capabilities.Merge(m1.Capabilities, m2.Capabilities)

	formFactor := m1.FormFactor
	if m2.FormFactor != "" {
		formFactor = m2.FormFactor
	}

	browserName := m1.BrowserName
	if m2.BrowserName != "" {
		browserName = m2.BrowserName
	}

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

	cropScreenshots := m1.CropScreenshots
	if m2.CropScreenshots != nil {
		cropScreenshots = m2.CropScreenshots
	}

	recordVideo := m1.RecordVideo
	if m2.RecordVideo != "" {
		recordVideo = m2.RecordVideo
	}

	healthyBeforeTest := m1.HealthyBeforeTest
	if m2.HealthyBeforeTest != nil {
		healthyBeforeTest = m2.HealthyBeforeTest
	}

	return Metadata{
		Capabilities:      capabilities,
		FormFactor:        formFactor,
		BrowserName:       browserName,
		Environment:       environment,
		BrowserLabel:      browserLabel,
		TestLabel:         testLabel,
		CropScreenshots:   cropScreenshots,
		RecordVideo:       recordVideo,
		HealthyBeforeTest: healthyBeforeTest,
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
		e.FormFactor == a.FormFactor &&
		e.BrowserName == a.BrowserName &&
		e.Environment == a.Environment &&
		e.BrowserLabel == a.BrowserLabel &&
		e.TestLabel == a.TestLabel &&
		e.CropScreenshots == a.CropScreenshots &&
		e.RecordVideo == a.RecordVideo &&
		e.HealthyBeforeTest == a.HealthyBeforeTest
}
