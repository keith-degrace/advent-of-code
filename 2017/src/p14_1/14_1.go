package p14_1

import (
	"fmt"
)

func reverse(array []int, pos int, length int) {
	for i := 0; i < length/2; i++ {
		a := (pos + i) % len(array)
		b := (pos + length - i - 1) % len(array)

		array[a], array[b] = array[b], array[a]
	}
}

func getSparseHash(numbers []int) []int {

	numbers = append(numbers, []int{17, 31, 73, 47, 23}...)

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

		denseHash += fmt.Sprintf("%08b", value)
	}

	return denseHash
}

func getKnotHash(value string) string {
	numbers := []int{}
	for _, char := range value {
		numbers = append(numbers, int(char))
	}

	sparseHash := getSparseHash(numbers)

	return getDenseHash(sparseHash)
}

func Solve() {
	input := "vbqugkhl"
	// input := "flqrgnkx"

	totalUsed := 0

	for i := 0; i < 128; i++ {
		rowValue := fmt.Sprintf("%v-%v", input, i)
		hash := getKnotHash(rowValue)

		for _, char := range hash {
			if char == '1' {
				totalUsed++
			}
		}
	}

	fmt.Println(totalUsed)
}
