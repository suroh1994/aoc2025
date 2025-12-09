package main

import (
	"aoc2025/lib"
	"fmt"
	"strings"
)

func main() {
	lines := lib.ReadInputAsLines(9)

	redTiles := make([]lib.Point2D, len(lines))
	for i, line := range lines {
		values := strings.Split(line, ",")
		redTiles[i] = lib.Point2D{X: lib.MustParseToInt(values[0]), Y: lib.MustParseToInt(values[1])}
	}

	maxArea := 0
	for i := range redTiles {
		for j := range redTiles[i+1:] {
			X := lib.Abs(redTiles[i].X-redTiles[j].X) + 1
			Y := lib.Abs(redTiles[i].Y-redTiles[j].Y) + 1
			maxArea = max(maxArea, X*Y)
		}
	}

	fmt.Println(maxArea)
}
