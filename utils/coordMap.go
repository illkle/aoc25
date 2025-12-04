package utils

import (
	"fmt"
	"strings"
)

type Map struct {
	coords map[string]rune
	maxX   int
	maxY   int
}

func NewEmpty() *Map {
	return &Map{
		coords: make(map[string]rune),
		maxX:   0,
		maxY:   0,
	}
}

func NewFromString(text string) *Map {
	m := NewEmpty()

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

func (c *Map) GetMaxX() int {
	return c.maxX
}

func (c *Map) GetMaxY() int {
	return c.maxY
}

func (c *Map) Set(x int, y int, value rune) {
	if x > c.maxX {
		c.maxX = x
	}
	if y > c.maxY {
		c.maxY = y
	}

	c.coords[ToMapString(x, y)] = value
}

func (c *Map) Get(x int, y int) (rune, bool) {
	value, found := c.coords[ToMapString(x, y)]
	return value, found
}

func (c *Map) GetAdjacent8(x int, y int) []rune {
	res := make([]rune, 0, 8)

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
