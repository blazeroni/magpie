// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package core

import (
	"image"
	"reflect"
	"testing"
)

type _intersectFn = func(dst image.Rectangle, r image.Rectangle, src image.Rectangle, sp image.Point, out *image.Rectangle, op image.Point) image.Rectangle

func TestIntersectNRGBA(t *testing.T) {
	_testIntersect(t, func(dst image.Rectangle, r image.Rectangle, src image.Rectangle, sp image.Point, out *image.Rectangle, op image.Point) image.Rectangle {
		var outImg *image.NRGBA
		if out != nil {
			outImg = image.NewNRGBA(*out)
		}
		return IntersectNRGBA(image.NewNRGBA(dst), r, image.NewNRGBA(src), sp, outImg, op)
	})
}

func TestIntersectRGBA(t *testing.T) {
	_testIntersect(t, func(dst image.Rectangle, r image.Rectangle, src image.Rectangle, sp image.Point, out *image.Rectangle, op image.Point) image.Rectangle {
		var outImg *image.RGBA
		if out != nil {
			outImg = image.NewRGBA(*out)
		}
		return IntersectRGBA(image.NewRGBA(dst), r, image.NewRGBA(src), sp, outImg, op)
	})
}

func _testIntersect(t *testing.T, intersectFn _intersectFn) {
	testCases := []struct {
		name     string
		dst      image.Rectangle
		r        image.Rectangle
		src      image.Rectangle
		sp       image.Point
		out      *image.Rectangle
		op       image.Point
		expected image.Rectangle
	}{
		{
			name:     "Fully contained",
			dst:      image.Rect(0, 0, 200, 200),
			r:        image.Rect(50, 50, 150, 150),
			src:      image.Rect(0, 0, 200, 200),
			sp:       image.Point{X: 10, Y: 10},
			out:      _rectPtr(0, 0, 200, 200),
			op:       image.Point{X: 20, Y: 20},
			expected: image.Rect(50, 50, 150, 150),
		},
		{
			name:     "Clipped by dst",
			dst:      image.Rect(0, 0, 100, 100),
			r:        image.Rect(50, 50, 150, 150),
			src:      image.Rect(0, 0, 200, 200),
			sp:       image.Point{X: 10, Y: 10},
			out:      _rectPtr(0, 0, 200, 200),
			op:       image.Point{X: 20, Y: 20},
			expected: image.Rect(50, 50, 100, 100),
		},
		{
			name:     "Clipped by src",
			dst:      image.Rect(0, 0, 200, 200),
			r:        image.Rect(50, 50, 150, 150),
			src:      image.Rect(0, 0, 50, 50),
			sp:       image.Point{X: 10, Y: 10},
			out:      _rectPtr(0, 0, 200, 200),
			op:       image.Point{X: 20, Y: 20},
			expected: image.Rect(50, 50, 90, 90),
		},
		{
			name:     "Clipped by out",
			dst:      image.Rect(0, 0, 200, 200),
			r:        image.Rect(50, 50, 150, 150),
			src:      image.Rect(0, 0, 200, 200),
			sp:       image.Point{X: 10, Y: 10},
			out:      _rectPtr(0, 0, 50, 50),
			op:       image.Point{X: 20, Y: 20},
			expected: image.Rect(50, 50, 80, 80),
		},
		{
			name:     "Nil out",
			dst:      image.Rect(0, 0, 100, 100),
			r:        image.Rect(50, 50, 150, 150),
			src:      image.Rect(0, 0, 50, 50),
			sp:       image.Point{X: 10, Y: 10},
			out:      nil,
			op:       image.Point{},
			expected: image.Rect(50, 50, 90, 90),
		},
		{
			name:     "Empty intersection",
			dst:      image.Rect(0, 0, 10, 10),
			r:        image.Rect(50, 50, 150, 150),
			src:      image.Rect(0, 0, 200, 200),
			sp:       image.Point{X: 10, Y: 10},
			out:      _rectPtr(0, 0, 200, 200),
			op:       image.Point{X: 20, Y: 20},
			expected: image.Rect(0, 0, 0, 0), // Zero rectangle
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := intersectFn(tc.dst, tc.r, tc.src, tc.sp, tc.out, tc.op)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func _rectPtr(x0, y0, x1, y1 int) *image.Rectangle {
	return &image.Rectangle{
		Min: image.Point{X: x0, Y: y0},
		Max: image.Point{X: x1, Y: y1},
	}
}
