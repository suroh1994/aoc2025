package main

import (
	"aoc2025/lib"
	"fmt"
	"slices"
	"strings"
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

	startLabel := "you"
	pathsToEnd := findPathToEnd(connections, startLabel, nil)

	fmt.Println(pathsToEnd)
}

func findPathToEnd(connections map[string][]string, currentLabel string, labelsVisited []string) int {
	if slices.Contains(labelsVisited, currentLabel) {
		return 0
	}

	if currentLabel == "out" {
		return 1
	}

	pathsToEnd := 0
	for _, connection := range connections[currentLabel] {
		newLabelsVisited := append(labelsVisited, currentLabel)
		pathsToEnd += findPathToEnd(connections, connection, newLabelsVisited)
	}

	return pathsToEnd
}
