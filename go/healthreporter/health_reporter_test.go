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

package healthreporter

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/errors"
)

func TestWaitForHealthyTimeout(t *testing.T) {
	failing := &fakeHealthReporter{errors.New("fake", "error")}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := WaitForHealthy(ctx, failing)
	if err == nil {
		t.Fatal("expected err to be non-nil")
	}
	if !strings.Contains(err.Error(), ctx.Err().Error()) {
		t.Errorf(`expected err to contain "%s", but got %s`, ctx.Err().Error(), err.Error())
	}
	if !strings.Contains(err.Error(), failing.health.Error()) {
		t.Errorf(`expected err to contain "%s", but got %s`, failing.health.Error(), err.Error())
	}
	if errors.Component(err) != "fake" {
		t.Errorf(`expected err.Component to be "%s", but got %s`, "fake", errors.Component(err))
	}
}

func TestWaitForHealthyCancelled(t *testing.T) {
	failing := &fakeHealthReporter{errors.New("fake", "error")}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := WaitForHealthy(ctx, failing)
	if err == nil {
		t.Fatal("expected err to be non-nil")
	}
	if !strings.Contains(err.Error(), ctx.Err().Error()) {
		t.Errorf(`expected err to contain "%s", but got %s`, ctx.Err().Error(), err.Error())
	}
	if errors.Component(err) != "fake" {
		t.Errorf(`expected err.Component to be "%s", but got %s`, "fake", errors.Component(err))
	}
}

func TestWaitForHealthySuccess(t *testing.T) {
	successful := &fakeHealthReporter{}

	err := WaitForHealthy(context.Background(), successful)
	if err != nil {
		t.Fatalf("expected err to be nil, but got %v", err)
	}
}

type fakeHealthReporter struct {
	health error
}

func (f *fakeHealthReporter) Name() string {
	return "fakeHealthReporter"
}

func (f *fakeHealthReporter) Healthy(context.Context) error {
	return f.health
}
