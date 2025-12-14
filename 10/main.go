package main

import (
	"aoc2025/lib"
	"fmt"
	"math"
	"strings"
	"sync"
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

	shortestPaths := 0
	completedCounter := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, machine := range machines {
		wg.Go(func() {
			buttonsPressed := make([]bool, len(machine.switches))
			stateMap := make(map[int]int)
			path := TracePathToTargetState(0, machine.targetState, machine.switches, buttonsPressed, stateMap)
			mu.Lock()
			shortestPaths += path
			completedCounter += 1
			fmt.Println(completedCounter)
			mu.Unlock()
		})
	}
	wg.Wait()

	fmt.Println(shortestPaths)
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

func TracePathToTargetState(currentState, targetState int, switches []int, pressed []bool, stateMap map[int]int) int {
	pressesSoFar := countPresses(pressed)
	if currentState == targetState {
		return pressesSoFar
	}

	if shortestPath, exists := stateMap[currentState]; exists && shortestPath < pressesSoFar {
		return math.MaxInt
	}
	stateMap[currentState] = pressesSoFar

	shortestPath := math.MaxInt
	for idx, button := range switches {
		if pressed[idx] {
			continue
		}

		buttonsPressed := make([]bool, len(pressed))
		copy(buttonsPressed, pressed)
		buttonsPressed[idx] = true
		shortestPath = min(shortestPath, TracePathToTargetState(currentState^button, targetState, switches, buttonsPressed, stateMap))
	}
	return shortestPath
}

func countPresses(pressedButtons []bool) int {
	count := 0
	for _, button := range pressedButtons {
		if button {
			count++
		}
	}
	return count
}
