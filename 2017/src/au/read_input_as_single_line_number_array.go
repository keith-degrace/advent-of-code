package au

import (
	"strings"
)

func ReadInputAsSingleLineNumberArray(filename string) ([]int) {
	var numberInputs []int;

	for _, input := range strings.Split(ReadInputAsString(filename), " ") {
		numberInputs = append(numberInputs, ToNumber(input))
	}

	return numberInputs;
}

