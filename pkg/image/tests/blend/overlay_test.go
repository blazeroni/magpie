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

func TestBlendOverlay(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x00_81_ff_ff),
				op.CompositeBlendAndSrc: c(0x00_81_ff_ff),
				op.CompositeBlendAndDst: c(0x00_81_ff_ff),
				op.CompositeBlendOnly:   c(0x00_81_ff_ff),
			},
			tolerance: 1,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x60_41_c1_ff),
				op.CompositeBlendAndSrc: c(0x60_41_c1_ff),
				op.CompositeBlendAndDst: c(0x60_41_c1_ff),
				op.CompositeBlendOnly:   c(0x60_41_c1_ff),
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
				op.CompositeAll:         c(0x75_55_aa_c0),
				op.CompositeBlendAndSrc: c(0x80_41_ab_80),
				op.CompositeBlendAndDst: c(0x55_56_c1_80),
				op.CompositeBlendOnly:   c(0x60_41_c1_40),
			},
			tolerance: 1,
		},
		{
			colors: testColors{
				name: "Dst_Dark",
				dst:  c(0x40_40_40_ff),
				src:  c(0xb0_b0_b0_ff),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0x58_58_58_ff),
			},
			tolerance: 1,
		},
		{
			colors: testColors{
				name: "Dst_Light",
				dst:  c(0xd0_d0_d0_ff),
				src:  c(0xb0_b0_b0_ff),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0xe2_e2_e2_ff),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "Overlay", testCases, nrgba.BlendOverlay, rgba.BlendOverlay)
}
