package p16_2

import (
	"au"
	"fmt"
	"time"
)

func mutate(a string) string {
	result := make([]byte, len(a) * 2 + 1)

	result[len(result) / 2] = '0'

	for i := 0; i < len(a); i++ {
		result[i] = a[i]

		if a[i] == '0' {
			result[len(result) - 1 - i] = '1'
		} else {
			result[len(result) - 1 - i] = '0'
		}
	}

	return string(result)
}

func getChecksumSubString(input string) string {
	substring := make([]byte, len(input) / 2)

	i := 0
	for i < len(input) {
		if input[i] == input[i + 1] {
			substring[i / 2] = '1'
		} else {
			substring[i / 2] = '0'
		}

		i += 2
	}

	return string(substring)
}

func getChecksum(input string) string {
	subString := getChecksumSubString(input)
	if len(subString) % 2 == 0 {
		return getChecksum(subString)
	} else {
		return subString
	}
}

func Solve() {
	fmt.Println("Starting\n")
	startTime := time.Now()

	input, diskSize := "11011110011011101", 35651584
	// input, diskSize  = "10000", 20

	au.AssertStringsEqual(mutate("1"), "100")
	au.AssertStringsEqual(mutate("0"), "001")
	au.AssertStringsEqual(mutate("11111"), "11111000000")
	au.AssertStringsEqual(mutate("111100001010"), "1111000010100101011110000")

	au.AssertStringsEqual(getChecksum("110010110100"), "100")

	for len(input) < diskSize {
		input = mutate(input)
	}

	fmt.Println(getChecksum(input[:diskSize]))

	fmt.Println("\nCompleted in", time.Now().Sub(startTime))
}