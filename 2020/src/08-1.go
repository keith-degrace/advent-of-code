package main

import (
	"au"
	"fmt"
	"strconv"
	"strings"
)

type Puzzle08_1 struct {
}

type Line08_1 struct {
	opcode  string
	value   int
	visited bool
}

func (p Puzzle08_1) parse(input []string) []Line08_1 {
	var lines []Line08_1

	for _, line := range input {
		parts := strings.Split(line, " ")

		opcode := parts[0]
		value, _ := strconv.Atoi(parts[1])

		lines = append(lines, Line08_1{opcode, value, false})
	}

	return lines
}

func (p Puzzle08_1) run() {
	input := au.ReadInputAsStringArray("08")

	lines := p.parse(input)

	current := 0
	acc := 0

	for !lines[current].visited {
		lines[current].visited = true

		if lines[current].opcode == "jmp" {
			current += lines[current].value
		} else {
			if lines[current].opcode == "acc" {
				acc += lines[current].value
			}

			current++
		}
	}

	fmt.Println(acc)
}
