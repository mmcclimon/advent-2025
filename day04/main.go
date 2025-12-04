package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2025/advent/input"
)

type Grid map[rc]struct{}
type rc struct{ r, c int }

func main() {
	grid := make(Grid)
	for r, line := range input.New().EnumerateLines() {
		for c, char := range line {
			if char == '@' {
				grid[rc{r, c}] = struct{}{}
			}
		}
	}

	toRemove := grid.removableRolls()
	fmt.Println("part 1:", len(toRemove))

	part2 := 0

	for len(toRemove) > 0 {
		part2 += len(toRemove)

		for _, pos := range toRemove {
			delete(grid, pos)
		}

		toRemove = grid.removableRolls()
	}

	fmt.Println("part 2:", part2)
}

func (g Grid) removableRolls() []rc {
	var ret []rc
	for pos := range g {
		neighbors := g.neighborsFor(pos)
		if len(neighbors) < 4 {
			ret = append(ret, pos)
		}
	}

	return ret
}

// returns a list of positions occupied in the grid
func (g Grid) neighborsFor(pos rc) []rc {
	ret := make([]rc, 0, 8)

	for _, r := range []int{pos.r - 1, pos.r, pos.r + 1} {
		for _, c := range []int{pos.c - 1, pos.c, pos.c + 1} {
			check := rc{r, c}
			if _, ok := g[check]; ok && check != pos {
				ret = append(ret, check)
			}
		}
	}

	return ret
}
