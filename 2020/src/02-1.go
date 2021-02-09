package main

import (
	"au"
	"fmt"
	"regexp"
	"strconv"
)

type Puzzle02_1 struct {
}

func (p Puzzle02_1) run() {
	validCount := 0

	for _, input := range au.ReadInputAsStringArray("02") {

		r := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): (.*)")
		matches := r.FindStringSubmatch(input)

		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		letter := matches[3]
		password := matches[4]

		letterCount := 0

		for _, rune := range password {
			if string(rune) == letter {
				letterCount++
			}
		}

		if letterCount >= min && letterCount <= max {
			validCount++
		}
	}

	fmt.Println(validCount)
}
