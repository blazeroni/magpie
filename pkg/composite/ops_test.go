// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package composite_test

import (
	"testing"

	"github.com/blazeroni/magpie/pkg/composite"
	"github.com/blazeroni/magpie/pkg/op"
)

func TestCompositeOps(t *testing.T) {
	tests := []struct {
		name         string
		opFunc       func() op.CompositeOp
		expectedMode op.CompositeMode
	}{
		{
			name:         "SourceOver",
			opFunc:       composite.SourceOver,
			expectedMode: op.SourceOver,
		},
		{
			name:         "Source",
			opFunc:       composite.Source,
			expectedMode: op.Source,
		},
		{
			name:         "SourceIn",
			opFunc:       composite.SourceIn,
			expectedMode: op.SourceIn,
		},
		{
			name:         "SourceOut",
			opFunc:       composite.SourceOut,
			expectedMode: op.SourceOut,
		},
		{
			name:         "SourceAtop",
			opFunc:       composite.SourceAtop,
			expectedMode: op.SourceAtop,
		},
		{
			name:         "DestinationOver",
			opFunc:       composite.DestinationOver,
			expectedMode: op.DestinationOver,
		},
		{
			name:         "Destination",
			opFunc:       composite.Destination,
			expectedMode: op.Destination,
		},
		{
			name:         "DestinationIn",
			opFunc:       composite.DestinationIn,
			expectedMode: op.DestinationIn,
		},
		{
			name:         "DestinationOut",
			opFunc:       composite.DestinationOut,
			expectedMode: op.DestinationOut,
		},
		{
			name:         "DestinationAtop",
			opFunc:       composite.DestinationAtop,
			expectedMode: op.DestinationAtop,
		},
		{
			name:         "Xor",
			opFunc:       composite.Xor,
			expectedMode: op.Xor,
		},
		{
			name:         "Clear",
			opFunc:       composite.Clear,
			expectedMode: op.Clear,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.opFunc()
			if result.Mode != tt.expectedMode {
				t.Errorf("%s() Mode got = %v, want %v", tt.name, result.Mode, tt.expectedMode)
			}
		})
	}
}
