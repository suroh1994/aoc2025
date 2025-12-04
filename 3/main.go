package main

import (
	"aoc2025/lib"
	"fmt"
)

func main() {
	batteryBankRunes := lib.ReadInputAsRuneMap()
	batteryBanks := runesToInt(batteryBankRunes)

	totalPowerOne := 0
	totalPowerTwo := 0
	for _, batteryBank := range batteryBanks {
		powerOne := partOneCalculation(batteryBank)
		powerTwo := partTwoCalculation(batteryBank)
		totalPowerOne += powerOne
		totalPowerTwo += powerTwo
	}
	fmt.Println(totalPowerOne)
	fmt.Println(totalPowerTwo)
}

func partOneCalculation(batteryBank []int) int {
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
	return power
}

func partTwoCalculation(batteryBank []int) int {
	digits := make([]int, 12)
	copy(digits, batteryBank)

	for _, joltage := range batteryBank[12:] {
		digits = append(digits, joltage)
		leastValuableDigit := findLeastValuableIndex(digits)
		digits = append(digits[:leastValuableDigit], digits[leastValuableDigit+1:]...)
	}

	bankPower := 0
	for i, digit := range digits {
		bankPower += digit * lib.Pow(10, len(digits)-i-1)
	}

	return bankPower
}

func findLeastValuableIndex(digits []int) int {
	for i := 0; i < len(digits)-1; i++ {
		if digits[i+1] > digits[i] {
			return i
		}
	}
	return len(digits) - 1
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
