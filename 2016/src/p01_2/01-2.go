package p01_2

import (
	"au"
	"fmt"
)

type Instruction struct {
	direction string
	steps     int
}

func parseInstructions(inputs []string) []Instruction {
	var instructions []Instruction
	for _, input := range inputs {
		direction := string(input[0])
		steps := au.ToNumber(input[1:])
		instructions = append(instructions, Instruction{direction, steps})
	}

	return instructions
}

func readInstructions() []Instruction {
	return parseInstructions(au.ReadInputAsCSVStringArray("01"))
}

const North = 0
const East = 1
const South = 2
const West = 3

func turn(heading int, dir string) int {
	if dir == "R" {
		if heading == 3 {
			return 0
		} else {
			return heading + 1
		}
	}

	if heading == 0 {
		return 3
	} else {
		return heading - 1
	}
}

func advance(x int, y int, heading int) (int, int) {
	if heading == North {
		y += 1
	} else if heading == East {
		x += 1
	} else if heading == South {
		y -= 1
	} else {
		x -= 1
	}

	return x, y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Solve() {
	instructions := readInstructions()
	// instructions := parseInstructions([]string { "R8", "R4", "R4", "R8" });

	x := 0
	y := 0
	heading := North

	history := map[string]int{}

	for _, instruction := range instructions {
		heading = turn(heading, instruction.direction)

		for i := 0; i < instruction.steps; i++ {
			x, y = advance(x, y, heading)

			historyKey := fmt.Sprintf("%v,%v", x, y)

			distance, ok := history[historyKey]
			if ok {
				fmt.Println(distance)
				return
			}

			history[historyKey] = abs(x) + abs(y)
		}
	}
}
