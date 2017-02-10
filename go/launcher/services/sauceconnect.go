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

// Package sauceconnect provides a service.Service for managing an instance of sauceconnect.
package sauceconnect

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/services/service"
	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
	"github.com/bazelbuild/rules_webtesting/go/util/portpicker"
)

const readySignature = "Sauce Connect is up, you may start your tests."

type SauceConnect struct {
	*service.Cmd
	tunnelID string

	mu    sync.RWMutex
	ready bool
}

// New creates a new service.Server instance that manages chromedriver.
func New(d diagnostics.Diagnostics, m *metadata.Metadata) (*SauceConnect, error) {
	scPath, err := m.GetFilePath("SAUCE_CONNECT")
	if err != nil {
		return nil, errors.New("SauceConnect", err)
	}

	caps, err := m.ResolvedCapabilities()
	if err != nil {
		return nil, errors.New("SauceConnect", err)
	}

	id, _ := caps["tunnel-identifier"].(string)

	port, err := portpicker.PickUnusedPort()
	if err != nil {
		return nil, errors.New("SauceConnect", err)
	}

	cmd, err := service.NewCmd(
		"SauceConnect",
		d,
		scPath,
		false,
		nil,
		"--tunnel-identifier", id, "--se-port", fmt.Sprintf("%d", port))
	if err != nil {
		return nil, err
	}
	return &SauceConnect{Cmd: cmd, tunnelID: id}, nil
}

func (s *SauceConnect) Start(ctx context.Context) error {
	stdout, err := s.StdoutPipe()
	if err != nil {
		return err
	}

	done := make(chan interface{})

	go func() {
		b := bufio.NewReader(io.TeeReader(stdout, os.Stdout))
		defer close(done)
		for {
			line, err := b.ReadString('\n')
			if strings.Contains(line, readySignature) {
				s.mu.Lock()
				s.ready = true
				s.mu.Unlock()
				done <- nil
				break
			}
			if err != nil {
				return
			}
		}
		ioutil.ReadAll(b)
	}()

	if err := s.Cmd.Start(ctx); err != nil {
		return err
	}

	timeout := time.After(30 * time.Second)

	select {
	case <-done:
		return nil
	case <-timeout:
		return errors.New(s.Name(), fmt.Errorf("timeout after waiting 30 seconds for %q", readySignature))
	}
}

func (s *SauceConnect) Healthy(ctx context.Context) error {
	if err := s.Cmd.Healthy(ctx); err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if !s.ready {
		return errors.New(s.Name(), "is not ready")
	}
	return nil
}

func (s *SauceConnect) TunnelID() string {
	return s.tunnelID
}
