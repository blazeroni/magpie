// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package internal

import (
	"testing"
)

func TestMd255(t *testing.T) {
	tests := []struct {
		a, b, want uint32
	}{
		{0, 0, 0},
		{255, 255, 255},
		{128, 128, 64},
		{100, 200, 78},
	}

	for _, tt := range tests {
		if got := Md255(tt.a, tt.b); got != tt.want {
			t.Errorf("Md255(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		a, want uint32
	}{
		{0, 0},
		{255, 255},
		{64, 128},
		{128, 181},
	}

	for _, tt := range tests {
		if got := Sqrt(tt.a); got != tt.want {
			t.Errorf("Sqrt(%d) = %d, want %d", tt.a, got, tt.want)
		}
	}
}

func TestDiv255(t *testing.T) {
	tests := []struct {
		x, want uint32
	}{
		{0, 0},
		{255, 1},
		{127, 0},
		{128, 1},
		{65535, 257},
	}

	for _, tt := range tests {
		if got := Div255(tt.x); got != tt.want {
			t.Errorf("Div255(%d) = %d, want %d", tt.x, got, tt.want)
		}
	}
}

func TestUnpremultiply(t *testing.T) {
	tests := []struct {
		color, alpha, want uint32
	}{
		{0, 255, 0},
		{255, 255, 255},
		{128, 255, 128},
		{64, 128, 128},
		{0, 0, 0}, // alpha = 0 should not panic
	}

	for _, tt := range tests {
		if got := Unpremultiply(tt.color, tt.alpha); got != tt.want {
			var expected uint32
			if tt.alpha > 0 {
				expected = (tt.color*255 + tt.alpha/2) / tt.alpha
			}
			t.Errorf("Unpremultiply(%d, %d) = %d, want %d (calculated: %d)", tt.color, tt.alpha, got, tt.want, expected)
		}
	}
}
