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

// Package cropper provides an image cropping function.
package cropper

import (
	"errors"
	"image"
	"image/draw"
)

type subImageSupported interface {
	SubImage(r image.Rectangle) image.Image
}

// Crop crops an image to bounds.
// The cropped image may be a view on the original image or a completely distinct image.
func Crop(src image.Image, bounds image.Rectangle) (image.Image, error) {
	s, ok := src.(subImageSupported)
	if !ok {
		return nil, errors.New("crop only works with images that support subimage")
	}

	return s.SubImage(bounds), nil
}

// Blackout replaces a given rectangle in the image with a black rectangle.
// The function may modify the original image.
func Blackout(src image.Image, bounds image.Rectangle) (image.Image, error) {
	s, ok := src.(draw.Image)
	if !ok {
		return nil, errors.New("crop only works with images that implement draw.Image")
	}

	draw.Draw(s, bounds, image.Black, image.ZP, draw.Src)
	return s, nil
}
