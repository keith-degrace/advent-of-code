package main

import (
	"au"
	"fmt"
)

type Puzzle03_1 struct {
}

func (p Puzzle03_1) isTree(input []string, x int, y int) bool {
	line := input[y]

	return string(line[x%len(line)]) == "#"
}

func (p Puzzle03_1) run() {
	input := au.ReadInputAsStringArray("03")

	treeCount := 0
	for y := 0; y < len(input); y++ {
		x := y * 3

		if p.isTree(input, x, y) {
			treeCount++
		}
	}

	fmt.Println(treeCount)
}
