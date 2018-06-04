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

// Package portpicker provides methods for picking unused TCP ports.
package portpicker

import (
	"errors"
	"io"
	"net"
	"strconv"
	"sync"
)

var (
	mu           sync.Mutex
	claimedPorts = map[int]bool{}
)

// PickUnusedPort picks an unused TCP port.
func PickUnusedPort() (int, error) {
	mu.Lock()
	defer mu.Unlock()
	var listeners []io.Closer
	defer func() {
		for _, c := range listeners {
			c.Close()
		}
	}()

	for i := 0; i <= len(claimedPorts); i++ {
		l, err := net.Listen("tcp", ":0")
		if err != nil {
			return 0, err
		}
		listeners = append(listeners, l)

		_, p, err := net.SplitHostPort(l.Addr().String())
		if err != nil {
			return 0, err
		}

		port, err := strconv.Atoi(p)
		if err != nil {
			return 0, err
		}

		if !claimedPorts[port] {
			claimedPorts[port] = true
			return port, nil
		}
	}
	return 0, errors.New("unable to get a port")
}

// RecycleUnusedPort makes a port claimable by a call to PickUnusedPort.
func RecycleUnusedPort(port int) error {
	mu.Lock()
	defer mu.Unlock()
	delete(claimedPorts, port)
	return nil
}
