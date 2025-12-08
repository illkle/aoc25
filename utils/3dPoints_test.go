package utils_test

import (
	"aoc-in-go/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_Distance(t *testing.T) {
	tests := []struct {
		name string
		a    utils.Point
		b    utils.Point
		want float64
	}{
		{
			a: utils.Point{
				X: 819,
				Y: 987,
				Z: 18,
			},
			b: utils.Point{
				X: 906,
				Y: 360,
				Z: 560,
			},
			want: 833.3438666000968,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.a.Distance(tt.b)

			assert.Equal(t, tt.want, got)

		})
	}
}
