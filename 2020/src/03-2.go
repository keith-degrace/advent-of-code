package main

import (
	"au"
	"fmt"
)

type Puzzle03_2 struct {
}

func (p Puzzle03_2) isTree(input []string, x int, y int) bool {
	line := input[y]

	return string(line[x%len(line)]) == "#"
}

func (p Puzzle03_2) getTreeCount(input []string, xSlope int, ySlope int) int {
	treeCount := 0

	x := 0
	y := 0

	for y < len(input) {
		if p.isTree(input, x, y) {
			treeCount++
		}

		x += xSlope
		y += ySlope
	}

	return treeCount
}

func (p Puzzle03_2) run() {
	input := au.ReadInputAsStringArray("03")

	result := 1

	result *= p.getTreeCount(input, 1, 1)
	result *= p.getTreeCount(input, 3, 1)
	result *= p.getTreeCount(input, 5, 1)
	result *= p.getTreeCount(input, 7, 1)
	result *= p.getTreeCount(input, 1, 2)

	fmt.Println(result)
}
