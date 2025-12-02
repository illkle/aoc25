package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func splitStringToPieces(s string, pieceLen int) []string {
	pieces := make([]string, 0, len(s)/pieceLen)

	for sta, end := 0, pieceLen; end <= len(s); sta, end = sta+pieceLen, end+pieceLen {
		pieces = append(pieces, s[sta:end])
	}

	return pieces
}

func allPiecesEqual(sss []string) bool {
	for i := range sss {
		if i == 0 {
			continue
		}

		if sss[i] != sss[i-1] {
			return false
		}
	}
	return true
}

func isInvalidOne(num int) bool {
	stringified := strconv.Itoa(num)

	if len(stringified)%2 != 0 {
		return false
	}
	split := splitStringToPieces(stringified, len(stringified)/2)

	return allPiecesEqual(split)
}

func isInvalidTwo(num int) bool {
	stringified := strconv.Itoa(num)
	sLen := len(stringified)

	for sn := 1; sn <= (sLen / 2); sn++ {
		if len(stringified)%sn != 0 {
			continue
		}

		split := splitStringToPieces(stringified, sn)

		if allPiecesEqual(split) {
			return true
		}

	}
	return false

}

type sequence struct {
	start int
	end   int
}

func p1(input []sequence) any {
	invalidSum := 0

	for _, v := range input {
		for i := v.start; i <= v.end; i++ {
			if isInvalidOne(i) {
				invalidSum += i
			}
		}
	}

	return invalidSum
}

func p2(input []sequence) any {

	resChan := make(chan int, len(input))
	var wg sync.WaitGroup

	for _, v := range input {
		wg.Add(1)
		go func(v sequence) {
			defer wg.Done()
			for i := v.start; i <= v.end; i++ {
				if isInvalidTwo(i) {
					resChan <- i
				}
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	total := 0
	for r := range resChan {
		total += r
	}

	return total
}

func run(part2 bool, input string) any {
	spl := strings.Split(input, ",")
	seqs := make([]sequence, 0, len(spl))

	for _, line := range spl {
		byDash := strings.Split(line, "-")
		first, err := strconv.Atoi(byDash[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(byDash[1])
		if err != nil {
			panic(err)
		}

		seqs = append(seqs, sequence{
			start: first,
			end:   second,
		})

	}

	if part2 {
		return p2(seqs)
	}
	return p1(seqs)
}
