// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package blend_test

import (
	"testing"

	"github.com/blazeroni/magpie/pkg/blend"
	"github.com/blazeroni/magpie/pkg/op"
)

func TestBlendOps(t *testing.T) {
	tests := []struct {
		name         string
		opFunc       func() op.BlendOp
		expectedMode op.BlendMode
	}{
		{
			name:         "Multiply",
			opFunc:       blend.Multiply,
			expectedMode: op.Multiply,
		},
		{
			name:         "Subtract",
			opFunc:       blend.Subtract,
			expectedMode: op.Subtract,
		},
		{
			name:         "Divide",
			opFunc:       blend.Divide,
			expectedMode: op.Divide,
		},
		{
			name:         "Screen",
			opFunc:       blend.Screen,
			expectedMode: op.Screen,
		},
		{
			name:         "Overlay",
			opFunc:       blend.Overlay,
			expectedMode: op.Overlay,
		},
		{
			name:         "Darken",
			opFunc:       blend.Darken,
			expectedMode: op.Darken,
		},
		{
			name:         "Lighten",
			opFunc:       blend.Lighten,
			expectedMode: op.Lighten,
		},
		{
			name:         "ColorDodge",
			opFunc:       blend.ColorDodge,
			expectedMode: op.ColorDodge,
		},
		{
			name:         "ColorBurn",
			opFunc:       blend.ColorBurn,
			expectedMode: op.ColorBurn,
		},
		{
			name:         "HardLight",
			opFunc:       blend.HardLight,
			expectedMode: op.HardLight,
		},
		{
			name:         "SoftLight",
			opFunc:       blend.SoftLight,
			expectedMode: op.SoftLight,
		},
		{
			name:         "Difference",
			opFunc:       blend.Difference,
			expectedMode: op.Difference,
		},
		{
			name:         "Exclusion",
			opFunc:       blend.Exclusion,
			expectedMode: op.Exclusion,
		},
		{
			name:         "LinearDodge",
			opFunc:       blend.LinearDodge,
			expectedMode: op.LinearDodge,
		},
		{
			name:         "LinearBurn",
			opFunc:       blend.LinearBurn,
			expectedMode: op.LinearBurn,
		},
		{
			name:         "VividLight",
			opFunc:       blend.VividLight,
			expectedMode: op.VividLight,
		},
		{
			name:         "LinearLight",
			opFunc:       blend.LinearLight,
			expectedMode: op.LinearLight,
		},
		{
			name:         "PinLight",
			opFunc:       blend.PinLight,
			expectedMode: op.PinLight,
		},
		{
			name:         "HardMix",
			opFunc:       blend.HardMix,
			expectedMode: op.HardMix,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.opFunc()
			if result.Mode != tt.expectedMode {
				t.Errorf("%s() Mode got = %v, want %v", tt.name, result.Mode, tt.expectedMode)
			}
			if result.Compositing != op.CompositeAll {
				t.Errorf("%s() Compositing got = %v, want %v", tt.name, result.Compositing, op.CompositeAll)
			}
		})
	}
}
