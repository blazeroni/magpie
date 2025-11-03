// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package core

import (
	"image"
	"runtime"
	"sync"
	"testing"
)

// --- Mock Implementations ---

// mockPixRowCalculator is a mock implementation of PixRowCalculator for testing.
type mockPixRowCalculator struct {
	rect      image.Rectangle
	processed []bool
	mutex     sync.Mutex
}

func newMockPixRowCalculator(rect image.Rectangle) *mockPixRowCalculator {
	return &mockPixRowCalculator{
		rect:      rect,
		processed: make([]bool, rect.Dy()),
	}
}

func (m *mockPixRowCalculator) Rect() image.Rectangle {
	return m.rect
}

func (m *mockPixRowCalculator) Calculate(y int) ([]uint8, []uint8, []uint8) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// y is absolute from image origin, but our processed slice is 0-indexed from rect.Min.Y
	index := y - m.rect.Min.Y
	if index >= 0 && index < len(m.processed) {
		m.processed[index] = true
	}
	return nil, nil, nil
}

func (m *mockPixRowCalculator) AllProcessed() bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, p := range m.processed {
		if !p {
			return false
		}
	}
	return true
}

// mockPixCalculator is a mock for the generic Iterate function test.
type mockPixCalculator[T image.Image] struct {
	PixRowCalculator

	img T
}

func (m *mockPixCalculator[T]) Result() T {
	return m.img
}

func newMockPixCalculator[T image.Image](calc PixRowCalculator, resultImg T) *mockPixCalculator[T] {
	return &mockPixCalculator[T]{
		PixRowCalculator: calc,
		img:              resultImg,
	}
}

// --- Tests ---

func TestNewPixelIterator(t *testing.T) {
	t.Run("Serial for concurrency <= 1", func(t *testing.T) {
		if _, ok := NewPixelIterator(1).(SerialPixelIterator); !ok {
			t.Error("NewPixelIterator(1) should return SerialPixelIterator")
		}
		if _, ok := NewPixelIterator(0).(SerialPixelIterator); !ok {
			t.Error("NewPixelIterator(0) should return SerialPixelIterator")
		}
		if _, ok := NewPixelIterator(-1).(SerialPixelIterator); !ok {
			t.Error("NewPixelIterator(-1) should return SerialPixelIterator")
		}
	})

	t.Run("Parallel for concurrency > 1", func(t *testing.T) {
		if _, ok := NewPixelIterator(2).(ParallelPixelIterator); !ok {
			t.Error("NewPixelIterator(2) should return ParallelPixelIterator")
		}
		if _, ok := NewPixelIterator(runtime.GOMAXPROCS(0)).(ParallelPixelIterator); !ok {
			t.Errorf("NewPixelIterator(%d) should return ParallelPixelIterator", runtime.GOMAXPROCS(0))
		}
	})
}

func TestSerialPixelIterator_Iterate(t *testing.T) {
	rect := image.Rect(0, 0, 1, 100)
	mockCalc := newMockPixRowCalculator(rect)
	iterator := NewSerialPixelIterator()

	iterator.Iterate(mockCalc, func(_, _, _ []uint8) {
		// a no-op function for this test
	})

	if !mockCalc.AllProcessed() {
		t.Errorf("SerialPixelIterator did not process all rows")
	}
}

func TestParallelPixelIterator_Iterate(t *testing.T) {
	tests := []struct {
		name        string
		rect        image.Rectangle
		concurrency int
	}{
		{"100 rows, 4 goroutines", image.Rect(0, 0, 1, 100), 4},
		{"100 rows, 1 goroutine", image.Rect(0, 10, 1, 110), 1},
		{"7 rows, 3 goroutines", image.Rect(0, 0, 1, 7), 3},
		{"10 rows, 12 goroutines", image.Rect(0, 0, 1, 10), 12},
		{"0 rows", image.Rect(0, 0, 1, 0), 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCalc := newMockPixRowCalculator(tt.rect)
			iterator := NewParallelPixelIterator(tt.concurrency)

			iterator.Iterate(mockCalc, func(_, _, _ []uint8) {
				// a no-op function for this test
			})

			if tt.rect.Dy() > 0 && !mockCalc.AllProcessed() {
				t.Errorf("ParallelPixelIterator did not process all rows for rect %v", tt.rect)
			}
			if tt.rect.Dy() == 0 && !mockCalc.AllProcessed() {
				t.Errorf("ParallelPixelIterator should handle zero-height rect correctly")
			}
		})
	}
}

func TestIterate(t *testing.T) {
	rect := image.Rect(0, 0, 1, 10)
	resultImg := image.NewRGBA(rect)
	mockRowCalc := newMockPixRowCalculator(rect)
	mockCalc := newMockPixCalculator(mockRowCalc, resultImg)

	iterator := NewSerialPixelIterator()

	result := Iterate(iterator, mockCalc, func(_, _, _ []uint8) {})

	if !mockRowCalc.AllProcessed() {
		t.Error("Iterate helper function did not process all rows")
	}

	if result != resultImg {
		t.Error("Iterate helper function did not return the correct result image")
	}
}
