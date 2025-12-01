package main

import (
	"aoc2025/lib"
	"fmt"
)

func main() {
	instructions := lib.ReadInputAsLines()

	timesLandedOnZero := 0

	position := 50
	for _, instruction := range instructions {
		// parse instruction
		amount := lib.MustParseToInt(instruction[1:])
		if instruction[0] == 'L' {
			amount *= -1
		}

		// apply rotation and reset into value range 0-99
		position = (position + amount) % 100

		// count
		if position == 0 {
			timesLandedOnZero++
		}
	}

	fmt.Println(timesLandedOnZero)
}
