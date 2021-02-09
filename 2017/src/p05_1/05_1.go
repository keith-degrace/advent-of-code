package p05_1

import (
	"au"
	"fmt"
)

func Solve() {
	input := au.ReadInputAsNumberArray("05")

	index := 0
	step := 0

	for index < len(input) {

		newIndex := index + input[index]
		input[index] += 1
		index = newIndex

		step++
	}

	fmt.Println(step)
}
