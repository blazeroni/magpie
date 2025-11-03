// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package magpie

import (
	"image"
	"image/color"
	"testing"

	"github.com/blazeroni/magpie/pkg/core"
	"github.com/blazeroni/magpie/pkg/internal"
	"github.com/blazeroni/magpie/pkg/op"
)

// --- Mocks ---

type mockOp struct {
	applyRGBACalled  bool
	applyNRGBACalled bool
}

func (m *mockOp) IsValid() bool { return true }

func (m *mockOp) ApplyRGBA(_ core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	m.applyRGBACalled = true
	return calc.Result()
}

func (m *mockOp) ApplyNRGBA(_ core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	m.applyNRGBACalled = true
	return calc.Result()
}

// --- Tests ---

func TestContext_Draw2_InvalidOp(t *testing.T) {
	ctx := NewContext()
	dst := image.NewNRGBA(image.Rect(0, 0, 10, 10))
	src := image.NewNRGBA(image.Rect(0, 0, 10, 10))

	invalidOp := op.BlendOp{Mode: -1, Compositing: -1}
	_, err := ctx.Blend(dst, dst.Bounds(), src, image.Point{}, invalidOp, ToDst())
	if err == nil {
		t.Error("Blend with invalid op should return an error")
	}
}

func TestContext_Composite_InvalidOp(t *testing.T) {
	ctx := NewContext()
	dst := image.NewNRGBA(image.Rect(0, 0, 10, 10))
	src := image.NewNRGBA(image.Rect(0, 0, 10, 10))

	invalidOp := op.CompositeOp{Mode: -1}
	_, err := ctx.Composite(dst, dst.Bounds(), src, image.Point{}, invalidOp, ToDst())
	if err == nil {
		t.Error("Composite with invalid op should return an error")
	}
}

func TestContextGetters(t *testing.T) {
	ctx := NewContext(
		WithPixelIterator(4),
		WithDefaultToDst(),
		WithDefaultColorModelRGBA(),
	)

	if _, ok := ctx.PixelIterator().(core.ParallelPixelIterator); !ok {
		t.Error("PixelIterator() did not return a ParallelPixelIterator")
	}

	if ctx.DefaultOutputMode() != core.DefaultOutputToDst {
		t.Errorf("DefaultOutputMode() = %v, want %v", ctx.DefaultOutputMode(), core.DefaultOutputToDst)
	}

	if ctx.DefaultColorModel() != color.RGBAModel {
		t.Errorf("DefaultColorModel() = %v, want %v", ctx.DefaultColorModel(), color.RGBAModel)
	}
}

func TestAsNRGBA(t *testing.T) {
	t.Run("NRGBA input", func(t *testing.T) {
		inImg := image.NewNRGBA(image.Rect(0, 0, 1, 1))
		inImg.SetNRGBA(0, 0, color.NRGBA{R: 10, G: 20, B: 30, A: 40})
		outImg := AsNRGBA(inImg)
		if outImg != inImg {
			t.Error("AsNRGBA should return the same pointer for NRGBA input")
		}
	})

	t.Run("RGBA input", func(t *testing.T) {
		inImg := image.NewRGBA(image.Rect(0, 0, 1, 1))
		inImg.SetRGBA(0, 0, color.RGBA{R: 10, G: 20, B: 30, A: 255})
		outImg := AsNRGBA(inImg)
		// Check if pixel data is converted correctly (straight alpha)
		expected := color.NRGBA{R: 10, G: 20, B: 30, A: 255}
		if outImg.NRGBAAt(0, 0) != expected {
			t.Errorf("AsNRGBA conversion failed. Got %v, want %v", outImg.NRGBAAt(0, 0), expected)
		}
	})
}

func TestAsRGBA(t *testing.T) {
	t.Run("RGBA input", func(t *testing.T) {
		inImg := image.NewRGBA(image.Rect(0, 0, 1, 1))
		inImg.SetRGBA(0, 0, color.RGBA{R: 10, G: 20, B: 30, A: 40})
		outImg := AsRGBA(inImg)
		if outImg != inImg {
			t.Error("AsRGBA should return the same pointer for RGBA input")
		}
	})

	t.Run("NRGBA input", func(t *testing.T) {
		inImg := image.NewNRGBA(image.Rect(0, 0, 1, 1))
		// Use non-opaque color to test premultiplication
		inImg.SetNRGBA(0, 0, color.NRGBA{R: 128, G: 128, B: 128, A: 128})
		outImg := AsRGBA(inImg)
		// Check if pixel data is converted correctly (premultiplied alpha)
		// Expected: R=128 * 128/255 = 64, G=64, B=64, A=128
		expected := color.RGBA{R: 64, G: 64, B: 64, A: 128}
		if outImg.RGBAAt(0, 0) != expected {
			t.Errorf("AsRGBA conversion failed. Got %v, want %v", outImg.RGBAAt(0, 0), expected)
		}
	})
}

func TestColorModel(t *testing.T) {
	rgbaImg := image.NewRGBA(image.Rect(0, 0, 1, 1))
	nrgbaImg := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	unsupportedImg := image.NewGray(image.Rect(0, 0, 1, 1))

	// Set a non-default to test preference
	SetDefaultContext(NewContext(WithDefaultColorModelNRGBA()))

	tests := []struct {
		name string
		dst  image.Image
		src  image.Image
		out  core.Output
		want color.Model
	}{
		{"Output has model", rgbaImg, nrgbaImg, core.ToNewRGBAImage(), color.RGBAModel},
		{"Output nil, Dst supported", nrgbaImg, unsupportedImg, nil, color.NRGBAModel},
		{"Output/Dst unsupported, Src supported", unsupportedImg, rgbaImg, ToNewImage(), color.RGBAModel},
		{"All unsupported", unsupportedImg, unsupportedImg, nil, color.NRGBAModel}, // Falls back to default
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := colorModel(tt.dst, tt.src, tt.out)
			if got != tt.want && err == nil {
				t.Errorf("colorModel() = %v, want %v", got, tt.want)
			}
		})
	}
	// Reset default for other tests
	SetDefaultContext(NewContext())
}

func TestNewImage(t *testing.T) {
	rect := image.Rect(0, 0, 10, 10)

	t.Run("RGBA", func(t *testing.T) {
		img, pt := newImage(color.RGBAModel, rect)
		if _, ok := img.(*image.RGBA); !ok {
			t.Error("newImage did not return *image.RGBA")
		}
		if pt != rect.Min {
			t.Errorf("newImage point = %v, want %v", pt, rect.Min)
		}
	})

	t.Run("NRGBA", func(t *testing.T) {
		img, pt := newImage(color.NRGBAModel, rect)
		if _, ok := img.(*image.NRGBA); !ok {
			t.Error("newImage did not return *image.NRGBA")
		}
		if pt != rect.Min {
			t.Errorf("newImage point = %v, want %v", pt, rect.Min)
		}
	})

	t.Run("Unsupported", func(t *testing.T) {
		img, _ := newImage(color.GrayModel, rect)
		if img != nil {
			t.Error("newImage should return nil for unsupported models")
		}
	})
}

func TestDraw2(t *testing.T) {
	ctx := &context{
		config: internal.DefaultConfig,
	}
	rect := image.Rect(0, 0, 1, 1)
	src := image.NewNRGBA(rect)
	dst := image.NewNRGBA(rect)

	t.Run("NRGBA model", func(t *testing.T) {
		mock := &mockOp{}
		_, err := ctx.draw2(dst, rect, src, image.Point{}, mock, nil)
		if err != nil {
			t.Fatalf("draw2 failed: %v", err)
		}
		if !mock.applyNRGBACalled {
			t.Error("ApplyNRGBA was not called for NRGBA model")
		}
		if mock.applyRGBACalled {
			t.Error("ApplyRGBA was called for NRGBA model")
		}
	})

	t.Run("RGBA model", func(t *testing.T) {
		mock := &mockOp{}
		dstRGBA := image.NewRGBA(rect)
		_, err := ctx.draw2(dstRGBA, rect, src, image.Point{}, mock, nil)
		if err != nil {
			t.Fatalf("draw2 failed: %v", err)
		}
		if !mock.applyRGBACalled {
			t.Error("ApplyRGBA was not called for RGBA model")
		}
		if mock.applyNRGBACalled {
			t.Error("ApplyNRGBA was called for RGBA model")
		}
	})

	t.Run("OutputToDst", func(t *testing.T) {
		mock := &mockOp{}
		result, err := ctx.draw2(dst, rect, src, image.Point{}, mock, ToDst())
		if err != nil {
			t.Fatalf("draw2 failed: %v", err)
		}
		if result != dst {
			t.Error("OutputToDst should return the destination image")
		}
	})

	t.Run("OutputToNewImage", func(t *testing.T) {
		mock := &mockOp{}
		result, err := ctx.draw2(dst, rect, src, image.Point{}, mock, ToNewImage())
		if err != nil {
			t.Fatalf("draw2 failed: %v", err)
		}
		if result == dst {
			t.Error("OutputToNewImage should return a new image, not the destination")
		}
		if _, ok := result.(*image.NRGBA); !ok {
			t.Error("OutputToNewImage should have created an NRGBA image by default")
		}
	})

	t.Run("OutputToProvidedImage", func(t *testing.T) {
		mock := &mockOp{}
		providedImg := image.NewNRGBA(rect)
		result, err := ctx.draw2(dst, rect, src, image.Point{}, mock, ToImage(providedImg, image.Point{}))
		if err != nil {
			t.Fatalf("draw2 failed: %v", err)
		}
		if result != providedImg {
			t.Error("OutputToProvidedImage should return the provided image")
		}
	})

	t.Run("Unsupported color model", func(t *testing.T) {
		mock := &mockOp{}
		unsupportedImg := image.NewGray(rect)
		// Set default to unsupported as well
		SetDefaultContext(NewContext(func(c *context) {
			_ = c.config.SetDefaultColorModel(color.GrayModel)
		}))
		_, err := ctx.draw2(unsupportedImg, rect, unsupportedImg, image.Point{}, mock, nil)
		if err == nil {
			t.Error("draw2 should return an error for unsupported color models")
		}
		SetDefaultContext(NewContext()) // reset
	})
}
