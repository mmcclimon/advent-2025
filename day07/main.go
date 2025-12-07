package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2025/advent/collections"
	"github.com/mmcclimon/advent-2025/advent/input"
)

type Grid struct {
	g    map[rc]rune
	maxR int
	maxC int
}

type rc struct{ r, c int }

func main() {
	grid := Grid{g: make(map[rc]rune)}

	for r, line := range input.New().EnumerateLines() {
		for c, char := range line {
			grid.g[rc{r, c}] = char
			grid.maxC = c
		}

		grid.maxR = r
	}

	numSplits := 0

	for r := range grid.maxR + 1 {
		nextLine := collections.NewSet[rc]()

		for c := range grid.maxC + 1 {
			switch char := grid.Get(r, c); char {
			case '.':
				// nothing to do
			case 'S', '|':
				nextLine.Add(rc{r + 1, c})
			case '^':
				if grid.Get(r-1, c) != '|' {
					// this splitter only splits if hit by a beam
					continue
				}

				numSplits++
				grid.AddBeam(r, c-1)
				grid.AddBeam(r, c+1)
				nextLine.Add(rc{r + 1, c - 1}, rc{r + 1, c + 1})
			}
		}

		for next := range nextLine.Iter() {
			grid.AddBeam(next.r, next.c)
		}
	}

	fmt.Println("part 1:", numSplits)
}

func (g Grid) Print() {
	for r := range g.maxR + 1 {
		for c := range g.maxC + 1 {
			fmt.Print(string(g.g[rc{r, c}]))
		}
		fmt.Print("\n")
	}
}

func (g Grid) Get(r, c int) rune {
	return g.g[rc{r, c}]
}

func (g Grid) AddBeam(r, c int) {
	if r < 0 || r > g.maxR || c < 0 || c > g.maxC {
		return
	}

	pos := rc{r, c}
	switch char := g.g[pos]; char {
	case '^', 'S':
	case '.', '|':
		g.g[pos] = '|'
	}
}
