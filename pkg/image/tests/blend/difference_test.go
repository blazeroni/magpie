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

func TestBlendDifference(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0xff_00_ff_ff),
				op.CompositeBlendAndSrc: c(0xff_00_ff_ff),
				op.CompositeBlendAndDst: c(0xff_00_ff_ff),
				op.CompositeBlendOnly:   c(0xff_00_ff_ff),
			},
			tolerance: 0,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x80_40_40_ff),
				op.CompositeBlendAndSrc: c(0x80_40_40_ff),
				op.CompositeBlendAndDst: c(0x80_40_40_ff),
				op.CompositeBlendOnly:   c(0x80_40_40_ff),
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
				op.CompositeAll:         c(0x80_55_80_c0),
				op.CompositeBlendAndSrc: c(0x95_40_55_80),
				op.CompositeBlendAndDst: c(0x6a_55_6a_80),
				op.CompositeBlendOnly:   c(0x80_40_40_40),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "Difference", testCases, nrgba.BlendDifference, rgba.BlendDifference)
}
