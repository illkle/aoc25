package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	els := strings.Split(input, "\n\n")

	hasSpace := 0
	// bruh.....
	for _, l := range strings.Split(els[len(els)-1], "\n") {

		var x, y, g0, g1, g2, g3, g4, g5 int
		_, err := fmt.Sscanf(l, "%dx%d: %d %d %d %d %d %d", &x, &y, &g0, &g1, &g2, &g3, &g4, &g5)

		if err != nil {
			panic(err)
		}

		spaceNeededMin := g0*7 + g1*7 + g2*7 + g3*7 + g4*7 + g5*7
		spaceAvailable := x * y

		if spaceAvailable > spaceNeededMin {
			hasSpace++
		}
	}

	return hasSpace
}
