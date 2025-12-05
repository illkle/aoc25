package main

import (
	"aoc-in-go/utils"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type rangeData struct {
	from int
	to   int
}

func (r *rangeData) InRange(v int) bool {
	return v >= r.from && v <= r.to
}

type processedInput struct {
	ingredients []int
	ranges      []rangeData
}

func p1(input processedInput) any {
	totalFresh := 0

	for _, ing := range input.ingredients {
		for _, r := range input.ranges {
			if r.InRange(ing) {
				totalFresh++
				break
			}
		}
	}

	return totalFresh
}

func p2(input processedInput) any {
	total := 0

	for _, r := range input.ranges {
		total += r.to - r.from + 1
	}

	return total
}

func compactRanges(ranges []rangeData) []rangeData {
	slices.SortFunc(ranges, func(a rangeData, b rangeData) int {
		return a.from - b.from
	})

	processed := []rangeData{}

	for _, v := range ranges {
		if len(processed) < 1 {
			processed = append(processed, v)
		}
		lastOne := processed[len(processed)-1]

		if v.from < lastOne.from {
			panic("something is wrong")
		}

		if v.from > lastOne.to+1 {
			processed = append(processed, v)
		} else {
			lastOne.to = utils.If(v.to > lastOne.to, v.to, lastOne.to)
			processed[len(processed)-1] = lastOne
		}
	}

	return processed
}

func run(part2 bool, input string) any {

	pi := processedInput{}

	for _, line := range strings.Fields(input) {
		if strings.Contains(line, "-") {
			split := strings.Split(line, "-")
			from, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			pi.ranges = append(pi.ranges, rangeData{
				from: from,
				to:   to,
			})
		} else {
			ing, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			pi.ingredients = append(pi.ingredients, ing)
		}
	}

	pi.ranges = compactRanges(pi.ranges)

	if part2 {
		return p2(pi)
	}
	return p1(pi)
}
