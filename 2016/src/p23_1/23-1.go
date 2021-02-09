package p23_1

import (
	"au"
	"fmt"
	"strings"
)

func testInputs() []string {
	return []string{
		"cpy 2 a",
		"tgl a",
		"tgl a",
		"tgl a",
		"cpy 1 a",
		"dec a",
		"dec a",
	}
}

func isRegister(value string) bool {
	return value == "a" || value == "b" || value == "c" || value == "d"
}

func getValue(registers map[string]int, value string) int {
	if isRegister(value) {
		return registers[value]
	}

	return au.ToNumber(value)
}

func Solve() {
	inputs := au.ReadInputAsStringArray("23")
	// inputs := testInputs()

	registers := map[string]int{}
	registers["a"] = 7

	index := 0
	for index < len(inputs) {
		tokens := strings.Split(inputs[index], " ")

		op := tokens[0]

		if op == "cpy" {
			if isRegister(tokens[2]) {
				registers[tokens[2]] = getValue(registers, tokens[1])
			}
			index++
		} else if op == "inc" {
			if isRegister(tokens[1]) {
				registers[tokens[1]] += 1
			}
			index++
		} else if op == "dec" {
			if isRegister(tokens[1]) {
				registers[tokens[1]] -= 1
			}
			index++
		} else if op == "jnz" {
			if getValue(registers, tokens[1]) != 0 {
				index += getValue(registers, tokens[2])
			} else {
				index++
			}
		} else if op == "tgl" {
			targetIndex := index + getValue(registers, tokens[1])

			if targetIndex >= 0 && targetIndex < len(inputs) {
				targetTokens := strings.Split(inputs[targetIndex], " ")

				if len(targetTokens) == 2 {
					if targetTokens[0] == "inc" {
						targetTokens[0] = "dec"
					} else {
						targetTokens[0] = "inc"
					}
				} else {
					if targetTokens[0] == "jnz" {
						targetTokens[0] = "cpy"
					} else {
						targetTokens[0] = "jnz"
					}
				}

				inputs[targetIndex] = strings.Join(targetTokens, " ")
			}

			index++
		}

		// fmt.Println(tokens)
		// fmt.Println(inputs)
		// fmt.Println(registers)
		// fmt.Println()
	}

	fmt.Println(registers["a"])
}
