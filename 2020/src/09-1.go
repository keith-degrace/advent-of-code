package main

import (
	"au"
	"fmt"
)

type Puzzle09_1 struct {
}

func (p Puzzle09_1) hasSum(input []int, index int, preambleLength int) bool {
	start := index - preambleLength

	for i := 0; i < preambleLength-1; i++ {
		for j := i + 1; j < preambleLength; j++ {
			if (input[start+i] + input[start+j]) == input[index] {
				return true
			}
		}
	}

	return false
}

func (p Puzzle09_1) run() {
	input := au.ReadInputAsNumberArray("09")

	for i := 25; i < len(input); i++ {
		if !p.hasSum(input, i, 25) {
			fmt.Println(input[i])
		}
	}
}
