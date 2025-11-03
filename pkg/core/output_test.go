// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package core

import (
	"image"
	"image/color"
	"testing"
)

func TestDefaultOutputMode_ToOutputMode(t *testing.T) {
	tests := []struct {
		name string
		m    DefaultOutputMode
		want OutputMode
	}{
		{"DefaultOutputToDst", DefaultOutputToDst, OutputToDst},
		{"DefaultOutputToNewImage", DefaultOutputToNewImage, OutputToNewImage},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ToOutputMode(); got != tt.want {
				t.Errorf("ToOutputMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDst(t *testing.T) {
	if got := ToDst(); got != nil {
		t.Errorf("ToDst() = %v, want nil", got)
	}
}

func TestToImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	pt := image.Point{X: 10, Y: 20}
	out := ToImage(img, pt)

	// This test will fail with the current implementation because the mode is not set correctly.
	if out.OutputMode() != OutputToProvidedImage {
		t.Errorf("ToImage().OutputMode() = %v, want %v", out.OutputMode(), OutputToProvidedImage)
	}

	pImg, pPt := out.ProvidedImage()
	if pImg != img {
		t.Errorf("ToImage().ProvidedImage() image = %v, want %v", pImg, img)
	}
	if pPt != pt {
		t.Errorf("ToImage().ProvidedImage() point = %v, want %v", pPt, pt)
	}

	if out.ColorModel() != img.ColorModel() {
		t.Errorf("ToImage().ColorModel() = %v, want %v", out.ColorModel(), img.ColorModel())
	}
}

func TestToNewImage(t *testing.T) {
	out := ToNewImage()
	if out.OutputMode() != OutputToNewImage {
		t.Errorf("ToNewImage().OutputMode() = %v, want %v", out.OutputMode(), OutputToNewImage)
	}
	if out.ColorModel() != nil {
		t.Errorf("ToNewImage().ColorModel() should be nil, got %v", out.ColorModel())
	}
}

func TestToNewRGBAImage(t *testing.T) {
	out := ToNewRGBAImage()
	if out.OutputMode() != OutputToNewImage {
		t.Errorf("ToNewRGBAImage().OutputMode() = %v, want %v", out.OutputMode(), OutputToNewImage)
	}
	if out.ColorModel() != color.RGBAModel {
		t.Errorf("ToNewRGBAImage().ColorModel() = %v, want %v", out.ColorModel(), color.RGBAModel)
	}
}

func TestToNewNRGBAImage(t *testing.T) {
	out := ToNewNRGBAImage()
	if out.OutputMode() != OutputToNewImage {
		t.Errorf("ToNewNRGBAImage().OutputMode() = %v, want %v", out.OutputMode(), OutputToNewImage)
	}
	if out.ColorModel() != color.NRGBAModel {
		t.Errorf("ToNewNRGBAImage().ColorModel() = %v, want %v", out.ColorModel(), color.NRGBAModel)
	}
}

func TestOutputImageMethods(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	pt := image.Point{X: 5, Y: 5}
	m := color.RGBAModel
	mode := OutputToProvidedImage

	out := outputImage{
		image: img,
		pt:    pt,
		model: m,
		mode:  mode,
	}

	if out.OutputMode() != mode {
		t.Errorf("OutputMode() = %v, want %v", out.OutputMode(), mode)
	}

	pImg, pPt := out.ProvidedImage()
	if pImg != img {
		t.Errorf("ProvidedImage() image = %v, want %v", pImg, img)
	}
	if pPt != pt {
		t.Errorf("ProvidedImage() point = %v, want %v", pPt, pt)
	}

	if out.ColorModel() != m {
		t.Errorf("ColorModel() = %v, want %v", out.ColorModel(), m)
	}
}
