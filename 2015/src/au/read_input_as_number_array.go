package au

func ReadInputAsNumberArray(filename string) ([]int) {
	numberInputs := make([]int, 0);

	for _, input := range ReadInputAsStringArray(filename) {
		numberInputs = append(numberInputs, ToNumber(input))
	}

	return numberInputs;
}

