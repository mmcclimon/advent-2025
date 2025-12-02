package main

import (
	"fmt"
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

		for num := conv.Atoi(left); num <= conv.Atoi(right); num++ {
			s := fmt.Sprint(num)
			var good int

		Chunk:
			for size := len(s) / 2; size >= 1; size-- {
				chunks := split(s, size)
				numChunks := len(chunks)
				if numChunks == 0 {
					continue
				}

				for i := 1; i < numChunks; i++ {
					if chunks[i] != chunks[0] {
						continue Chunk
					}
				}

				good = numChunks
				break
			}

			switch good {
			case 0:
			case 2:
				part1 += num
				fallthrough
			default:
				part2 += num
			}
		}
	}

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func split(s string, size int) []string {
	if len(s)%size != 0 {
		return nil
	}

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
