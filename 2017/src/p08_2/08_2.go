package p08_2

import (
	"au"
	"fmt"
	"regexp"
)

type Registers struct {
	registers map[string]int
}

func (r Registers) Set(register string, value int) {
	r.registers[register] = value
}

func (r Registers) Get(register string) int {
	if r.registers == nil {
		return 0
	}

	value, ok := r.registers[register]
	if !ok {
		value = 0
	}

	return value
}

type Instruction struct {
	register          string
	operator          string
	value             int
	conditionRegister string
	conditionOperator string
	conditionValue    int
}

func load(input []string) []Instruction {
	instructions := []Instruction{}

	re := regexp.MustCompile("([^ ]+) ([^ ]+) ([^ ]+) if ([^ ]+) ([^ ]+) ([^ ]+)")
	for _, line := range input {
		m := re.FindStringSubmatch(line)

		instruction := Instruction{}

		instruction.register = m[1]
		instruction.operator = m[2]
		instruction.value = au.ToNumber(m[3])
		instruction.conditionRegister = m[4]
		instruction.conditionOperator = m[5]
		instruction.conditionValue = au.ToNumber(m[6])

		instructions = append(instructions, instruction)
	}

	return instructions
}

func shouldExecute(registers *Registers, instruction Instruction) bool {
	conditionRegisterValue := registers.Get(instruction.conditionRegister)

	if instruction.conditionOperator == "==" {
		return conditionRegisterValue == instruction.conditionValue
	} else if instruction.conditionOperator == "!=" {
		return conditionRegisterValue != instruction.conditionValue
	} else if instruction.conditionOperator == ">" {
		return conditionRegisterValue > instruction.conditionValue
	} else if instruction.conditionOperator == ">=" {
		return conditionRegisterValue >= instruction.conditionValue
	} else if instruction.conditionOperator == "<" {
		return conditionRegisterValue < instruction.conditionValue
	} else if instruction.conditionOperator == "<=" {
		return conditionRegisterValue <= instruction.conditionValue
	} else {
		au.Assert(false)
	}

	return false
}

func execute(registers *Registers, instruction Instruction) {
	registerValue := registers.Get(instruction.register)

	if instruction.operator == "inc" {
		registerValue += instruction.value
	} else if instruction.operator == "dec" {
		registerValue -= instruction.value
	} else {
		au.Assert(false)
	}

	registers.Set(instruction.register, registerValue)
}

func Solve() {
	input := au.ReadInputAsStringArray("08")

	instructions := load(input)

	registers := new(Registers)
	registers.registers = make(map[string]int)

	maxSeen := 0

	for _, instruction := range instructions {

		if shouldExecute(registers, instruction) {
			execute(registers, instruction)
		}

		// Get the highest map value
		for _, v := range registers.registers {
			maxSeen = au.MaxInt(maxSeen, v)
		}
	}

	fmt.Println(maxSeen)
}
