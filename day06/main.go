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

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var cols [][]int

	first := conv.ToInts(strings.Fields(lines[0]))
	for _, n := range first {
		cols = append(cols, []int{n})
	}

	for _, line := range lines[1 : len(lines)-1] {
		nums := conv.ToInts(strings.Fields(line))
		for i, n := range nums {
			cols[i] = append(cols[i], n)
		}
	}

	total := 0
	ops := strings.Fields(lines[len(lines)-1])
	for i, op := range ops {
		f := reducerFor(op)
		total += f(slices.Values(cols[i]))
	}

	fmt.Println("part 1:", total)

}

func part2(lines []string) {
	total := 0
	ops := strings.Fields(lines[len(lines)-1])

	var buf []int

	flush := func() {
		f := reducerFor(ops[len(ops)-1])
		total += f(slices.Values(buf))

		buf = nil
		ops = ops[:len(ops)-1]
	}

	for i := len(lines[0]) - 1; i >= 0; i-- {
		var s []byte
		for _, line := range lines[:len(lines)-1] {
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

func reducerFor(op string) func(iter.Seq[int]) int {
	switch op {
	case "+":
		return mathx.Sum
	case "*":
		return mathx.Product
	}

	panic("bad op: " + op)
}
