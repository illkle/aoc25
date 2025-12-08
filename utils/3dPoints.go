package utils

import (
	"math"
)

type Point struct {
	X int
	Y int
	Z int
}

func (a Point) Distance(b Point) float64 {
	calc := math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2) + math.Pow(float64(b.Z-a.Z), 2)

	return math.Sqrt(math.Abs(calc))
}
