package p02_2

import (
	"au"
	"fmt"
	"strings"
)

func Solve() {
	input := au.ReadInputAsStringArray("02")

	checksum := 0
	for _, line := range input {

		numbers := []int{}
		for _, token := range strings.Split(line, "\t") {
			numbers = append(numbers, au.ToNumber(token))
		}

		value := 0

		for i := 0; i < len(numbers); i++ {
			for j := 0; j < len(numbers); j++ {
				if i != j {
					if numbers[i]%numbers[j] == 0 {
						value = numbers[i] / numbers[j]
					}
				}
			}
		}

		checksum += value
	}

	fmt.Println(checksum)
}
