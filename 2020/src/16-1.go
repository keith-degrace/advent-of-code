package main

import (
	"au"
	"fmt"
	"regexp"
	"strings"
)

type Puzzle16_1 struct {
}

type Rule16_1 struct {
	name string
	min1 int
	max1 int
	min2 int
	max2 int
}

func (r Rule16_1) isValid(number int) bool {
	return (number >= r.min1 && number <= r.max1) ||
		(number >= r.min2 && number <= r.max2)
}

type Log16_1 struct {
	rules         []Rule16_1
	yourTicket    []int
	nearbyTickets [][]int
}

func (l Log16_1) isValid(number int) bool {
	for _, rule := range l.rules {
		if rule.isValid(number) {
			return true
		}
	}

	return false
}

func (p Puzzle16_1) parse(input []string) Log16_1 {
	log := Log16_1{}

	part := 1
	for _, line := range input {
		if len(line) == 0 {
			part++
			continue
		}

		if part == 1 {
			re := regexp.MustCompile("(.+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")

			m := re.FindStringSubmatch(line)
			name := m[1]
			min1 := au.ToNumber(m[2])
			max1 := au.ToNumber(m[3])
			min2 := au.ToNumber(m[4])
			max2 := au.ToNumber(m[5])

			log.rules = append(log.rules, Rule16_1{name, min1, max1, min2, max2})
		}

		if part == 2 {
			if line != "your ticket:" {
				log.yourTicket = au.ToNumbers(strings.Split(line, ","))
			}
		}

		if part == 3 {
			if line != "nearby tickets:" {
				log.nearbyTickets = append(log.nearbyTickets, au.ToNumbers(strings.Split(line, ",")))
			}
		}
	}

	return log
}

func (p Puzzle16_1) run() {
	input := au.ReadInputAsStringArray("16")

	log := p.parse(input)

	errorRate := 0

	for _, ticket := range log.nearbyTickets {
		for _, number := range ticket {
			if !log.isValid(number) {
				errorRate += number
			}
		}
	}

	fmt.Println(errorRate)
}
