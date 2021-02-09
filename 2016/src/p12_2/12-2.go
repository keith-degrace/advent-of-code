package p12_2

import (
	"au"
	"fmt"
	"strings"
)

func getValue(registers map [string] int, value string) int {
	if value == "a" || value == "b" || value == "c" || value == "d" {
		return registers[value]
	}

	return au.ToNumber(value)
}

func Solve() {
	inputs := au.ReadInputAsStringArray("12")

	registers := map [string] int{}
	registers["c"] = 1

	index := 0
	for index < len(inputs) {
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
