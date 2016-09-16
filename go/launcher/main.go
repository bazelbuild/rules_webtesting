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
package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/bazelbuild/rules_web/go/launcher/cmdhelper"
	"github.com/bazelbuild/rules_web/go/launcher/environments/environment"
	"github.com/bazelbuild/rules_web/go/launcher/environments/external"
	"github.com/bazelbuild/rules_web/go/launcher/environments/native"
	"github.com/bazelbuild/rules_web/go/launcher/environments/phantomjs"
	"github.com/bazelbuild/rules_web/go/launcher/proxy/proxy"
	"github.com/bazelbuild/rules_web/go/metadata/metadata"
	"github.com/bazelbuild/rules_web/go/util/bazel"
)

var (
	test             = flag.String("test", "", "Test script to launch")
	metadataFileFlag = flag.String("metadata", "", "metadata file")
)

func main() {
	flag.Parse()

	os.Exit(run())
}

func run() int {
	metadataFile, err := bazel.Runfile(*metadataFileFlag)
	if err != nil {
		log.Printf("Error locating metadata file: %v", err)
		return 127
	}

	m, err := metadata.FromFile(metadataFile)
	if err != nil {
		log.Printf("Error reading metadata file: %v", err)
		return 127
	}

	env, err := buildEnv(m)
	if err != nil {
		log.Printf("Error building environment: %v", err)
		return 127
	}

	if err := env.SetUp(context.Background()); err != nil {
		log.Printf("Error setting up environment: %v", err)
		return 127
	}

	defer func() {
		if err := env.TearDown(context.Background()); err != nil {
			log.Printf("Error tearing down environment: %v", err)
		}
	}()

	p, err := proxy.New(env)
	if err != nil {
		log.Printf("Error creating proxy: %v", err)
		return 127
	}

	if err := p.Start(context.Background()); err != nil {
		log.Printf("Error starting proxy: %v", err)
		return 127
	}

	testExe, err := bazel.Runfile(*test)
	if err != nil {
		log.Printf("unable to find %s", *test)
		return 127
	}

	// Temporary directory where WEB_TEST infrastructure writes it tmp files.
	webTestTmpDir, _ := bazel.TestTmpDir()

	// Make an isolated temp directory for the test.
	tmpDir, err := ioutil.TempDir(webTestTmpDir, "test")
	if err != nil {
		log.Printf("Unable to create new temp dir for test: %v", err)
		// Fallback to previous value.
		tmpDir = webTestTmpDir
	} else {
		// cleanup tmpDir after test is done.
		defer os.RemoveAll(tmpDir)
	}

	testCmd := exec.Command(testExe, flag.Args()...)
	testCmd.Env = cmdhelper.BulkUpdateEnv(os.Environ(), map[string]string{
		"WEB_TEST_BROWSER_METADATA": *metadataFileFlag,
		"REMOTE_WEBDRIVER_SERVER":   p.Address,
		"TEST_TMPDIR":               tmpDir,
		"WEB_TEST_TMPDIR":           webTestTmpDir,
		"WEB_TEST_TARGET":           *test,
	})
	testCmd.Stdout = os.Stdout
	testCmd.Stderr = os.Stderr
	testCmd.Stdin = os.Stdin

	if status := testCmd.Run(); status != nil {
		log.Printf("test failed %v", status)
		if ee, ok := err.(*exec.ExitError); ok {
			if ws, ok := ee.Sys().(syscall.WaitStatus); ok {
				return ws.ExitStatus()
			}
		}
		return 1
	}
	return 0
}

func buildEnv(m *metadata.Metadata) (environment.Env, error) {
	switch m.Environment {
	case "external":
		return external.NewEnv(m)
	case "native":
		return native.NewEnv(m)
	case "phantomjs":
		return phantomjs.NewEnv(m)
	}
	return nil, fmt.Errorf("unknown environment: %s", m.Environment)
}
