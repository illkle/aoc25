package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitStringToPieces(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		pieceLen int
		want     []string
	}{
		{
			name:     "2",
			s:        "123123",
			pieceLen: 2,
			want:     []string{"12", "31", "23"},
		},

		{
			name:     "3",
			s:        "123123",
			pieceLen: 3,
			want:     []string{"123", "123"},
		},

		{
			name:     "1",
			s:        "123123",
			pieceLen: 1,
			want:     []string{"1", "2", "3", "1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitStringToPieces(tt.s, tt.pieceLen)
			assert.Equal(t, got, tt.want)
		})
	}
}
