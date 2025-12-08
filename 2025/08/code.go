package main

import (
	"aoc-in-go/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type PossibleConn struct {
	index1   int
	index2   int
	distance float64
}

func getGroups(points []utils.Point, conns map[int][]int) [][]int {
	groups := [][]int{}
	seenP := map[int]bool{}

	for pi, _ := range points {
		if seenP[pi] {
			continue
		}

		newGroup := []int{}

		q := []int{pi}

		for len(q) > 0 {
			el := q[0]
			q = q[1:]

			if seenP[el] {
				continue
			} else {
				seenP[el] = true
			}

			newGroup = append(newGroup, el)

			q = append(q, conns[el]...)
		}

		groups = append(groups, newGroup)
	}
	return groups

}

func getPossibleConns(points []utils.Point) []PossibleConn {
	possibleConns := []PossibleConn{}

	seen := map[string]bool{}

	for a, ap := range points {
		for b, bp := range points {
			if a == b {
				continue
			}

			if seen[fmt.Sprintf("%d-%d", a, b)] {
				continue
			} else {
				seen[fmt.Sprintf("%d-%d", a, b)] = true
				seen[fmt.Sprintf("%d-%d", b, a)] = true
			}

			possibleConns = append(possibleConns, PossibleConn{
				distance: ap.Distance(bp),
				index1:   a,
				index2:   b,
			})
		}
	}
	slices.SortFunc(possibleConns, func(a PossibleConn, b PossibleConn) int {
		return int(a.distance - b.distance)
	})

	return possibleConns
}

func p1(points []utils.Point) any {
	possibleConns := getPossibleConns(points)

	connsToMake := utils.If(len(points) > 20, 1000, 10)

	conns := map[int][]int{}

	for _, c := range possibleConns[:connsToMake] {
		conns[c.index1] = append(conns[c.index1], c.index2)
		conns[c.index2] = append(conns[c.index2], c.index1)
	}

	groups := getGroups(points, conns)

	slices.SortFunc(groups, func(g1 []int, g2 []int) int {
		return len(g2) - len(g1)
	})

	return len(groups[0]) * len(groups[1]) * len(groups[2])
}

func p2(points []utils.Point) any {

	possibleConns := getPossibleConns(points)

	conns := map[int][]int{}

	for _, c := range possibleConns {
		conns[c.index1] = append(conns[c.index1], c.index2)
		conns[c.index2] = append(conns[c.index2], c.index1)

		g := getGroups(points, conns)

		if len(g) == 1 {
			return points[c.index1].X * points[c.index2].X
		}

	}

	return 1
}

func run(part2 bool, input string) any {

	points := []utils.Point{}

	for _, line := range strings.Fields(input) {
		split := strings.Split(line, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		points = append(points, utils.Point{X: x, Y: y, Z: z})
	}

	if part2 {
		return p2(points)
	}
	return p1(points)
}
