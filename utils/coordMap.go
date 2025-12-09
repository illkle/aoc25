package utils

import (
	"fmt"
	"strings"
)

type Map[T any] struct {
	coords map[string]T
	maxX   int
	maxY   int
}

func NewEmpty[T any]() *Map[T] {
	return &Map[T]{
		coords: make(map[string]T),
		maxX:   0,
		maxY:   0,
	}
}

func NewFromString(text string) *Map[rune] {
	m := NewEmpty[rune]()

	for y, line := range strings.Fields(text) {
		for x, r := range line {
			m.Set(x, y, r)
		}

	}
	return m
}

func ToMapString(x int, y int) string {
	return fmt.Sprintf("%d-%d", y, x)
}

func (c *Map[T]) GetMaxX() int {
	return c.maxX
}

func (c *Map[T]) GetMaxY() int {
	return c.maxY
}

func (c *Map[T]) Set(x int, y int, value T) {

	if x > c.maxX {
		c.maxX = x
	}
	if y > c.maxY {
		c.maxY = y
	}

	c.coords[ToMapString(x, y)] = value
}

func (c *Map[T]) Get(x int, y int) (T, bool) {
	value, found := c.coords[ToMapString(x, y)]
	return value, found
}

func (c *Map[T]) GetAdjacent8(x int, y int) []T {
	res := make([]T, 0, 8)

	for xx := x - 1; xx <= x+1; xx++ {
		for yy := y - 1; yy <= y+1; yy++ {
			if xx == x && yy == y {
				continue
			}
			v, e := c.Get(xx, yy)
			if e {
				res = append(res, v)
			}
		}
	}

	return res
}

func (m *Map[T]) FindFirst(tester func(v T) bool) (int, int, bool) {
	for x := 0; x <= m.GetMaxX(); x++ {
		for y := 0; y <= m.GetMaxY(); y++ {
			val, found := m.Get(x, y)
			if !found {
				panic("NF")
			} else if tester(val) {
				return x, y, true
			}
		}
	}

	return 0, 0, false
}

func (m *Map[T]) Display(fff func(v T) string) {
	for y := 0; y <= m.GetMaxY(); y++ {
		for x := 0; x <= m.GetMaxX(); x++ {
			v, _ := m.Get(x, y)
			fmt.Print(fff(v))
		}
		fmt.Print("\n")
	}

	fmt.Print("\n\n")
}

func (c *Map[T]) Update(x int, y int, function func(before T) T) {
	value, _ := c.Get(x, y)
	c.Set(x, y, function(value))
}
