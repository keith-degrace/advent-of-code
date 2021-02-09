package p23_1

import (
	"au"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getValue(registers map[string]int, value string) int {
	numberValue, err := strconv.Atoi(value)
	if err == nil {
		return numberValue
	}

	registerValue, ok := registers[value]
	if !ok {
		return 0
	}

	return registerValue
}

func Solve() {
	input := au.ReadInputAsStringArray("23")

	registers := make(map[string]int)

	mulCallCount := 0

	current := 0
	for current < len(input) {
		parts := strings.Split(input[current], " ")

		opcode := parts[0]

		if opcode == "set" {
			x := parts[1]
			y := getValue(registers, parts[2])
			fmt.Printf("set %v %v\n", x, y)
			registers[x] = y
		} else if opcode == "sub" {
			x := parts[1]
			y := getValue(registers, parts[2])
			fmt.Printf("sub %v %v\n", x, y)
			registers[x] = getValue(registers, x) - y
		} else if opcode == "mul" {
			x := parts[1]
			y := getValue(registers, parts[2])
			fmt.Printf("mul %v %v\n", x, y)
			mulCallCount++
			registers[x] = getValue(registers, x) * y
		} else if opcode == "jnz" {
			x := getValue(registers, parts[1])
			y := getValue(registers, parts[2])
			if x != 0 {
				fmt.Printf("jnz %v %v (yes)\n", x, y)
				current = current + y - 1
			} else {
				fmt.Printf("jnz %v %v (no)\n", x, y)
			}
		}

		fmt.Println(registers)

		current++

		time.Sleep(2000)
	}

	fmt.Println(mulCallCount)
}

// Not 2512599
