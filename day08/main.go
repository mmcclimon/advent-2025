package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
)

type Point struct{ x, y, z int }

type Box struct {
	pos         Point
	dists       map[Point]float64
	connections []Point
}

func main() {
	boxes := make(map[Point]*Box)

	for line := range input.New().Lines() {
		ints := conv.ToInts(strings.Split(line, ","))
		pos := Point{ints[0], ints[1], ints[2]}
		boxes[pos] = &Box{
			pos:   pos,
			dists: make(map[Point]float64),
		}
	}

	makeConnection(boxes)
	makeConnection(boxes)
	makeConnection(boxes)
	makeConnection(boxes)

}

func makeConnection(boxes map[Point]*Box) {
	best := math.Inf(1)
	var bestPair []*Box

	for _, src := range boxes {
		for _, dst := range boxes {
			if src == dst || slices.Contains(src.connections, dst.pos) {
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
	a.connectTo(b)
	b.connectTo(a)

	fmt.Println("connect:", a, b)
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

func (b Box) String() string {
	return fmt.Sprintf("box@%s, l=%d", b.pos, len(b.dists))
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

func (b *Box) connectTo(other *Box) {
	delete(b.dists, other.pos)
	b.connections = append(b.connections, other.pos)
	b.connections = append(b.connections, other.connections...)

	fmt.Println(b, b.connections)
}
