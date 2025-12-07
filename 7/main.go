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
}
