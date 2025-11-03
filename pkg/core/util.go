// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package core

import (
	"cmp"
	"image"
	"image/color"
)

func Clamp[T cmp.Ordered](value, mn, mx T) T {
	return max(mn, min(value, mx))
}

func IntersectNRGBA(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, out *image.NRGBA, op image.Point) image.Rectangle {
	orig := r.Min
	r = r.Intersect(dst.Bounds())
	r = r.Intersect(src.Bounds().Add(orig.Sub(sp)))
	if out != nil {
		r = r.Intersect(out.Bounds().Add(orig.Sub(op)))
	}
	return r
}

func IntersectRGBA(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, out *image.RGBA, op image.Point) image.Rectangle {
	orig := r.Min
	r = r.Intersect(dst.Bounds())
	r = r.Intersect(src.Bounds().Add(orig.Sub(sp)))
	if out != nil {
		r = r.Intersect(out.Bounds().Add(orig.Sub(op)))
	}
	return r
}

func IsColorModelSupported(model color.Model) bool {
	return model == color.RGBAModel || model == color.NRGBAModel
}
