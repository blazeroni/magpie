// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package blend

import (
	"github.com/blazeroni/magpie/pkg/op"
)

type Compositing = op.BlendCompositing

func ColorBurn() op.BlendOp {
	return op.BlendOp{Mode: op.ColorBurn, Compositing: op.CompositeAll}
}

func ColorDodge() op.BlendOp {
	return op.BlendOp{Mode: op.ColorDodge, Compositing: op.CompositeAll}
}

func Darken() op.BlendOp {
	return op.BlendOp{Mode: op.Darken, Compositing: op.CompositeAll}
}

func Difference() op.BlendOp {
	return op.BlendOp{Mode: op.Difference, Compositing: op.CompositeAll}
}

func Divide() op.BlendOp {
	return op.BlendOp{Mode: op.Divide, Compositing: op.CompositeAll}
}

func Exclusion() op.BlendOp {
	return op.BlendOp{Mode: op.Exclusion, Compositing: op.CompositeAll}
}

func HardLight() op.BlendOp {
	return op.BlendOp{Mode: op.HardLight, Compositing: op.CompositeAll}
}

func HardMix() op.BlendOp {
	return op.BlendOp{Mode: op.HardMix, Compositing: op.CompositeAll}
}

func Lighten() op.BlendOp {
	return op.BlendOp{Mode: op.Lighten, Compositing: op.CompositeAll}
}

func LinearBurn() op.BlendOp {
	return op.BlendOp{Mode: op.LinearBurn, Compositing: op.CompositeAll}
}

func LinearDodge() op.BlendOp {
	return op.BlendOp{Mode: op.LinearDodge, Compositing: op.CompositeAll}
}

func LinearLight() op.BlendOp {
	return op.BlendOp{Mode: op.LinearLight, Compositing: op.CompositeAll}
}

func Multiply() op.BlendOp {
	return op.BlendOp{Mode: op.Multiply, Compositing: op.CompositeAll}
}

func Overlay() op.BlendOp {
	return op.BlendOp{Mode: op.Overlay, Compositing: op.CompositeAll}
}

func PinLight() op.BlendOp {
	return op.BlendOp{Mode: op.PinLight, Compositing: op.CompositeAll}
}

func Screen() op.BlendOp {
	return op.BlendOp{Mode: op.Screen, Compositing: op.CompositeAll}
}

func SoftLight() op.BlendOp {
	return op.BlendOp{Mode: op.SoftLight, Compositing: op.CompositeAll}
}

func Subtract() op.BlendOp {
	return op.BlendOp{Mode: op.Subtract, Compositing: op.CompositeAll}
}

func VividLight() op.BlendOp {
	return op.BlendOp{Mode: op.VividLight, Compositing: op.CompositeAll}
}
