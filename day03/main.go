package main

import (
	"fmt"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
)

func main() {
	part1 := 0

	for line := range input.New().Lines() {

	Tens:
		for tens := '9'; tens >= '1'; tens-- {
			idx := strings.IndexRune(line[:len(line)-1], tens)
			if idx == -1 {
				continue
			}

			for ones := '9'; ones >= '0'; ones-- {
				if strings.IndexRune(line[idx+1:], ones) != -1 {
					part1 += conv.Atoi(fmt.Sprintf("%c%c", tens, ones))
					break Tens
				}
			}
		}
	}

	fmt.Println(part1)
}
