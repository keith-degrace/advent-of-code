package main

import (
	"au"
	"fmt"
	"strings"
)

func isUpperCase(value byte) bool {
	return value < 97;
}

func removeUnit(input string, unit int) string {
	result := strings.Replace(input, string(unit), "", -1)
	result = strings.Replace(result, strings.ToUpper(string(unit)), "", -1)
	return result
}

func reactSegment(a byte, b byte) bool {
	if strings.ToLower(string(a)) != strings.ToLower(string(b)) {
		return false
	}

	if (isUpperCase(a) != isUpperCase(b)) {
		return true;
	}

	return false
}

func react(input string) string {
	index := 0
	for index < len(input) - 1 {
		if (reactSegment(input[index], input[index+1])) {
			input = input[:index] + input[index+2:]
			index = index - 1;
			if index < 0 {
				index = 0
			}
		} else {
			index++;
		}
	}

	return input
}

func main() {
	input := au.ReadInputAsString("05")
	// input := "dabAcCaCBAcCcaDA"

	shortest := 100000
	for i := 97; i < 123; i++ {
		polymer := react(removeUnit(input, i))

		if (len(polymer) < shortest) {
			shortest = len(polymer)
		}
	}

	fmt.Println(shortest)
}
