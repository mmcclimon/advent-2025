package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mmcclimon/advent-2025/advent/assert"
	"github.com/mmcclimon/advent-2025/advent/collections"
	"github.com/mmcclimon/advent-2025/advent/input"
)

type Graph map[string][]string

func main() {
	graph := make(Graph)

	for line := range input.New().Lines() {
		fields := strings.Fields(line)
		src := strings.Trim(fields[0], ":")
		graph[src] = fields[1:]
	}

	topo := kahn(graph)

	fmt.Println("part 1:", walk(graph, topo, "you", "out"))

	var wg sync.WaitGroup
	var a, b, c int

	wg.Go(func() {
		a = walk(graph, topo, "svr", "fft")
		fmt.Println("svr -> fft:", a)
	})

	wg.Go(func() {
		b = walk(graph, topo, "fft", "dac")
		fmt.Println("fft -> dac:", b)
	})

	wg.Go(func() {
		c = walk(graph, topo, "dac", "out")
		fmt.Println("dac -> out:", c)
	})

	wg.Wait()

	fmt.Println("part 2:", a*b*c)
}

//nolint:unused
func printDot(graph Graph) {
	fmt.Println("digraph {")

	for in, out := range graph {
		for _, node := range out {
			fmt.Printf("%s -> %s\n", in, node)
		}
	}

	fmt.Println("}")
}

func walk(g Graph, topo map[string]int, start, end string) int {
	s := collections.NewDeque[string]()
	seen := collections.NewSet[string]()

	s.Append(start)
	numPaths := 0

	for s.Len() > 0 {
		node, err := s.Pop()
		assert.Nil(err)

		// prune: we can't ever get there from here.
		if topo[node] > topo[end] {
			continue
		}

		if node == end {
			numPaths++
			continue
		}

		if !seen.Contains(node) {
			seen.Add(node)
		}

		for _, other := range g[node] {
			s.Append(other)
		}
	}

	return numPaths
}

// This is a topological sort so we can prune the tree later.
func kahn(g Graph) map[string]int {
	rev := make(map[string]collections.Set[string])

	for k, v := range g {
		for _, node := range v {
			if len(rev[node]) == 0 {
				rev[node] = collections.NewSet[string]()
			}

			rev[node].Add(k)
		}
	}

	var l []string
	s := collections.NewSet("svr")

	for len(s) > 0 {
		node := s.Pop()
		l = append(l, node)

		for _, other := range g[node] {
			rev[other].Delete(node)
			if len(rev[other]) == 0 {
				s.Add(other)
			}
		}
	}

	ret := make(map[string]int)
	for i, n := range l {
		ret[n] = i
	}

	return ret
}
