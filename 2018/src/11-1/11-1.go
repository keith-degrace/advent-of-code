package main

import (
	// "au"
	"fmt"
)

func getPowerLevel(x int, y int, gridSerialNumber int) int {
	rackID := x + 10
	power := rackID * y
	power += gridSerialNumber
	power *= rackID
	power = (power % 1000) / 100
	power -= 5
	return power
}

func getSquareLevel(grid [][]int, x int, y int) int {
	level := 0

	for xx := 0; xx < 3; xx++ {
		for yy := 0; yy < 3; yy++ {
			level += grid[xx+x][yy+y]
		}
	}

	return level
}

func main() {
	input := 9306
	// input := 18
	// input := 42

	// fmt.Println(getPowerLevel(3, 5, 8))
	// fmt.Println(getPowerLevel(122, 79, 57))
	// fmt.Println(getPowerLevel(217, 196, 39))
	// fmt.Println(getPowerLevel(101, 153, 71))

	grid := make([][]int, 300)
	for x := 0; x < 300; x++ {
		grid[x] = make([]int, 300)
	}

	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			grid[x][y] = getPowerLevel(x+1, y+1, input)
		}
	}

	maxXY := ""
	maxSquareLevel := 0
	for x := 0; x < 297; x++ {
		for y := 0; y < 297; y++ {
			squareLevel := getSquareLevel(grid, x, y)
			if squareLevel > maxSquareLevel {
				maxSquareLevel = squareLevel
				maxXY = fmt.Sprintf("%v,%v", x+1, y+1)
			}
		}
	}

	fmt.Println(maxXY)
}
