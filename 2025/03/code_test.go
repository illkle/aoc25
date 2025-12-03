package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_finMaxOptimized(t *testing.T) {
	tests := []struct {
		name        string
		line        string
		needNumbers int
		want        int
	}{
		{
			name:        "test",
			line:        "978",
			needNumbers: 2,
			want:        98,
		},
		{
			name:        "test",
			line:        "987654321111111",
			needNumbers: 12,
			want:        987654321111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := finMaxOptimized(tt.line, tt.needNumbers)
			assert.Equal(t, tt.want, got)
		})
	}
}
