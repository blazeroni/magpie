// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package composite_test

import (
	"testing"

	"github.com/blazeroni/magpie/pkg/image/nrgba"
	"github.com/blazeroni/magpie/pkg/image/rgba"
)

func TestCompositeSourceOver(t *testing.T) {
	testCases := []compositeTestCase{
		{
			colors:    Opaque1,
			expected:  c(0xff_80_00_ff),
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
			expected:  c(0x96_55_95_bf),
			tolerance: 2,
		},
	}

	runCompositeTest(t, "SourceOver", testCases, nrgba.CompositeSourceOver, rgba.CompositeSourceOver)
}
