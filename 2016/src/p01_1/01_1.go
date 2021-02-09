package p01_1

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

func advance(x int, y int, heading int, steps int) (int, int) {
	if heading == North {
		y += steps
	} else if heading == East {
		x += steps
	} else if heading == South {
		y -= steps
	} else {
		x -= steps
	}

	return x, y
}

func move(x int, y int, heading int, instruction Instruction) (int, int, int) {
	heading = turn(heading, instruction.direction)
	x, y = advance(x, y, heading, instruction.steps)
	return x, y, heading
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Solve() {
	instructions := readInstructions()
	// instructions := parseInstructions([]string { "R2", "L3" });
	// instructions := parseInstructions([]string { "R2", "R2", "R2" });
	// instructions := parseInstructions([]string { "R5", "L5", "R5", "R3" });

	x := 0
	y := 0
	heading := North

	for _, instruction := range instructions {
		x, y, heading = move(x, y, heading, instruction)
	}

	fmt.Println(abs(x) + abs(y))
}
