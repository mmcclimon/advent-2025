package main

import (
	"cmp"
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/collections"
	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
)

type Point struct{ x, y, z int }

type Pair struct{ left, right Point }

type Box struct {
	pos    Point
	dists  map[Point]float64
	direct collections.Set[Point]
}

type Registry map[Point]*Box

func main() {
	registry := readBoxes()
	toConnect := registry.makePairs()

	limit := 1000
	// special-case the example
	if len(registry) == 20 {
		limit = 10
	}

	for range limit {
		registry.connect(toConnect[0])
		toConnect = toConnect[1:]
	}

	// Part 1: find the 3 biggest circuits
	circuits := registry.makeCircuits()
	slices.Sort(circuits)
	slices.Reverse(circuits)

	part1 := circuits[0] * circuits[1] * circuits[2]
	fmt.Println("part 1:", part1)

	// Part 2: go unti there's only one
	var part2 int

	for len(circuits) > 1 {
		pair := toConnect[0]
		toConnect = toConnect[1:]

		registry.connect(pair)
		part2 = pair.left.x * pair.right.x
		circuits = registry.makeCircuits()
	}

	fmt.Println("part 2:", part2)
}

func readBoxes() Registry {
	boxes := make(map[Point]*Box)

	for line := range input.New().Lines() {
		ints := conv.ToInts(strings.Split(line, ","))
		pos := Point{ints[0], ints[1], ints[2]}
		boxes[pos] = &Box{
			pos:    pos,
			dists:  make(map[Point]float64),
			direct: collections.NewSet[Point](),
		}
	}
	return boxes
}

func (r Registry) GetPair(pair Pair) (*Box, *Box) {
	return r[pair.left], r[pair.right]
}

func (r Registry) makePairs() []Pair {
	dists := make(map[Pair]float64)

	// calculate distances
	slice := slices.Collect(maps.Values(r))
	for i := range slice {
		for j := i + 1; j < len(slice); j++ {
			left, right := slice[i], slice[j]
			pair := newPair(left.pos, right.pos)

			dist := left.distanceTo(right)
			dists[pair] = dist
		}
	}

	// then sort the pairs by distance
	return slices.SortedFunc(maps.Keys(dists), func(a, b Pair) int {
		return cmp.Compare(dists[a], dists[b])
	})
}

func (r Registry) makeCircuits() []int {
	var circuits []int

	todo := collections.NewSetFromIter(maps.Keys(r))
	for len(todo) > 0 {
		box := r[todo.Pop()]

		// walkCircuit returns the set of points in the circuit; we can delete
		// those from the todo list because we've already seen them all.
		circuit := box.walkCircuit(r)
		circuits = append(circuits, len(circuit))
		todo.DeleteIter(circuit.Iter())
	}

	return circuits
}

func (r Registry) connect(pair Pair) {
	r[pair.left].connectTo(r[pair.right])
}

func newPair(a, b Point) Pair {
	if sortPoints(a, b) < 0 {
		return Pair{a, b}
	}

	return Pair{b, a}
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

func (b Box) String() string {
	return fmt.Sprintf("box@%s", b.pos)
}

func (b *Box) distanceTo(other *Box) float64 {
	if b.pos == other.pos {
		return 0
	}

	if dist, ok := b.dists[other.pos]; ok {
		return dist
	}

	dist := math.Sqrt(
		math.Pow(float64(b.pos.x-other.pos.x), 2) +
			math.Pow(float64(b.pos.y-other.pos.y), 2) +
			math.Pow(float64(b.pos.z-other.pos.z), 2),
	)

	b.dists[other.pos] = dist
	return dist
}

func (src *Box) connectTo(dst *Box) {
	src.direct.Add(dst.pos)
	dst.direct.Add(src.pos)
}

func (box *Box) walkCircuit(boxes map[Point]*Box) collections.Set[Point] {
	// dfs on the edges
	s := collections.NewDeque[Point]()
	seen := collections.NewSet[Point]()

	s.Append(box.pos)

	for s.Len() > 0 {
		v, _ := s.Pop()
		if seen.Contains(v) {
			continue
		}

		seen.Add(v)
		for pos := range boxes[v].direct.Iter() {
			s.Append(pos)
		}
	}

	return seen
}

func sortPoints(a, b Point) int {
	if x := cmp.Compare(a.x, b.x); x != 0 {
		return x
	}

	if y := cmp.Compare(a.y, b.y); y != 0 {
		return y
	}

	return cmp.Compare(a.z, b.z)
}
