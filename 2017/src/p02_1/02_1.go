package p02_1

import (
	"au"
	"fmt"
	"math"
	"strings"
)

func Solve() {
	input := au.ReadInputAsStringArray("02")

	checksum := 0
	for _, line := range input {

		min := math.MaxInt16
		max := 0
		for _, token := range strings.Split(line, "\t") {
			number := au.ToNumber(token)
			min = au.MinInt(min, number)
			max = au.MaxInt(max, number)
		}

		checksum += max - min
	}

	fmt.Println(checksum)
}
