package p12_1

import (
	"au"
	"fmt"
	"strings"
)

func testInputs() []string {
	return []string {
		"cpy 41 a",
		"inc a",
		"inc a",
		"dec a",
		"jnz a 2",
		"dec a",
	}
}

func getValue(registers map [string] int, value string) int {
	if value == "a" || value == "b" || value == "c" || value == "d" {
		return registers[value]
	}

	return au.ToNumber(value)
}

func Solve() {
	inputs := au.ReadInputAsStringArray("12")
	// inputs = testInputs()

	registers := map [string] int{}

	index := 0
	for index < len(inputs) {
		// fmt.Println(index, inputs[index], registers)

		tokens := strings.Split(inputs[index], " ")

		op := tokens[0]

		if op == "cpy" {
			registers[tokens[2]] = getValue(registers, tokens[1])
			index++
		} else if op == "inc" {
			registers[tokens[1]] += 1
			index++
		} else if op == "dec" {
			registers[tokens[1]] -= 1
			index++
		} else if op == "jnz" {
			if getValue(registers, tokens[1]) != 0 {
				index += au.ToNumber(tokens[2])
			} else {
				index++
			}
		}
	}

	fmt.Println(registers["a"])
}
