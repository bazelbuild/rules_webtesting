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
	"log"
	"os"

	"github.com/bazelbuild/rules_webtesting/go/wsl"
)

var (
	port          = flag.Int("port", 4444, "Port to start WSL on.")
	downloadRoot  = flag.String("download_root", "", "Directory served at /google/staticfile/.")
	uploadRoot    = flag.String("upload_root", "", "Directory to which files sent to /session/<id>/upload will be uploaded.")
	logFile       = flag.String("log_file", "", "File for WSL logs.")
	localHostname = flag.String("local_hostname", "localhost", "Name to use for localhost.")
)

func main() {
	flag.Parse()

	if *logFile != "" {
		out, err := os.Create(*logFile)
		if err != nil {
			log.Fatalf("unable to open log file: %v", err)
		}
		defer out.Close()
		log.SetOutput(out)
	}

	log.SetFlags(log.Flags() | log.Lmicroseconds)
	wsl.Run(*localHostname, *port, *downloadRoot, *uploadRoot)
}
