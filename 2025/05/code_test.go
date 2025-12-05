package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compactRanges(t *testing.T) {
	tests := []struct {
		name   string
		ranges []rangeData
		want   []rangeData
	}{
		{
			ranges: []rangeData{{from: 10, to: 14}, {from: 16, to: 20}, {from: 12, to: 18}},
			want:   []rangeData{{from: 10, to: 20}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compactRanges(tt.ranges)
			assert.Equal(t, tt.want, got)
		})
	}
}
