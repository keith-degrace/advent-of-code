package p06_2

import (
	"au"
	"fmt"
	"strings"
)

func getLargestBankIndex(banks []int) int {
	largestIndex := 0

	for i := 1; i < len(banks); i++ {
		if banks[i] > banks[largestIndex] {
			largestIndex = i
		}
	}

	return largestIndex
}

func getHash(banks []int) string {
	hash := ""

	for _, bank := range banks {
		hash += fmt.Sprintf("%v,", bank)
	}

	return hash
}

func Solve() {
	input := au.ReadInputAsString("06")
	banks := au.ToNumbers(strings.Split(input, "\t"))

	iteration := 0
	seen := make(map[string]int)

	for {
		iteration++

		index := getLargestBankIndex(banks)

		amountToDistribute := banks[index]
		banks[index] = 0

		for ; amountToDistribute > 0; amountToDistribute-- {
			index = (index + 1) % len(banks)
			banks[index]++
		}

		hash := getHash(banks)
		when, ok := seen[hash]
		if ok {
			fmt.Println(iteration - when)
			break
		}

		seen[hash] = iteration
	}
}
