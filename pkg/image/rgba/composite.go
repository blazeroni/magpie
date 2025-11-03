// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package rgba

import (
	"image"

	"github.com/blazeroni/magpie/pkg/core"
)

// CompositeSourceOver performs a "Source Over" compositing operation on premultiplied RGBA images.
func CompositeSourceOver(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sR, sG, sB, sA := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2]), uint32(src[i+3])
			dR, dG, dB, dA := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2]), uint32(dst[i+3])

			if sA == 255 {
				// Source is fully opaque, it covers destination
				out[i], out[i+1], out[i+2], out[i+3] = src[i], src[i+1], src[i+2], src[i+3]
				continue
			}
			if sA == 0 {
				// Source is fully transparent, output is destination
				out[i], out[i+1], out[i+2], out[i+3] = dst[i], dst[i+1], dst[i+2], dst[i+3]
				continue
			}

			invSA := 255 - sA
			oA := sA + md255(dA, invSA)
			oR := sR + md255(dR, invSA)
			oG := sG + md255(dG, invSA)
			oB := sB + md255(dB, invSA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeSourceIn performs a "Source In" compositing operation on premultiplied RGBA images.
func CompositeSourceIn(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sR, sG, sB, sA := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2]), uint32(src[i+3])
			dA := uint32(dst[i+3])

			oA := md255(sA, dA)
			oR := md255(sR, dA)
			oG := md255(sG, dA)
			oB := md255(sB, dA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeSourceAtop performs a "Source Atop" compositing operation on premultiplied RGBA images.
func CompositeSourceAtop(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sR, sG, sB, sA := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2]), uint32(src[i+3])
			dR, dG, dB, dA := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2]), uint32(dst[i+3])

			invSA := 255 - sA
			oA := dA
			oR := md255(sR, dA) + md255(dR, invSA)
			oG := md255(sG, dA) + md255(dG, invSA)
			oB := md255(sB, dA) + md255(dB, invSA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeSourceOut performs a "Source Out" compositing operation on premultiplied RGBA images.
func CompositeSourceOut(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sR, sG, sB, sA := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2]), uint32(src[i+3])
			dA := uint32(dst[i+3])

			invDA := 255 - dA
			oA := md255(sA, invDA)
			oR := md255(sR, invDA)
			oG := md255(sG, invDA)
			oB := md255(sB, invDA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeSource performs a "Source" (or "Copy") compositing operation.
func CompositeSource(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(_, src, out []uint8) {
		copy(out, src)
	})
}

// CompositeDestinationOver performs a "Destination Over" compositing operation on premultiplied RGBA images.
func CompositeDestinationOver(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sR, sG, sB, sA := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2]), uint32(src[i+3])
			dR, dG, dB, dA := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2]), uint32(dst[i+3])

			invDA := 255 - dA
			oA := dA + md255(sA, invDA)
			oR := dR + md255(sR, invDA)
			oG := dG + md255(sG, invDA)
			oB := dB + md255(sB, invDA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeDestinationIn performs a "Destination In" compositing operation on premultiplied RGBA images.
func CompositeDestinationIn(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			dR, dG, dB, dA := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2]), uint32(dst[i+3])
			sA := uint32(src[i+3])

			oA := md255(dA, sA)
			oR := md255(dR, sA)
			oG := md255(dG, sA)
			oB := md255(dB, sA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeDestinationAtop performs a "Destination Atop" compositing operation on premultiplied RGBA images.
func CompositeDestinationAtop(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sR, sG, sB, sA := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2]), uint32(src[i+3])
			dR, dG, dB, dA := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2]), uint32(dst[i+3])

			invDA := 255 - dA
			oA := sA
			oR := md255(dR, sA) + md255(sR, invDA)
			oG := md255(dG, sA) + md255(sG, invDA)
			oB := md255(dB, sA) + md255(sB, invDA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeDestinationOut performs a "Destination Out" compositing operation on premultiplied RGBA images.
func CompositeDestinationOut(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			dR, dG, dB, dA := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2]), uint32(dst[i+3])
			sA := uint32(src[i+3])

			invSA := 255 - sA
			oA := md255(dA, invSA)
			oR := md255(dR, invSA)
			oG := md255(dG, invSA)
			oB := md255(dB, invSA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeXor performs a "Xor" compositing operation on premultiplied RGBA images.
func CompositeXor(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sR, sG, sB, sA := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2]), uint32(src[i+3])
			dR, dG, dB, dA := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2]), uint32(dst[i+3])

			invSA := 255 - sA
			invDA := 255 - dA

			oA := md255(sA, invDA) + md255(dA, invSA)
			oR := md255(sR, invDA) + md255(dR, invSA)
			oG := md255(sG, invDA) + md255(dG, invSA)
			oB := md255(sB, invDA) + md255(dB, invSA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeClear performs a "Clear" compositing operation.
func CompositeClear(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(_, _, out []uint8) {
		for i := range out {
			out[i] = 0
		}
	})
}

// CompositeDestination performs a "Destination" compositing operation.
func CompositeDestination(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA {
	return core.Iterate(pixIter, calc, func(dst, _, out []uint8) {
		// For "Destination", the output is simply the destination.
		copy(out, dst)
	})
}
