// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package core

import (
	"image"
	"runtime"
	"sync"
	"sync/atomic"
)

var _ PixelIterator = (*SerialPixelIterator)(nil)
var _ PixelIterator = (*ParallelPixelIterator)(nil)

type PixelIterator interface {
	Iterate(pixCalc PixRowCalculator, fn func(dst, src, out []uint8))
}

func NewPixelIterator(concurrency int) PixelIterator {
	if concurrency <= 1 {
		return SerialPixelIterator{}
	}
	return NewParallelPixelIterator(concurrency)
}

func Iterate[T image.Image](pixIter PixelIterator, pixCalc PixCalculator[T], fn func(dst, src, out []uint8)) T {
	pixIter.Iterate(pixCalc, fn)
	return pixCalc.Result()
}

// --- PixelIterator implementations ---

// region SerialPixelIterator

// SerialPixelIterator iterates over a PixRowCalculator in serial with no concurrency.
type SerialPixelIterator struct{}

// NewSerialPixelIterator creates a new SerialPixelIterator.
func NewSerialPixelIterator() SerialPixelIterator {
	return SerialPixelIterator{}
}

func (i SerialPixelIterator) Iterate(pixCalc PixRowCalculator, fn func(dst, src, out []uint8)) {
	rect := pixCalc.Rect()
	for y := range rect.Dy() {
		fn(pixCalc.Calculate(y))
	}
}

// endregion SerialPixelIterator

// region ParallelPixelIterator

// ParallelPixelIterator iterates over a PixRowCalculator in parallel with the specified concurrency.
type ParallelPixelIterator struct {
	concurrency int
}

// NewParallelPixelIterator creates a new ParallelPixelIterator.
// Concurrency specifies the number of goroutines to use for parallel processing.
// Each iteration will create its own set of goroutines, so concurrently processing multiple
// images will result in multiple sets of goroutines.
func NewParallelPixelIterator(concurrency int) ParallelPixelIterator {
	return ParallelPixelIterator{
		concurrency: concurrency,
	}
}

func (ppi ParallelPixelIterator) Iterate(pixCalc PixRowCalculator, fn func(dst, src, out []uint8)) {
	numGoroutines := Clamp(ppi.concurrency, 1, runtime.GOMAXPROCS(0))
	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)

	var y int32 = -1

	rect := pixCalc.Rect()

	for range numGoroutines {
		go func() {
			defer wg.Done()
			for {
				row := atomic.AddInt32(&y, 1)
				if int(row) >= rect.Dy() {
					break
				}
				fn(pixCalc.Calculate(rect.Min.Y + int(row)))
			}
		}()
	}

	wg.Wait()
}

// endregion ParallelPixelIterator
