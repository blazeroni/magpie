// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package core

import (
	"image"
	"image/color"
)

type DefaultOutputMode int

const (
	DefaultOutputToDst DefaultOutputMode = iota
	DefaultOutputToNewImage
)

type OutputMode int

const (
	OutputToDst OutputMode = iota
	OutputToNewImage
	OutputToProvidedImage
)

func (m DefaultOutputMode) ToOutputMode() OutputMode {
	var mode OutputMode
	switch m {
	case DefaultOutputToDst:
		mode = OutputToDst
	case DefaultOutputToNewImage:
		mode = OutputToNewImage
	}
	return mode
}

type Output interface {
	OutputMode() OutputMode
	ColorModel() color.Model
	ProvidedImage() (image.Image, image.Point)
}

type outputImage struct {
	image image.Image
	model color.Model
	pt    image.Point
	mode  OutputMode
}

func (out outputImage) OutputMode() OutputMode {
	return out.mode
}

func (out outputImage) ProvidedImage() (image.Image, image.Point) {
	return out.image, out.pt
}

func (out outputImage) ColorModel() color.Model {
	return out.model
}

func ToDst() Output {
	return zeroValue[Output]()
}

func ToImage(img image.Image, pt image.Point) Output {
	return outputImage{
		image: img,
		pt:    pt,
		model: img.ColorModel(),
		mode:  OutputToProvidedImage,
	}
}

func ToNewImage() Output {
	return outputImage{
		mode: OutputToNewImage,
	}
}

func ToNewRGBAImage() Output {
	return outputImage{
		mode:  OutputToNewImage,
		model: color.RGBAModel,
	}
}

func ToNewNRGBAImage() Output {
	return outputImage{
		mode:  OutputToNewImage,
		model: color.NRGBAModel,
	}
}

func zeroValue[T any]() T {
	var zero T
	return zero
}
