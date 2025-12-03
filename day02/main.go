package main

import (
	"fmt"
	"strings"

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

			for size := len(s) / 2; size >= 1; size-- {
				numChunks := len(s) / size
				if strings.Repeat(s[0:size], numChunks) != s {
					continue
				}

				if numChunks == 2 {
					part1 += num
				}
				part2 += num
				break
			}
		}
	}

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}
