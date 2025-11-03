// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package magpie

import (
	"fmt"
	"image"

	"github.com/blazeroni/magpie/pkg/core"
	"github.com/blazeroni/magpie/pkg/internal"
	"github.com/blazeroni/magpie/pkg/op"
)

var defaultContext *context

func init() {
	defaultContext = &context{
		config: internal.DefaultConfig,
	}
}

// DefaultContext returns the default Magpie context.
// This context is used by the top-level Blend and Composite functions.
func DefaultContext() Context {
	return defaultContext
}

// SetDefaultContext sets the default configuration for the project.
// This will affect all subsequent operations that use the default context.
func SetDefaultContext(ctx Context) {
	internal.DefaultConfig = internal.NewConfig(
		ctx.PixelIterator(),
		ctx.DefaultOutputMode(),
		ctx.DefaultColorModel(),
	)
	defaultContext.config = internal.DefaultConfig
}

// Draw is the primary function for applying image manipulation operations (both blend and composite) using the default context.
// It follows similar semantics as Go's draw.Draw function. However, rather than always writing to the destination image,
// this function writes to a specified output which may be the destination image, a provided image, or a new image.
// It returns the modified output image.
func Draw(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, oper internal.Op, output core.Output) (image.Image, error) {
	switch opType := oper.(type) {
	case op.BlendOp:
		return Blend(dst, r, src, sp, opType, output)
	case op.CompositeOp:
		return Composite(dst, r, src, sp, opType, output)
	}
	return nil, fmt.Errorf("unsupported operation type: %T", oper)
}

// DrawToDst is a convenience function that applies an operation using the default context
// and writes the result directly to the destination image (dst).
// It is equivalent to calling Draw with ToDst() as the output.
func DrawToDst(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op internal.Op) (image.Image, error) {
	return Draw(dst, r, src, sp, op, ToDst())
}

// DrawToNewImage is a convenience function that applies an operation using the default context
// and writes the result to a newly created image.
// It is equivalent to calling Draw with ToNewImage() as the output.
func DrawToNewImage(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op internal.Op) (image.Image, error) {
	return Draw(dst, r, src, sp, op, ToNewImage())
}

// DrawToImage is a convenience function that applies an operation using the default context
// and writes the result to a user-provided image at a specified point.
// It is equivalent to calling Draw with ToImage(out, outPt) as the output.
func DrawToImage(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op internal.Op, out image.Image, outPt image.Point) (image.Image, error) {
	return Draw(dst, r, src, sp, op, ToImage(out, outPt))
}

// Blend performs a blend operation using the default context.
// See Context.Blend for details.
func Blend(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op op.BlendOp, output core.Output) (image.Image, error) {
	return defaultContext.Blend(dst, r, src, sp, op, output)
}

// Composite performs a composite operation using the default context.
// See Context.Composite for details.
func Composite(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op op.CompositeOp, output core.Output) (image.Image, error) {
	return defaultContext.Composite(dst, r, src, sp, op, output)
}

// ToDst returns an Output that writes to the destination image.
func ToDst() core.Output {
	return core.ToDst()
}

// ToNewImage returns an Output that creates a new image.
func ToNewImage() core.Output {
	return core.ToNewImage()
}

// ToImage returns an Output that writes to the provided image.
func ToImage(img image.Image, pt image.Point) core.Output {
	return core.ToImage(img, pt)
}
