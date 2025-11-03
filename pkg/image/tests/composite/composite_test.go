// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package composite_test

import (
	"fmt"
	"image"
	"image/color"
	"testing"

	"github.com/blazeroni/magpie/pkg/core"
)

var Opaque1 = testColors{
	name: "Opaque1",
	dst:  c(0x00_80_ff_ff),
	src:  c(0xff_80_00_ff),
}
var Opaque2 = testColors{
	name: "Opaque2",
	dst:  c(0x40_80_c0_ff),
	src:  c(0xc0_40_80_ff),
}
var TransparentSrc = testColors{
	name: "TransparentSrc",
	dst:  c(0x00_80_ff_ff),
	src:  c(0x00_00_00_00),
}
var TransparentDst = testColors{
	name: "TransparentDst",
	dst:  c(0x00_00_00_00),
	src:  c(0x00_80_ff_ff),
}
var Translucent = testColors{
	name: "Translucent",
	dst:  c(0x40_80_c0_80),
	src:  c(0xc0_40_80_80),
}

// compositeTestCase defines a standard test case for a composite operation.
// Colors are defined as NRGBA, and will be converted for RGBA tests.
type compositeTestCase struct {
	colors    testColors
	expected  color.NRGBA
	tolerance uint8
}

// nrgbaCompositeFunc is the signature for any NRGBA Composite<Mode> function.
type nrgbaCompositeFunc func(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA

// rgbaCompositeFunc is the signature for any RGBA Composite<Mode> function.
type rgbaCompositeFunc func(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA

// runCompositeTest provides a common runner for all composite mode tests.
func runCompositeTest(t *testing.T, mode string, testCases []compositeTestCase, nrgbaCf nrgbaCompositeFunc, rgbaCf rgbaCompositeFunc) {
	// NRGBA tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("NRGBA/%s/%s", mode, tc.colors.name), func(t *testing.T) {
			dst := newNRGBA(tc.colors.dst)
			src := newNRGBA(tc.colors.src)

			out := image.NewNRGBA(image.Rect(0, 0, 1, 1))

			pixIter := core.SerialPixelIterator{}
			calc := core.NewPixCalculatorNRGBA(dst, dst.Bounds(), src, src.Rect.Min, out, out.Rect.Min)
			nrgbaCf(pixIter, calc)

			actual := out.NRGBAAt(0, 0)

			if !colorsAlmostEqual(actual, tc.expected, tc.tolerance) {
				t.Errorf("Expected color %v [%s], but got %v [%s]", tc.expected, hexNRGBA(tc.expected), actual, hexNRGBA(actual))
			}
		})
	}

	// RGBA tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("RGBA/%s/%s", mode, tc.colors.name), func(t *testing.T) {
			dstRGBA := toRGBA(tc.colors.dst)
			srcRGBA := toRGBA(tc.colors.src)
			expectedRGBA := toRGBA(tc.expected)

			src := newRGBA(srcRGBA)
			dst := newRGBA(dstRGBA)

			out := image.NewRGBA(image.Rect(0, 0, 1, 1))

			pixIter := core.SerialPixelIterator{}
			calc := core.NewPixCalculatorRGBA(dst, dst.Bounds(), src, src.Rect.Min, out, out.Rect.Min)
			rgbaCf(pixIter, calc)

			actual := out.RGBAAt(0, 0)

			if !colorsAlmostEqual(actual, expectedRGBA, tc.tolerance) {
				t.Errorf("Expected color %v [%s] / %v [%s], but got %v [%s]; orig colors: %v [%s], %v [%s]",
					expectedRGBA, hex(expectedRGBA), tc.expected, hexNRGBA(tc.expected),
					actual, hex(actual), dstRGBA, hex(dstRGBA), srcRGBA, hex(srcRGBA))
			}
		})
	}
}

type testColors struct {
	name string
	dst  color.NRGBA
	src  color.NRGBA
}

func c(val uint32) color.NRGBA {
	return color.NRGBA{
		R: uint8(val >> 24),
		G: uint8(val >> 16),
		B: uint8(val >> 8),
		A: uint8(val),
	}
}

func toRGBA(c color.NRGBA) color.RGBA {
	return color.RGBAModel.Convert(c).(color.RGBA)
}

func hexNRGBA(c color.NRGBA) string {
	return fmt.Sprintf("0x%02x_%02x_%02x_%02x", c.R, c.G, c.B, c.A)
}

func hex(c color.Color) string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf("0x%02x_%02x_%02x_%02x", uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
}

func absDiff(a, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return b - a
}

func colorsAlmostEqual(c1, c2 color.Color, tolerance uint8) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	r1_8, g1_8, b1_8, a1_8 := uint8(r1>>8), uint8(g1>>8), uint8(b1>>8), uint8(a1>>8)
	r2_8, g2_8, b2_8, a2_8 := uint8(r2>>8), uint8(g2>>8), uint8(b2>>8), uint8(a2>>8)

	return absDiff(r1_8, r2_8) <= tolerance &&
		absDiff(g1_8, g2_8) <= tolerance &&
		absDiff(b1_8, b2_8) <= tolerance &&
		absDiff(a1_8, a2_8) <= tolerance
}

func newNRGBA(c color.NRGBA) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	img.SetNRGBA(0, 0, c)
	return img
}

func newRGBA(c color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.SetRGBA(0, 0, c)
	return img
}
