// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package internal

import (
	"image"

	"github.com/blazeroni/magpie/pkg/core"
)

// Op defines a drawing operation.
type Op interface {
	IsValid() bool
	ApplyNRGBA(pixIter core.PixelIterator, calc core.PixCalculator[*image.NRGBA]) *image.NRGBA
	ApplyRGBA(pixIter core.PixelIterator, calc core.PixCalculator[*image.RGBA]) *image.RGBA
}
