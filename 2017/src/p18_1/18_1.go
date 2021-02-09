package p18_1

import (
	"au"
	"fmt"
	"strconv"
	"strings"
)

func getValue(registers map[string]int, x string) int {
	numberValue, err := strconv.Atoi(x)
	if err == nil {
		return numberValue
	}

	registerValue, ok := registers[x]
	if ok {
		return registerValue
	} else {
		return 0
	}
}

func Solve() {
	input := au.ReadInputAsStringArray("18")

	registers := make(map[string]int)

	lastPlayed := 0
	lastRecovered := 0

	for index := 0; index < len(input); index++1 {
		tokens := strings.Split(input[index], " ")

		instruction := tokens[0]

		if instruction == "snd" {
			x := getValue(registers, tokens[1])
			lastPlayed = x
		} else if instruction == "set" {
			x := tokens[1]
			y := getValue(registers, tokens[2])

			registers[x] = y
		} else if instruction == "add" {
			x := tokens[1]
			y := getValue(registers, tokens[2])

			registers[x] += y
		} else if instruction == "mul" {
			x := tokens[1]
			y := getValue(registers, tokens[2])

			registers[x] *= y
		} else if instruction == "mod" {
			x := tokens[1]
			y := getValue(registers, tokens[2])

			registers[x] %= y
		} else if instruction == "rcv" {
			x := getValue(registers, tokens[1])
			if x != 0 {
				lastRecovered = lastPlayed
				break
			}
		} else if instruction == "jgz" {
			x := getValue(registers, tokens[1])
			y := getValue(registers, tokens[2])

			if x > 0 {
				index += (y - 1)
			}
		}
	}

	fmt.Println(lastRecovered)
}
