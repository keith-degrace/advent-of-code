package p18_1

import (
	"au"
	"fmt"
	"time"
)

func isTrap(input string, pos int) bool {
	if pos < 0 || pos >= len(input) {
		return false
	}

	return input[pos] == '^'
}

func isTrapInNext(input string, pos int) bool {
	leftTrap := isTrap(input, pos - 1)
	centerTrap := isTrap(input, pos)
	rightTrap := isTrap(input, pos + 1)

	// Its left and center tiles are traps, but its right tile is not.
	if leftTrap && centerTrap && !rightTrap {
		return true
	}

	// Its center and right tiles are traps, but its left tile is not.
	if !leftTrap && centerTrap && rightTrap {
		return true
	}

	// Only its left tile is a trap.
	if leftTrap && !centerTrap && !rightTrap {
		return true
	}

	// Only its right tile is a trap.
	if !leftTrap && !centerTrap && rightTrap {
		return true
	}

	return false
}

func getNext(input string) string {
	next := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		if isTrapInNext(input, i) {
			next[i] = '^'
		} else {
			next[i] = '.'
		}
	}

	return string(next)
}

func getSafeTileCount(result []string) int {
	count := 0

	for _,line := range result {
		for _,char := range line {
			if char == '.' {
				count++
			}
		}
	}

	return count
}

func Solve() {
	fmt.Println("Starting\n")
	startTime := time.Now()

	input,rows := au.ReadInputAsString("18"), 40
	// input,rows  = "..^^.", 4
	// input,rows  = ".^^.^.^^^^", 10

	result := []string { input }
	for i := 0; i < rows - 1; i++ {
		result = append(result, getNext(result[len(result)-1]))
	}

	fmt.Println(getSafeTileCount(result))

	fmt.Println("\nCompleted in", time.Now().Sub(startTime))
}