// Binary merger takes multiple metadata files and merges them to produce a single
// metadata file.
package main

import (
	"flag"
	"log"

	"github.com/web_test_launcher/metadata/metadata"
)

const defaultRecordVideo = "never"

var (
	output       = flag.String("output", "", "output file for generated metadata")
	testLabel    = flag.String("test_label", "", "test label for generated metadata")
	browserLabel = flag.String("browser_label", "", "browser label for generated metadata")
)

func main() {
	flag.Parse()

	data := metadata.Metadata{
		TestLabel:    *testLabel,
		BrowserLabel: *browserLabel,
		RecordVideo:  defaultRecordVideo,
	}

	for _, file := range flag.Args() {
		m, err := metadata.FromFile(file)
		if err != nil {
			log.Fatalf("Error reading %s: %v", file, err)
		}
		data = metadata.Merge(data, m)
	}

	if err := data.ToFile(*output); err != nil {
		log.Fatalf("Error writing %s: %v", *output, err)
	}
}
