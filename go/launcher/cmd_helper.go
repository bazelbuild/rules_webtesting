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

// Package cmdhelper provides functions to make working os/exec Command easier.
package cmdhelper

import (
	"os"
	"strings"
)

// UpdateEnv takes an environment array, an env var name, and an environment
// var value and updates the current definition of the env var or adds a new
// environment variable definition.
func UpdateEnv(env []string, name, value string) []string {
	prefix := name + "="

	for i := 0; i < len(env); {
		if strings.HasPrefix(env[i], prefix) {
			env = append(env[:i], env[i+1:]...)
		} else {
			i++
		}
	}

	return append(env, prefix+value)
}

// BulkUpdateEnv returns a new slice suitable for use with exec.Cmd.Env
// based on env but with environment variables for the keys in update
// added/changed to the values in update.
func BulkUpdateEnv(env []string, update map[string]string) []string {
	for name, value := range update {
		env = UpdateEnv(env, name, value)
	}
	return env
}

// IsTruthyEnv return true if name is set in the environment and has value other than
// 0 or false (case-insensitive).
func IsTruthyEnv(name string) bool {
	value := os.Getenv(name)
	return value != "" && value != "0" && strings.ToLower(value) != "false"
}
