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

// Binary launcher is used to manage the envrionment for web tests and start the underlying test.
package main

import (
	"flag"
	"os"

	"github.com/bazelbuild/rules_webtesting/go/wtl"
	"github.com/bazelbuild/rules_webtesting/go/wtl/diagnostics"
)

var (
	test             = flag.String("test", "", "Test to run.")
	metadataFileFlag = flag.String("metadata", "", "Metadata file for the browser.")
	debuggerPort     = flag.Int("debugger_port", 0, "Port to start WTL debugger on.")
	httpPort         = flag.Int("http_port", 0, "Port to start WTL HTTP Proxy on.")
	httpsPort        = flag.Int("https_port", 0, "Port to start WTL HTTPS Proxy on.")
)

func main() {
	flag.Parse()

	d := diagnostics.NoOP()

	status := wtl.Run(d, *test, *metadataFileFlag, *httpPort, *httpsPort, *debuggerPort)

	d.Close()
	os.Exit(status)
}
