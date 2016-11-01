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
// Binary merger takes multiple metadata files and merges them to produce a single
// metadata file.
package main

import (
	"flag"
	"log"

	"github.com/bazelbuild/rules_webtesting/go/metadata/metadata"
)

var output = flag.String("output", "", "output file for generated metadata")

func main() {
	flag.Parse()

	data := &metadata.Metadata{}

	for _, file := range flag.Args() {
		m, err := metadata.FromFile(file)
		if err != nil {
			log.Fatalf("Error reading %q: %v", file, err)
		}
		data, err = metadata.Merge(data, m)
		if err != nil {
			log.Fatalf("Error merging file %q: %v", file, err)
		}
	}

	if err := data.ToFile(*output); err != nil {
		log.Fatalf("Error writing %s: %v", *output, err)
	}
}
