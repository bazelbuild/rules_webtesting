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

package cropper

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/util/bazel"
)

func TestCrop(t *testing.T) {
	p, err := bazel.Runfile("go/util/testdata/example.png")
	if err != nil {
		t.Fatal(err)
	}

	in, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}

	src, err := png.Decode(in)
	in.Close()
	if err != nil {
		t.Fatal(err)
	}

	dst, err := Crop(src, image.Rect(10, 20, 100, 90))
	if err != nil {
		t.Fatal(err)
	}

	tmpDir, err := bazel.NewTmpDir("crop_test")
	if err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(tmpDir, "cropped.png")

	out, err := os.Create(outPath)
	if err != nil {
		t.Fatal(err)
	}

	err = png.Encode(out, dst)
	out.Close()
	if err != nil {
		t.Fatal(err)
	}

	check, err := os.Open(outPath)
	if err != nil {
		t.Fatal(err)
	}

	config, err := png.DecodeConfig(check)
	check.Close()
	if err != nil {
		t.Fatal(err)
	}

	if config.Width != 90 || config.Height != 70 {
		t.Errorf("got size == %d, %d, expected 90, 70", config.Width, config.Height)
	}
}

func TestBlackout(t *testing.T) {
	p, err := bazel.Runfile("go/util/testdata/example.png")
	if err != nil {
		t.Fatal(err)
	}

	in, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}

	src, err := png.Decode(in)
	in.Close()
	if err != nil {
		t.Fatal(err)
	}

	dst, err := Blackout(src, image.Rect(10, 20, 200, 90))
	if err != nil {
		t.Fatal(err)
	}

	tmpDir, err := bazel.NewTmpDir("crop_test")
	if err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(tmpDir, "blackedout.png")

	log.Printf("Output file: %s", outPath)

	out, err := os.Create(outPath)
	if err != nil {
		t.Fatal(err)
	}

	err = png.Encode(out, dst)
	out.Close()
	if err != nil {
		t.Fatal(err)
	}

	check, err := os.Open(outPath)
	if err != nil {
		t.Fatal(err)
	}

	i, err := png.Decode(check)
	check.Close()
	if err != nil {
		t.Fatal(err)
	}

	img, ok := i.(*image.NRGBA)
	if !ok {
		t.Fatalf("got %T, expected image to be RGBA", i)
	}

	px := color.RGBAModel.Convert(img.At(20, 30))

	if px != color.RGBAModel.Convert(color.Black) {
		t.Fatalf("got %v, expected RGBA(R: 0, G: 0, B: 0, A: 255)", px)
	}
}
