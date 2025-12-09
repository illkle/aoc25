package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pointWithinLineOnY(t *testing.T) {
	tests := []struct {
		name string
		p    point
		y1   int
		y2   int
		want bool
	}{
		{
			name: "in1",
			p:    point{X: 99, Y: 10},
			y1:   9,
			y2:   11,
			want: true,
		},
		{
			name: "in1",
			p:    point{X: 99, Y: 10},
			y1:   12,
			y2:   8,
			want: true,
		},
		{
			name: "bottom",
			p:    point{X: 99, Y: 10},
			y1:   12,
			y2:   10,
			want: true,
		},
		{
			name: "top",
			p:    point{X: 99, Y: 12},
			y1:   12,
			y2:   10,
			want: true,
		},
		{
			name: "bottomR",
			p:    point{X: 99, Y: 10},
			y1:   10,
			y2:   12,
			want: true,
		},
		{
			name: "topR",
			p:    point{X: 99, Y: 12},
			y1:   10,
			y2:   12,
			want: true,
		},
		{
			name: "out 1",
			p:    point{X: 99, Y: 13},
			y1:   10,
			y2:   12,
			want: false,
		},
		{
			name: "out 2",
			p:    point{X: 99, Y: 9},
			y1:   10,
			y2:   12,
			want: false,
		},
		{
			name: "hor",
			p:    point{X: 99, Y: 9},
			y1:   9,
			y2:   9,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pointWithinLineOnY(tt.p, tt.y1, tt.y2)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_pointInPolygon(t *testing.T) {
	tests := []struct {
		name   string
		p      point
		points []point
		want   bool
	}{
		{
			name:   "inside",
			p:      point{X: 10, Y: 10},
			points: []point{{5, 5}, {15, 5}, {15, 15}, {5, 15}},
			want:   true,
		},

		{
			name:   "on left edge",
			p:      point{X: 5, Y: 5},
			points: []point{{5, 5}, {15, 5}, {15, 15}, {5, 15}},
			want:   true,
		},
		{
			name:   "on right edge",
			p:      point{X: 15, Y: 5},
			points: []point{{5, 5}, {15, 5}, {15, 15}, {5, 15}},
			want:   true,
		},
		{
			name:   "before",
			p:      point{X: 4, Y: 5},
			points: []point{{5, 5}, {15, 5}, {15, 15}, {5, 15}},
			want:   false,
		},
		{
			name:   "after",
			p:      point{X: 16, Y: 5},
			points: []point{{5, 5}, {15, 5}, {15, 15}, {5, 15}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pointInPolygon(tt.p, tt.points)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_linesCross(t *testing.T) {
	tests := []struct {
		name string
		p1   point
		p2   point
		c1   point
		c2   point
		want bool
	}{
		{
			p1: point{
				X: 1,
				Y: 1,
			},
			p2: point{
				X: 1,
				Y: 10,
			},
			c1: point{
				X: 1,
				Y: 5,
			},
			c2: point{
				X: 2,
				Y: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := linesCross(tt.p1, tt.p2, tt.c1, tt.c2)
			assert.Equal(t, tt.want, got)
		})
	}
}
