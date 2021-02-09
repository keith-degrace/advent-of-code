package p23_2

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

const Inc = 0
const Dec = 1
const Cpy = 2
const Jnz = 3
const Tgl = 4

type Instruction struct {
	op int

	argumentCount int

	isRegisterIn1 bool
	register1     byte
	value1        int

	isRegisterIn2 bool
	register2     byte
	value2        int
}

func (i Instruction) getValue1(registers map[byte]int) int {
	if i.isRegisterIn1 {
		return registers[i.register1]
	}

	return i.value1
}

func (i Instruction) getValue2(registers map[byte]int) int {
	if i.isRegisterIn2 {
		return registers[i.register2]
	}

	return i.value2
}

func load(inputs []string) []Instruction {
	var instructions []Instruction

	for _, line := range inputs {
		tokens := strings.Split(line, " ")

		var op int

		if tokens[0] == "cpy" {
			op = Cpy
		} else if tokens[0] == "inc" {
			op = Inc
		} else if tokens[0] == "dec" {
			op = Dec
		} else if tokens[0] == "jnz" {
			op = Jnz
		} else {
			op = Tgl
		}

		argumentCount := len(tokens) - 1

		isRegisterIn1 := tokens[1] == "a" || tokens[1] == "b" || tokens[1] == "c" || tokens[1] == "d"
		var register1 byte
		var value1 int
		if isRegisterIn1 {
			register1 = tokens[1][0]
		} else {
			value1 = au.ToNumber(tokens[1])
		}

		isRegisterIn2 := false
		var register2 byte
		var value2 int

		if argumentCount == 2 {
			isRegisterIn2 = tokens[2] == "a" || tokens[2] == "b" || tokens[2] == "c" || tokens[2] == "d"

			if isRegisterIn2 {
				register2 = tokens[2][0]
			} else {
				value2 = au.ToNumber(tokens[2])
			}
		}

		instructions = append(instructions, Instruction{op, argumentCount, isRegisterIn1, register1, value1, isRegisterIn2, register2, value2})
	}

	return instructions
}

func Solve() {
	inputs := au.ReadInputAsStringArray("23")
	// inputs := testInputs()

	instructions := load(inputs)

	registers := map[byte]int{}
	registers['a'] = 12

	index := 0
	for index < len(instructions) {

		switch instructions[index].op {

		case Cpy:
			if instructions[index].isRegisterIn2 {
				registers[instructions[index].register2] = instructions[index].getValue1(registers)
			}
			index++

			break

		case Inc:
			if instructions[index].isRegisterIn1 {
				registers[instructions[index].register1] += 1
			}
			index++

			break

		case Dec:
			if instructions[index].isRegisterIn1 {
				registers[instructions[index].register1] -= 1
			}
			index++

			break

		case Jnz:
			if instructions[index].getValue1(registers) != 0 {
				index += instructions[index].getValue2(registers)
			} else {
				index++
			}

			break

		case Tgl:
			targetIndex := index + instructions[index].getValue1(registers)

			if targetIndex >= 0 && targetIndex < len(inputs) {
				targetInstruction := instructions[targetIndex]

				if targetInstruction.argumentCount == 1 {
					if targetInstruction.op == Inc {
						targetInstruction.op = Dec
					} else {
						targetInstruction.op = Inc
					}
				} else {
					if targetInstruction.op == Jnz {
						targetInstruction.op = Cpy
					} else {
						targetInstruction.op = Jnz
					}
				}

				instructions[targetIndex] = targetInstruction
			}

			index++

			break

		}
	}

	fmt.Println(registers['a'])
}
