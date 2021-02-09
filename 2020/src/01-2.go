package main

import (
	"au"
	"fmt"
)

type Puzzle01_2 struct {
}

func (p Puzzle01_2) run() {
	inputs := au.ReadInputAsNumberArray("01")

	for i := 0; i < len(inputs); i++ {
		for j := i + 1; j < len(inputs); j++ {
			for k := j + 1; k < len(inputs); k++ {
				if inputs[i]+inputs[j]+inputs[k] == 2020 {
					fmt.Println(inputs[i] * inputs[j] * inputs[k])
				}
			}
		}
	}
}
