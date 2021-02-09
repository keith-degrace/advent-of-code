package main

import (
	// "au"
	"fmt"
)

type Cell struct {
	power int
	squareLevels [300]int
}

func getPowerLevel(x int, y int, gridSerialNumber int) int {
	rackID := x + 10
	power := rackID * y
	power += gridSerialNumber
	power *= rackID
	power = (power % 1000) / 100
	power -= 5
	return power
}

func getSquareLevel(grid [][]Cell, x int, y int, size int) int {
	if size == 1 {
		return grid[x][y].power
	}

	level := grid[x][y].squareLevels[size-2]

	for offsetX := 0; offsetX < size - 1; offsetX++ {
		level += grid[x + offsetX][y + size - 1].power
	}

	for offsetY := 0; offsetY < size; offsetY++ {
		level += grid[x + size - 1][y + offsetY].power
	}

	grid[x][y].squareLevels[size-1] = level

	return level
}

func getMaxSquareLevel(grid [][]Cell, size int) (int, string) {
	maxResult := ""
	maxSquareLevel := 0

	for x := 0; x < 300 - size; x++ {
		for y := 0; y < 300 - size; y++ {
			squareLevel := getSquareLevel(grid, x, y, size)
			if squareLevel > maxSquareLevel {
				maxSquareLevel = squareLevel
				maxResult = fmt.Sprintf("%v,%v,%v", x+1, y+1, size)
			}
		}
	}

	return maxSquareLevel, maxResult
}

func main() {
	input := 9306
	// input := 18
	// input := 42

	grid := make([][]Cell, 300)
	for x := 0; x < 300; x++ {
		grid[x] = make([]Cell, 300)
	}

	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			grid[x][y].power = getPowerLevel(x+1, y+1, input)
			grid[x][y].squareLevels[0] = grid[x][y].power
		}
	}

	maxResult := ""
	maxSquareLevel := 0

	for size := 1; size <= 300; size++ {
		squareLevel, results := getMaxSquareLevel(grid, size)
		if squareLevel > maxSquareLevel {
			maxSquareLevel = squareLevel
			maxResult = results
		}
}

	fmt.Println(maxResult)
}
