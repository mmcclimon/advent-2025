package main

import (
	"bytes"
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
	"github.com/mmcclimon/advent-2025/advent/mathx"
)

func main() {
	lines := slices.Collect(input.New().Lines())
	ops := makeOps(lines[len(lines)-1])
	lines = lines[:len(lines)-1]

	part1(lines, ops)
	part2(lines, ops)
}

type Op struct {
	repr  string
	apply func(iter.Seq[int]) int
}

func (op Op) String() string {
	return op.repr
}

func makeOps(line string) []Op {
	var ret []Op

	for _, s := range strings.Fields(line) {
		op := Op{repr: s}

		switch s {
		case "+":
			op.apply = mathx.Sum
		case "*":
			op.apply = mathx.Product
		default:
			panic("bad op: " + s)
		}

		ret = append(ret, op)
	}

	return ret
}

func part1(lines []string, ops []Op) {
	total := 0

	split := make([][]int, len(lines))
	for i, line := range lines {
		split[i] = conv.ToInts(strings.Fields(line))
	}

	for fieldNum := range len(split[0]) {
		total += ops[fieldNum].apply(func(yield func(n int) bool) {
			for _, line := range split {
				yield(line[fieldNum])
			}
		})
	}

	fmt.Println("part 1:", total)
}

// The problem says that this math should be done right to left, but this
// doesn't matter even a little bit, so we just do it left to right.
func part2(lines []string, ops []Op) {
	total := 0

	var buf []int

	flush := func() {
		total += ops[0].apply(slices.Values(buf))
		buf = nil
		ops = ops[1:]
	}

	for i := range lines[0] {
		var s []byte
		for _, line := range lines {
			s = append(s, line[i])
		}

		s = bytes.TrimSpace(s)

		if len(s) == 0 {
			flush()
			continue
		}

		buf = append(buf, conv.Atoi(string(s)))
	}

	flush()

	fmt.Println("part 2:", total)
}
