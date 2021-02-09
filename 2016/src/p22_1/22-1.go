package p22_1

import (
	"au"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Node struct {
	x     int
	y     int
	size  int
	used  int
	avail int
	use   int
}

func parse(input []string) []Node {
	var nodes []Node

	re := regexp.MustCompile("x([0-9]+)-y([0-9]+) *([0-9]+)T *([0-9]+)T *([0-9]+)T *([0-9]+)%")
	for _, line := range input {
		m := re.FindStringSubmatch(line)
		if m != nil {
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])
			size, _ := strconv.Atoi(m[3])
			used, _ := strconv.Atoi(m[4])
			avail, _ := strconv.Atoi(m[5])
			use, _ := strconv.Atoi(m[6])

			nodes = append(nodes, Node{x, y, size, used, avail, use})
		}
	}

	return nodes
}

func isViablePair(a Node, b Node) bool {
	// Node A is not empty (its Used is not zero).
	if a.used == 0 {
		return false
	}

	// Nodes A and B are not the same node.
	if a.x == b.x && a.y == b.y {
		return false
	}

	if a.used > b.avail {
		return false
	}

	// The data on node A (its Used) would fit on node B (its Avail).
	return true
}

func Solve() {
	fmt.Println("Starting\n")
	startTime := time.Now()

	input := au.ReadInputAsStringArray("22")

	nodes := parse(input)

	count := 0

	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes); j++ {
			if isViablePair(nodes[i], nodes[j]) {
				count++
			}
		}
	}

	fmt.Println(count)

	fmt.Println("\nCompleted in", time.Now().Sub(startTime))
}
