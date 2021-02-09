package p06_1

import (
	"au"
	"fmt"
)

func testInputs() []string {
	return []string {
		"eedadn",
		"drvtee",
		"eandsr",
		"raavrd",
		"atevrs",
		"tsrnev",
		"sdttsa",
		"rasrtv",
		"nssdts",
		"ntnada",
		"svetve",
		"tesnvt",
		"vntsnd",
		"vrdear",
		"dvrsen",
		"enarar",
	};
}

func getMostCommonCharacter(inputs []string, position int) string {
	countMap := map[byte] int {}

	for _, input := range inputs {
		character := input[position]
		countMap[character]++
	}

	var mostCommonCharacter byte
	var mostCommonCharacterCount int

	for character, count := range countMap {
		if count > mostCommonCharacterCount {
			mostCommonCharacter = character
			mostCommonCharacterCount = count
		}
	}

	return string(mostCommonCharacter)
}

func Solve() {
	inputs := au.ReadInputAsStringArray("06")
	// inputs := testInputs()

	version := ""
	for position := 0; position < len(inputs[0]); position++ {
		version += getMostCommonCharacter(inputs, position)
	}

	fmt.Println(version)
}
