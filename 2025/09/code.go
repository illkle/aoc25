package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type point struct {
	X int
	Y int
}

func calcArea(a, b point) float64 {
	return (math.Abs(float64(a.X-b.X)) + 1) * (math.Abs(float64(a.Y-b.Y)) + 1)
}

func p1(points []point) any {
	best := float64(0)

	for _, p := range points {
		for _, p2 := range points {
			area := calcArea(p, p2)
			if area > best {
				best = area
			}
		}
	}

	return int(best)
}

func pointWithinLineOnY(p point, y1, y2 int) bool {
	if y2 > y1 {
		return p.Y >= y1 && p.Y <= y2
	}
	return p.Y >= y2 && p.Y <= y1
}

func pointWithinLineOnX(p point, x1, x2 int) bool {
	if x2 > x1 {
		return p.X >= x1 && p.X <= x2
	}
	return p.X >= x2 && p.X <= x1
}

func pointInPolygon(p point, points []point) bool {
	pLen := len(points)
	inside := false
	for cur := 0; cur < pLen; cur++ {
		next := (cur + 1) % pLen
		x1, y1 := points[cur].X, points[cur].Y
		x2, y2 := points[next].X, points[next].Y

		withinY := pointWithinLineOnY(p, y1, y2)
		withinX := pointWithinLineOnX(p, x1, x2)
		if x1 == x2 && p.X == x1 && withinY {
			return true
		}

		if y1 == y2 && p.Y == y1 && withinX {
			return true
		}

		if x1 != x2 {
			continue
		}

		crossesX := p.X <= x1
		intersect := (y1 > p.Y) != (y2 > p.Y) && crossesX
		if intersect {
			inside = !inside
		}
	}

	return inside
}

type Pair struct {
	a    point
	b    point
	area float64
}

func allPointWithin(p Pair, points []point) bool {
	minX, maxX := min(p.a.X, p.b.X), max(p.a.X, p.b.X)
	minY, maxY := min(p.a.Y, p.b.Y), max(p.a.Y, p.b.Y)

	allPoints := []point{
		{X: minX, Y: minY},
		{X: minX, Y: maxY},
		{X: maxX, Y: minY},
		{X: maxX, Y: maxY},
	}
	for _, p := range allPoints {
		if !pointInPolygon(p, points) {
			return false
		}
	}
	return true
}

func linesCross(p1, p2, c1, c2 point) bool {
	pHorizontal := p1.Y == p2.Y
	cHorizontal := c1.Y == c2.Y

	if pHorizontal == cHorizontal {
		return false
	}

	if pHorizontal {
		minX, maxX := min(p1.X, p2.X), max(p1.X, p2.X)
		minY, maxY := min(c1.Y, c2.Y), max(c1.Y, c2.Y)
		return c1.X > minX && c1.X < maxX && p1.Y > minY && p1.Y < maxY
	} else {
		minX, maxX := min(c1.X, c2.X), max(c1.X, c2.X)
		minY, maxY := min(p1.Y, p2.Y), max(p1.Y, p2.Y)
		return p1.X > minX && p1.X < maxX && c1.Y > minY && c1.Y < maxY
	}
}

func noEdgeCrosses(p Pair, points []point) bool {
	minX, maxX := min(p.a.X, p.b.X), max(p.a.X, p.b.X)
	minY, maxY := min(p.a.Y, p.b.Y), max(p.a.Y, p.b.Y)

	rectEdges := [][2]point{
		{{X: minX, Y: minY}, {X: maxX, Y: minY}},
		{{X: maxX, Y: minY}, {X: maxX, Y: maxY}},
		{{X: maxX, Y: maxY}, {X: minX, Y: maxY}},
		{{X: minX, Y: maxY}, {X: minX, Y: minY}},
	}

	pLen := len(points)
	for _, e := range rectEdges {
		for cur := 0; cur < pLen; cur++ {
			next := (cur + 1) % pLen

			edgeStart := points[cur]
			edgeEnd := points[next]

			if linesCross(e[0], e[1], edgeStart, edgeEnd) {
				return false
			}

		}
	}

	return true
}
func allPointWithinFully(p Pair, points []point) bool {
	minX, maxX := min(p.a.X, p.b.X), max(p.a.X, p.b.X)
	minY, maxY := min(p.a.Y, p.b.Y), max(p.a.Y, p.b.Y)

	for yy := minY; yy <= maxY; yy++ {
		if !pointInPolygon(point{X: minX, Y: yy}, points) {
			return false
		}
		if !pointInPolygon(point{X: maxX, Y: yy}, points) {
			return false
		}
	}

	for xx := minX; xx <= maxX; xx++ {
		if !pointInPolygon(point{X: xx, Y: minY}, points) {
			return false
		}
		if !pointInPolygon(point{X: xx, Y: maxY}, points) {
			return false
		}
	}

	return true
}

func p2(points []point) any {
	areas := []Pair{}

	for _, sp := range points {
		for _, sp2 := range points {
			areas = append(areas, Pair{a: sp, b: sp2, area: calcArea(sp, sp2)})
		}
	}

	best := float64(0)
	for _, a := range areas {
		if a.area < best {
			continue
		}
		if allPointWithin(a, points) && noEdgeCrosses(a, points) {
			best = a.area
		}
	}

	return int(best)

	/* Stupid and parallel, takes 1min on m2 pro. Only after writing both versions I discovered that my calcArea was incorrectly calculating area in p2 💀💀💀💀💀
	var wg sync.WaitGroup
	var bestM sync.RWMutex
	best := 0
	ch := make(chan int)

	for _, ar := range areas {
		wg.Add(1)
		go func() {
			defer wg.Done()

			bestM.RLock()
			vvv := float64(best)
			bestM.RUnlock()
			if ar.area <= vvv {
				return
			}

			if !allPointWithin(ar, points) {
				return
			}
			if allPointWithinFully(ar, points) {
				ch <- int(ar.area)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		bestM.Lock()

		if v > int(best) {
			fmt.Println("NEW BEST", v)
			best = v
		}

		bestM.Unlock()

	}
	return best
	*/

}

func run(part2 bool, input string) any {
	points := []point{}

	for _, p := range strings.Fields(input) {
		var x, y int
		_, err := fmt.Sscanf(p, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		points = append(points, point{X: x, Y: y})
	}

	if part2 {
		return p2(points)
	}
	return p1(points)
}
