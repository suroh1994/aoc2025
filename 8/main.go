package main

import (
	"aoc2025/lib"
	"fmt"
	"math"
	"slices"
	"strings"
	"sync"
)

func main() {
	coords := lib.ReadInputAsLines(8)

	type Point3D struct {
		X, Y, Z int
	}

	parsedCoords := make([]Point3D, 0, len(coords))
	for _, coord := range coords {
		dimensions := strings.Split(coord, ",")
		parsedCoords = append(parsedCoords, Point3D{
			X: lib.MustParseToInt(dimensions[0]),
			Y: lib.MustParseToInt(dimensions[1]),
			Z: lib.MustParseToInt(dimensions[2]),
		})
	}

	type Distance struct {
		idxA, idxB int
		distance   float64
	}

	distances := make([]Distance, 0)
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < len(parsedCoords); i++ {
		wg.Go(func() {
			for j := i + 1; j < len(parsedCoords); j++ {
				if i == j {
					continue
				}
				distance := math.Sqrt(
					float64(
						lib.Pow(parsedCoords[i].X-parsedCoords[j].X, 2) +
							lib.Pow(parsedCoords[i].Y-parsedCoords[j].Y, 2) +
							lib.Pow(parsedCoords[i].Z-parsedCoords[j].Z, 2),
					),
				)
				m.Lock()
				distances = append(distances, Distance{
					idxA:     i,
					idxB:     j,
					distance: distance,
				})
				m.Unlock()
			}
		})
	}
	wg.Wait()

	slices.SortFunc(distances, func(a, b Distance) int {
		if a.distance-b.distance < 0 {
			return -1
		}
		return 1
	})

	circuits := make([][]int, len(coords))
	for i := 0; i < len(coords); i++ {
		circuits[i] = []int{i}
	}

	for i, distance := range distances {
		if i == 1000 {
			break
		}

		var circuitsToMerge []int
		// look for id in circuit list
		for j, circuit := range circuits {
			if slices.Contains(circuit, distance.idxA) || slices.Contains(circuit, distance.idxB) {
				circuitsToMerge = append(circuitsToMerge, j)
			}
		}

		if len(circuitsToMerge) != 2 {
			continue
		}
		circuits[circuitsToMerge[0]] = append(circuits[circuitsToMerge[0]], circuits[circuitsToMerge[1]]...)
		circuits = append(circuits[:circuitsToMerge[1]], circuits[circuitsToMerge[1]+1:]...)
	}

	slices.SortFunc(circuits, func(a, b []int) int {
		return len(b) - len(a)
	})

	fmt.Println(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))
}
