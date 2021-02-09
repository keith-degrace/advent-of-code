package main

import (
	"au"
	"fmt"
	"strings"
)

func isUpperCase(value byte) bool {
	return value < 97;
}

func react(a byte, b byte) bool {
	if strings.ToLower(string(a)) != strings.ToLower(string(b)) {
		return false
	}

	if (isUpperCase(a) != isUpperCase(b)) {
		return true;
	}

	return false
}

func main() {
	input := au.ReadInputAsString("05")
	// input := "dabAcCaCBAcCcaDA"

	index := 0
	for index < len(input) - 1 {
		if (react(input[index], input[index+1])) {
			input = input[:index] + input[index+2:]
			index = index - 1;
			if index < 0 {
				index = 0
			}
		} else {
			index++;
		}
	}

	//11312

	fmt.Println(len(input))
}
