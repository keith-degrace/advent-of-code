package au

import (
	"bufio"
	"os"
	"strings"
)

func ReadInputAsStringArray(filename string) ([]string) {
	file, err := os.Open("inputs\\" + filename + ".txt");
	FatalOnError(err);
	defer file.Close();

	var inputs []string;
	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		input := scanner.Text()
		if (len(strings.TrimSpace(input)) > 0) {
			inputs = append(inputs, input)
		}
	}

	FatalOnError(scanner.Err());

	return inputs;
}

