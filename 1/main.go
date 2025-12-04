package main

import (
	"aoc2025/lib"
	"fmt"
)

func main() {
	instructions := lib.ReadInputAsLines(1)

	timesLandedOnZero := 0
	timesPassedOrLandedOnZero := 0

	position := 50
	for _, instruction := range instructions {
		// parse instruction
		amount := lib.MustParseToInt(instruction[1:])
		fullRotations := amount / 100
		amount %= 100
		timesPassedOrLandedOnZero += fullRotations

		if instruction[0] == 'L' {
			amount *= -1
		}

		// apply rotation
		position += amount
		if (position <= 0 && position != amount) || position >= 100 {
			timesPassedOrLandedOnZero++
		}

		// reset into value range 0-99
		position = (position + 100) % 100

		// count zero position
		if position == 0 {
			timesLandedOnZero++
		}
	}

	fmt.Println(timesLandedOnZero)
	fmt.Println(timesPassedOrLandedOnZero)
}
