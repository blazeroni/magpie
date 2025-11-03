// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package magpie

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/blazeroni/magpie/pkg/core"
	"github.com/blazeroni/magpie/pkg/internal"
	"github.com/blazeroni/magpie/pkg/op"
)

// PixelIterator returns the pixel iterator used by the context.
func (ctx *context) PixelIterator() PixelIterator {
	return ctx.config.PixelIterator()
}

// DefaultOutputMode returns the default output mode for operations.
func (ctx *context) DefaultOutputMode() core.DefaultOutputMode {
	return ctx.config.DefaultOutputMode()
}

// DefaultColorModel returns the default color model used for operations
// when the provided images are not a supported color model.
func (ctx *context) DefaultColorModel() color.Model {
	return ctx.config.DefaultColorModel()
}

func (ctx *context) Blend(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op op.BlendOp, output core.Output) (image.Image, error) {
	return ctx.draw2(dst, r, src, sp, op, output)
}

func (ctx *context) Composite(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op op.CompositeOp, output core.Output) (image.Image, error) {
	return ctx.draw2(dst, r, src, sp, op, output)
}

// draw2 is the internal drawing function that handles both blend and composite operations.
func (ctx *context) draw2(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op internal.Op, output core.Output) (image.Image, error) {
	if !op.IsValid() {
		return nil, fmt.Errorf("invalid operation")
	}

	// Decide on a color model
	clrModel, err := colorModel(dst, src, output)
	if err != nil {
		return nil, err
	}
	var outputMode core.OutputMode
	if output != nil {
		outputMode = output.OutputMode()
	} else {
		outputMode = ctx.DefaultOutputMode().ToOutputMode()
	}

	var out image.Image
	var outPt image.Point
	switch outputMode {
	case core.OutputToDst:
		out = dst
		outPt = r.Min
	case core.OutputToNewImage:
		out, outPt = newImage(clrModel, r)
	case core.OutputToProvidedImage:
		out, outPt = output.ProvidedImage()
	default:
		return nil, fmt.Errorf("unsupported output mode %v", outputMode)
	}

	switch clrModel {
	case color.RGBAModel:
		dstRGBA, srcRGBA, outRGBA := AsRGBA(dst), AsRGBA(src), AsRGBA(out)
		calc := core.NewPixCalculatorRGBA(dstRGBA, r, srcRGBA, sp, outRGBA, outPt)
		return op.ApplyRGBA(ctx.PixelIterator(), calc), nil
	case color.NRGBAModel:
		dstNRGBA, srcNRGBA, outNRGBA := AsNRGBA(dst), AsNRGBA(src), AsNRGBA(out)
		calc := core.NewPixCalculatorNRGBA(dstNRGBA, r, srcNRGBA, sp, outNRGBA, outPt)
		return op.ApplyNRGBA(ctx.PixelIterator(), calc), nil
	default:
		return nil, fmt.Errorf("unsupported color model %v", clrModel)
	}
}

// newImage creates a new image with the given color model and bounds.
func newImage(colorModel color.Model, bounds image.Rectangle) (image.Image, image.Point) {
	switch colorModel {
	case color.NRGBAModel:
		return image.NewNRGBA(bounds), bounds.Min
	case color.RGBAModel:
		return image.NewRGBA(bounds), bounds.Min
	default:
		return nil, image.Point{}
	}
}

// colorModel determines the best color model to use for an operation
// based on the destination, source, and output images.
// Preference is given first to the output image, then destination, source,
// and finally to the default color model.  Only supported color models are considered.
// See core.IsColorModelSupported for details on which color models are supported.
func colorModel(dst image.Image, src image.Image, out core.Output) (color.Model, error) {
	dstModel, srcModel := dst.ColorModel(), src.ColorModel()
	var outModel color.Model
	if out != nil {
		outModel = out.ColorModel()
	}
	switch {
	case core.IsColorModelSupported(outModel):
		return outModel, nil
	case core.IsColorModelSupported(dstModel):
		return dstModel, nil
	case core.IsColorModelSupported(srcModel):
		return srcModel, nil
	default:
		return nil, fmt.Errorf("unsupported color model")
	}
}

// AsNRGBA returns the image as an *image.NRGBA. If the image is already in
// this format, it is returned directly. Otherwise, a new NRGBA image is created
// and the content is drawn onto it.
func AsNRGBA(img image.Image) *image.NRGBA {
	if rgba, ok := img.(*image.NRGBA); ok {
		return rgba
	}
	bounds := img.Bounds()
	out := image.NewNRGBA(bounds)
	draw.Draw(out, bounds, img, bounds.Min, draw.Src)
	return out
}

// AsRGBA returns the image as an *image.RGBA. If the image is already in
// this format, it is returned directly. Otherwise, a new RGBA image is created
// and the content is drawn onto it.
func AsRGBA(img image.Image) *image.RGBA {
	if rgba, ok := img.(*image.RGBA); ok {
		return rgba
	}
	bounds := img.Bounds()
	out := image.NewRGBA(bounds)
	draw.Draw(out, bounds, img, bounds.Min, draw.Over)
	return out
}
