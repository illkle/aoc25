package main

import (
	"aoc-in-go/utils"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func p1(input string) int {
	dial := 50
	seenZero := 0

	for _, line := range strings.Fields(input) {
		runic := []rune(line)
		number, err := strconv.Atoi(string(runic[1:]))

		if err != nil {
			panic(err)
		}

		dial = utils.CycleNumber(dial+(number*utils.If(runic[0] == 'L', -1, 1)), 0, 99)

		if dial == 0 {
			seenZero++
		}
	}

	return seenZero
}

func p2(input string) int {
	dial := 50
	seenZero := 0
	for _, line := range strings.Fields(input) {
		runic := []rune(line)
		toRotate, err := strconv.Atoi(string(runic[1:]))

		if err != nil {
			panic(err)
		}

		numberDirection := utils.If(runic[0] == 'L', -1, 1)

		for i := 0; i < toRotate; i++ {
			dial = utils.CycleNumber(numberDirection+dial, 0, 99)
			if dial == 0 {
				seenZero++
			}
		}
	}

	return seenZero
}

func run(part2 bool, input string) any {
	if part2 {
		return p2(input)
	}

	return p1(input)
}
