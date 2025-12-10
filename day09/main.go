package main

import (
	"cmp"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/assert"
	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
	"github.com/mmcclimon/advent-2025/advent/mathx"
	"github.com/mmcclimon/advent-2025/advent/operator"
)

type xy struct{ x, y int }
type Rect struct{ l, r xy }

func main() {
	var points []xy
	for line := range input.New().Lines() {
		ints := conv.ToInts(strings.Split(line, ","))
		points = append(points, xy{ints[0], ints[1]})
	}

	rects := part1(points)
	// printSVG(points, "out.svg")
	part2(points, rects)
}

//nolint:unused
func part1(points []xy) []Rect {
	all := make(map[Rect]int)

	for i, src := range points {
		for j := i + 1; j < len(points); j++ {
			dst := points[j]

			rect := Rect{
				l: xy{min(src.x, dst.x), min(src.y, dst.y)},
				r: xy{max(src.x, dst.x), max(src.y, dst.y)},
			}

			all[rect] = rect.Area()
		}
	}

	rects := slices.SortedFunc(maps.Keys(all), func(a, b Rect) int {
		return cmp.Compare(all[b], all[a])
	})

	fmt.Println("part 1:", rects[0].Area())
	return rects
}

func (rect Rect) Area() int {
	return (1 + rect.r.x - rect.l.x) * (1 + rect.r.y - rect.l.y)
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

func part2(points []xy, rects []Rect) {
	// p := points[0]
	// minX, maxX, minY, maxY := p.x, p.x, p.y, p.y
rectLoop:
	for _, rect := range rects {
		fmt.Println("check rect", rect, rect.Area())

		for i, point := range points {
			idx := operator.CrummyTernary(i+1 == len(points), 0, i+1)
			next := points[idx]

			a := xy{min(point.x, next.x), min(point.y, next.y)}
			b := xy{max(point.x, next.x), max(point.y, next.y)}

			fmt.Println("  check", a, b)

			// this line intersects with our rectangle
			if a.x < rect.r.x && // line starts left of rect's right edge
				a.y < rect.r.x && // line starts above rect's bottom edge
				b.x > rect.l.x && // line ends right of rect's left edge
				b.y > rect.l.y { // line ends below rect's top edge
				continue rectLoop
			}

		}

		fmt.Println("good?", rect.Area())
		break
	}

}

//nolint:unused
func part2bis(points []xy) {
	var holeY []int

	for i, point := range points {
		idx := operator.CrummyTernary(i+1 == len(points), 0, i+1)
		next := points[idx]

		dist := math.Sqrt(
			math.Pow(float64(point.x-next.x), 2) +
				math.Pow(float64(point.y-next.y), 2),
		)
		// dumb special casing
		if dist > 2000 {
			holeY = append(holeY, point.y)
			fmt.Println("horizontal:", point, next)
		}
	}

	slices.Sort(holeY)
	top, bottom := holeY[0], holeY[1]
	fmt.Println(top, bottom)

	var best int

	for i, src := range points {
		for j := i + 1; j < len(points); j++ {
			dst := points[j]

			inTop := src.y <= top && dst.y <= top
			inBottom := src.y >= bottom && dst.y >= bottom

			if !inTop && !inBottom {
				continue
			}

			// if src.y <= top && dst.y >= bottom || src.y >= bottom && dst.y <= top {
			// 	continue
			// }

			square := (1 + mathx.Abs(dst.x-src.x)) * (1 + mathx.Abs(dst.y-src.y))
			best = max(best, square)

			if best == square {
				fmt.Println(best, src, dst)
			}
		}
	}

	// fmt.Println(minX, maxX, minY, maxY)

	// fmt.Println(minX, maxX)
	// fmt.Println(minY, maxY)
}

func shouldSwap(a, b xy) bool {
	if a.x < b.x && a.y < b.y {
		return false
	}

	return a.x > b.x
}

// part 1:   4738108384
// too high: 3067899254
// too high: 3004559280
// too low: 103294448
