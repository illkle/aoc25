package utils

func CycleNumber(number, min, max int) int {
	if min >= max {
		panic("Min cannot be over max")
	}

	rr := max - min + 1

	if number < min {
		return max - ((min - number - 1) % rr)
	} else if number > max {
		return min + ((number - max - 1) % rr)
	}

	return number

}

func If[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
