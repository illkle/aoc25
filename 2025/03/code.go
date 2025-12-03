package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func findMaxNaive(line string) int {
	split := strings.Split(line, "")

	max := 0

	for i1 := 0; i1 < len(split)-1; i1++ {
		for i2 := i1 + 1; i2 < len(split); i2++ {
			res, err := strconv.Atoi(split[i1] + split[i2])
			if err != nil {
				panic(err)
			}

			if res > max {
				max = res
			}
		}
	}
	return max
}

func finMaxOptimized(line string, needNumbers int) int {
	split := strings.Split(line, "")

	lineMap := map[string][]int{}

	for i := 0; i < len(split); i++ {
		lineMap[split[i]] = append(lineMap[split[i]], i)
	}

	res := ""
	nextIndex := 0

	for i := 0; i < needNumbers; i++ {
		blockPositionsOver := len(split) - (needNumbers - i)

	outerLoop:
		for n := 9; n > 0; n-- {
			letter := strconv.Itoa(n)
			target := lineMap[letter]

			for _, ii := range target {
				if ii <= blockPositionsOver && ii >= nextIndex {
					res += letter
					nextIndex = ii + 1
					break outerLoop
				}
			}

		}
	}

	c, err := strconv.Atoi(res)

	if err != nil {
		panic(err)
	}

	return c
}

func p1(input string) any {
	total := 0

	for _, line := range strings.Fields(input) {
		total += findMaxNaive(line)
	}

	return total
}

func p2(input string) any {
	total := 0

	for _, line := range strings.Fields(input) {
		total += finMaxOptimized(line, 12)
	}

	return total
}

func run(part2 bool, input string) any {
	if part2 {
		return p2(input)
	}
	return p1(input)
}
