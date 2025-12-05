package main

import (
	"aoc2025/lib"
	"fmt"
	"sort"
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
	sort.Slice(ranges, func(i, j int) bool { return ranges[i].Lower < ranges[j].Lower })

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
}
