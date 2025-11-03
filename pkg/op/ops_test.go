// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package op

import (
	"image"
	"testing"

	"github.com/blazeroni/magpie/pkg/core"
)

// --- Mocks ---

type mockPixelIterator struct{}

func (m mockPixelIterator) Iterate(core.PixRowCalculator, func(dst, src, out []uint8)) {}

type mockNRGBACalc struct{}

func (m mockNRGBACalc) Result() *image.NRGBA                      { return nil }
func (m mockNRGBACalc) Rect() image.Rectangle                     { return image.Rectangle{} }
func (m mockNRGBACalc) Calculate(int) ([]uint8, []uint8, []uint8) { return nil, nil, nil }

type mockRGBACalc struct{}

func (m mockRGBACalc) Result() *image.RGBA                       { return nil }
func (m mockRGBACalc) Rect() image.Rectangle                     { return image.Rectangle{} }
func (m mockRGBACalc) Calculate(int) ([]uint8, []uint8, []uint8) { return nil, nil, nil }

// --- Tests ---

func TestCompositeMode_IsValid(t *testing.T) {
	tests := []struct {
		name        string
		mode        CompositeMode
		wantIsValid bool
	}{
		{
			name:        "Valid CompositeMode - Clear",
			mode:        Clear,
			wantIsValid: true,
		},
		{
			name:        "Valid CompositeMode - SourceOver",
			mode:        SourceOver,
			wantIsValid: true,
		},
		{
			name:        "Invalid CompositeMode - out of bounds",
			mode:        _maxCompositeMode,
			wantIsValid: false,
		},
		{
			name:        "Invalid CompositeMode - negative",
			mode:        -1,
			wantIsValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := CompositeOp{Mode: tt.mode}
			if got := op.IsValid(); got != tt.wantIsValid {
				t.Errorf("CompositeMode.IsValid() = %v, want %v", got, tt.wantIsValid)
			}
		})
	}
}

func TestCompositeOp(t *testing.T) {
	originalNRGBA := nrgbaCompositeFuncs
	originalRGBA := rgbaCompositeFuncs
	defer func() {
		nrgbaCompositeFuncs = originalNRGBA
		rgbaCompositeFuncs = originalRGBA
	}()

	mockIterator := mockPixelIterator{}
	mockNRGBACalc := mockNRGBACalc{}
	mockRGBACalc := mockRGBACalc{}

	mode := SourceOver
	nrgbaCalled := false
	rgbaCalled := false

	// Replace with mocks
	nrgbaCompositeFuncs = make([]func(core.PixelIterator, core.PixCalculator[*image.NRGBA]) *image.NRGBA, len(originalNRGBA))
	rgbaCompositeFuncs = make([]func(core.PixelIterator, core.PixCalculator[*image.RGBA]) *image.RGBA, len(originalRGBA))

	nrgbaCompositeFuncs[mode] = func(core.PixelIterator, core.PixCalculator[*image.NRGBA]) *image.NRGBA {
		nrgbaCalled = true
		return nil
	}
	rgbaCompositeFuncs[mode] = func(core.PixelIterator, core.PixCalculator[*image.RGBA]) *image.RGBA {
		rgbaCalled = true
		return nil
	}

	op := CompositeOp{Mode: mode}
	op.ApplyNRGBA(mockIterator, mockNRGBACalc)
	op.ApplyRGBA(mockIterator, mockRGBACalc)

	if !nrgbaCalled {
		t.Errorf("NRGBA composite function for %v was not called", mode)
	}
	if !rgbaCalled {
		t.Errorf("RGBA composite function for %v was not called", mode)
	}
}

func TestBlendOp_IsValid(t *testing.T) {
	tests := []struct {
		name        string
		mode        BlendMode
		compositing BlendCompositing
		wantIsValid bool
	}{
		{
			name:        "Valid BlendOp - Multiply and CompositeAll",
			mode:        Multiply,
			compositing: CompositeAll,
			wantIsValid: true,
		},
		{
			name:        "Valid BlendOp - Screen and CompositeBlendOnly",
			mode:        Screen,
			compositing: CompositeBlendOnly,
			wantIsValid: true,
		},
		{
			name:        "Invalid BlendMode - out of bounds",
			mode:        _maxBlendMode,
			compositing: CompositeAll,
			wantIsValid: false,
		},
		{
			name:        "Invalid BlendMode - negative",
			mode:        -1,
			compositing: CompositeAll,
			wantIsValid: false,
		},
		{
			name:        "Invalid Compositing - out of bounds",
			mode:        Multiply,
			compositing: 999,
			wantIsValid: false,
		},
		{
			name:        "Invalid Compositing - in bounds, but does not exist",
			mode:        Multiply,
			compositing: 5,
			wantIsValid: false,
		},
		{
			name:        "Invalid Compositing - negative",
			mode:        Multiply,
			compositing: -1,
			wantIsValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := BlendOp{Mode: tt.mode, Compositing: tt.compositing}
			if got := op.IsValid(); got != tt.wantIsValid {
				t.Errorf("BlendOp.IsValid() = %v, want %v", got, tt.wantIsValid)
			}
		})
	}
}

func TestBlendOp(t *testing.T) {
	originalNRGBA := nrgbaBlendFuncs
	originalRGBA := rgbaBlendFuncs
	defer func() {
		nrgbaBlendFuncs = originalNRGBA
		rgbaBlendFuncs = originalRGBA
	}()

	mockIterator := mockPixelIterator{}
	mockNRGBACalc := mockNRGBACalc{}
	mockRGBACalc := mockRGBACalc{}

	mode := Multiply
	nrgbaCalled := false
	rgbaCalled := false

	// Replace with mocks
	nrgbaBlendFuncs = make([]func(core.PixelIterator, core.PixCalculator[*image.NRGBA], BlendCompositing) *image.NRGBA, len(originalNRGBA))
	rgbaBlendFuncs = make([]func(core.PixelIterator, core.PixCalculator[*image.RGBA], BlendCompositing) *image.RGBA, len(originalRGBA))

	nrgbaBlendFuncs[mode] = func(core.PixelIterator, core.PixCalculator[*image.NRGBA], BlendCompositing) *image.NRGBA {
		nrgbaCalled = true
		return nil
	}
	rgbaBlendFuncs[mode] = func(core.PixelIterator, core.PixCalculator[*image.RGBA], BlendCompositing) *image.RGBA {
		rgbaCalled = true
		return nil
	}

	op := BlendOp{Mode: mode}
	op.ApplyNRGBA(mockIterator, mockNRGBACalc)
	op.ApplyRGBA(mockIterator, mockRGBACalc)

	if !nrgbaCalled {
		t.Errorf("NRGBA blend function for %v was not called", mode)
	}
	if !rgbaCalled {
		t.Errorf("RGBA blend function for %v was not called", mode)
	}
}
