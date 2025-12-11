package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type qEl struct {
	current string
	fft     bool
	dac     bool
}

func findPathCount(paths map[string][]string, memo map[qEl]int, me qEl, to string, isValid func(e qEl) bool) int {
	memoized, found := memo[me]
	if found {
		return memoized
	}

	if me.current == to {
		if isValid(me) {
			memo[me] = 1
			return 1
		} else {
			memo[me] = 0
			return 0
		}
	}

	total := 0

	if me.current == "fft" {
		me.fft = true
	}
	if me.current == "dac" {
		me.dac = true
	}

	for _, p := range paths[me.current] {
		newMe := me
		newMe.current = p
		total += findPathCount(paths, memo, newMe, to, isValid)
	}

	memo[me] = total

	return total
}

func p1(paths map[string][]string) any {
	memo := make(map[qEl]int)
	return findPathCount(paths, memo, qEl{current: "you"}, "out", func(e qEl) bool { return true })
}

func p2(paths map[string][]string) any {
	memo := make(map[qEl]int)
	return findPathCount(paths, memo, qEl{current: "svr"}, "out", func(el qEl) bool { return el.dac && el.fft })
}

func run(part2 bool, input string) any {

	mm := map[string][]string{}

	for _, l := range strings.Split(input, "\n") {
		spl1 := strings.Split(l, ": ")
		mm[spl1[0]] = strings.Split(spl1[1], " ")
	}

	if part2 {
		return p2(mm)
	}
	return p1(mm)
}
