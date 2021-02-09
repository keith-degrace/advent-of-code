package au

func ToNumbers(array []string) []int {
	numbers := []int{}

	for _, string := range array {
		numbers = append(numbers, ToNumber(string))
	}

	return numbers
}
