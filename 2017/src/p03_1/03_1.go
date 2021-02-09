package p03_1

import (
	"fmt"
	"math"
)

func Solve() {
	input := 347991

	x := 0
	y := 0

	direction := 1

	value := 1

	for width := 1; value < input; width++ {

		for dx := 0; dx < width; dx++ {
			x += direction
			value++

			if value == input {
				fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
				break
			}
		}

		direction = -1 * direction

		for dy := 0; dy < width; dy++ {
			y += direction
			value++

			if value == input {
				fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
				break
			}
		}
	}

}
