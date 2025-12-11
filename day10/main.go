package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

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

	var wg sync.WaitGroup

	var part1 int
	var part2 atomic.Int64
	for _, line := range lines {
		part1 += doLights(line)
		wg.Go(func() {
			n := doJolts(line)
			part2.Add(int64(n))
		})
	}

	fmt.Println("part 1:", part1)
	wg.Wait()

	fmt.Println("part 2:", part2.Load())
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

func doLights(line Line) int {
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

func doJolts(line Line) int {
	// fmt.Println(line.jolts, line.buttons)

	maxes := make([]int, len(line.buttons))

	pressButton := func(buttonIdx int, numTimes int, jolts []int) {
		for _, idx := range line.buttons[buttonIdx] {
			jolts[idx] += numTimes
		}
	}

	for i, button := range line.buttons {
		jolts := make([]int, len(line.jolts))
		// fmt.Printf("  %d: determining max for %v\n", i, button)
	Max:
		for numPresses := 0; true; numPresses++ {
			for _, idx := range button {
				jolts[idx]++
				// fmt.Printf("  inc %d, now %v\n", idx, jolts)
				if jolts[idx] >= line.jolts[idx] {
					maxes[i] = numPresses + 1
					break Max
				}
			}
		}
	}

	var all [][]int

	// fmt.Println(maxes)

	for _, max := range maxes {
		var working []int
		for n := range max + 1 {
			working = append(working, n)
		}

		all = append(all, working)
	}

	var best int

	for seq := range Product(all...) {
		numPresses := mathx.Sum(slices.Values(seq))
		if best > 0 && numPresses > best {
			continue
		}

		jolts := make([]int, len(line.jolts))
		for i, n := range seq {
			pressButton(i, n, jolts)
		}

		if slices.Equal(jolts, line.jolts) {
			// fmt.Println("jolts", jolts, "presses=", seq)
			best = numPresses
		}
	}

	fmt.Println(best, line.jolts)
	return best
}

// Stolen from https://github.com/Skarlso/goitertools/blob/99fdc18feb0e/itertools/itertools.go
func Product(args ...[]int) iter.Seq[[]int] {

	pools := args
	npools := len(pools)
	indices := make([]int, npools)

	result := make([]int, npools)
	for i := range result {
		if len(pools[i]) == 0 {
			return nil
		}
		result[i] = pools[i][0]
	}

	// results := [][]int{result}

	return func(yield func([]int) bool) {
		for {
			i := npools - 1
			for ; i >= 0; i -= 1 {
				pool := pools[i]
				indices[i] += 1

				if indices[i] == len(pool) {
					indices[i] = 0
					result[i] = pool[0]
				} else {
					result[i] = pool[indices[i]]
					break
				}
			}

			if i < 0 {
				return
			}

			if !yield(result) {
				break
			}

			// newresult := make([]int, npools)
			// copy(newresult, result)
			// results = append(results, newresult)
		}
	}

	return nil
}
