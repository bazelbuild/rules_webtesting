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

// WSL (Webdriver Server Light) is a lightweight replacement for Selenium Server.
package main

import (
	"flag"

	"github.com/bazelbuild/rules_webtesting/go/wsl"
)

var port = flag.Int("port", 4444, "Start WSL on given port")

func main() {
	flag.Parse()

	wsl.Run(*port)
}
