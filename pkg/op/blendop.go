// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package op

import (
	"image"

	"github.com/blazeroni/magpie/pkg/core"
	"github.com/blazeroni/magpie/pkg/image/nrgba"
	"github.com/blazeroni/magpie/pkg/image/rgba"
	"github.com/blazeroni/magpie/pkg/internal"
)

var _ internal.Op = (*BlendOp)(nil)

type BlendOp struct {
	Mode        BlendMode
	Compositing BlendCompositing
}

type BlendMode int
type BlendCompositing = internal.BlendCompositing

const (
	ColorBurn BlendMode = iota
	ColorDodge
	Darken
	Difference
	Divide
	Exclusion
	HardLight
	HardMix
	Lighten
	LinearBurn
	LinearDodge
	LinearLight
	Multiply
	Overlay
	PinLight
	Screen
	SoftLight
	Subtract
	VividLight

	_maxBlendMode
)

// Aliases for BlendModes defined above.
const (
	Add = LinearDodge
)

const (
	CompositeAll         = internal.CompositeAll
	CompositeBlendOnly   = internal.CompositeBlendOnly
	CompositeBlendAndDst = internal.CompositeBlendAndDst
	CompositeBlendAndSrc = internal.CompositeBlendAndSrc
)

func (o BlendOp) ApplyNRGBA(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	f := nrgbaBlendFuncs[o.Mode]
	if f != nil {
		f(pixIter, calc, o.Compositing)
	}
	return calc.Result()
}

func (o BlendOp) ApplyRGBA(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	f := rgbaBlendFuncs[o.Mode]
	if f != nil {
		f(pixIter, calc, o.Compositing)
	}
	return calc.Result()
}

func (o BlendOp) IsValid() bool {
	if o.Mode < 0 || o.Mode >= _maxBlendMode {
		return false
	}
	switch o.Compositing {
	case CompositeAll, CompositeBlendOnly, CompositeBlendAndDst, CompositeBlendAndSrc:
		return true
	default:
		return false
	}
}

var nrgbaBlendFuncs = []func(core.PixelIterator, core.PixCalculator[*image.NRGBA], BlendCompositing) *image.NRGBA{
	ColorBurn:   nrgba.BlendColorBurn,
	ColorDodge:  nrgba.BlendColorDodge,
	Darken:      nrgba.BlendDarken,
	Difference:  nrgba.BlendDifference,
	Divide:      nrgba.BlendDivide,
	Exclusion:   nrgba.BlendExclusion,
	HardLight:   nrgba.BlendHardLight,
	HardMix:     nrgba.BlendHardMix,
	Lighten:     nrgba.BlendLighten,
	LinearBurn:  nrgba.BlendLinearBurn,
	LinearDodge: nrgba.BlendLinearDodge,
	LinearLight: nrgba.BlendLinearLight,
	Multiply:    nrgba.BlendMultiply,
	Overlay:     nrgba.BlendOverlay,
	PinLight:    nrgba.BlendPinLight,
	Screen:      nrgba.BlendScreen,
	SoftLight:   nrgba.BlendSoftLight,
	Subtract:    nrgba.BlendSubtract,
	VividLight:  nrgba.BlendVividLight,
}

var rgbaBlendFuncs = []func(core.PixelIterator, core.PixCalculator[*image.RGBA], BlendCompositing) *image.RGBA{
	ColorBurn:   rgba.BlendColorBurn,
	ColorDodge:  rgba.BlendColorDodge,
	Darken:      rgba.BlendDarken,
	Difference:  rgba.BlendDifference,
	Divide:      rgba.BlendDivide,
	Exclusion:   rgba.BlendExclusion,
	HardLight:   rgba.BlendHardLight,
	HardMix:     rgba.BlendHardMix,
	Lighten:     rgba.BlendLighten,
	LinearBurn:  rgba.BlendLinearBurn,
	LinearDodge: rgba.BlendLinearDodge,
	LinearLight: rgba.BlendLinearLight,
	Multiply:    rgba.BlendMultiply,
	Overlay:     rgba.BlendOverlay,
	PinLight:    rgba.BlendPinLight,
	Screen:      rgba.BlendScreen,
	SoftLight:   rgba.BlendSoftLight,
	Subtract:    rgba.BlendSubtract,
	VividLight:  rgba.BlendVividLight,
}
