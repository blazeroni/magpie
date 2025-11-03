// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package composite

import "github.com/blazeroni/magpie/pkg/op"

func SourceOver() op.CompositeOp {
	return op.CompositeOp{Mode: op.SourceOver}
}

func Source() op.CompositeOp {
	return op.CompositeOp{Mode: op.Source}
}

func SourceIn() op.CompositeOp {
	return op.CompositeOp{Mode: op.SourceIn}
}

func SourceOut() op.CompositeOp {
	return op.CompositeOp{Mode: op.SourceOut}
}

func SourceAtop() op.CompositeOp {
	return op.CompositeOp{Mode: op.SourceAtop}
}

func DestinationOver() op.CompositeOp {
	return op.CompositeOp{Mode: op.DestinationOver}
}

func Destination() op.CompositeOp {
	return op.CompositeOp{Mode: op.Destination}
}

func DestinationIn() op.CompositeOp {
	return op.CompositeOp{Mode: op.DestinationIn}
}

func DestinationOut() op.CompositeOp {
	return op.CompositeOp{Mode: op.DestinationOut}
}

func DestinationAtop() op.CompositeOp {
	return op.CompositeOp{Mode: op.DestinationAtop}
}

func Xor() op.CompositeOp {
	return op.CompositeOp{Mode: op.Xor}
}

func Clear() op.CompositeOp {
	return op.CompositeOp{Mode: op.Clear}
}
