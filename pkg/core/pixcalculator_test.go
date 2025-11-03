// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package core

import (
	"image"
	"image/color"
	"testing"
)

func TestPixCalculatorNRGBA_PixelRange(t *testing.T) {
	// 1. Setup with gradients to check specific pixel values
	dst := image.NewNRGBA(image.Rect(0, 0, 20, 20))
	src := image.NewNRGBA(image.Rect(0, 0, 20, 20))
	out := image.NewNRGBA(image.Rect(0, 0, 20, 20))

	for y := range 20 {
		for x := range 20 {
			dst.SetNRGBA(x, y, color.NRGBA{R: uint8(x), G: uint8(y), A: 255})
			src.SetNRGBA(x, y, color.NRGBA{R: uint8(100 + x), G: uint8(100 + y), A: 255})
			out.SetNRGBA(x, y, color.NRGBA{R: uint8(200 + x), G: uint8(200 + y), A: 255})
		}
	}

	r := image.Rect(5, 5, 15, 15) // 10x10 rect in dst
	srcPt := image.Pt(2, 3)
	outPt := image.Pt(4, 6)

	// 2. Create calculator
	calc := NewPixCalculatorNRGBA(dst, r, src, srcPt, out, outPt)

	// 3. Test Rect() - assuming IntersectNRGBA is correct and tested elsewhere
	// For this setup, all images are large enough not to clip r.
	expectedRect := r
	if !calc.Rect().Eq(expectedRect) {
		t.Errorf("Expected rect %v, got %v", expectedRect, calc.Rect())
	}

	// 4. Test Result()
	if calc.Result() != out {
		t.Errorf("Expected result to be the 'out' image")
	}

	// 5. Test Calculate() for correct pixel mapping
	for yOffset := range calc.Rect().Dy() {
		dstRow, srcRow, outRow := calc.Calculate(yOffset)

		for xOffset := range calc.Rect().Dx() {
			// Expected coordinates
			expectedDstX, expectedDstY := r.Min.X+xOffset, r.Min.Y+yOffset
			expectedSrcX, expectedSrcY := srcPt.X+xOffset, srcPt.Y+yOffset
			expectedOutX, expectedOutY := outPt.X+xOffset, outPt.Y+yOffset

			// Check dst pixel
			pixOffset := xOffset * 4
			if dstRow[pixOffset] != uint8(expectedDstX) || dstRow[pixOffset+1] != uint8(expectedDstY) {
				t.Errorf("Dst pixel mismatch at row %d, col %d", yOffset, xOffset)
			}

			// Check src pixel
			if srcRow[pixOffset] != uint8(100+expectedSrcX) || srcRow[pixOffset+1] != uint8(100+expectedSrcY) {
				t.Errorf("Src pixel mismatch at row %d, col %d", yOffset, xOffset)
			}

			// Check out pixel
			if outRow[pixOffset] != uint8(200+expectedOutX) || outRow[pixOffset+1] != uint8(200+expectedOutY) {
				t.Errorf("Out pixel mismatch at row %d, col %d", yOffset, xOffset)
			}
		}
	}
}

func TestPixCalculatorRGBA_PixelRange(t *testing.T) {
	// 1. Setup with gradients to check specific pixel values
	dst := image.NewRGBA(image.Rect(0, 0, 20, 20))
	src := image.NewRGBA(image.Rect(0, 0, 20, 20))
	out := image.NewRGBA(image.Rect(0, 0, 20, 20))

	for y := range 20 {
		for x := range 20 {
			dst.SetRGBA(x, y, color.RGBA{R: uint8(x), G: uint8(y), A: 255})
			src.SetRGBA(x, y, color.RGBA{R: uint8(100 + x), G: uint8(100 + y), A: 255})
			out.SetRGBA(x, y, color.RGBA{R: uint8(200 + x), G: uint8(200 + y), A: 255})
		}
	}

	r := image.Rect(5, 5, 15, 15) // 10x10 rect in dst
	srcPt := image.Pt(2, 3)
	outPt := image.Pt(4, 6)

	// 2. Create calculator
	calc := NewPixCalculatorRGBA(dst, r, src, srcPt, out, outPt)

	// 3. Test Rect()
	expectedRect := r
	if !calc.Rect().Eq(expectedRect) {
		t.Errorf("Expected rect %v, got %v", expectedRect, calc.Rect())
	}

	// 4. Test Result()
	if calc.Result() != out {
		t.Errorf("Expected result to be the 'out' image")
	}

	// 5. Test Calculate() for correct pixel mapping
	for yOffset := range calc.Rect().Dy() {
		dstRow, srcRow, outRow := calc.Calculate(yOffset)

		for xOffset := range calc.Rect().Dx() {
			// Expected coordinates
			expectedDstX, expectedDstY := r.Min.X+xOffset, r.Min.Y+yOffset
			expectedSrcX, expectedSrcY := srcPt.X+xOffset, srcPt.Y+yOffset
			expectedOutX, expectedOutY := outPt.X+xOffset, outPt.Y+yOffset

			// Check dst pixel
			pixOffset := xOffset * 4
			if dstRow[pixOffset] != uint8(expectedDstX) || dstRow[pixOffset+1] != uint8(expectedDstY) {
				t.Errorf("Dst pixel mismatch at row %d, col %d", yOffset, xOffset)
			}

			// Check src pixel
			if srcRow[pixOffset] != uint8(100+expectedSrcX) || srcRow[pixOffset+1] != uint8(100+expectedSrcY) {
				t.Errorf("Src pixel mismatch at row %d, col %d", yOffset, xOffset)
			}

			// Check out pixel
			if outRow[pixOffset] != uint8(200+expectedOutX) || outRow[pixOffset+1] != uint8(200+expectedOutY) {
				t.Errorf("Out pixel mismatch at row %d, col %d", yOffset, xOffset)
			}
		}
	}
}
