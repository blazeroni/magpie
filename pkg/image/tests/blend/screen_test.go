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

func TestBlendScreen(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0xff_c0_ff_ff),
				op.CompositeBlendAndSrc: c(0xff_c0_ff_ff),
				op.CompositeBlendAndDst: c(0xff_c0_ff_ff),
				op.CompositeBlendOnly:   c(0xff_c0_ff_ff),
			},
			tolerance: 1,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0xd0_a0_e0_ff),
				op.CompositeBlendAndSrc: c(0xd0_a0_e0_ff),
				op.CompositeBlendAndDst: c(0xd0_a0_e0_ff),
				op.CompositeBlendOnly:   c(0xd0_a0_e0_ff),
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
				op.CompositeAll:         c(0x9a_75_b5_c0),
				op.CompositeBlendAndSrc: c(0xca_80_bf_80),
				op.CompositeBlendAndDst: c(0x9f_95_d5_80),
				op.CompositeBlendOnly:   c(0xd0_a0_e0_40),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "Screen", testCases, nrgba.BlendScreen, rgba.BlendScreen)
}
