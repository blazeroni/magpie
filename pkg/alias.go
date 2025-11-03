// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package magpie

import (
	"image"

	"github.com/blazeroni/magpie/pkg/core"
)

// PixelIterator is an alias for core.PixelIterator.
// It iterates over pixels in an image an passes each row to a PixRowCalculator.
type PixelIterator = core.PixelIterator

// PixCalculator calculates the set of Pix slices for a given row.
type PixCalculator[T image.Image] = core.PixCalculator[T]
type PixRowCalculator = core.PixRowCalculator

// NewPixelIterator creates a new PixelIterator.
// Concurrency specifies the number of goroutines to use for parallel processing.
func NewPixelIterator(concurrency int) PixelIterator {
	return core.NewPixelIterator(concurrency)
}
