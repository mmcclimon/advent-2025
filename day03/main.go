package main

import (
	"fmt"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
)

func main() {
	part1, part2 := 0, 0

	for line := range input.New().Lines() {
		part1 += maximizeLine(line, 2)
		part2 += maximizeLine(line, 12)
	}

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func maximizeLine(line string, numDigits int) int {
	ret := make([]byte, numDigits)

	for i := range numDigits {
		line = maximize(line, numDigits-i)
		ret[i], line = line[0], line[1:]
	}

	return conv.Atoi(string(ret))
}

// digitLen is how much space we need to make sure to leave at the end
func maximize(s string, digitLen int) string {
	for i := '9'; i >= '0'; i-- {
		idx := strings.IndexRune(s[:len(s)-(digitLen-1)], i)
		if idx >= 0 {
			return s[idx:]
		}
	}

	panic("unreachable")
}
