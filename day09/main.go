package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/assert"
	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
	"github.com/mmcclimon/advent-2025/advent/operator"
)

type xy struct{ x, y int }
type Rect struct{ left, right, top, bottom int }
type Edge struct{ left, right, top, bottom int }

func main() {
	var points []xy
	for line := range input.New().Lines() {
		ints := conv.ToInts(strings.Split(line, ","))
		points = append(points, xy{ints[0], ints[1]})
	}

	rects, edges := part1(points)
	part2(rects, edges)
}

//nolint:unused
func part1(points []xy) ([]Rect, []Edge) {
	var rects []Rect
	var edges []Edge
	for i, src := range points {
		// first, stash the edge since we're looping already and will need this later
		idx := operator.CrummyTernary(i+1 == len(points), 0, i+1)
		dst := points[idx]

		edges = append(edges, Edge{
			left:   min(src.x, dst.x),
			right:  max(src.x, dst.x),
			top:    min(src.y, dst.y),
			bottom: max(src.y, dst.y),
		})

		// then, grab all possible rectangles
		for j := i + 1; j < len(points); j++ {
			dst := points[j]

			rect := Rect{
				left:   min(src.x, dst.x),
				right:  max(src.x, dst.x),
				top:    min(src.y, dst.y),
				bottom: max(src.y, dst.y),
			}

			rects = append(rects, rect)
		}
	}

	slices.SortFunc(rects, func(a, b Rect) int {
		return cmp.Compare(b.Area(), a.Area())
	})

	slices.SortFunc(edges, func(a, b Edge) int {
		return cmp.Compare(b.Length(), a.Length())
	})

	fmt.Println("part 1:", rects[0].Area())
	return rects, edges
}

func (rect Rect) Area() int {
	return (1 + rect.right - rect.left) * (1 + rect.bottom - rect.top)
}

func (edge Edge) Length() int {
	return (edge.right - edge.left) + (edge.bottom - edge.top)
}

func part2(rects []Rect, edges []Edge) {
RECT:
	for _, rect := range rects {
		for _, edge := range edges {
			if edge.left < rect.right && edge.right > rect.left &&
				edge.top < rect.bottom && edge.bottom > rect.top {
				continue RECT
			}
		}

		fmt.Println("part 2:", rect.Area())
		return
	}
}

// nolint
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
