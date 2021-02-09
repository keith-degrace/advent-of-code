package p02_1

import (
	"au"
	"fmt"
)

func applyInstructions(currentButton int, instruction string) int {

	for _, move := range instruction {
		switch move {
		case 'U':
			if currentButton > 3 {
				currentButton -= 3
			}
		case 'D':
			if currentButton < 7 {
				currentButton += 3
			}
		case 'L':
			if currentButton != 1 && currentButton != 4 && currentButton != 7 {
				currentButton -= 1
			}
		case 'R':
			if currentButton != 3 && currentButton != 6 && currentButton != 9 {
				currentButton += 1
			}
		}
	}

	return currentButton
}

func Solve() {
	instructions := au.ReadInputAsStringArray("02")
	// instructions := []string { "ULL", "RRDDD", "LURDL", "UUUUD" };

	currentButton := 5

	code := []int{}

	for _, instruction := range instructions {
		currentButton = applyInstructions(currentButton, instruction)
		code = append(code, currentButton)
	}

	fmt.Println(code)
}
