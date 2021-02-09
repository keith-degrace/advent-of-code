package main

import (
	"au"
	"fmt"
	"regexp"
	"strings"
)

type Register [4]int
type Parameters [3]int

type Instruction struct {
	opcode int
	parameters Parameters
}

type Sample struct {
	instruction Instruction
	before Register
	after Register
}

func parseInstruction(inputs string) Instruction {
	instruction := Instruction{}

	tokens := strings.Split(inputs, " ")

	instruction.opcode = au.ToNumber(tokens[0])
	instruction.parameters[0] = au.ToNumber(tokens[1])
	instruction.parameters[1] = au.ToNumber(tokens[2])
	instruction.parameters[2] = au.ToNumber(tokens[3])

	return instruction
}

func parseInputs(inputs []string) ([]Sample, []Instruction) {
	samples := make([]Sample, 0)

	re := regexp.MustCompile(".*\\[([0-9]*), ([0-9]*), ([0-9]*), ([0-9]*)]")

	index := 0
	for index < len(inputs) {
		sample := Sample {}

		beforeMatches := re.FindStringSubmatch(inputs[index])
		if len(beforeMatches) == 0 {
			break
		}

		sample.before[0] = au.ToNumber(beforeMatches[1])
		sample.before[1] = au.ToNumber(beforeMatches[2])
		sample.before[2] = au.ToNumber(beforeMatches[3])
		sample.before[3] = au.ToNumber(beforeMatches[4])

		sample.instruction = parseInstruction(inputs[index + 1])

		afterMatches := re.FindStringSubmatch(inputs[index + 2])
		sample.after[0] = au.ToNumber(afterMatches[1])
		sample.after[1] = au.ToNumber(afterMatches[2])
		sample.after[2] = au.ToNumber(afterMatches[3])
		sample.after[3] = au.ToNumber(afterMatches[4])

		samples = append(samples, sample)

		index += 3
	}

	program := []Instruction{}
	for index < len(inputs) {
		program = append(program, parseInstruction(inputs[index]))
		index++
	}

	return samples, program
}

func registersEqual(register1 Register, register2 Register) bool {
	return register1[0] == register2[0] &&
		     register1[1] == register2[1] &&
		     register1[2] == register2[2] &&
		     register1[3] == register2[3]
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

func executeOp(opName string, register Register, parameters Parameters) Register {
	switch opName {
		case "addr": return addr(register, parameters)
		case "addi": return addi(register, parameters)
		case "mulr": return mulr(register, parameters)
		case "muli": return muli(register, parameters)
		case "banr": return banr(register, parameters)
		case "bani": return bani(register, parameters)
		case "borr": return borr(register, parameters)
		case "bori": return bori(register, parameters)
		case "setr": return setr(register, parameters)
		case "seti": return seti(register, parameters)
		case "gtir": return gtir(register, parameters)
		case "gtri": return gtri(register, parameters)
		case "gtrr": return gtrr(register, parameters)
		case "eqir": return eqir(register, parameters)
		case "eqri": return eqri(register, parameters)
		case "eqrr": return eqrr(register, parameters)
	}

	panic(opName)
}

func getOpNames() []string {
	return []string {
		"addr", "addi",
		"mulr", "muli",
		"banr", "bani",
		"borr", "bori",
		"setr", "seti",
		"gtir", "gtri", "gtrr",
		"eqir", "eqri", "eqrr",
	}
}

func getOpMap(samples []Sample) map[int] string {
	opPossibilityMap := map[string] map [int] bool{}

	for _,sample := range samples {
		for _,opName := range getOpNames() {
			result := executeOp(opName, sample.before, sample.instruction.parameters)
			if registersEqual(result, sample.after) {
				if opPossibilityMap[opName] == nil {
					opPossibilityMap[opName] = map [int] bool{}
				}
				opPossibilityMap[opName][sample.instruction.opcode] = true
			}
		}
	}

	opMap := map[int] string{}

	for len(opPossibilityMap) > 0 {
		resolvedOpCode := -1
		for opName,opCodes := range opPossibilityMap {
			if len(opCodes) == 1 {
				for opCode := range opCodes {
					opMap[opCode] = opName
					delete(opPossibilityMap, opName)
					resolvedOpCode = opCode
				}
				break;
			}
		}

		for _,opCodes := range opPossibilityMap {
			delete(opCodes, resolvedOpCode)
		}
	}

	for k,v := range opPossibilityMap {
		fmt.Println(k, v)
	}

	return opMap
}

func executeProgram(program []Instruction, opMap map[int] string) Register {
	register := Register{}

	for _,instruction := range program {
		register = executeOp(opMap[instruction.opcode], register, instruction.parameters)
	}

	return register
}


func main() {
	inputs := au.ReadInputAsStringArray("16")

	samples, program := parseInputs(inputs)

	opMap := getOpMap(samples)

	register := executeProgram(program, opMap)
	fmt.Println(register[0])
}