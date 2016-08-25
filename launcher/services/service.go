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

// Package service provides the Service interface for managing the life-cycle
// of a single binary.
package service

import (
	"context"
	"sync"

	"github.com/web_test_launcher/launcher/errors"
	"github.com/web_test_launcher/launcher/healthreporter"
)

// Service provides is an interface for an individual service. These are usually
// composed to create an environment.
type Service interface {
	healthreporter.HealthReporter
	// Start starts the service, returning when the service is healthy or has failed.
	Start(context.Context) error
	// Stop shutdowns the service.
	Stop(context.Context) error
}

// Base is a base struct for defining a service.
type Base struct {
	name    string
	mu      sync.RWMutex
	started bool
	stopped bool
}

// NewBase creates a new Base service with the given component name.
func NewBase(name string) *Base {
	return &Base{name: name}
}

// Start makes this service as started, failing if it has already been started.
func (b *Base) Start(context.Context) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.started {
		return errors.NewPermanent(b.Name(), "has already started")
	}
	b.started = true
	return nil
}

// Stop makes this service as stopped, failing if it hasn't been started or has already been stopped.
func (b *Base) Stop(context.Context) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if !b.started {
		return errors.NewPermanent(b.Name(), "has not been started")
	}
	if b.stopped {
		return errors.NewPermanent(b.Name(), "has already stopped")
	}
	b.stopped = true
	return nil
}

// Healthy returns nil if this service has been started and has not been stopped.
func (b *Base) Healthy(context.Context) error {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if !b.started {
		return errors.NewPermanent(b.Name(), "has not been started")
	}
	if b.stopped {
		return errors.NewPermanent(b.Name(), "has been stopped")
	}
	return nil
}

// Component returns the component name used in error messages.
func (b *Base) Name() string {
	return b.name
}
