package main

import (
	"au"
	"fmt"
	"strings"
)

type Register [6]int
type Parameters [3]int

type Instruction struct {
	opName string
	parameters Parameters
}

func parseInstruction(inputs string) Instruction {
	instruction := Instruction{}

	tokens := strings.Split(inputs, " ")

	instruction.opName = tokens[0]
	instruction.parameters[0] = au.ToNumber(tokens[1])
	instruction.parameters[1] = au.ToNumber(tokens[2])
	instruction.parameters[2] = au.ToNumber(tokens[3])

	return instruction
}

func parseInputs(inputs []string) (int, []Instruction) {
	ip := au.ToNumber(strings.Split(inputs[0], " ")[1])

	instructions := []Instruction{}

	for i := 1; i < len(inputs); i++ {
		instructions = append(instructions, parseInstruction(inputs[i]))
	}

	return ip, instructions
}

func addr(register Register, parameters Parameters) Register {
	register[parameters[2]] = register[parameters[0]] + register[parameters[1]]
	return register
}

func addi(register Register, parameters Parameters) Register {
	register[parameters[2]] = register[parameters[0]] + parameters[1]
	return register
}

func mulr(register Register, parameters Parameters) Register {
	register[parameters[2]] = register[parameters[0]] * register[parameters[1]]
	return register
}

func muli(register Register, parameters Parameters) Register {
	register[parameters[2]] = register[parameters[0]] * parameters[1]
	return register
}

func banr(register Register, parameters Parameters) Register {
	register[parameters[2]] = register[parameters[0]] & register[parameters[1]]
	return register
}

func bani(register Register, parameters Parameters) Register {
	register[parameters[2]] = register[parameters[0]] & parameters[1]
	return register
}

func borr(register Register, parameters Parameters) Register {
	register[parameters[2]] = register[parameters[0]] | register[parameters[1]]
	return register
}

func bori(register Register, parameters Parameters) Register {
	register[parameters[2]] = register[parameters[0]] | parameters[1]
	return register
}

func setr(register Register, parameters Parameters) Register {
	register[parameters[2]] = register[parameters[0]]
	return register
}

func seti(register Register, parameters Parameters) Register {
	register[parameters[2]] = parameters[0]
	return register
}

func gtir(register Register, parameters Parameters) Register {
	if parameters[0] > register[parameters[1]] {
		register[parameters[2]] = 1
	} else {
		register[parameters[2]] = 0
	}
	return register
}

func gtri(register Register, parameters Parameters) Register {
	if register[parameters[0]] > parameters[1] {
		register[parameters[2]] = 1
	} else {
		register[parameters[2]] = 0
	}
	return register
}

func gtrr(register Register, parameters Parameters) Register {
	if register[parameters[0]] > register[parameters[1]] {
		register[parameters[2]] = 1
	} else {
		register[parameters[2]] = 0
	}
	return register
}

func eqir(register Register, parameters Parameters) Register {
	if parameters[0] == register[parameters[1]] {
		register[parameters[2]] = 1
	} else {
		register[parameters[2]] = 0
	}
	return register
}

func eqri(register Register, parameters Parameters) Register {
	if register[parameters[0]] == parameters[1] {
		register[parameters[2]] = 1
	} else {
		register[parameters[2]] = 0
	}
	return register
}

func eqrr(register Register, parameters Parameters) Register {
	if register[parameters[0]] == register[parameters[1]] {
		register[parameters[2]] = 1
	} else {
		register[parameters[2]] = 0
	}
	return register
}

func execute(instruction Instruction, register Register) Register {
	switch instruction.opName {
		case "addr": return addr(register, instruction.parameters)
		case "addi": return addi(register, instruction.parameters)
		case "mulr": return mulr(register, instruction.parameters)
		case "muli": return muli(register, instruction.parameters)
		case "banr": return banr(register, instruction.parameters)
		case "bani": return bani(register, instruction.parameters)
		case "borr": return borr(register, instruction.parameters)
		case "bori": return bori(register, instruction.parameters)
		case "setr": return setr(register, instruction.parameters)
		case "seti": return seti(register, instruction.parameters)
		case "gtir": return gtir(register, instruction.parameters)
		case "gtri": return gtri(register, instruction.parameters)
		case "gtrr": return gtrr(register, instruction.parameters)
		case "eqir": return eqir(register, instruction.parameters)
		case "eqri": return eqri(register, instruction.parameters)
		case "eqrr": return eqrr(register, instruction.parameters)
	}

	panic(instruction)
}

func main() {
	inputs := au.ReadInputAsStringArray("19")

	ipRegister, instructions := parseInputs(inputs)

	register := Register{}

	ip := 0
	for ip < len(instructions) {
		register[ipRegister] = ip

		register = execute(instructions[ip], register)

		ip = register[ipRegister]
		ip++
	}

	fmt.Println(register[0])
}