package main

import (
	"aoc2025/lib"
	"fmt"
	"strings"
)

func main() {
	lines := lib.ReadInputAsLines(10)

	type MachineDescription struct {
		targetState int
		switches    []int
	}

	machines := make([]MachineDescription, len(lines))
	for i, line := range lines {
		segments := strings.Split(line, " ")
		var lightCount int
		machines[i].targetState, lightCount = ParseMachineTargetState(segments[0])
		machines[i].switches = ParseSwitches(lightCount, segments[1:len(segments)-1])
	}

	for _, machine := range machines {
		//build statemachine
		//generate all paths from targetState to initialState and return shortest path
	}

	fmt.Println(machines)
}

func ParseMachineTargetState(state string) (int, int) {
	//discard [ ]
	state = state[1 : len(state)-1]
	target := 0
	for i := range state {
		if state[i] == '#' {
			target += 1 << (len(state) - 1 - i)
		}
	}
	return target, len(state)
}

func ParseSwitches(lightCount int, switches []string) []int {
	parsedSwitches := make([]int, len(switches))
	for i, singleSwitch := range switches {
		//discard ( )
		singleSwitch = singleSwitch[1 : len(singleSwitch)-1]
		lightsSwitched := strings.Split(singleSwitch, ",")
		for _, light := range lightsSwitched {
			parsedSwitches[i] += 1 << (lightCount - 1 - lib.MustParseToInt(light))
		}
	}
	return parsedSwitches
}
