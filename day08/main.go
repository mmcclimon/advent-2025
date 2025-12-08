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
	pos         Point
	dists       map[Point]float64
	connections collections.Set[Point]
}

func main() {
	boxes := make(map[Point]*Box)

	for line := range input.New().Lines() {
		ints := conv.ToInts(strings.Split(line, ","))
		pos := Point{ints[0], ints[1], ints[2]}
		boxes[pos] = &Box{
			pos:         pos,
			dists:       make(map[Point]float64),
			connections: collections.NewSet[Point](),
		}
	}

	for range 4 {
		makeConnection(boxes)
	}

	fmt.Println("---")

	// calculate the circuits
	circuits := make(map[string]int)

	// this is really dumb
	for _, box := range boxes {
		box.trimConnections()
		// fmt.Println(box)
	}

	for _, box := range boxes {
		circuits[box.circuitID()] = box.circuitLen()
		// fmt.Println(box)
	}

	ordered := slices.SortedFunc(maps.Keys(circuits), func(a, b string) int {
		return cmp.Compare(circuits[b], circuits[a])
	})

	for _, k := range ordered {
		fmt.Println(circuits[k], k)
	}

}

func makeConnection(boxes map[Point]*Box) {
	best := math.Inf(1)
	var bestPair []*Box

	for _, src := range boxes {
		for _, dst := range boxes {
			if src == dst {
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
	fmt.Println("connect:", a, b)
	a.connectTo(boxes, b)
	// b.connectTo(a)

	fmt.Printf("  connections for %s: %v\n", a.pos, a.connections.ToSlice())
	fmt.Printf("  connections for %s: %v\n", b.pos, b.connections.ToSlice())
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

func (src *Box) connectTo(boxes map[Point]*Box, dst *Box) {
	delete(src.dists, dst.pos)
	delete(dst.dists, src.pos)

	for pos := range dst.connections.Iter() {
		conn := boxes[pos]
		src.addConnectionTo(conn)
		// conn.addConnectionTo(src)
		// b.connections.Add(conn.pos)
		// fmt.Printf("    adding %s to connections for %s\n", b.pos, conn.pos)
		// conn.connections.Add(b.pos)
	}

	for pos := range src.connections.Iter() {
		conn := boxes[pos]
		conn.addConnectionTo(dst)
	}

	src.addConnectionTo(dst)

	// b.connections.Add(other.pos)
	// other.connections.Add(b.pos)
	// b.connections.AddIter(other.connections.Iter())

	// other.connections.AddIter(b.connections.Iter())
	// b.trimConnections()
	// other.trimConnections()
}

func (b *Box) addConnectionTo(other *Box) {
	fmt.Printf("    %s: adding connection to %s\n", b.pos, other.pos)
	b.connections.Add(other.pos)
	other.connections.Add(b.pos)
}

func (b *Box) trimConnections() {
	b.connections.Delete(b.pos)
}

// NB this changes as connections change
func (b *Box) circuitID() string {
	members := append([]Point{b.pos}, b.connections.ToSlice()...)
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
	return 1 + len(b.connections)
}
