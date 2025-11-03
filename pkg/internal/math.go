// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package internal

import "math"

var _md255 [256][256]uint8
var _sqrt [256]uint8
var _unpremult [256][256]uint8

func init() {
	for i := range 256 {
		for j := range 256 {
			_md255[i][j] = uint8((uint32(i)*uint32(j) + 127) / 255)
			if j > 0 {
				_unpremult[i][j] = uint8((uint32(i)*255 + uint32(j)/2) / uint32(j))
			}
		}
		_sqrt[i] = uint8(math.Sqrt(float64(i)/255)*255 + 0.5)
	}
}

// Md255 performs a multiplication then division by 255.
// Parameters a and b must be less than 256 due to using a lookup table.
// Approximates: (a * b / 255).
func Md255(a, b uint32) uint32 {
	return uint32(_md255[uint8(a)][uint8(b)])
}

// Sqrt performs a square root.
// Parameter x must be less than 256 due to using a lookup table.
// Approximates: sqrt(x).
func Sqrt(x uint32) uint32 {
	return uint32(_sqrt[uint8(x)])
}

// Div255 performs a division by 255.
// Approximates: (x / 255).
func Div255(x uint32) uint32 {
	return (x + (x >> 8) + 128) >> 8
}

// Unpremultiply performs an RGBA unpremultiplication.
// Color and alpha must be less than 256 due to using a lookup table.
// Approximates: (color / alpha).
func Unpremultiply(color, alpha uint32) uint32 {
	return uint32(_unpremult[uint8(color)][uint8(alpha)])
}
