package main

import (
	"aoc2025/lib"
	"fmt"
	"slices"
	"strings"
	"sync"
)

func main() {
	lines := lib.ReadInputAsLines(11)

	connections := map[string][]string{}
	for _, line := range lines {
		segments := strings.Split(line, ": ")
		connections[segments[0]] = []string{}
		for _, connection := range strings.Split(segments[1], " ") {
			connections[segments[0]] = append(connections[segments[0]], connection)
		}
	}

	youLabel := "you"
	endLabel := "out"
	pathsToEnd := findPathToTarget(connections, youLabel, endLabel, nil, map[string]int{})

	fmt.Println(pathsToEnd)

	//~~~~~ Part 2 ~~~~~//
	serverRackLabel := "svr"
	dacLabel := "dac"
	fftLabel := "fft"
	connectionsFromSvrToDac := findPathToTarget(connections, serverRackLabel, dacLabel, nil, map[string]int{})
	connectionsFromSvrToFft := findPathToTarget(connections, serverRackLabel, fftLabel, nil, map[string]int{})
	connectionsFromDacToFft := findPathToTarget(connections, dacLabel, fftLabel, nil, map[string]int{})
	connectionsFromFftToDac := findPathToTarget(connections, fftLabel, dacLabel, nil, map[string]int{})
	connectionsFromDacToEnd := findPathToTarget(connections, dacLabel, endLabel, nil, map[string]int{})
	connectionsFromFftToEnd := findPathToTarget(connections, fftLabel, endLabel, nil, map[string]int{})

	connectionsThroughDACAndFFT := connectionsFromSvrToDac*connectionsFromDacToFft*connectionsFromFftToEnd +
		connectionsFromSvrToFft*connectionsFromFftToDac*connectionsFromDacToEnd
	fmt.Println(connectionsThroughDACAndFFT)
}

var (
	knownMutex = sync.Mutex{}
)

func findPathToTarget(connections map[string][]string, currentLabel, targetLabel string, labelsVisited []string, knownRoutes map[string]int) int {
	knownMutex.Lock()
	if routes, exists := knownRoutes[currentLabel]; exists {
		knownMutex.Unlock()
		return routes
	}
	knownMutex.Unlock()

	if slices.Contains(labelsVisited, currentLabel) {
		return 0
	}

	if currentLabel == targetLabel {
		return 1
	}

	pathsToEnd := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for _, connection := range connections[currentLabel] {
		wg.Go(func() {
			newLabelsVisited := append(labelsVisited, currentLabel)
			pathsFound := findPathToTarget(connections, connection, targetLabel, newLabelsVisited, knownRoutes)
			m.Lock()
			pathsToEnd += pathsFound
			m.Unlock()
		})
	}
	wg.Wait()

	knownMutex.Lock()
	knownRoutes[currentLabel] = pathsToEnd
	knownMutex.Unlock()
	return pathsToEnd
}
