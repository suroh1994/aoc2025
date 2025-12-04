package lib

import "math"

// DigitsInNum returns the number of digits in a number. Only works on positive numbers.
func DigitsInNum(num int) int {
	// log10 only works from 2 onwards, assume we need to move at least one position
	if num < 2 {
		return 1
	}
	// add one to push 10 and 100 over to the next level
	return int(math.Ceil(math.Log10(float64(num + 1))))
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Pow(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}
