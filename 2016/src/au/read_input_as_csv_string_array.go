package au

import (
	"strings"
)

func ReadInputAsCSVStringArray(filename string) ([]string) {
	input := ReadInputAsStringArray(filename);

	array := []string{}
	
	for _, s := range strings.Split(input[0], ",") {
		array = append(array, strings.TrimSpace(s))
	}

	return array;
}
