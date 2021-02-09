package main

import (
	"au"
	"fmt"
	"strings"
)

type Register [6]int
type Parameters struct {
	a int
	b int
	c int
}

type Instruction struct {
	opName string
	parameters Parameters
}

func parseInstruction(inputs string) Instruction {
	instruction := Instruction{}

	tokens := strings.Split(inputs, " ")

	instruction.opName = tokens[0]
	instruction.parameters.a = au.ToNumber(tokens[1])
	instruction.parameters.b = au.ToNumber(tokens[2])
	instruction.parameters.c = au.ToNumber(tokens[3])

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

func addr(register *Register, parameters Parameters) {
	register[parameters.c] = register[parameters.a] + register[parameters.b]
}

func addi(register *Register, parameters Parameters) {
	register[parameters.c] = register[parameters.a] + parameters.b
}

func mulr(register *Register, parameters Parameters) {
	register[parameters.c] = register[parameters.a] * register[parameters.b]
}

func muli(register *Register, parameters Parameters) {
	register[parameters.c] = register[parameters.a] * parameters.b
}

func banr(register *Register, parameters Parameters) {
	register[parameters.c] = register[parameters.a] & register[parameters.b]
}

func bani(register *Register, parameters Parameters) {
	register[parameters.c] = register[parameters.a] & parameters.b
}

func borr(register *Register, parameters Parameters) {
	register[parameters.c] = register[parameters.a] | register[parameters.b]
}

func bori(register *Register, parameters Parameters) {
	register[parameters.c] = register[parameters.a] | parameters.b
}

func setr(register *Register, parameters Parameters) {
	register[parameters.c] = register[parameters.a]
}

func seti(register *Register, parameters Parameters) {
	register[parameters.c] = parameters.a
}

func gtir(register *Register, parameters Parameters) {
	if parameters.a > register[parameters.b] {
		register[parameters.c] = 1
	} else {
		register[parameters.c] = 0
	}
}

func gtri(register *Register, parameters Parameters) {
	if register[parameters.a] > parameters.b {
		register[parameters.c] = 1
	} else {
		register[parameters.c] = 0
	}
}

func gtrr(register *Register, parameters Parameters) {
	if register[parameters.a] > register[parameters.b] {
		register[parameters.c] = 1
	} else {
		register[parameters.c] = 0
	}
}

func eqir(register *Register, parameters Parameters) {
	if parameters.a == register[parameters.b] {
		register[parameters.c] = 1
	} else {
		register[parameters.c] = 0
	}
}

func eqri(register *Register, parameters Parameters) {
	if register[parameters.a] == parameters.b {
		register[parameters.c] = 1
	} else {
		register[parameters.c] = 0
	}
}

func eqrr(register *Register, parameters Parameters) {
	if register[parameters.a] == register[parameters.b] {
		register[parameters.c] = 1
	} else {
		register[parameters.c] = 0
	}
}

func execute(instruction Instruction, register *Register) {
	switch instruction.opName {
		case "addr": addr(register, instruction.parameters)
		case "addi": addi(register, instruction.parameters)
		case "mulr": mulr(register, instruction.parameters)
		case "muli": muli(register, instruction.parameters)
		case "banr": banr(register, instruction.parameters)
		case "bani": bani(register, instruction.parameters)
		case "borr": borr(register, instruction.parameters)
		case "bori": bori(register, instruction.parameters)
		case "setr": setr(register, instruction.parameters)
		case "seti": seti(register, instruction.parameters)
		case "gtir": gtir(register, instruction.parameters)
		case "gtri": gtri(register, instruction.parameters)
		case "gtrr": gtrr(register, instruction.parameters)
		case "eqir": eqir(register, instruction.parameters)
		case "eqri": eqri(register, instruction.parameters)
		case "eqrr": eqrr(register, instruction.parameters)
	}
}

func formatRegister(register Register) string {
	return fmt.Sprintf("%10d %10d %10d %10d %10d", register[0], register[1], register[2], register[3], register[4])
}

func main() {
	inputs := au.ReadInputAsStringArray("19")

	ipRegister, instructions := parseInputs(inputs)

	register := Register{}
	register[0] = 1

	ip := 0
	count := 0

	for ip < len(instructions) {
		register[ipRegister] = ip

		// 
		/*
			Optimized this section

    		mulr 5 3 2
    		eqrr 2 4 2
    		addr 2 1 1
    		addi 1 1 1
    		addr 5 0 0
    		addi 3 1 3
    		gtrr 3 4 2
    		addr 1 2 1
    		seti 2 6 1
 
			Some  pseudo code
			{
					if (r3 * r5) == r4 {
							r0 = r0 + r5
					}

					r3 = r3 + 3

					if r3 > r4 {
							r2 = 1
							r1 = 11
					} else {
							r2 = 0
							r1 = 2
					}
			}

		*/
		if ip == 3 {
			if register[4] % register[5] == 0 {
				register[0] += register[5]
			}

			register[2] = 0
			register[3] = register[4]
			register[1] = 11

			// fmt.Printf("*%5d %3d %3v %10d %10d %10d %10d %10d %10d\n", count, ip, instructions[ip], register[0], register[1], register[2], register[3], register[4], register[5])
		} else {
			execute(instructions[ip], &register)
			// fmt.Printf(" %5d %3d %3v %10d %10d %10d %10d %10d %10d\n", count, ip, instructions[ip], register[0], register[1], register[2], register[3], register[4], register[5])
		}

		ip = register[ipRegister]
		ip++

		count++
	}

	fmt.Println(register[0])
}