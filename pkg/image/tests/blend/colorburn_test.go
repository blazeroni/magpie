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

func TestBlendColorBurn(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x00_02_00_ff),
				op.CompositeBlendAndSrc: c(0x00_02_00_ff),
				op.CompositeBlendAndDst: c(0x00_02_00_ff),
				op.CompositeBlendOnly:   c(0x00_02_00_ff),
			},
			tolerance: 1,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x02_00_82_ff),
				op.CompositeBlendAndSrc: c(0x02_00_82_ff),
				op.CompositeBlendAndDst: c(0x02_00_82_ff),
				op.CompositeBlendOnly:   c(0x02_00_82_ff),
			},
			tolerance: 1,
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
				op.CompositeAll:         c(0x56_40_96_c0),
				op.CompositeBlendAndSrc: c(0x41_15_81_80),
				op.CompositeBlendAndDst: c(0x17_2b_96_80),
				op.CompositeBlendOnly:   c(0x02_00_82_40),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "ColorBurn", testCases, nrgba.BlendColorBurn, rgba.BlendColorBurn)
}
