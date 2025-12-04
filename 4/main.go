package main

import (
	"aoc2025/lib"
	"fmt"
)

func main() {
	input := lib.ReadInputAsRuneMap(4)

	accessibleRollCount := 0
	for x, row := range input {
		for y, val := range row {
			if val != '@' {
				continue
			}

			neighbouringRollCount := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}

					if lib.IsInBounds(input, x+i, y+j) && input[x+i][y+j] == '@' {
						neighbouringRollCount++
					}
				}
			}

			if neighbouringRollCount < 4 {
				accessibleRollCount++
			}
		}
	}

	fmt.Println(accessibleRollCount)
}
