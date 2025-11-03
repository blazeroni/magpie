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

func TestBlendPinLight(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0xff_80_00_ff),
				op.CompositeBlendAndSrc: c(0xff_80_00_ff),
				op.CompositeBlendAndDst: c(0xff_80_00_ff),
				op.CompositeBlendOnly:   c(0xff_80_00_ff),
			},
			tolerance: 0,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x81_80_c0_ff),
				op.CompositeBlendAndSrc: c(0x81_80_c0_ff),
				op.CompositeBlendAndDst: c(0x81_80_c0_ff),
				op.CompositeBlendOnly:   c(0x81_80_c0_ff),
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
				op.CompositeAll:         c(0x80_6a_aa_c0),
				op.CompositeBlendAndSrc: c(0x96_6a_aa_80),
				op.CompositeBlendAndDst: c(0x6c_80_bf_80),
				op.CompositeBlendOnly:   c(0x81_80_c0_40),
			},
			tolerance: 1,
		},
		{
			colors: testColors{
				name: "DarkSrc_DstDarker",
				dst:  c(0x40_40_40_ff),
				src:  c(0x60_60_60_ff),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0x40_40_40_ff),
			},
		},
		{
			colors: testColors{
				name: "DarkSrc_DstLighter",
				dst:  c(0xe0_e0_e0_ff),
				src:  c(0x60_60_60_ff),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0xc0_c0_c0_ff),
			},
		},
		{
			colors: testColors{
				name: "LightSrc_DstLighter",
				dst:  c(0xe0_e0_e0_ff),
				src:  c(0xa0_a0_a0_ff),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0xe0_e0_e0_ff),
			},
		},
		{
			colors: testColors{
				name: "LightSrc_DstDarker",
				dst:  c(0x20_20_20_ff),
				src:  c(0xa0_a0_a0_ff),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0x41_41_41_ff),
			},
			tolerance: 1,
		},
		{
			colors: testColors{
				name: "DarkSrc_DstDarker_Translucent",
				dst:  c(0x40_40_40_80),
				src:  c(0x60_60_60_80),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0x4b_4b_4b_c0),
			},
			tolerance: 1,
		},
		{
			colors: testColors{
				name: "DarkSrc_DstLighter_Translucent",
				dst:  c(0xe0_e0_e0_80),
				src:  c(0x60_60_60_80),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0xab_ab_ab_c0),
			},
			tolerance: 1,
		},
		{
			colors: testColors{
				name: "LightSrc_DstLighter_Translucent",
				dst:  c(0xe0_e0_e0_80),
				src:  c(0xa0_a0_a0_80),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0xca_ca_ca_c0),
			},
			tolerance: 1,
		},
		{
			colors: testColors{
				name: "LightSrc_DstDarker_Translucent",
				dst:  c(0x20_20_20_80),
				src:  c(0xa0_a0_a0_80),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0x56_56_56_c0),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "PinLight", testCases, nrgba.BlendPinLight, rgba.BlendPinLight)
}
