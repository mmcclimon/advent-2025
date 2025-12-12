package main

import (
	"fmt"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
)

// specific to my input: index => numPixels
var sizes = map[int]int{
	0: 7,
	1: 7,
	2: 6,
	3: 7,
	4: 5,
	5: 7,
}

func main() {
	var lines []string
	for hunk := range input.New().Hunks() {
		lines = hunk
	}

	good := 0

	// This was a stupid heuristic that turned out to be right in the end.
	for _, line := range lines {
		fields := strings.Fields(line)

		dimensions := conv.ToInts(strings.Split(strings.Trim(fields[0], ":"), "x"))
		maxSize := dimensions[0] * dimensions[1]

		var totalSize int
		for i, num := range conv.ToInts(fields[1:]) {
			totalSize += num * sizes[i]
		}

		if totalSize < maxSize {
			good++
		}
	}

	fmt.Println("part 1:", good)
}
