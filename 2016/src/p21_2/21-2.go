package p21_2

import (
	"au"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Instruction struct {
	op      string
	number1 int
	number2 int
	letter1 byte
	letter2 byte
}

func parseInputs(inputs []string) []Instruction {
	instructions := make([]Instruction, 0)

	re1 := regexp.MustCompile("swap position ([0-9]+) with position ([0-9]+)")
	re2 := regexp.MustCompile("swap letter (.+) with letter (.+)")
	re3 := regexp.MustCompile("reverse positions ([0-9]+) through ([0-9]+)")
	re4 := regexp.MustCompile("rotate (left|right) ([0-9]+) step")
	re5 := regexp.MustCompile("move position ([0-9]+) to position ([0-9]+)")
	re6 := regexp.MustCompile("rotate based on position of letter (.+)")

	for _, input := range inputs {
		matches1 := re1.FindStringSubmatch(input)
		if len(matches1) > 0 {
			instructions = append(instructions, Instruction{
				op:      "swap_position",
				number1: au.ToNumber(matches1[1]),
				number2: au.ToNumber(matches1[2]),
			})
			continue
		}

		matches2 := re2.FindStringSubmatch(input)
		if len(matches2) > 0 {
			instructions = append(instructions, Instruction{
				op:      "swap_letter",
				letter1: matches2[1][0],
				letter2: matches2[2][0],
			})
			continue
		}

		matches3 := re3.FindStringSubmatch(input)
		if len(matches3) > 0 {
			instructions = append(instructions, Instruction{
				op:      "reverse",
				number1: au.ToNumber(matches3[1]),
				number2: au.ToNumber(matches3[2]),
			})
			continue
		}

		matches4 := re4.FindStringSubmatch(input)
		if len(matches4) > 0 {
			negative := matches4[1] == "left"
			steps := au.ToNumber(matches4[2])
			if negative {
				steps = -steps
			}

			instructions = append(instructions, Instruction{
				op:      "rotate_by_step",
				number1: steps,
			})
			continue
		}

		matches5 := re5.FindStringSubmatch(input)
		if len(matches5) > 0 {
			instructions = append(instructions, Instruction{
				op:      "move",
				number1: au.ToNumber(matches5[1]),
				number2: au.ToNumber(matches5[2]),
			})
			continue
		}

		matches6 := re6.FindStringSubmatch(input)
		if len(matches6) > 0 {
			instructions = append(instructions, Instruction{
				op:      "rotate_by_position",
				letter1: matches6[1][0],
			})
			continue
		}

		panic(input)
	}

	return instructions
}

func setStringChar(input string, pos int, newValue byte) string {
	newString := input[:pos]
	newString += string(newValue)
	if pos < len(input)-1 {
		newString += input[pos+1:]
	}

	return newString
}

func insertStringChar(input string, pos int, newValue byte) string {
	newString := input[:pos]
	newString += string(newValue)
	if pos <= len(input)-1 {
		newString += input[pos:]
	}

	return newString
}

func deleteChar(input string, pos int) string {
	newString := input[:pos]
	if pos < len(input)-1 {
		newString += input[pos+1:]
	}

	return newString
}

func swapPosition(instruction Instruction, input string) string {
	letter1 := input[instruction.number1]
	letter2 := input[instruction.number2]

	input = setStringChar(input, instruction.number1, letter2)
	input = setStringChar(input, instruction.number2, letter1)

	return input
}

func swapLetters(instruction Instruction, input string) string {
	newString := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		if input[i] == instruction.letter1 {
			newString[i] = instruction.letter2
		} else if input[i] == instruction.letter2 {
			newString[i] = instruction.letter1
		} else {
			newString[i] = input[i]
		}
	}

	return string(newString)
}

func rotateByPosition(instruction Instruction, input string) string {
	newString := make([]byte, len(input))

	index := strings.Index(input, string(instruction.letter1))

	step := 0
	if index == 0 {
		step = -1
	} else if index <= 7 {
		if index%2 == 1 {
			step = -(index-1)/2 - 1
		} else {
			step = -(index+len(input))/2 - 1
		}
	} else {
		step = -(index-1)/2 - 2
	}

	for i := 0; i < len(input); i++ {
		newString[i] = input[(i-step)%len(input)]
	}

	return string(newString)
}

func rotateByStep(instruction Instruction, input string) string {
	newString := make([]byte, len(input))

	step := -instruction.number1
	if step > 0 {
		for i := 0; i < len(input); i++ {
			newString[(i+step)%len(input)] = input[i]
		}
	} else {
		for i := 0; i < len(input); i++ {
			newString[i] = input[(i-step)%len(input)]
		}
	}

	return string(newString)
}

func move(instruction Instruction, input string) string {
	letter := input[instruction.number2]

	if instruction.number2 < instruction.number1 {
		input = deleteChar(input, instruction.number2)
		input = insertStringChar(input, instruction.number1, letter)
	} else {
		input = insertStringChar(input, instruction.number1, letter)
		input = deleteChar(input, instruction.number2+1)
	}

	return input
}

func reverse(instruction Instruction, input string) string {
	subString := input[instruction.number1 : instruction.number2+1]
	subString = au.ReverseString(subString)

	newString := input[:instruction.number1]
	newString += subString
	if instruction.number2 < len(input)-1 {
		newString += input[instruction.number2+1:]
	}

	return newString
}

func apply(instruction Instruction, input string) string {
	// fmt.Println(instruction)
	switch instruction.op {
	case "swap_position":
		return swapPosition(instruction, input)
	case "swap_letter":
		return swapLetters(instruction, input)
	case "reverse":
		return reverse(instruction, input)
	case "rotate_by_step":
		return rotateByStep(instruction, input)
	case "move":
		return move(instruction, input)
	case "rotate_by_position":
		return rotateByPosition(instruction, input)
	}

	panic("Error!")
}
func Solve() {
	fmt.Println("Starting\n")
	startTime := time.Now()

	// inputs, value := au.ReadInputAsStringArray("21"), "bfheacgd"
	inputs, value := au.ReadInputAsStringArray("21"), "fbgdceah"

	instructions := parseInputs(inputs)

	for i := len(instructions) - 1; i >= 0; i-- {
		value = apply(instructions[i], value)
	}

	fmt.Println(value)

	fmt.Println("\nCompleted in", time.Now().Sub(startTime))
}
