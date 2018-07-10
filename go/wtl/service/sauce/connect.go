// Copyright 2018 Google Inc.
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

// Package sauce provides a Service for managing Sauce Connect.
package sauce

import (
	"bufio"
	"context"
	"os/exec"
	"strings"
	"sync"

	"github.com/bazelbuild/rules_webtesting/go/errors"
	"github.com/bazelbuild/rules_webtesting/go/metadata"
)

const (
	compName    = "Sauce Connect Service"
	scNamedFile = "SAUCE_CONNECT"
)

type Connect struct {
	cmd *exec.Cmd

	mu    sync.Mutex
	ready bool
	err   error
}

// New creates a new service that manages Sauce Connect.
func New(m *metadata.Metadata, username, accessKey, tunnelID string) (*Connect, error) {
	scPath, err := m.GetFilePath(scNamedFile)
	if err != nil {
		return nil, errors.New(compName, err)
	}

	cmd := exec.Command(
		scPath,
		"--user", username,
		"--api-key", accessKey,
		"--tunnel-identifier", tunnelID)

	return &Connect{cmd: cmd}, nil
}

// Start starts Sauce Connect and waits for it to be ready for use.
func (c *Connect) Start(_ context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := c.cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Sauce Connect is up, you may start your tests") {
			c.ready = true
			go c.monitor()
			return nil
		}
	}

	return errors.New(c.Name(), "terminated without becoming healthy")
}

// Stop stops a running Sauce Connect.
func (c *Connect) Stop(_ context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ready || c.err != nil {
		return nil
	}
	c.ready = false

	if c.cmd.Process != nil {
		c.err = c.cmd.Process.Kill()
	}
	return nil
}

// Healthy returns nil if Sauce Connect is running and ready for use, otherwise it returns an error.
func (c *Connect) Healthy(_ context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.err != nil {
		return c.err
	}

	if !c.ready {
		return errors.New(c.Name(), "has not been started")
	}

	return nil
}

func (c *Connect) monitor() {
	err := c.cmd.Wait()

	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ready && c.err == nil {
		c.ready = false
		c.err = err
	}
}

func (c *Connect) Name() string {
	return compName
}
