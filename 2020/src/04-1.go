package main

import (
	"au"
	"fmt"
	"regexp"
)

type Puzzle04_1 struct {
}

func (p Puzzle04_1) parse(input []string) []map[string]string {
	var result []map[string]string

	var currentPassword = make(map[string]string)

	for _, line := range input {
		if len(line) == 0 {
			result = append(result, currentPassword)
			currentPassword = make(map[string]string)
		} else {
			r := regexp.MustCompile("([^ ]+):([^ ]+)")

			for _, match := range r.FindAllStringSubmatch(line, -1) {
				fieldName := match[1]
				fieldValue := match[2]

				currentPassword[fieldName] = fieldValue
			}

		}
	}

	result = append(result, currentPassword)

	return result
}

func (p Puzzle04_1) hasField(passport map[string]string, field string) bool {
	_, ok := passport[field]
	return ok
}

func (p Puzzle04_1) isValid(passport map[string]string) bool {
	var requiredFields = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, requiredField := range requiredFields {
		if !p.hasField(passport, requiredField) {
			return false
		}
	}

	return true
}

func (p Puzzle04_1) run() {
	input := au.ReadInputAsStringArray("04")

	passports := p.parse(input)

	validCount := 0
	for _, passport := range passports {
		if p.isValid(passport) {
			validCount++
		}
	}

	fmt.Println(validCount)
}
