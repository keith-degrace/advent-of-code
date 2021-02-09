package p09_1

import (
	"au"
	"fmt"
)

func testInputs() []string {
	return []string {
		"ADVENT",
		"A(1x5)BC",
		"(3x3)XYZ",
		"A(2x2)BCD(2x2)EFG",
		"(6x1)(1x3)A",
		"X(8x2)(3x3)ABCY",
	};
}

func readUpTo(value string, index int, stopChar byte) (int, string) {
	buffer := ""

	for index < len(value) && value[index] != stopChar {
		buffer += string(value[index])
		index++;
	}

	return index, buffer
}

func getDecompressedLength(compressedValue string) int {
	length := 0

	index := 0
	buffer := ""
	for  {
		index, buffer = readUpTo(compressedValue, index, '(')
		length += len(buffer);

		if index == len(compressedValue) {
			break;
		}

		index, buffer = readUpTo(compressedValue, index + 1, 'x')
		repeatLength := au.ToNumber(buffer)

		index, buffer = readUpTo(compressedValue, index + 1, ')')
		repeatCount := au.ToNumber(buffer)

		length += repeatLength * repeatCount;
		index += repeatLength + 1
	}

	return length;
}

func Solve() {
	inputs := au.ReadInputAsStringArray("09")
	// inputs := testInputs()

	for _, input := range inputs {
		fmt.Println(getDecompressedLength(input))
	}
}
