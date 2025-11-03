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

func TestBlendHardLight(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0xff_81_00_ff),
				op.CompositeBlendAndSrc: c(0xff_81_00_ff),
				op.CompositeBlendAndDst: c(0xff_81_00_ff),
				op.CompositeBlendOnly:   c(0xff_81_00_ff),
			},
			tolerance: 1,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0xa1_40_c1_ff),
				op.CompositeBlendAndSrc: c(0xa1_40_c1_ff),
				op.CompositeBlendAndDst: c(0xa1_40_c1_ff),
				op.CompositeBlendOnly:   c(0xa1_40_c1_ff),
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
				op.CompositeAll:         c(0x8a_55_aa_c0),
				op.CompositeBlendAndSrc: c(0xab_40_ab_80),
				op.CompositeBlendAndDst: c(0x81_55_c1_80),
				op.CompositeBlendOnly:   c(0xa1_40_c1_40),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "HardLight", testCases, nrgba.BlendHardLight, rgba.BlendHardLight)
}
