package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2025/advent/input"
)

const (
	Blank    = 0
	Splitter = -1
)

type Grid struct {
	g    map[rc]int
	maxR int
	maxC int
}

type rc struct{ r, c int }

func main() {
	grid := Grid{g: make(map[rc]int)}

	for r, line := range input.New().EnumerateLines() {
		for c, char := range line {
			var n int
			switch char {
			case 'S':
				n = 1
			case '^':
				n = Splitter
			}

			grid.g[rc{r, c}] = n
			grid.maxC = c
		}

		grid.maxR = r
	}

	numSplits, numPaths := 0, 0

	// Every position in the grid track the number of paths to that position. As
	// we go, we add up the number of splitters we hit (for part 1), and then at
	// the end we sum up all the paths to the bottom.
	for r := range grid.maxR + 1 {
		for c := range grid.maxC + 1 {
			incoming := grid.Get(r-1, c)

			// If there's nothing coming into this position, it's not interesting.
			if incoming <= 0 {
				continue
			}

			if grid.Get(r, c) == Splitter {
				// Each splitter splits the beam into its left and right
				numSplits++
				grid.AddBeam(r, c-1, incoming)
				grid.AddBeam(r, c+1, incoming)
			} else {
				grid.AddBeam(r, c, incoming)
			}
		}
	}

	// Now, sum up the number of paths.
	for c := range grid.maxC + 1 {
		numPaths += grid.Get(grid.maxR, c)
	}

	fmt.Println("part 1:", numSplits)
	fmt.Println("part 2:", numPaths)
}

func (g Grid) Print() {
	for r := range g.maxR + 1 {
		for c := range g.maxC + 1 {
			switch item := g.Get(r, c); item {
			case Blank:
				fmt.Print(".")
			case Splitter:
				fmt.Print("^")
			default:
				if item < 16 {
					fmt.Printf("%x", item)
				} else {
					fmt.Print("x")
				}
			}
		}
		fmt.Print("\n")
	}
}

func (g Grid) Get(r, c int) int {
	return g.g[rc{r, c}]
}

func (g Grid) AddBeam(r, c int, howMuch int) {
	if r < 0 || r > g.maxR || c < 0 || c > g.maxC {
		return
	}

	pos := rc{r, c}
	switch item := g.g[pos]; item {
	case Splitter:
	default:
		g.g[pos] += howMuch
	}
}
