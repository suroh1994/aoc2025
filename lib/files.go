package lib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(day int) string {
	content, err := os.ReadFile(fmt.Sprintf("%d/input", day))
	if err != nil {
		panic(fmt.Sprintf("'input' file missing! %v", err))
	}

	return string(content)
}

func ReadInputAsLines(day int) []string {
	input := ReadInput(day)
	separator := "\n"
	if strings.Contains(input, "\r") {
		separator = "\r\n"
	}
	return strings.Split(input, separator)
}

func ReadInputAsRuneMap(day int) [][]rune {
	lines := ReadInputAsLines(day)
	return LinesToRuneMap(lines)
}

func LinesToRuneMap(lines []string) [][]rune {
	runeMap := make([][]rune, len(lines))
	for i, line := range lines {
		runeMap[i] = []rune(line)
	}
	return runeMap
}

func ReadMultipleIntValuesPerLine(day int, delimiter string) [][]int {
	lines := ReadInputAsLines(day)
	values := make([][]int, len(lines))
	for idx, line := range lines {
		valuesInLine := strings.Split(line, delimiter)
		values[idx] = make([]int, len(valuesInLine))
		var err error
		for secIdx, singleValue := range valuesInLine {
			values[idx][secIdx], err = strconv.Atoi(singleValue)
			if err != nil {
				panic(err)
			}
		}
	}

	return values
}
