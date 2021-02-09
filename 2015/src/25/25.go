package main

import (
	"fmt"
)

func getCode(x int, y int) int {
	code := 0

	for row := 1; ; row++ {
		for i := 0; i < row; i++ {
			xx := i + 1
			yy := row - i

			if xx == 1 && yy == 1 {
				code = 20151125
			} else {
				code = (code  * 252533) % 33554393
			}

			if xx == x && yy == y {
				return code
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(getCode(3075, 2981))
}
