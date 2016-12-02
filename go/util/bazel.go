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

// Package bazel provides utilities for interacting with the surrounding Bazel environment.
package bazel

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// DefaultWorkspace is the name of the default Bazel workspace.
var DefaultWorkspace = "io_bazel_rules_webtesting"

// Runfile returns an absolute path to the specified file in the runfiles directory of the running target.
// It searches in both RunfilesPath() directory and in RunfilesPath()/TestWorkspace().
// Returns an error if unable to locate RunfilesPath() or if the file does not exist.
func Runfile(path string) (string, error) {
	runfiles, err := RunfilesPath()
	if err != nil {
		return "", err
	}

	filename := filepath.Join(runfiles, path)
	if _, err := os.Stat(filename); err == nil {
		// found at RunfilesPath()/path
		return filename, nil
	}

	filename = filepath.Join(runfiles, TestWorkspace(), path)
	if _, err := os.Stat(filename); err != nil {
		// not found at RunfilesPath()/TestWorkspace()/path
		return "", err
	}

	return filename, nil
}

// RunfilesPath return the path to the run files tree for this test.
// It returns an error if TEST_SRCDIR does not exist.
func RunfilesPath() (string, error) {
	const srcEnv = "TEST_SRCDIR"
	src, ok := os.LookupEnv(srcEnv)
	if !ok {
		return "", fmt.Errorf("environment variable %q is not defined, are you running with bazel test", srcEnv)
	}
	return src, nil
}

// NewTmpDir creates a new temporary directory in TestTmpDir().
func NewTmpDir(prefix string) (string, error) {
	return ioutil.TempDir(TestTmpDir(), prefix)
}

// TestTmpDir returns the path the Bazel test temp directory.
// If TEST_TMPDIR is not defined, it returns the OS default temp dir.
func TestTmpDir() string {
	const tmpEnv = "TEST_TMPDIR"
	tmp, ok := os.LookupEnv(tmpEnv)
	if !ok {
		return os.TempDir()
	}
	return tmp
}

// TestWorkspace returns the name of the Bazel workspace for this test.
// If TEST_WORKSPACE is not defined, it returns DefaultWorkspace.
func TestWorkspace() string {
	const wsEnv = "TEST_WORKSPACE"
	ws, ok := os.LookupEnv(wsEnv)
	if !ok {
		return DefaultWorkspace
	}
	return ws
}
