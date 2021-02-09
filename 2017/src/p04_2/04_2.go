package p04_2

import (
	"au"
	"fmt"
	"sort"
	"strings"
)

func isValid(words []string) bool {
	dict := make(map[string]bool)

	for _, word := range words {

		s := strings.Split(word, "")
		sort.Strings(s)
		word := strings.Join(s, "")

		_, ok := dict[word]
		if ok {
			return false
		}

		dict[word] = true
	}

	return true
}

func Solve() {
	input := au.ReadInputAsStringArray("04")

	validCount := 0

	for _, passphrase := range input {
		words := strings.Split(passphrase, " ")

		if isValid(words) {
			validCount++
		}
	}

	fmt.Println(validCount)
}
