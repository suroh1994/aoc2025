package lib

import "strconv"

func MustParseToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func RuneToInt(r rune) int {
	return int(r) - 48
}

func IntToRune(i int) rune {
	return rune(i + 48)
}
