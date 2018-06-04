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

// Package resolver resolves WSL, WSLPORT, and WSLENV variables in capabilities.
package resolver

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"

	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/portpicker"
)

// A WSLResolver reolves WSL, WSLENV, and WSLPORT capabilities variables.
type WSLResolver struct {
	mu        sync.Mutex
	sessionID string
	ports     map[string]int
}

// New returns a new WSLResolver struct ready for use.
func New(sessionID string) *WSLResolver {
	return &WSLResolver{
		sessionID: sessionID,
		ports:     map[string]int{},
	}
}

// Resolve resolves a WSL, WSLENV, and WSLPORT capabilities variable.
func (w *WSLResolver) Resolve(p, n string) (string, error) {
	switch p {
	case "WSLPORT":
		w.mu.Lock()
		defer w.mu.Unlock()
		port, ok := w.ports[n]
		if ok {
			return strconv.Itoa(port), nil
		}
		port, err := portpicker.PickUnusedPort()
		if err != nil {
			return "", err
		}
		w.ports[n] = port
		return strconv.Itoa(port), nil
	case "WSLENV":
		env, ok := os.LookupEnv(n)
		if !ok {
			return "", fmt.Errorf("environment variable %s is undefined", n)
		}
		return env, nil
	case "WSL":
		switch n {
		case "SESSION_ID":
			return w.sessionID, nil
		case "HOST_IP":
			ips, err := net.LookupIP("localhost")
			if err != nil {
				return "", err
			}
			if len(ips) == 0 {
				return "", errors.New("no ip found for localhost")
			}
			return ips[0].String(), nil
		default:
			return "", fmt.Errorf("unknown variable WSL:%s", n)
		}
	default:
		return capabilities.NoOPResolver(p, n)
	}
}

// RecyclePorts returns the ports allocated by Resolve to the portpicker.
func (w *WSLResolver) RecyclePorts() error {
	var errs []error
	for _, port := range w.ports {
		if err := portpicker.RecycleUnusedPort(port); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return fmt.Errorf("errors recycling ports: %v", errs)
}
