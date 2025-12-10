package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2025/advent/conv"
	"github.com/mmcclimon/advent-2025/advent/input"
	"github.com/mmcclimon/advent-2025/advent/mathx"
)

type Line struct {
	lights  []int
	buttons [][]int
	jolts   []int
}

func main() {
	var lines []Line
	for line := range input.New().Lines() {
		lines = append(lines, parseLine(line))
	}

	var total int
	for _, line := range lines {
		total += doLine(line)
	}

	fmt.Println("part 1:", total)
}

func parseLine(line string) Line {
	fields := strings.Fields(line)

	// first field is lights
	var lights []int
	for _, c := range strings.Trim(fields[0], "[]") {
		switch c {
		case '.':
			lights = append(lights, 0)
		case '#':
			lights = append(lights, 1)
		default:
			panic("bad char")
		}
	}

	// last field is joltage
	var jolts []int
	field := strings.Trim(fields[len(fields)-1], "{}")
	jolts = conv.ToInts(strings.Split(field, ","))

	var buttons [][]int
	for _, raw := range fields[1 : len(fields)-1] {
		button := strings.Trim(raw, "()")
		buttons = append(buttons, conv.ToInts(strings.Split(button, ",")))
	}

	return Line{
		lights:  lights,
		buttons: buttons,
		jolts:   jolts,
	}
}

func doLine(line Line) int {
	numBits := len(line.buttons)
	allOnes := mathx.Pow(2, numBits)
	var best int

	for n := range allOnes {
		s := fmt.Sprintf("%0*b", numBits, n)
		numPresses := strings.Count(s, "1")

		if best > 0 && numPresses >= best {
			continue
		}

		lights := make([]int, len(line.lights))

		for i, bit := range s {
			if bit == '0' {
				continue
			}

			button := line.buttons[i]
			for _, light := range button {
				lights[light] ^= 1
			}
		}

		if slices.Equal(lights, line.lights) {
			// fmt.Println("hit!", lights, s, numPresses)
			best = numPresses
		}
	}

	return best
}
