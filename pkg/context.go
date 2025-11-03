// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package magpie

import (
	"image"
	"image/color"

	"github.com/blazeroni/magpie/pkg/core"
	"github.com/blazeroni/magpie/pkg/internal"
	"github.com/blazeroni/magpie/pkg/op"
)

var _ core.Config = (*context)(nil)
var _ Context = (*context)(nil)

// Context defines the interface for image manipulation operations.
// It holds configuration settings that control how blending and compositing
// operations are performed.
type Context interface {
	core.Config

	// Blend applies a blend operation and optional compositing.
	// It follows similar semantics as draw.Draw. However, rather than always writing to the destination image,
	// this function writes to output which may be the destination image, a provided image, or a new image.
	// Returns the modified output image.
	Blend(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op op.BlendOp, out core.Output) (image.Image, error)

	// Composite applies a composite operation.
	// It follows similar semantics as draw.Draw. However, rather than always writing to the destination image,
	// this function writes to output which may be the destination image, a provided image, or a new image.
	// Returns the modified output image.
	Composite(dst image.Image, r image.Rectangle, src image.Image, sp image.Point, op op.CompositeOp, out core.Output) (image.Image, error)
}

// context implements the Context interface.
type context struct {
	config *internal.Config
}

// NewContext creates a new Context with the given options.
// The default configuration is used as a base, with the options overriding
// default values.
func NewContext(options ...func(*context)) Context {
	ctx := &context{
		config: internal.DefaultConfig,
	}
	for _, option := range options {
		option(ctx)
	}
	return ctx
}

// Option is a function that configures a Context.
type Option func(*context)

// WithPixelIterator sets the pixel iterator to be used by the context.
// The concurrency parameter specifies the number of goroutines to use for parallel processing.
func WithPixelIterator(concurrency int) Option {
	return func(p *context) {
		p.config.SetPixelIterator(NewPixelIterator(concurrency))
	}
}

// WithPixelIteratorInstance sets the pixel iterator to be used by the context.
// The provided instance will be used instead of creating a new one.
func WithPixelIteratorInstance(pixIter core.PixelIterator) Option {
	return func(p *context) {
		p.config.SetPixelIterator(pixIter)
	}
}

// WithDefaultToDst sets the default output mode to write to the destination image.
func WithDefaultToDst() Option {
	return func(p *context) {
		p.config.SetDefaultOutputMode(core.DefaultOutputToDst)
	}
}

// WithDefaultToNewImage sets the default output mode to create a new image.
func WithDefaultToNewImage() Option {
	return func(p *context) {
		p.config.SetDefaultOutputMode(core.DefaultOutputToNewImage)
	}
}

// WithDefaultColorModelNRGBA sets the default color model to NRGBA.
// The default color model is only used when none of the images
// use a supported color model.
func WithDefaultColorModelNRGBA() Option {
	return func(p *context) {
		_ = p.config.SetDefaultColorModel(color.NRGBAModel) //nolint:errcheck
	}
}

// WithDefaultColorModelRGBA sets the default color model to RGBA.
// The default color model is only used when none of the images
// use a supported color model.
func WithDefaultColorModelRGBA() Option {
	return func(p *context) {
		_ = p.config.SetDefaultColorModel(color.RGBAModel) //nolint:errcheck
	}
}
