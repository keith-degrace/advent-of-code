package main

import (
	"au"
	"fmt"
)

type Puzzle06_1 struct {
}

func (p Puzzle06_1) getAllGroupAnswers(input []string) []map[string]bool {
	var result []map[string]bool

	var currentGroup = make(map[string]bool)

	for _, line := range input {

		if len(line) == 0 {
			result = append(result, currentGroup)
			currentGroup = make(map[string]bool)
		} else {
			for _, char := range line {
				currentGroup[string(char)] = true
			}
		}
	}

	result = append(result, currentGroup)

	return result
}

func (p Puzzle06_1) run() {
	input := au.ReadInputAsStringArray("06")

	sum := 0

	for _, groupAnswers := range p.getAllGroupAnswers(input) {
		sum += len(groupAnswers)
	}

	fmt.Println(sum)
}
