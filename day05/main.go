package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
)

type Span struct{ low, high int }

func main() {
	hunks := slices.Collect(input.New().Hunks())
	var spans []*Span

	for _, line := range hunks[0] {
		nums := conv.ToInts(strings.Split(line, "-"))
		spans = append(spans, &Span{nums[0], nums[1]})
	}

	part1(spans, hunks[1])
	part2(spans)
}

func (s Span) Contains(n int) bool {
	return s.low <= n && n <= s.high
}

func (s Span) Total() int {
	return 1 + s.high - s.low
}

func (s Span) String() string {
	return fmt.Sprintf("%d-%d", s.low, s.high)
}

func part1(spans []*Span, lines []string) {
	total := 0

	for _, line := range lines {
		n := conv.Atoi(line)
		for _, r := range spans {
			if r.Contains(n) {
				total++
				break
			}
		}
	}

	fmt.Println("part 1:", total)
}

func part2(spans []*Span) {
	slices.SortFunc(spans, func(a, b *Span) int {
		if ret := cmp.Compare(a.low, b.low); ret != 0 {
			return ret
		}

		return cmp.Compare(a.high, b.high)
	})

	idx := 0

	for idx < len(spans)-1 {
		left, right := spans[idx], spans[idx+1]

		switch {
		case left.high < right.low:
			// we can't do anything with these
			idx++

		case right.low <= left.high:
			// the max here is in case the right span is entirely within the left
			left.high = max(left.high, right.high)
			spans = slices.Delete(spans, idx+1, idx+2)

		default:
			panic("wat")
		}
	}

	total := 0
	for _, span := range spans {
		total += span.Total()
	}

	fmt.Println("part 2:", total)
}
