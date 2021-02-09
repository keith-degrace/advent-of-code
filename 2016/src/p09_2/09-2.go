package p09_2

import (
	"au"
	"fmt"
)

func testInputs() []string {
	return []string{
		"(3x3)XYZ",
		"X(8x2)(3x3)ABCY",
		"(27x12)(20x12)(13x14)(7x10)(1x12)A",
		"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
	}
}

func read(value string, index int, length int) (int, string) {
	return index + length, value[index : index+length]
}

func readUpTo(value string, index int, stopChar byte) (int, string) {
	buffer := ""

	for index < len(value) && value[index] != stopChar {
		buffer += string(value[index])
		index++
	}

	return index, buffer
}

func getDecompressedLength(compressedValue string) int {
	length := 0

	index := 0
	buffer := ""
	for {
		index, buffer = readUpTo(compressedValue, index, '(')
		length += len(buffer)

		if index == len(compressedValue) {
			break
		}

		index, buffer = readUpTo(compressedValue, index+1, 'x')
		repeatLength := au.ToNumber(buffer)

		index, buffer = readUpTo(compressedValue, index+1, ')')
		repeatCount := au.ToNumber(buffer)

		index, buffer = read(compressedValue, index+1, repeatLength)
		length += getDecompressedLength(buffer) * repeatCount
	}

	return length
}

func Solve() {
	inputs := au.ReadInputAsStringArray("09")
	// ${fileDirname}inputs := testInputs()

	for _, input := range inputs {
		fmt.Println(getDecompressedLength(input))
	}
}
