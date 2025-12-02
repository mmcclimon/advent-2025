package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/assert"
	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
)

func main() {
	line := input.New().Slurp()

	part1, part2 := 0, 0

	for chunk := range strings.SplitSeq(line, ",") {
		nums := strings.Split(strings.TrimSpace(chunk), "-")
		left, right := nums[0], nums[1]

		for i := conv.Atoi(left); i <= conv.Atoi(right); i++ {
			s := fmt.Sprint(i)

			mid := len(s) / 2

			var ok []int

		Chunk:
			for size := 1; size <= mid; size++ {
				if len(s)%size != 0 {
					continue
				}

				chunked := split(s, size)

				for j := 1; j < len(chunked); j++ {
					if chunked[j] != chunked[0] {
						continue Chunk
					}
				}

				ok = append(ok, len(chunked))
			}

			if len(ok) > 0 {
				part2 += i
			}

			if slices.Contains(ok, 2) {
				part1 += i
			}
		}
	}

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func split(s string, size int) []string {
	assert.True(len(s)%size == 0, "string must divide evenly by size")

	var ret []string
	var cur string

	for _, c := range s {
		cur += string(c)
		if len(cur) == size {
			ret = append(ret, cur)
			cur = ""
		}
	}

	assert.True(cur == "", "ended up with extra chars??")

	return ret
}

// 11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124
