// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package composite_test

import (
	"testing"

	"github.com/blazeroni/magpie/pkg/image/nrgba"
	"github.com/blazeroni/magpie/pkg/image/rgba"
)

func TestCompositeXor(t *testing.T) {
	testCases := []compositeTestCase{
		{
			colors:    Opaque1,
			expected:  c(0x00_00_00_00),
			tolerance: 0,
		},
		{
			colors:    TransparentSrc,
			expected:  c(0x00_80_ff_ff),
			tolerance: 0,
		},
		{
			colors:    TransparentDst,
			expected:  c(0x00_80_ff_ff),
			tolerance: 0,
		},
		{
			colors:    Translucent,
			expected:  c(0x80_60_a0_80),
			tolerance: 1,
		},
	}

	runCompositeTest(t, "Xor", testCases, nrgba.CompositeXor, rgba.CompositeXor)
}
