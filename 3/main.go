package main

import (
	"aoc2025/lib"
	"fmt"
)

func main() {
	batteryBankRunes := lib.ReadInputAsRuneMap()
	batteryBanks := runesToInt(batteryBankRunes)

	totalPower := 0
	for _, batteryBank := range batteryBanks {
		tens, ones := batteryBank[0], batteryBank[1]
		for _, joltage := range batteryBank[2:] {
			switch {
			case tens*10+ones < ones*10+joltage:
				tens = ones
				ones = joltage
			case tens*10+ones < tens*10+joltage:
				ones = joltage
			}
		}
		power := tens*10 + ones
		totalPower += power
	}
	fmt.Println(totalPower)
}

func runesToInt(runes [][]rune) [][]int {
	result := make([][]int, len(runes))
	for i, line := range runes {
		result[i] = make([]int, len(line))
		for j, r := range line {
			result[i][j] = int(r - 48)
		}
	}
	return result
}
