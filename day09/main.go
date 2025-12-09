package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/assert"
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

	// part1(points)
	printSVG(points, "out.svg")
	part2(points)
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

func printSVG(points []xy, filename string) {
	f, err := os.Create(filename)
	assert.Nil(err)
	defer f.Close()

	fmt.Fprintln(f, `<svg xmlns="http://www.w3.org/2000/svg">`)
	fmt.Fprintln(f, `<polygon stroke="none" fill="red" points="`)
	for _, point := range points {
		fmt.Fprintf(f, "%d,%d\n", point.x, point.y)
	}
	fmt.Fprintln(f, `" />`)
	fmt.Fprintln(f, `</svg>`)
}

func part2(points []xy) {
	return
	// p := points[0]
	// minX, maxX, minY, maxY := p.x, p.x, p.y, p.y

	for _, point := range points {
		fmt.Printf("%d,%d\n", point.x, point.y)

		/*
			idx := operator.CrummyTernary(i+1 == len(points), 0, i+1)
			next := points[idx]

			left, right := point, next
			if point.x == next.x && next.y < point.y {
				left, right = right, left
			} else if point.y == next.y && next.x < point.x {
				left, right = right, left
			}

			dist := math.Sqrt(
				math.Pow(float64(left.x-right.x), 2) +
					math.Pow(float64(left.y-right.y), 2),
			)
			// dumb special casing
			if dist > 2000 {
				fmt.Println(dist, left, right)

			}

			// This is stupid, but one of these loops will always be a no-op so who cares.
			for x := left.x; x <= right.x; x++ {
				for y := left.y; y <= right.y; y++ {
					allPoints.Add(xy{x, y})
				}
			}
		*/
	}

	// fmt.Println(minX, maxX, minY, maxY)

	return

	// fmt.Println(minX, maxX)
	// fmt.Println(minY, maxY)
}

// 9_363_561_880
