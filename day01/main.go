package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
	"github.com/mmcclimon/advent-2025/advent/mathx"
	"github.com/mmcclimon/advent-2025/advent/operator"
)

func main() {
	in := input.New()
	part1, part2 := 0, 0

	pos := 50

	for line := range in.Lines() {
		mult := operator.CrummyTernary(line[0] == 'L', -1, 1)
		n := conv.Atoi(line[1:])

		zerosPassed := n / 100
		rem := n % 100
		newPos := pos + (mult * rem)

		// The `pos != 0` check is so that we don't double-count zeros
		if pos != 0 && (newPos < 0 || newPos > 100) {
			zerosPassed++
		}

		pos = mathx.Mod(newPos, 100)
		part2 += zerosPassed

		if pos == 0 {
			part1++
			part2++
		}
	}

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}
