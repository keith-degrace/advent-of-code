package main

import (
	"au"
	"fmt"
)

type Puzzle06_2 struct {
}

func (p Puzzle06_2) getAllGroupAnswers(input []string) []map[string]int {
	var result []map[string]int

	var currentGroup = make(map[string]int)

	for _, line := range input {

		if len(line) == 0 {
			result = append(result, currentGroup)
			currentGroup = make(map[string]int)
		} else {
			currentGroup[" "]++
			for _, char := range line {
				currentGroup[string(char)]++
			}
		}
	}

	result = append(result, currentGroup)

	return result
}

func (p Puzzle06_2) run() {
	input := au.ReadInputAsStringArray("06")

	sum := 0

	for _, groupAnswers := range p.getAllGroupAnswers(input) {
		groupSize := groupAnswers[" "]
		for answer, count := range groupAnswers {
			if answer != " " && count == groupSize {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
