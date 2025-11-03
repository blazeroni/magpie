// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package nrgba

import (
	"image"

	"github.com/blazeroni/magpie/pkg/core"
)

func CompositeSourceOver(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sA := uint32(src[i+3])

			// Fast path optimizations for Source Over
			switch sA {
			case 0:
				// Source is fully transparent, so output is just the destination.
				out[i], out[i+1], out[i+2], out[i+3] = dst[i], dst[i+1], dst[i+2], dst[i+3]
			case 255:
				// Source is fully opaque, so it completely covers the destination.
				out[i], out[i+1], out[i+2], out[i+3] = src[i], src[i+1], src[i+2], src[i+3]
			default:
				sR, sG, sB := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2])
				dR, dG, dB, dA := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2]), uint32(dst[i+3])

				var oR, oG, oB uint32
				var oA uint32

				// General Source Over calculation for partial alpha.
				invSA := 255 - sA
				wt := md255(dA, invSA)
				oA = sA + wt

				// Use direct straight-alpha formula, keeping precision high
				// before the final division.
				// Formula: oR = (sR*sA + dR*dA*(1-sA)) / oA
				oR = (sR*sA + dR*wt + oA/2) / oA
				oG = (sG*sA + dG*wt + oA/2) / oA
				oB = (sB*sA + dB*wt + oA/2) / oA

				out[i] = uint8(oR)
				out[i+1] = uint8(oG)
				out[i+2] = uint8(oB)
				out[i+3] = uint8(oA)
			}
		}
	})
}

// CompositeSourceIn performs a "Source In" compositing operation.
func CompositeSourceIn(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sA := uint32(src[i+3])
			dA := uint32(dst[i+3])

			if sA == 0 || dA == 0 {
				out[i], out[i+1], out[i+2], out[i+3] = 0, 0, 0, 0
				continue
			}

			var oA uint32
			if dA == 255 {
				oA = sA
			} else {
				oA = md255(sA, dA)
			}

			// Color is from source
			out[i] = src[i]
			out[i+1] = src[i+1]
			out[i+2] = src[i+2]
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeSourceAtop performs a "Source Atop" compositing operation.
func CompositeSourceAtop(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sA := uint32(src[i+3])
			dA := uint32(dst[i+3])

			out[i+3] = uint8(dA)

			if dA == 0 {
				out[i], out[i+1], out[i+2] = 0, 0, 0
				continue
			}

			if sA == 0 {
				out[i] = dst[i]
				out[i+1] = dst[i+1]
				out[i+2] = dst[i+2]
				continue
			}

			sR, sG, sB := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2])

			if sA == 255 {
				out[i] = uint8(sR)
				out[i+1] = uint8(sG)
				out[i+2] = uint8(sB)
				continue
			}

			// General case: oC = (sC*sA + dC*(1-sA) + 127) / 255
			dR, dG, dB := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2])
			invSA := 255 - sA
			// Could be faster with two md255
			oR := div255(sR*sA + dR*invSA)
			oG := div255(sG*sA + dG*invSA)
			oB := div255(sB*sA + dB*invSA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
		}
	})
}

// CompositeSourceOut performs a "Source Out" compositing operation.
func CompositeSourceOut(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sA := uint32(src[i+3])
			dA := uint32(dst[i+3])

			var oA uint32
			if sA == 0 || dA == 255 {
				out[i], out[i+1], out[i+2], out[i+3] = 0, 0, 0, 0
				continue
			}

			if dA == 0 {
				oA = sA
			} else {
				invDA := 255 - dA
				oA = md255(sA, invDA)
			}

			// Color is from source
			out[i] = src[i]
			out[i+1] = src[i+1]
			out[i+2] = src[i+2]
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeSource performs a "Source" (or "Copy") compositing operation.
func CompositeSource(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(_, src, out []uint8) {
		// For "Source", the output is simply the source.
		copy(out, src)
	})
}

// CompositeDestinationOver performs a "Destination Over" compositing operation.
func CompositeDestinationOver(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sA := uint32(src[i+3])
			dA := uint32(dst[i+3])

			if dA == 255 || sA == 0 {
				// If dest is opaque or source is transparent, output is destination.
				out[i] = dst[i]
				out[i+1] = dst[i+1]
				out[i+2] = dst[i+2]
				out[i+3] = uint8(dA)
				continue
			}

			sR, sG, sB := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2])
			dR, dG, dB := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2])
			var oR, oG, oB, oA uint32

			// General case: oA = dA + sA*(1-dA)
			invDA := 255 - dA
			wt := md255(sA, invDA)
			oA = dA + wt

			// oR = (dR*dA + sR*sA*(1-dA)) / oA
			oR = (dR*dA + sR*wt + oA/2) / oA
			oG = (dG*dA + sG*wt + oA/2) / oA
			oB = (dB*dA + sB*wt + oA/2) / oA

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeDestinationIn performs a "Destination In" compositing operation.
func CompositeDestinationIn(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sA := uint32(src[i+3])
			dA := uint32(dst[i+3])

			var oA uint32
			if sA == 0 || dA == 0 {
				out[i], out[i+1], out[i+2], out[i+3] = 0, 0, 0, 0
				continue
			}

			if sA == 255 {
				oA = dA
			} else {
				oA = md255(sA, dA)
			}

			// Color is from destination
			out[i] = dst[i]
			out[i+1] = dst[i+1]
			out[i+2] = dst[i+2]
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeDestinationAtop performs a "Destination Atop" compositing operation.
func CompositeDestinationAtop(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sA := uint32(src[i+3])
			dA := uint32(dst[i+3])

			out[i+3] = uint8(sA) // oA = sA

			if sA == 0 {
				out[i], out[i+1], out[i+2] = 0, 0, 0
				continue
			}

			if dA == 0 {
				out[i], out[i+1], out[i+2] = src[i], src[i+1], src[i+2]
				continue
			}

			if dA == 255 {
				out[i], out[i+1], out[i+2] = dst[i], dst[i+1], dst[i+2]
				continue
			}

			// General case: oC = (dC*dA + sC*(1-dA) + 127) / 255
			dR, dG, dB := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2])
			sR, sG, sB := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2])
			invDA := 255 - dA

			// Could be faster with two md255
			oR := div255(dR*dA + sR*invDA)
			oG := div255(dG*dA + sG*invDA)
			oB := div255(dB*dA + sB*invDA)

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
		}
	})
}

// CompositeDestinationOut performs a "Destination Out" compositing operation.
func CompositeDestinationOut(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sA := uint32(src[i+3])
			dA := uint32(dst[i+3])

			var oA uint32
			if dA == 0 || sA == 255 {
				out[i], out[i+1], out[i+2], out[i+3] = 0, 0, 0, 0
				continue
			}

			if sA == 0 {
				oA = dA
			} else {
				invSA := 255 - sA
				oA = md255(dA, invSA)
			}

			// Color is from destination
			out[i] = dst[i]
			out[i+1] = dst[i+1]
			out[i+2] = dst[i+2]
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeXor performs an "Xor" compositing operation.
func CompositeXor(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, src, out []uint8) {
		for i := 0; i < len(src); i += 4 {
			sA := uint32(src[i+3])
			dA := uint32(dst[i+3])
			sR, sG, sB := uint32(src[i]), uint32(src[i+1]), uint32(src[i+2])
			dR, dG, dB := uint32(dst[i]), uint32(dst[i+1]), uint32(dst[i+2])

			var oR, oG, oB, oA uint32

			if sA == 0 {
				out[i], out[i+1], out[i+2], out[i+3] = dst[i], dst[i+1], dst[i+2], dst[i+3]
				continue
			}
			if dA == 0 {
				out[i], out[i+1], out[i+2], out[i+3] = src[i], src[i+1], src[i+2], src[i+3]
				continue
			}

			if (sA & dA) == 255 {
				out[i], out[i+1], out[i+2], out[i+3] = 0, 0, 0, 0
				continue
			}

			invSA := 255 - sA
			invDA := 255 - dA

			sWt := md255(sA, invDA)
			dWt := md255(dA, invSA)

			oA = sWt + dWt
			oR = (sR*sWt + dR*dWt + oA/2) / oA
			oG = (sG*sWt + dG*dWt + oA/2) / oA
			oB = (sB*sWt + dB*dWt + oA/2) / oA

			out[i] = uint8(oR)
			out[i+1] = uint8(oG)
			out[i+2] = uint8(oB)
			out[i+3] = uint8(oA)
		}
	})
}

// CompositeClear performs a "Clear" compositing operation.
func CompositeClear(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(_, _, out []uint8) {
		// For "Clear", the output is always transparent black.
		for i := range out {
			out[i] = 0
		}
	})
}

// CompositeDestination performs a "Destination" compositing operation.
func CompositeDestination(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA {
	return core.Iterate(pixIter, calc, func(dst, _, out []uint8) {
		// For "Destination", the output is simply the destination.
		copy(out, dst)
	})
}
