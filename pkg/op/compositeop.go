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

var _ internal.Op = (*CompositeOp)(nil)

type CompositeOp struct {
	Mode CompositeMode
}

func (o CompositeOp) ApplyNRGBA(p core.PixelIterator, c core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	f := nrgbaCompositeFuncs[o.Mode]
	if f != nil {
		f(p, c)
	}
	return c.Result()
}

func (o CompositeOp) ApplyRGBA(p core.PixelIterator, c core.PixCalculator[*image.RGBA]) *image.RGBA {
	f := rgbaCompositeFuncs[o.Mode]
	if f != nil {
		f(p, c)
	}
	return c.Result()
}

type CompositeMode int

const (
	Clear CompositeMode = iota
	Source
	SourceOver
	SourceIn
	SourceOut
	SourceAtop
	Destination
	DestinationOver
	DestinationIn
	DestinationOut
	DestinationAtop
	Xor

	_maxCompositeMode
)

func (o CompositeOp) IsValid() bool {
	return o.Mode >= 0 && o.Mode < _maxCompositeMode
}

var nrgbaCompositeFuncs = []func(core.PixelIterator, core.PixCalculator[*image.NRGBA]) *image.NRGBA{
	Clear:           nrgba.CompositeClear,
	Source:          nrgba.CompositeSource,
	SourceOver:      nrgba.CompositeSourceOver,
	SourceIn:        nrgba.CompositeSourceIn,
	SourceOut:       nrgba.CompositeSourceOut,
	SourceAtop:      nrgba.CompositeSourceAtop,
	Destination:     nrgba.CompositeDestination,
	DestinationOver: nrgba.CompositeDestinationOver,
	DestinationIn:   nrgba.CompositeDestinationIn,
	DestinationOut:  nrgba.CompositeDestinationOut,
	DestinationAtop: nrgba.CompositeDestinationAtop,
	Xor:             nrgba.CompositeXor,
}

var rgbaCompositeFuncs = []func(core.PixelIterator, core.PixCalculator[*image.RGBA]) *image.RGBA{
	Clear:           rgba.CompositeClear,
	Source:          rgba.CompositeSource,
	SourceOver:      rgba.CompositeSourceOver,
	SourceIn:        rgba.CompositeSourceIn,
	SourceOut:       rgba.CompositeSourceOut,
	SourceAtop:      rgba.CompositeSourceAtop,
	Destination:     rgba.CompositeDestination,
	DestinationOver: rgba.CompositeDestinationOver,
	DestinationIn:   rgba.CompositeDestinationIn,
	DestinationOut:  rgba.CompositeDestinationOut,
	DestinationAtop: rgba.CompositeDestinationAtop,
	Xor:             rgba.CompositeXor,
}
