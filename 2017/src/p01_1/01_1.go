package p01_1

import (
	"au"
	"fmt"
)

func Solve() {
	input := au.ReadInputAsString("01")

	sum := 0

	for i := 0; i < len(input); i++ {
		number := au.ToNumber(string(input[i]))
		nextNumber := au.ToNumber(string(input[(i+1)%len(input)]))

		if number == nextNumber {
			sum += number
		}
	}

	fmt.Println(sum)
}
