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

// Package selenium provides a service.Server for managing an instance of Selenium Standalone server.
package selenium

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/services/service"
	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
)

type Selenium struct {
	*service.Server
}

func NewSelenium(m *metadata.Metadata, xvfb bool) (*Selenium, error) {
	seleniumPath, err := m.GetFilePath("SELENIUM")
	if err != nil {
		return nil, errors.New("SeleniumServer", err)
	}
	log.Printf("Selenium found at at: %s", seleniumPath)

	javaPath, err := m.GetFilePath("JAVA")
	if err != nil {
		log.Print("did not find provided java")
		javaPath, err = exec.LookPath("java")
		if err != nil {
			return nil, errors.New("SeleniumServer", "unable to find a suitable java runtime environment")
		}
	}
	log.Printf("Java found at at: %s", javaPath)

	args := []string{"-jar", seleniumPath}

	if chromedriverPath, err := m.GetFilePath("CHROMEDRIVER"); err == nil {
		log.Printf("ChromeDriver found at: %q", chromedriverPath)
		args = append(args, fmt.Sprintf("--jvm_flag=-Dwebdriver.chrome.driver=%s", chromedriverPath))
	}
	if geckodriverPath, err := m.GetFilePath("GECKODRIVER"); err == nil {
		log.Printf("GeckoDriver found at: %q", geckodriverPath)
		args = append(args, fmt.Sprintf("--jvm_flag=-Dwebdriver.gecko.driver=%s", geckodriverPath))
	}
	if firefoxPath, err := m.GetFilePath("FIREFOX"); err == nil {
		log.Printf("Firefox found at: %q", firefoxPath)
		args = append(args, fmt.Sprintf("--jvm_flag=-Dwebdriver.firefox.bin=%s", firefoxPath))
	}
	args = append(args, "-port", "{port}")
	server, err := service.NewServer(
		"SeleniumServer",
		javaPath,
		"http://%s/wd/hub/status",
		xvfb,
		60*time.Second,
		nil, args...)
	if err != nil {
		return nil, err
	}
	return &Selenium{server}, nil
}

func (s *Selenium) Address() string {
	return fmt.Sprintf("http://%s/wd/hub/", s.Server.Address())
}
