package main

import (
	"au"
	"fmt"
	"sort"
)

type Puzzle10_1 struct {
}

func (p Puzzle10_1) run() {
	input := au.ReadInputAsNumberArray("10")

	sort.Ints(input)

	oneDifference := 0
	threeDifference := 1

	if input[0] == 1 {
		oneDifference++
	} else if input[0] == 3 {
		threeDifference++
	}

	for i := 1; i < len(input); i++ {
		difference := input[i] - input[i-1]
		if difference == 1 {
			oneDifference++
		} else if difference == 3 {
			threeDifference++
		}
	}

	fmt.Println(oneDifference * threeDifference)
}
