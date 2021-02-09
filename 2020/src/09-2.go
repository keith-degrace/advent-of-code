package main

import (
	"au"
	"fmt"
)

type Puzzle09_2 struct {
}

func (p Puzzle09_2) run() {
	input := au.ReadInputAsNumberArray("09")

	number := 90433990

	for i := 0; i < len(input); i++ {

		sum := input[i]
		min := input[i]
		max := input[i]
		for j := i + 1; j < len(input); j++ {
			sum += input[j]

			if input[j] < min {
				min = input[j]
			}

			if input[j] > max {
				max = input[j]
			}

			if sum == number {
				fmt.Println(min + max)
			} else if sum > number {
				break
			}
		}
	}
}
