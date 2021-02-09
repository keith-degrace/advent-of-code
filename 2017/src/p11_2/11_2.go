package p11_2

import (
	"au"
	"fmt"
	"strings"
)

func Solve() {
	input := au.ReadInputAsString("11")
	// input := "ne,ne,sw,sw"

	moves := strings.Split(input, ",")

	maxDistance := 0

	current := []int{0, 0}
	for _, direction := range moves {
		if direction == "n" {
			current[1] -= 10
		}

		if direction == "ne" {
			current[0] += 5
			current[1] -= 5
		}

		if direction == "se" {
			current[0] += 5
			current[1] += 5
		}

		if direction == "s" {
			current[1] += 10
		}

		if direction == "sw" {
			current[0] -= 5
			current[1] += 5
		}

		if direction == "nw" {
			current[0] -= 5
			current[1] -= 5
		}

		distance := au.AbsInt(current[0]/10) + au.AbsInt(current[1]/10)
		if (current[0]%10 != 0) || (current[1]%10 != 0) {
			distance++
		}

		maxDistance = au.MaxInt(distance, maxDistance)
	}

	fmt.Println(maxDistance)
}
