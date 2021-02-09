package main

import (
	"au"
	"fmt"
)

type Puzzle01_1 struct {
}

func (p Puzzle01_1) run() {
	inputs := au.ReadInputAsNumberArray("01")

	for i := 0; i < len(inputs); i++ {
		for j := i + 1; j < len(inputs); j++ {
			if inputs[i]+inputs[j] == 2020 {
				fmt.Println(inputs[i] * inputs[j])
			}
		}
	}
}
