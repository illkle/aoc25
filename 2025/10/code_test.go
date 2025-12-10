package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestButton_updateState(t *testing.T) {
	tests := []struct {
		name      string
		corrLight []int
		someState []int
		want      []int
	}{
		{
			corrLight: []int{0, 1, 2},
			someState: []int{0, 0, 0, 1},
			want:      []int{1, 1, 1, 1},
		},
		{
			corrLight: []int{3},
			someState: []int{0, 0, 0, 1},
			want:      []int{0, 0, 0, 2},
		},
		{
			corrLight: []int{2, 3},
			someState: []int{0, 0, 0, 1},
			want:      []int{0, 0, 1, 2},
		},
		{
			corrLight: []int{1, 3},
			someState: []int{0, 0, 0, 1},
			want:      []int{0, 1, 0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Button{
				corrLight: tt.corrLight,
			}
			got := b.updateState(tt.someState)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMachine_isDesired(t *testing.T) {
	tests := []struct {
		name         string
		desiredState []bool
		someState    []int
		want         bool
	}{
		{
			desiredState: []bool{true, false, true, false},
			someState:    []int{1, 0, 1, 0},
			want:         true,
		},
		{
			desiredState: []bool{true, false, true, false},
			someState:    []int{0, 0, 1, 0},
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Machine{
				desiredState: tt.desiredState,
			}
			got := m.isDesiredOne(tt.someState)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSequential(t *testing.T) {
	t.Run("test", func(t *testing.T) {

		b1 := Button{
			corrLight: []int{1, 3},
		}

		bb := Button{
			corrLight: []int{3},
		}

		l := []int{0, 0, 0, 0}

		l = b1.updateState(l)
		assert.Equal(t, []int{0, 1, 0, 1}, l)
		l = bb.updateState(l)
		assert.Equal(t, []int{0, 1, 0, 2}, l)
		l = b1.updateState(l)
		assert.Equal(t, []int{0, 2, 0, 3}, l)

	})
}
