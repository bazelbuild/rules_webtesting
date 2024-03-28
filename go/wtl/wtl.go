// Copyright 2016 Google LLC.
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

// Package wtl is used to manage the environment for web tests and start the underlying test.
package wtl

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/bazel"
	"github.com/bazelbuild/rules_webtesting/go/cmdhelper"
	"github.com/bazelbuild/rules_webtesting/go/errors"
	"github.com/bazelbuild/rules_webtesting/go/metadata"
	"github.com/bazelbuild/rules_webtesting/go/portpicker"
	"github.com/bazelbuild/rules_webtesting/go/wtl/diagnostics"
	"github.com/bazelbuild/rules_webtesting/go/wtl/environment"
	"github.com/bazelbuild/rules_webtesting/go/wtl/environment/external"
	"github.com/bazelbuild/rules_webtesting/go/wtl/environment/local"
	"github.com/bazelbuild/rules_webtesting/go/wtl/environment/sauce"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy/driverhub"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy/driverhub/drivermu"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy/driverhub/quithandler"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy/driverhub/scripttimeout"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy/healthz"
)

type envProvider func(m *metadata.Metadata, d diagnostics.Diagnostics) (environment.Env, error)

var envProviders = map[string]envProvider{}

func init() {
	// Configure Environments.
	RegisterEnvProviderFunc("external", external.NewEnv)
	RegisterEnvProviderFunc("local", local.NewEnv)
	RegisterEnvProviderFunc("sauce", sauce.NewEnv)

	// Configure HTTP Handlers
	proxy.AddHTTPHandlerProvider("/wd/hub/", driverhub.HTTPHandlerProvider)
	proxy.AddHTTPHandlerProvider("/healthz", healthz.HTTPHandlerProvider)

	// Configure WebDriver handlers.
	driverhub.HandlerProviderFunc(quithandler.ProviderFunc)
	driverhub.HandlerProviderFunc(scripttimeout.ProviderFunc)

	// drivermu should always be last.
	driverhub.HandlerProviderFunc(drivermu.ProviderFunc)
}

// RegisterEnvProviderFunc adds a new env provider.
func RegisterEnvProviderFunc(name string, p envProvider) {
	envProviders[name] = p
}

// Run runs the test.
func Run(d diagnostics.Diagnostics, testPath, mdPath string, httpPort, httpsPort, debuggerPort int) int {
	ctx := context.Background()

	testTerminated := make(chan os.Signal, 1)
	signal.Notify(testTerminated, syscall.SIGTERM, syscall.SIGINT)

	proxyStarted := make(chan error)
	envStarted := make(chan error)
	testFinished := make(chan int)
	envShutdown := make(chan error)

	mdFile, err := bazel.Runfile(mdPath)
	if err != nil {
		log.Print(err)
		return 127
	}

	md, err := metadata.FromFile(mdFile, nil)
	if err != nil {
		d.Severe(err)
		return 127
	}

	if debuggerPort != 0 {
		md.DebuggerPort = debuggerPort
	}

	testExe, err := bazel.Runfile(testPath)
	if err != nil {
		d.Severe(err)
		return 127
	}

	env, err := buildEnv(md, d)
	if err != nil {
		d.Severe(err)
		return 127
	}

	if httpPort == 0 {
		p, err := portpicker.PickUnusedPort()
		if err != nil {
			d.Severe(err)
			return 127
		}
		httpPort = p
		defer portpicker.RecycleUnusedPort(httpPort)
	}

	if httpsPort == 0 {
		p, err := portpicker.PickUnusedPort()
		if err != nil {
			d.Severe(err)
			return 127
		}
		httpsPort = p
		defer portpicker.RecycleUnusedPort(httpsPort)
	}

	p, err := proxy.New(env, md, d, httpPort, httpsPort)
	if err != nil {
		d.Severe(err)
		return 127
	}

	tmpDir, err := bazel.NewTmpDir("test")
	if err != nil {
		d.Severe(err)
		return 127
	}

	testCmd := exec.Command(testExe, flag.Args()...)

	envVars := map[string]string{
		"WEB_TEST_HTTP_SERVER":      fmt.Sprintf("http://%s", p.HTTPAddress),
		"WEB_TEST_WEBDRIVER_SERVER": fmt.Sprintf("http://%s/wd/hub", p.HTTPAddress),
		"TEST_TMPDIR":               tmpDir,
		"WEB_TEST_TMPDIR":           bazel.TestTmpDir(),
		"WEB_TEST_TARGET":           testPath,
	}

	if p.HTTPSAddress != "" {
		envVars["WEB_TEST_HTTPS_SERVER"] = fmt.Sprintf("https://%s", p.HTTPSAddress)
	}

	testCmd.Env = cmdhelper.BulkUpdateEnv(os.Environ(), envVars)

	testCmd.Stdout = os.Stdout
	testCmd.Stderr = os.Stderr
	testCmd.Stdin = os.Stdin

	go func() {
		envStarted <- env.SetUp(ctx)
	}()

	shutdownFunc := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// When the environment shutdowns or fails to shutdown, a message will be sent to envShutdown.
		go func() {
			var errors []error

			if err := p.Shutdown(ctx); err != nil {
				errors = append(errors, err)
			}
			if err := env.TearDown(ctx); err != nil {
				errors = append(errors, err)
			}
			switch len(errors) {
			case 0:
				envShutdown <- nil
			case 1:
				envShutdown <- errors[0]
			default:
				envShutdown <- fmt.Errorf("errors shutting down environment: %v", errors)
			}
		}()

		select {
		case <-testTerminated:
			d.Warning(errors.New("WTL", "test timed out during environment shutdown."))
		case <-ctx.Done():
			d.Warning(errors.New("WTL", "environment shutdown took longer than 5 seconds."))
		case err := <-envShutdown:
			if err != nil {
				d.Warning(err)
			}
		}
	}

	go func() {
		proxyStarted <- p.Start(ctx)
	}()

	for done := false; !done; {
		select {
		case <-testTerminated:
			return 0x8f
		case err := <-proxyStarted:
			if err != nil {
				d.Severe(err)
				return 127
			}
			done = true
		case err := <-envStarted:
			if err != nil {
				d.Severe(err)
				return 127
			}
			defer shutdownFunc()
		}
	}

	go func() {
		if status := testCmd.Run(); status != nil {
			log.Printf("test failed %v", status)
			if ee, ok := err.(*exec.ExitError); ok {
				if ws, ok := ee.Sys().(syscall.WaitStatus); ok {
					testFinished <- ws.ExitStatus()
					return
				}
			}
			testFinished <- 1
			return
		}
		testFinished <- 0
	}()

	for {
		select {
		case <-testTerminated:
			return 0x8f
		case err := <-envStarted:
			if err != nil {
				d.Severe(err)
				return 127
			}
			defer shutdownFunc()
		case status := <-testFinished:
			return status
		}
	}
}

func buildEnv(m *metadata.Metadata, d diagnostics.Diagnostics) (environment.Env, error) {
	p, ok := envProviders[m.Environment]
	if !ok {
		return nil, fmt.Errorf("unknown environment: %s", m.Environment)
	}
	return p(m, d)
}
