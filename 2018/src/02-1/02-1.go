package main

import (
	"fmt"
	"au"
)

func hasTwosAndThrees(id string) (bool, bool) {
	counts := make(map[rune]int)

	for _, char := range id {
		counts[char] += 1;
	}

	hasTwos := false
	hasThrees := false
	for _, v := range counts {
		if (v == 2) {
			hasTwos = true;
		}

		if (v == 3) {
			hasThrees = true;
		}

		if (hasTwos && hasThrees) {
			break;
		}
	}

	return hasTwos, hasThrees
}

func main() {
	inputs := au.ReadInputAsStringArray("02")
	// inputs := []string { "abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab" }

	twos := 0;
	threes := 0;
	for _, input := range inputs {
		hasTwos, hasThrees := hasTwosAndThrees(input);
		if (hasTwos) {
			twos += 1;
		}

		if (hasThrees) {
			threes += 1;
		}
	}

	checksum := twos * threes;

	fmt.Println(checksum)
}
