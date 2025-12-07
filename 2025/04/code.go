package main

import (
	"aoc-in-go/utils"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func findAccessibleRolls(m *utils.Map[rune], alsoRemove bool) int {
	total := 0

	for x := 0; x <= m.GetMaxX(); x++ {
		for y := 0; y <= m.GetMaxY(); y++ {
			val, found := m.Get(x, y)
			if !found {
				panic("NF")
			} else if val != '@' {
				continue
			}

			adj := m.GetAdjacent8(x, y)

			adjacentRolls := 0

			for _, v := range adj {
				if v == '@' {
					adjacentRolls++
				}
			}

			if adjacentRolls < 4 {
				total++
				if alsoRemove {
					m.Set(x, y, '.')
				}
			}
		}
	}

	return total

}

func p1(input string) any {
	m := utils.NewFromString(input)

	return findAccessibleRolls(m, false)
}

func p2(input string) any {

	m := utils.NewFromString(input)

	total := 0
	for true {
		remove := findAccessibleRolls(m, true)
		if remove > 0 {
			total += remove
		} else {
			break
		}
	}

	return total
}

func run(part2 bool, input string) any {

	if part2 {
		return p2(input)
	}
	return p1(input)
}
