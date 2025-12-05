package main

import (
	"aoc2025/lib"
	"fmt"
	"slices"
	"strings"
)

type Range struct {
	Lower int
	Upper int
}

func main() {
	lines := lib.ReadInputAsLines(5)
	ranges := make([]Range, 0)

	i := 0
	for ; lines[i] != ""; i++ {
		boundaries := strings.Split(lines[i], "-")
		lower := lib.MustParseToInt(boundaries[0])
		upper := lib.MustParseToInt(boundaries[1])
		ranges = append(ranges, Range{Lower: lower, Upper: upper})
	}
	i++
	slices.SortFunc(ranges, SortRanges)

	freshIngredientCount := 0
	for ; i < len(lines); i++ {
		ingredient := lib.MustParseToInt(lines[i])

		isFresh := false
		for _, boundary := range ranges {
			if ingredient < boundary.Lower {
				continue
			}

			if boundary.Lower > ingredient {
				break
			}

			if boundary.Upper >= ingredient {
				isFresh = true
				break
			}
		}

		if isFresh {
			freshIngredientCount++
		}
	}

	fmt.Println(freshIngredientCount)

	//~~~~~~ Part 2 ~~~~~~
	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i].Upper >= ranges[i+1].Lower {
			ranges[i].Lower = min(ranges[i].Lower, ranges[i+1].Lower)
			ranges[i].Upper = max(ranges[i].Upper, ranges[i+1].Upper)
			ranges = append(ranges[:i+1], ranges[i+2:]...)
			i--
		}
	}

	validIdCount := 0
	for _, boundary := range ranges {
		validIdCount += boundary.Upper - boundary.Lower + 1
	}

	fmt.Println(validIdCount)
}

func SortRanges(i, j Range) int {
	return i.Lower - j.Lower
}
