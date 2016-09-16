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
//
////////////////////////////////////////////////////////////////////////////////
//
package bazel

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	runfileEnv = "TEST_SRCDIR"
	tmpDirEnv  = "TEST_TMPDIR"
)

// Runfile returns an absolute path to the specified file in the runfiles directory of the running target.
// Returns an error if TEST_SRCDIR is not set or if the file does not exist.
func Runfile(path string) (string, error) {
	runfileDir, ok := os.LookupEnv(runfileEnv)
	if !ok {
		return "", fmt.Errorf("environment variable %q is not defined, are you running with bazel test", runfileEnv)
	}
	filename := filepath.Join(runfileDir, path)
	if _, err := os.Stat(filename); err != nil {
		return "", err
	}
	return filename, nil
}

func NewTmpDir(prefix string) (string, error) {
	testTmpDir, err := TestTmpDir()
	if err != nil {
		return "", err
	}
	return ioutil.TempDir(testTmpDir, prefix)
}

func TestTmpDir() (string, error) {
	testTmpDir, ok := os.LookupEnv(tmpDirEnv)
	if !ok {
		return "", fmt.Errorf("environment variable %q is not defined, are you running with bazel test", tmpDirEnv)
	}
	return testTmpDir, nil
}
