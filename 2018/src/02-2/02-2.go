package main

import (
	"fmt"
	"au"
)

func isMatch(id1 string, id2 string) (bool) {
	oneDifferenceFound := false

	for index, _  := range id1 {
		if id1[index] != id2[index] {
			if oneDifferenceFound {
				return false;
			}

			oneDifferenceFound = true;
		}
	}

	return oneDifferenceFound;
}

func diff(id1 string, id2 string) (string) {
	diff := "";

	for index, _  := range id1 {
		if id1[index] == id2[index] {
			diff += string(id1[index]);
		}
	}

	return diff;
}

func main() {
	inputs := au.ReadInputAsStringArray("02")
	// inputs := []string { "abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz" }

	for i := 0; i < len(inputs); i++ {
		for j := i+1; j < len(inputs); j++ {
			if (isMatch(inputs[i], inputs[j])) {
				fmt.Printf("%v\n", diff(inputs[i], inputs[j]))
			}
		}
	}
}
