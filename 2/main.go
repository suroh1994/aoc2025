package main

import (
	"aoc2025/lib"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	input := lib.ReadInput(2)
	ranges := strings.Split(input, ",")

	invalidTotal := 0

	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for _, idRange := range ranges {
		wg.Go(func() {
			boundaries := strings.Split(idRange, "-")
			lower := lib.MustParseToInt(boundaries[0])
			upper := lib.MustParseToInt(boundaries[1])

			for id := lower; id <= upper; id++ {
				// IDs with an odd number of digits cannot be invalid
				if lib.Log10Int(id)%2 == 1 {
					continue
				}

				idString := strconv.Itoa(id)
				invertedIdString := idString[len(idString)/2:] + idString[:len(idString)/2]
				if idString == invertedIdString {
					m.Lock()
					invalidTotal += id
					m.Unlock()
				}
			}
		})
	}
	wg.Wait()
	fmt.Println(invalidTotal)
}
