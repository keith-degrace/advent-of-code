package p10_1

import (
	"au"
	"fmt"
	"strings"
)

func reverse(array []int, pos int, length int) {
	for i := 0; i < length/2; i++ {
		a := (pos + i) % len(array)
		b := (pos + length - i - 1) % len(array)

		array[a], array[b] = array[b], array[a]
	}
}

func Solve() {
	input, arraySize := "230,1,2,221,97,252,168,169,57,99,0,254,181,255,235,167", 256
	//input, arraySize := "3,4,1,5", 5

	numbers := au.ToNumbers(strings.Split(input, ","))

	array := make([]int, arraySize)
	for i := 0; i < len(array); i++ {
		array[i] = i
	}

	current := 0
	skipSize := 0

	for _, number := range numbers {
		reverse(array, current, number)

		current += number%len(array) + skipSize
		skipSize++
	}

	fmt.Println(array[0] * array[1])
}
