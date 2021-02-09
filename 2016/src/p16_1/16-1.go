package p16_1

import (
	"au"
	"fmt"
	"strings"
	"time"
)

func mutate(a string) string {
	// Call the data you have at this point "a".

	// Make a copy of "a"; call this copy "b".
	b := a

	// Reverse the order of the characters in "b".
	b = au.ReverseString(b)

	// In "b", replace all instances of 0 with 1 and all 1s with 0.
	b = strings.Replace(b, "1", "a", -1)
	b = strings.Replace(b, "0", "1", -1)
	b = strings.Replace(b, "a", "0", -1)

	// The resulting data is "a", then a single 0, then "b".

	return a + "0" + b
}

func getChecksumSubString(input string) string {
	substring := ""

	i := 0
	for i < len(input) {
		if input[i] == input[i+1] {
			substring += "1"
		} else {
			substring += "0"
		}

		i += 2
	}

	return substring
}

func getChecksum(input string) string {
	subString := getChecksumSubString(input)
	if len(subString)%2 == 0 {
		return getChecksum(subString)
	} else {
		return subString
	}
}

func Solve() {
	fmt.Printf("Starting\n\n")
	startTime := time.Now()

	input, diskSize := "11011110011011101", 272
	//input, diskSize  = "10000", 20

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
