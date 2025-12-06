package main

import (
	"aoc2025/lib"
	"fmt"
	"sync"
)

func main() {
	input := lib.ReadInputAsRuneMap(4)

	buffer := make(chan lib.Point2D, len(input)*len(input[0]))

	accessibleRollCount := 0
	for x, row := range input {
		for y, val := range row {
			if val != '@' {
				continue
			}

			buffer <- lib.NewPoint2D(x, y)

			neighbouringRollCount := countNeighbouringRolls(input, x, y)
			if neighbouringRollCount < 4 {
				accessibleRollCount++
			}
		}
	}

	fmt.Println(accessibleRollCount)

	m := sync.Mutex{}
	wipMap := make(map[lib.Point2D]bool)
	wg := sync.WaitGroup{}

	workerTask := func() {
		for {
			if len(buffer) == 0 {
				return
			}

			select {
			case point := <-buffer:
				m.Lock()
				// drop queue entry if not a paper roll
				if input[point.X][point.Y] == '.' {
					m.Unlock()
					continue
				}

				// requeue if currently being processed
				if wipMap[point] {
					buffer <- point
					m.Unlock()
					continue
				}

				// mark this position as being processed
				wipMap[point] = true
				m.Unlock()

				// check whether this roll can be removed
				neighbouringRollCount := countNeighbouringRolls(input, point.X, point.Y)
				if neighbouringRollCount < 4 {
					// count roll as accessible and remove it from the map
					m.Lock()
					accessibleRollCount++
					m.Unlock()
					input[point.X][point.Y] = '.'

					// (re-)check all neighboring rolls
					for i := -1; i <= 1; i++ {
						for j := -1; j <= 1; j++ {
							neighbourPos := point.Add(lib.NewPoint2D(i, j))
							if lib.IsPosInBounds(input, neighbourPos) && input[neighbourPos.X][neighbourPos.Y] == '@' {
								buffer <- lib.NewPoint2D(point.X+i, point.Y+j)
							}
						}
					}
				}

				// work is done, release for the next check
				m.Lock()
				wipMap[point] = false
				m.Unlock()
			default:
				break
			}
		}
	}

	accessibleRollCount = 0
	for i := 0; i < 16; i++ {
		wg.Go(workerTask)
	}

	wg.Wait()
	fmt.Println(accessibleRollCount)
}

func countNeighbouringRolls(input [][]rune, x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if lib.IsInBounds(input, x+i, y+j) && input[x+i][y+j] == '@' {
				count++
			}
		}
	}
	return count
}
