// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package core

import "image/color"

type Config interface {
	PixelIterator() PixelIterator
	DefaultOutputMode() DefaultOutputMode
	DefaultColorModel() color.Model
}
