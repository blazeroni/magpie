// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package core

import (
	"image"
)

var _ PixCalculator[*image.NRGBA] = (*pixCalculator[*image.NRGBA])(nil)

type PixCalculator[T image.Image] interface {
	PixRowCalculator
	Result() T
}

type PixRowCalculator interface {
	Rect() image.Rectangle
	Calculate(row int) (dst, src, out []uint8)
}

type pixCalculator[T image.Image] struct {
	out                             T
	dstPix, srcPix, outPix          []uint8
	dstStride, srcStride, outStride int
	rect                            image.Rectangle
	srcStart, dstStart, outStart    int
	bytesPerPixel                   int
}

func (p *pixCalculator[T]) Result() T {
	return p.out
}

func NewPixCalculatorNRGBA(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, srcPt image.Point, out *image.NRGBA, outPt image.Point) PixCalculator[*image.NRGBA] {
	bounds := IntersectNRGBA(dst, r, src, srcPt, out, outPt)
	return &pixCalculator[*image.NRGBA]{
		out:           out,
		dstPix:        dst.Pix,
		srcPix:        src.Pix,
		outPix:        out.Pix,
		rect:          bounds,
		dstStart:      r.Min.Y*dst.Stride + (r.Min.X)*4,
		srcStart:      srcPt.Y*src.Stride + srcPt.X*4,
		outStart:      outPt.Y*out.Stride + outPt.X*4,
		dstStride:     dst.Stride,
		srcStride:     src.Stride,
		outStride:     out.Stride,
		bytesPerPixel: 4,
	}
}

func NewPixCalculatorRGBA(dst *image.RGBA, r image.Rectangle, src *image.RGBA, srcPt image.Point, out *image.RGBA, outPt image.Point) PixCalculator[*image.RGBA] {
	bounds := IntersectRGBA(dst, r, src, srcPt, out, outPt)
	return &pixCalculator[*image.RGBA]{
		out:           out,
		dstPix:        dst.Pix,
		srcPix:        src.Pix,
		outPix:        out.Pix,
		rect:          bounds,
		dstStart:      r.Min.Y*dst.Stride + (r.Min.X)*4,
		srcStart:      srcPt.Y*src.Stride + srcPt.X*4,
		outStart:      outPt.Y*out.Stride + outPt.X*4,
		dstStride:     dst.Stride,
		srcStride:     src.Stride,
		outStride:     out.Stride,
		bytesPerPixel: 4,
	}
}

func (p *pixCalculator[T]) Calculate(row int) ([]uint8, []uint8, []uint8) {
	// row is in rect coordinates and needs to be translated to dst, src, and out coordinates
	di := p.dstStart + (row * p.dstStride)
	si := p.srcStart + (row * p.srcStride)
	oi := p.outStart + (row * p.outStride)
	rowLength := p.rect.Dx() * p.bytesPerPixel
	return p.dstPix[di : di+rowLength],
		p.srcPix[si : si+rowLength],
		p.outPix[oi : oi+rowLength]
}

func (p *pixCalculator[T]) Rect() image.Rectangle {
	return p.rect
}
