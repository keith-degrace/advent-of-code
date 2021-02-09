package main

import (
	"au"
	"fmt"
	"strings"
)

func testInputs() []string {
	return []string {
		"inc a",
		"jio a, +2",
		"tpl a",
		"inc a",
	}
}

func parseInput(input string) (string, []string) {
	instruction := input[:3]

	params := []string{}
	for _,param := range strings.Split(input[4:], ",") {
		params = append(params, strings.TrimSpace(param))
	}
	
	return instruction, params
}

func main() {
	inputs := au.ReadInputAsStringArray("23")
	//inputs = testInputs()

	registers := map[string] int {}

	current := 0
	for current < len(inputs) {
		instruction, params := parseInput(inputs[current])

		if instruction == "hlf" {
			register := params[0]
			registers[register] = registers[register] / 2
			current++
		} else if instruction == "tpl" {
			register := params[0]
			registers[register] = registers[register] * 3
			current++
		} else if instruction == "inc" {
			register := params[0]
			registers[register]++
			current++
		} else if instruction == "jmp" {
			offset := au.ToNumber(params[0])
			current += offset
		} else if instruction == "jie" {
			register := params[0]
			if registers[register] % 2 == 0 {
				offset := au.ToNumber(params[1])
				current += offset
			} else {
				current++
			}
		} else if instruction == "jio" {
			register := params[0]
			if registers[register] == 1 {
				offset := au.ToNumber(params[1])
				current += offset
			} else {
				current++
			}
		}
	}

	fmt.Println(registers["b"])
}
