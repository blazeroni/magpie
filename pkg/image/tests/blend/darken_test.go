// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package blend_test

import (
	"image/color"
	"testing"

	"github.com/blazeroni/magpie/pkg/image/nrgba"
	"github.com/blazeroni/magpie/pkg/image/rgba"
	"github.com/blazeroni/magpie/pkg/op"
)

func TestBlendDarken(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x00_80_00_ff),
				op.CompositeBlendAndSrc: c(0x00_80_00_ff),
				op.CompositeBlendAndDst: c(0x00_80_00_ff),
				op.CompositeBlendOnly:   c(0x00_80_00_ff),
			},
			tolerance: 0,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x40_40_80_ff),
				op.CompositeBlendAndSrc: c(0x40_40_80_ff),
				op.CompositeBlendAndDst: c(0x40_40_80_ff),
				op.CompositeBlendOnly:   c(0x40_40_80_ff),
			},
			tolerance: 0,
		},
		{
			colors: TransparentSrc,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x00_80_ff_ff),
				op.CompositeBlendAndSrc: c(0x00_00_00_00),
				op.CompositeBlendAndDst: c(0x00_80_ff_ff),
				op.CompositeBlendOnly:   c(0x00_00_00_00),
			},
			tolerance: 0,
		},
		{
			colors: TransparentDst,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x00_80_ff_ff),
				op.CompositeBlendAndSrc: c(0x00_80_ff_ff),
				op.CompositeBlendAndDst: c(0x00_00_00_00),
				op.CompositeBlendOnly:   c(0x00_00_00_00),
			},
			tolerance: 0,
		},
		{
			colors: Translucent,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x6a_55_95_c0),
				op.CompositeBlendAndSrc: c(0x6a_40_80_80),
				op.CompositeBlendAndDst: c(0x40_55_95_80),
				op.CompositeBlendOnly:   c(0x40_40_80_40),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "Darken", testCases, nrgba.BlendDarken, rgba.BlendDarken)
}
