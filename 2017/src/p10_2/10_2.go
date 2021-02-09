package p10_2

import (
	"fmt"
)

func parse(input string) []int {
	numbers := []int{}
	for i := range input {
		numbers = append(numbers, int(input[i]))
	}

	numbers = append(numbers, []int{17, 31, 73, 47, 23}...)

	return numbers
}

func reverse(array []int, pos int, length int) {
	for i := 0; i < length/2; i++ {
		a := (pos + i) % len(array)
		b := (pos + length - i - 1) % len(array)

		array[a], array[b] = array[b], array[a]
	}
}

func getSparseHash(numbers []int) []int {

	sparseHash := make([]int, 256)
	for i := 0; i < len(sparseHash); i++ {
		sparseHash[i] = i
	}

	current := 0
	skipSize := 0

	for round := 0; round < 64; round++ {
		for _, number := range numbers {
			reverse(sparseHash, current, number)

			current += number%len(sparseHash) + skipSize
			skipSize++
		}
	}

	return sparseHash
}

func getDenseHash(sparseHash []int) string {
	denseHash := ""

	for i := 0; i < 16; i++ {

		value := sparseHash[i*16]
		for j := 1; j < 16; j++ {
			value ^= sparseHash[i*16+j]
		}

		denseHash += fmt.Sprintf("%02x", value)
	}

	return denseHash
}

func Solve() {
	input := "230,1,2,221,97,252,168,169,57,99,0,254,181,255,235,167"

	numbers := parse(input)

	sparseHash := getSparseHash(numbers)
	denseHash := getDenseHash(sparseHash)

	fmt.Println(denseHash)
}
