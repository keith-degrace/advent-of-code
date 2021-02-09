package p05_2

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

		if input[index] >= 3 {
			input[index] -= 1
		} else {
			input[index] += 1
		}

		index = newIndex

		step++
	}

	fmt.Println(step)
}
