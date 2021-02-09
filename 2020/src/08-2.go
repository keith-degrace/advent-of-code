package main

import (
	"au"
	"fmt"
	"strconv"
	"strings"
)

type Puzzle08_2 struct {
}

type Line08_2 struct {
	opcode string
	value  int
}

func (p Puzzle08_2) parse(input []string) []Line08_2 {
	var lines []Line08_2

	for _, line := range input {
		parts := strings.Split(line, " ")

		opcode := parts[0]
		value, _ := strconv.Atoi(parts[1])

		lines = append(lines, Line08_2{opcode, value})
	}

	return lines
}

func (p Puzzle08_2) execute(lines []Line08_2) bool {
	current := 0
	acc := 0

	visited := make(map[int]bool)

	for {
		visited[current] = true

		if lines[current].opcode == "jmp" {
			current += lines[current].value
		} else {
			if lines[current].opcode == "acc" {
				acc += lines[current].value
			}

			current++
		}

		if current >= len(lines) {
			fmt.Println(acc)
			return true
		}

		_, ok := visited[current]
		if ok {
			return false
		}
	}
}

func (p Puzzle08_2) run() {
	input := au.ReadInputAsStringArray("08")

	lines := p.parse(input)

	for i := 0; i < len(lines); i++ {
		if lines[i].opcode == "jmp" {
			fmt.Println("Flipping jmp")
			lines[i].opcode = "nop"
			if p.execute(lines) {
				break
			}
			lines[i].opcode = "jmp"
		} else if lines[i].opcode == "nop" {
			fmt.Println("Flipping nop")
			lines[i].opcode = "jmp"
			if p.execute(lines) {
				break
			}
			lines[i].opcode = "nop"
		}
	}
}
