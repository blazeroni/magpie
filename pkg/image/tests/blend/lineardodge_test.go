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

func TestBlendLinearDodge(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0xff_ff_ff_ff),
				op.CompositeBlendAndSrc: c(0xff_ff_ff_ff),
				op.CompositeBlendAndDst: c(0xff_ff_ff_ff),
				op.CompositeBlendOnly:   c(0xff_ff_ff_ff),
			},
			tolerance: 0,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0xff_c0_ff_ff),
				op.CompositeBlendAndSrc: c(0xff_c0_ff_ff),
				op.CompositeBlendAndDst: c(0xff_c0_ff_ff),
				op.CompositeBlendOnly:   c(0xff_c0_ff_ff),
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
				op.CompositeAll:         c(0xaa_80_c0_c0),
				op.CompositeBlendAndSrc: c(0xea_95_d5_80),
				op.CompositeBlendAndDst: c(0xbf_aa_ea_80),
				op.CompositeBlendOnly:   c(0xff_c0_ff_40),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "LinearDodge", testCases, nrgba.BlendLinearDodge, rgba.BlendLinearDodge)
}
