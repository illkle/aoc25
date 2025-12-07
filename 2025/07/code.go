package main

import (
	"aoc-in-go/utils"
	"fmt"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type coords struct {
	x int
	y int
}

func p1(input string) any {
	m := utils.NewFromString(input)

	startX, startY, found := m.FindFirst(func(v rune) bool {
		return v == 'S'
	})

	if !found {
		panic("not found start")
	}

	beamLocations := []coords{{
		x: startX,
		y: startY + 1,
	}}

	uniqSplits := map[string]bool{}
	seenCoords := map[string]bool{}

	for len(beamLocations) > 0 {
		toProcess := beamLocations[0]
		beamLocations = beamLocations[1:]

		myCoords := fmt.Sprintf("%d-%d", toProcess.x, toProcess.y)

		if seenCoords[myCoords] {
			continue
		} else {
			seenCoords[myCoords] = true
		}

		nextLoc, found := m.Get(toProcess.x, toProcess.y+1)

		if !found {
			continue
		}

		if nextLoc == '^' {
			uniqSplits[fmt.Sprintf("%d-%d", toProcess.x, toProcess.y+1)] = true

			left := coords{
				x: toProcess.x - 1,
				y: toProcess.y + 1,
			}
			right := coords{
				x: toProcess.x + 1,
				y: toProcess.y + 1,
			}

			m.Set(left.x, left.y, '|')
			m.Set(right.x, right.y, '|')
			beamLocations = append(beamLocations, left, right)
		} else {

			next := coords{x: toProcess.x, y: toProcess.y + 1}
			m.Set(next.x, next.y, '|')
			beamLocations = append(beamLocations, next)
		}

	}

	return len(uniqSplits)
}

func p2(input string) any {
	m := utils.NewFromString(input)

	startX, startY, found := m.FindFirst(func(v rune) bool {
		return v == 'S'
	})

	if !found {
		panic("not found start")
	}

	futures := utils.NewEmpty[int]()
	futures.Set(startX, startY, 1)

	for y := startY + 1; y <= m.GetMaxY(); y++ {
		for x := 0; x <= m.GetMaxX(); x++ {
			prev, _ := futures.Get(x, y-1)
			if prev > 0 {
				cur, found := m.Get(x, y)
				if !found {
					panic("nf")
				}
				switch cur {
				case '.':
					futures.Update(x, y, func(before int) int { return before + prev })
				case '^':
					futures.Update(x-1, y, func(before int) int { return before + prev })
					futures.Update(x+1, y, func(before int) int { return before + prev })
				}
			}
		}
	}

	my := m.GetMaxY()
	total := 0
	for x := 0; x <= m.GetMaxX(); x++ {
		vvv, _ := futures.Get(x, my)
		total += vvv
	}

	return total
}

func run(part2 bool, input string) any {

	if part2 {
		return p2(input)
	}
	return p1(input)
}
