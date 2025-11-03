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

func TestBlendVividLight(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0xff_80_00_ff),
				op.CompositeBlendAndSrc: c(0xff_80_00_ff),
				op.CompositeBlendAndDst: c(0xff_80_00_ff),
				op.CompositeBlendOnly:   c(0xff_80_00_ff),
			},
			tolerance: 1,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x81_02_c0_ff),
				op.CompositeBlendAndSrc: c(0x81_02_c0_ff),
				op.CompositeBlendAndDst: c(0x81_02_c0_ff),
				op.CompositeBlendOnly:   c(0x81_02_c0_ff),
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
				op.CompositeAll:         c(0x80_41_aa_c0),
				op.CompositeBlendAndSrc: c(0x96_17_aa_80),
				op.CompositeBlendAndDst: c(0x6c_2c_bf_80),
				op.CompositeBlendOnly:   c(0x81_02_c0_40),
			},
			tolerance: 1,
		},
		{
			colors: testColors{
				name: "DarkSrc",
				dst:  c(0x80_80_80_ff),
				src:  c(0x40_40_40_ff),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0x02_02_02_ff),
			},
			tolerance: 0,
		},
		{
			colors: testColors{
				name: "LightSrc",
				dst:  c(0x80_80_80_ff),
				src:  c(0xc0_c0_c0_ff),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0xff_ff_ff_ff),
			},
			tolerance: 0,
		},
		{
			colors: testColors{
				name: "DarkSrc_Translucent",
				dst:  c(0x80_80_80_80),
				src:  c(0x40_40_40_80),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0x40_40_40_c0),
			},
			tolerance: 1,
		},
		{
			colors: testColors{
				name: "LightSrc_Translucent",
				dst:  c(0x80_80_80_80),
				src:  c(0xc0_c0_c0_80),
			},
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll: c(0xc0_c0_c0_c0),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "VividLight", testCases, nrgba.BlendVividLight, rgba.BlendVividLight)
}
