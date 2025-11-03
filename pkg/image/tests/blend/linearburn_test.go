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

func TestBlendLinearBurn(t *testing.T) {
	testCases := []blendTestCase{
		{
			colors: Opaque1,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x00_01_00_ff),
				op.CompositeBlendAndSrc: c(0x00_01_00_ff),
				op.CompositeBlendAndDst: c(0x00_01_00_ff),
				op.CompositeBlendOnly:   c(0x00_01_00_ff),
			},
			tolerance: 1,
		},
		{
			colors: Opaque2,
			compositing: map[op.BlendCompositing]color.NRGBA{
				op.CompositeAll:         c(0x01_00_41_ff),
				op.CompositeBlendAndSrc: c(0x01_00_41_ff),
				op.CompositeBlendAndDst: c(0x01_00_41_ff),
				op.CompositeBlendOnly:   c(0x01_00_41_ff),
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
				op.CompositeAll:         c(0x55_40_80_c0),
				op.CompositeBlendAndSrc: c(0x41_15_56_80),
				op.CompositeBlendAndDst: c(0x17_2b_6c_80),
				op.CompositeBlendOnly:   c(0x01_00_41_40),
			},
			tolerance: 1,
		},
	}

	runBlendTest(t, "LinearBurn", testCases, nrgba.BlendLinearBurn, rgba.BlendLinearBurn)
}
