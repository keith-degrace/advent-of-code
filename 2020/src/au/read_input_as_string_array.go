package au

import (
	"bufio"
	"os"
)

func ReadInputAsStringArray(filename string) []string {
	file, err := os.Open("..\\inputs\\" + filename + ".txt")
	FatalOnError(err)
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		inputs = append(inputs, input)
	}

	FatalOnError(scanner.Err())

	return inputs
}
