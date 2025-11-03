// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package rgba

import "github.com/blazeroni/magpie/pkg/internal"

func md255(a, b uint32) uint32 {
	return internal.Md255(a, b)
}

func sqrt(a uint32) uint32 {
	return internal.Sqrt(a)
}

func div255(x uint32) uint32 {
	return internal.Div255(x)
}

func unpremultiply(color, alpha uint32) uint32 {
	return internal.Unpremultiply(color, alpha)
}
