package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var re = regexp.MustCompile(`\s+`)

func parseLine(input string) []string {
	input = strings.TrimSpace(input)
	parts := re.Split(input, -1)
	return parts
}

func parseLineNum(input string) []int {
	input = strings.TrimSpace(input)
	parts := re.Split(input, -1)
	ints := make([]int, 0, len(parts))
	for _, v := range parts {
		ii, e := strconv.Atoi(v)
		if e != nil {
			panic(e)
		}
		ints = append(ints, ii)
	}
	return ints
}

func performAction(action string, a int, b int) int {
	switch action {
	case "*":
		return a * b
	case "+":
		return a + b

	}

	panic("unknown action")
}

func p1(input string) any {
	lines := strings.Split(input, "\n")

	action := parseLine(lines[len(lines)-1])
	nums := [][]int{}

	for i := 0; i < (len(lines) - 1); i++ {
		nums = append(nums, parseLineNum(lines[i]))
	}

	total := 0

	for ia := 0; ia < len(action); ia++ {
		subTotal := nums[0][ia]
		act := action[ia]

		for i := 1; i < len(nums); i++ {
			n2 := nums[i][ia]
			subTotal = performAction(act, subTotal, n2)
		}

		total += subTotal

	}

	return total
}

func isSymbol(v rune) bool {
	return v == '*' || v == '+'
}

func p2(input string) any {
	lines := strings.Split(input, "\n")

	actionLine := lines[len(lines)-1]
	lines = lines[:len(lines)-1]

	rrrr := []rune(actionLine)

	total := 0
	subTotal := 0

	currentAction := ' '

	for i := 0; i < len(rrrr); i++ {
		current := rrrr[i]

		number := ""

		for _, v := range lines {
			number += string(v[i])
		}

		if i+1 < len(rrrr) && isSymbol(rrrr[i+1]) {
			continue
		}

		nConv, _ := strconv.Atoi(strings.TrimSpace(number))

		if i+1 == len(rrrr) {
			subTotal = performAction(string(currentAction), subTotal, nConv)
			total += subTotal
		} else if isSymbol(current) {
			total += subTotal
			currentAction = current
			subTotal = nConv
		} else {
			subTotal = performAction(string(currentAction), subTotal, nConv)
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
