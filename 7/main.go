package main

import (
	"aoc2025/lib"
	"fmt"
	"slices"
)

func main() {
	beamMap := lib.ReadInputAsRuneMap(7)

	start := lib.NewPoint2D(0, slices.Index(beamMap[0], 'S'))
	beams := []lib.Point2D{
		start.Add(lib.DOWN),
	}

	numOfSplits := 0
	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]
		if !lib.IsPosInBounds(beamMap, beam) {
			continue
		}

		if beamMap[beam.X][beam.Y] == '|' {
			//beam has already been followed
			continue
		}

		if beamMap[beam.X][beam.Y] == '.' {
			beamMap[beam.X][beam.Y] = '|'
			beams = append(beams, beam.Add(lib.DOWN))
			continue
		}

		if beamMap[beam.X][beam.Y] == '^' {
			beams = append(beams, beam.Add(lib.LEFT))
			beams = append(beams, beam.Add(lib.RIGHT))
			numOfSplits++
			continue
		}

		panic(fmt.Sprintf("Unknown rune encountered: %s", string(beamMap[beam.X][beam.Y])))
	}

	fmt.Println(numOfSplits)

	//~~~~~ Part 2 ~~~~~//
	timelineCount := followAllPaths(beamMap, start)
	fmt.Println(timelineCount)
}

var KnownTimelines = map[lib.Point2D]int{}

func followAllPaths(beamMap [][]rune, startPosition lib.Point2D) int {
	if timelineCount, isKnown := KnownTimelines[startPosition]; isKnown {
		return timelineCount
	}

	currentPosition := startPosition.Add(lib.DOWN)
	for lib.IsPosInBounds(beamMap, currentPosition) &&
		(beamMap[currentPosition.X][currentPosition.Y] == '.' ||
			beamMap[currentPosition.X][currentPosition.Y] == '|') {
		currentPosition = currentPosition.Add(lib.DOWN)
	}

	if !lib.IsPosInBounds(beamMap, currentPosition) {
		return 1
	}

	timelineCount := followAllPaths(beamMap, currentPosition.Add(lib.LEFT)) +
		followAllPaths(beamMap, currentPosition.Add(lib.RIGHT))

	KnownTimelines[startPosition] = timelineCount
	return timelineCount
}
