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

// Package diagnostics provides the Diagnostics interface for reporting various
// test statistics as well as a no-op implementation of the interface.
package diagnostics

import (
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
)

// Diagnostics manages and outputs a test diagnostics.
type Diagnostics interface {
	// Name is the name of the component used in error messages.
	Name() string
	// Timing reports timing info.
	Timing(component, description, detail string, begin, end time.Time) error
	// Severe reports an error that is highly likely to cause problems in tests.
	Severe(err error) error
	// Warning reports a error that is unlikely to cause problems in tests, but
	// that you want to track anyway.
	Warning(err error) error
	// Close closes this diagnostics object.
	Close() error
}

type noOPDiagnostics struct {
	closed bool
}

// NoOP creates a new empty Diagnostics object.
func NoOP() Diagnostics {
	return &noOPDiagnostics{}
}

// Name is the name of the component used in error messages.
func (d *noOPDiagnostics) Name() string {
	return "No-OP Diagnostics"
}

// Timing reports timing info.
func (d *noOPDiagnostics) Timing(_, _, _ string, _, _ time.Time) error {
	if d.closed {
		return errors.New(d.Name(), "cannot add timing data after diagnostics are closed")
	}
	return nil
}

// Severe reports an error that is highly likely to cause problems in tests.
func (d *noOPDiagnostics) Severe(err error) error {
	if d.closed {
		return errors.New(d.Name(), "cannot add errors after diagnostics are closed")
	}
	return nil
}

// Warning reports a error that is unlikely to cause problems in tests, but
// that you want to track anyway.
func (d *noOPDiagnostics) Warning(err error) error {
	if d.closed {
		return errors.New(d.Name(), "cannot add errors after diagnostics are closed")
	}
	return nil
}

// Close closes this diagnostics object.
func (d *noOPDiagnostics) Close() error {
	d.closed = true
	return nil
}

// String returns a string representation of the diagnostics.
func (d *noOPDiagnostics) String() string {
  return "No-OP Diagnostics"
}
