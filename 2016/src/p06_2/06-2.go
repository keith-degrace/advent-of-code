package p06_2

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

func getLeastCommonCharacter(inputs []string, position int) string {
	countMap := map[byte] int {}

	for _, input := range inputs {
		character := input[position]
		countMap[character]++
	}

	var leastCommonCharacter byte
	leastCommonCharacterCount := 10000000000

	for character, count := range countMap {
		if count < leastCommonCharacterCount {
			leastCommonCharacter = character
			leastCommonCharacterCount = count
		}
	}

	return string(leastCommonCharacter)
}

func Solve() {
	inputs := au.ReadInputAsStringArray("06")
	// inputs := testInputs()

	version := ""
	for position := 0; position < len(inputs[0]); position++ {
		version += getLeastCommonCharacter(inputs, position)
	}

	fmt.Println(version)
}
