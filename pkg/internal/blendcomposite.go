// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package internal

// BlendCompositing defines the blending mode for composite operations.
// Exposed for public use in the op package.
type BlendCompositing int

const (
	CompositeBlendOnly   BlendCompositing = 1
	CompositeBlendAndDst BlendCompositing = 2
	CompositeBlendAndSrc BlendCompositing = 4
	CompositeAll         BlendCompositing = 6 // CompositeBlendAndDst | CompositeBlendAndSrc
)
