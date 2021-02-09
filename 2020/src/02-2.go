package main

import (
	"au"
	"fmt"
	"regexp"
	"strconv"
)

type Puzzle02_2 struct {
}

func (p Puzzle02_2) run() {
	validCount := 0

	for _, input := range au.ReadInputAsStringArray("02") {

		r := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): (.*)")
		matches := r.FindStringSubmatch(input)

		pos1, _ := strconv.Atoi(matches[1])
		pos2, _ := strconv.Atoi(matches[2])
		letter := matches[3]
		password := matches[4]

		pos1Valid := string(password[pos1-1]) == letter
		pos2Valid := string(password[pos2-1]) == letter

		if pos1Valid != pos2Valid {
			validCount++
		}
	}

	fmt.Println(validCount)
}
