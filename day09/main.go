package main

import (
	"fmt"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
	"github.com/mmcclimon/advent-2025/advent/mathx"
)

type xy struct{ x, y int }

func main() {
	var points []xy
	for line := range input.New().Lines() {
		ints := conv.ToInts(strings.Split(line, ","))
		points = append(points, xy{ints[0], ints[1]})
	}

	part1(points)
}

func part1(points []xy) {
	var best int

	for i, src := range points {
		for j := i + 1; j < len(points); j++ {
			dst := points[j]

			square := (1 + mathx.Abs(dst.x-src.x)) * (1 + mathx.Abs(dst.y-src.y))
			best = max(best, square)
		}
	}

	fmt.Println("part 1:", best)
}
