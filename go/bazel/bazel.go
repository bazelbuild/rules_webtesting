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
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bazelbuild/rules_webtesting/go/cmdhelper"
)

// DefaultWorkspace is the name of the default Bazel workspace.
var DefaultWorkspace = "io_bazel_rules_webtesting"

// Runfile returns an absolute path to the specified file in the runfiles directory of the running target.
// It searches the current working directory, RunfilesPath() directory, and RunfilesPath()/TestWorkspace().
// Returns an error if unable to locate RunfilesPath() or if the file does not exist.
func Runfile(path string) (string, error) {
	if _, err := os.Stat(path); err == nil {
		// absolute path or found in current working directory
		return path, nil
	}

	if cmdhelper.IsTruthyEnv("RUNFILES_MANIFEST_ONLY") {
		return runfileInManifest(path)
	}

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
	if _, err := os.Stat(filename); err == nil {
		// found at RunfilesPath()/TestWorkspace()/path
		return filename, nil
	}

	return runfileInManifest(path)
}

// RunfilesPath return the path to the run files tree for this test.
// It returns an error if TEST_SRCDIR does not exist.
func RunfilesPath() (string, error) {
	const srcEnv = "TEST_SRCDIR"
	if src, ok := os.LookupEnv(srcEnv); ok {
		return src, nil
	}
	return "", fmt.Errorf("environment variable %q is not defined, are you running with bazel test", srcEnv)
}

// NewTmpDir creates a new temporary directory in TestTmpDir().
func NewTmpDir(prefix string) (string, error) {
	return ioutil.TempDir(TestTmpDir(), prefix)
}

// TestTmpDir returns the path the Bazel test temp directory.
// If TEST_TMPDIR is not defined, it returns the OS default temp dir.
func TestTmpDir() string {
	if tmp, ok := os.LookupEnv("TEST_TMPDIR"); ok {
		return tmp
	}
	return os.TempDir()
}

// TestWorkspace returns the name of the Bazel workspace for this test.
// If TEST_WORKSPACE is not defined, it returns DefaultWorkspace.
func TestWorkspace() string {
	if ws, ok := os.LookupEnv("TEST_WORKSPACE"); ok {
		return ws
	}
	return DefaultWorkspace
}

// RunfilesManifest returns the runfiles MANIFEST file.
func RunfilesManifest() (string, error) {
	mp, ok := os.LookupEnv("RUNFILES_MANIFEST_FILE")
	if !ok {
		return "", errors.New("environment variable RUNFILES_MANIFEST_FILE is not defined, are you running with bazel test")
	}

	_, err := os.Stat(mp)

	return mp, err
}

func runfileInManifest(path string) (string, error) {
	// TODO(DrMarcII): Consider support for finding for TestWorkspace()/path in the manifest.
	mp, err := RunfilesManifest()
	if err != nil {
		return "", err
	}

	f, err := os.Open(mp)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokens := strings.SplitN(scanner.Text(), " ", 2)

		if len(tokens) != 2 {
			continue
		}

		if tokens[0] == path {
			if _, err := os.Stat(tokens[1]); err != nil {
				continue
			}
			return tokens[1], nil
		}

		if strings.HasPrefix(path, tokens[0]) {
			rel, err := filepath.Rel(tokens[0], path)
			if err != nil {
				continue
			}
			f := filepath.Join(tokens[1], rel)
			if _, err := os.Stat(f); err != nil {
				continue
			}
			return f, nil
		}
	}

	return "", fmt.Errorf("cannot find runfile %q in manifest %q", path, mp)
}
