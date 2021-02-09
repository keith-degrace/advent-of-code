package main

import (
	"au"
	"fmt"
	"regexp"
	"strconv"
)

type Puzzle04_2 struct {
}

func (p Puzzle04_2) parse(input []string) []map[string]string {
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

func (p Puzzle04_2) hasField(passport map[string]string, field string) bool {
	_, ok := passport[field]
	return ok
}

func (p Puzzle04_2) isByrValid(passport map[string]string) bool {
	value, ok := passport["byr"]
	if !ok {
		return false
	}

	year, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return year >= 1920 && year <= 2002
}

func (p Puzzle04_2) isIyrValid(passport map[string]string) bool {
	value, ok := passport["iyr"]
	if !ok {
		return false
	}

	year, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return year >= 2010 && year <= 2020
}

func (p Puzzle04_2) isEyrValid(passport map[string]string) bool {
	value, ok := passport["eyr"]
	if !ok {
		return false
	}

	year, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return year >= 2020 && year <= 2030
}

func (p Puzzle04_2) isHgtValid(passport map[string]string) bool {
	value, ok := passport["hgt"]
	if !ok {
		return false
	}

	height, _ := strconv.Atoi(value[0 : len(value)-2])
	units := value[len(value)-2:]

	if units == "cm" {
		return height >= 150 && height <= 193
	} else if units == "in" {
		return height >= 59 && height <= 76
	} else {
		return false
	}
}

func (p Puzzle04_2) isHclValid(passport map[string]string) bool {
	value, ok := passport["hcl"]
	if !ok {
		return false
	}

	r := regexp.MustCompile("#[a-f0-9]{6}")
	return r.MatchString(value)
}

func (p Puzzle04_2) isEclValid(passport map[string]string) bool {
	value, ok := passport["ecl"]
	if !ok {
		return false
	}

	return value == "amb" ||
		value == "blu" ||
		value == "brn" ||
		value == "gry" ||
		value == "grn" ||
		value == "hzl" ||
		value == "oth"
}

func (p Puzzle04_2) isPidValid(passport map[string]string) bool {
	value, ok := passport["pid"]
	if !ok {
		return false
	}

	r := regexp.MustCompile("^[0-9]{9}$")
	return r.MatchString(value)
}

func (p Puzzle04_2) isValid(passport map[string]string) bool {
	return p.isByrValid(passport) &&
		p.isIyrValid(passport) &&
		p.isEyrValid(passport) &&
		p.isHgtValid(passport) &&
		p.isHclValid(passport) &&
		p.isEclValid(passport) &&
		p.isPidValid(passport)
}

func (p Puzzle04_2) run() {
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
