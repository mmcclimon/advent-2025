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

type Box struct {
	pos     Point
	dists   map[Point]float64
	direct  collections.Set[Point]
	circuit collections.Set[Point]
}

func main() {
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

	for i := range 1000 {
		fmt.Println("connection", i+1)
		makeConnection(boxes)
	}

	// calculate the circuits
	circuits := make(map[string]int)

	fmt.Println("walking circuits")
	// this is really dumb
	for _, box := range boxes {
		/*
			fmt.Printf("%q\n", box.pos)
			for pos := range box.direct.Iter() {
				fmt.Printf("%q -- %q\n", box.pos, pos)
			}
		*/
		box.walkCircuits(boxes)

		/*
			fmt.Printf("BOX %s\n", box.pos)
			for pos := range box.circuit.Iter() {
				fmt.Printf("  %s\n", pos)
			}
			fmt.Printf("  id: %s\n", box.circuitID())
		*/
	}

	for _, box := range boxes {
		circuits[box.circuitID()] = box.circuitLen()
		// fmt.Println(box)
	}

	ordered := slices.SortedFunc(maps.Keys(circuits), func(a, b string) int {
		return cmp.Compare(circuits[b], circuits[a])
	})

	total := 1

	for _, k := range ordered[:3] {
		total *= circuits[k]
	}

	fmt.Println("part 1:", total)

}

func makeConnection(boxes map[Point]*Box) {
	best := math.Inf(1)
	var bestPair []*Box

	for _, src := range boxes {
		for _, dst := range boxes {
			if src == dst || src.direct.Contains(dst.pos) {
				continue
			}

			dist := src.distanceTo(dst)
			if dist < best {
				best = dist
				bestPair = []*Box{src, dst}
			}
		}
	}

	a, b := bestPair[0], bestPair[1]
	// fmt.Println("connect:", a, b)
	a.connectTo(b)
	// b.connectTo(a)

	// fmt.Printf("  connections for %s: %v\n", a.pos, a.connections.ToSlice())
	// fmt.Printf("  connections for %s: %v\n", b.pos, b.connections.ToSlice())
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

// NB this changes as connections change
func (b *Box) circuitID() string {
	members := b.circuit.ToSlice()
	slices.SortFunc(members, func(a, b Point) int {
		if x := cmp.Compare(a.x, b.x); x != 0 {
			return x
		}

		if y := cmp.Compare(a.y, b.y); y != 0 {
			return y
		}

		return cmp.Compare(a.z, b.z)
	})

	strs := make([]string, len(members))
	for i, pos := range members {
		strs[i] = pos.String()
	}

	return strings.Join(strs, "|")
}

func (b *Box) circuitLen() int {
	return len(b.circuit)
}

func (box *Box) walkCircuits(boxes map[Point]*Box) {
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

	// fmt.Printf("for box %s:\n", box.pos)
	// for pos := range seen.Iter() {
	// 	fmt.Printf("  %s\n", pos)
	// }

	box.circuit = seen

	/*
			procedure DFS_iterative(G, v) is
		    let S be a stack
		    S.push(v)
		    while S is not empty do
		        v = S.pop()
		        if v is not labeled as discovered then
		            label v as discovered
		            for all edges from v to w in G.adjacentEdges(v) do
		                S.push(w)
	*/

}
