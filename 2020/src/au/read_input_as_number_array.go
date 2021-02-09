package au

func ReadInputAsNumberArray(filename string) ([]int) {
	var numberInputs []int;

	for _, input := range ReadInputAsStringArray(filename) {
		numberInputs = append(numberInputs, ToNumber(input))
	}

	return numberInputs;
}

