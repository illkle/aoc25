package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Machine struct {
	desiredState []bool
	desiredNums  []int
	buttons      []Button
}

func (m Machine) isDesiredOne(someState []int) bool {
	for i := range someState {
		if (someState[i]%2 == 0) == m.desiredState[i] {
			return false
		}
	}
	return true
}

func (m Machine) isDesiredTwo(someState []int) (bool, bool) {
	for i := range someState {
		if someState[i] > m.desiredNums[i] {
			return false, true
		}

	}

	for i := range someState {
		if someState[i] != m.desiredNums[i] {
			return false, false
		}

	}
	return true, false
}

type Button struct {
	corrLight []int
}

func (b Button) updateState(someState []int) []int {
	ns := make([]int, len(someState))
	copy(ns, someState)

	for _, v := range b.corrLight {
		ns[v] += 1
	}
	return ns
}

type tempState struct {
	lights  []int
	steps   int
	pressed string
}

func findLowestPresses(m Machine) int {
	ts := tempState{}

	for range m.desiredState {
		ts.lights = append(ts.lights, 0)
	}

	q := []tempState{ts}

	for len(q) > 0 {
		el := q[0]
		q = q[1:]

		if m.isDesiredOne(el.lights) {
			return el.steps
		}

		for i, b := range m.buttons {
			q = append(q, tempState{
				lights:  b.updateState(el.lights),
				steps:   el.steps + 1,
				pressed: el.pressed + strconv.Itoa(i),
			})

		}
	}

	return -1

}

func findLowestPressesSolver(m Machine) int64 {
	lp := golp.NewLP(0, len(m.buttons))

	obj := make([]float64, len(m.buttons))
	for i := range m.buttons {
		obj[i] = 1
		lp.SetInt(i, true)
		lp.SetBounds(i, 0, math.Inf(1))
	}
	lp.SetObjFn(obj)

	buttonsThatIncrementLight := map[int][]int{}
	for buttonIndex, b := range m.buttons {
		for _, corrLightIndex := range b.corrLight {
			buttonsThatIncrementLight[corrLightIndex] = append(buttonsThatIncrementLight[corrLightIndex], buttonIndex)
		}
	}

	for lightIndex, desiredButtonVal := range m.desiredNums {
		used := make([]float64, len(m.buttons))

		for _, btnI := range buttonsThatIncrementLight[lightIndex] {
			used[btnI] = 1
		}

		lp.AddConstraint(used, golp.EQ, float64(desiredButtonVal))
	}

	ret := lp.Solve()
	if ret != golp.OPTIMAL {
		log.Fatalf("Failed to solve LP, status code: %v", ret)
	}

	return int64(math.Round(lp.Objective()))
}

func p1(machines []Machine) any {
	total := 0
	for _, m := range machines {
		lowest := findLowestPresses(m)
		total += lowest
	}

	return total
}

func p2(machines []Machine) any {
	total := int64(0)
	for _, m := range machines {
		lowest := findLowestPressesSolver(m)
		total += lowest
	}

	return total
}

var regBtns = regexp.MustCompile(`\((.*?)\)`)
var regMachine = regexp.MustCompile(`\[(.*?)\]`)
var regVoltage = regexp.MustCompile(`{(.*?)}`)

func run(part2 bool, input string) any {

	machines := []Machine{}

	for _, l := range strings.Split(input, "\n") {
		mmm := regMachine.FindStringSubmatch(l)
		lll := regBtns.FindAllStringSubmatch(l, -1)
		vvv := regVoltage.FindStringSubmatch(l)

		ma := Machine{}

		for _, v := range mmm[1] {
			ma.desiredState = append(ma.desiredState, v == '#')
		}

		splitVoltage := strings.Split(vvv[1], ",")
		for _, vv := range splitVoltage {
			nnnn, err := strconv.Atoi(vv)
			if err != nil {
				panic("lol")
			}
			ma.desiredNums = append(ma.desiredNums, nnnn)
		}

		for _, v := range lll {
			b := Button{}
			fg := v[1]
			split := strings.Split(fg, ",")

			for _, vv := range split {
				nnnn, err := strconv.Atoi(vv)
				if err != nil {
					panic("lol")
				}
				b.corrLight = append(b.corrLight, nnnn)
			}

			ma.buttons = append(ma.buttons, b)
		}
		machines = append(machines, ma)
	}

	if part2 {
		return p2(machines)
	}
	return p1(machines)
}
