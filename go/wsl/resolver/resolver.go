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

	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/portpicker"
)

// Resolver returns a capabilities.Resolver for WSL, WSLPORT, and WSLENV capabilities variables.
func New(sessionID string) capabilities.Resolver {
	ports := map[string]string{}

	return func(p, n string) (string, error) {
		switch p {
		case "WSLPORT":
			portStr, ok := ports[n]
			if ok {
				return portStr, nil
			}
			port, err := portpicker.PickUnusedPort()
			if err != nil {
				return "", err
			}
			portStr = strconv.Itoa(port)
			ports[n] = portStr
			return portStr, nil
		case "WSLENV":
			env, ok := os.LookupEnv(n)
			if !ok {
				return "", fmt.Errorf("environment variable %s is undefined", n)
			}
			return env, nil
		case "WSL":
			switch n {
			case "SESSION_ID":
				return sessionID, nil
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
}
