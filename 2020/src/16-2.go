package main

import (
	"au"
	"fmt"
	"regexp"
	"strings"
)

type Puzzle16_2 struct {
}

type Field16_2 struct {
	name string
	min1 int
	max1 int
	min2 int
	max2 int
}

func (r Field16_2) isValid(number int) bool {
	return (number >= r.min1 && number <= r.max1) ||
		(number >= r.min2 && number <= r.max2)
}

type Log16_2 struct {
	fields        []Field16_2
	yourTicket    []int
	nearbyTickets [][]int
}

func (l Log16_2) isValid(number int) bool {
	for _, field := range l.fields {
		if field.isValid(number) {
			return true
		}
	}

	return false
}

func (p Puzzle16_2) parse(input []string) Log16_2 {
	log := Log16_2{}

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

			log.fields = append(log.fields, Field16_2{name, min1, max1, min2, max2})
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

func (p Puzzle16_2) isValidTicket(log Log16_2, ticket []int) bool {
	for _, number := range ticket {
		if !log.isValid(number) {
			return false
		}
	}

	return true
}

func (p Puzzle16_2) getValidTickets(log Log16_2) [][]int {
	validTickets := [][]int{}

	for _, ticket := range log.nearbyTickets {
		if p.isValidTicket(log, ticket) {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

func (p Puzzle16_2) getCandidateFields(fields []Field16_2, validTickets [][]int, position int) []Field16_2 {
	var candidateFields []Field16_2

	for _, field := range fields {

		matchesAllTickets := true

		for _, ticket := range validTickets {
			if !field.isValid(ticket[position]) {
				matchesAllTickets = false
				break
			}
		}

		if matchesAllTickets {
			candidateFields = append(candidateFields, field)
		}
	}

	return candidateFields
}

func (p Puzzle16_2) hasField(fields []Field16_2, name string) bool {
	for _, field := range fields {
		if field.name == name {
			return true
		}
	}

	return false
}

func (p Puzzle16_2) removeField(fields []Field16_2, name string) []Field16_2 {
	newFields := []Field16_2{}

	for _, field := range fields {
		if field.name != name {
			newFields = append(newFields, field)
		}
	}

	return newFields
}

func (p Puzzle16_2) run() {
	input := au.ReadInputAsStringArray("16")

	log := p.parse(input)

	validTickets := p.getValidTickets(log)
	validTickets = append(validTickets, log.yourTicket)

	matches := [][]Field16_2{}
	for i := 0; i < 20; i++ {
		matches = append(matches, p.getCandidateFields(log.fields, validTickets, i))
	}

	noMoreFound := false
	for !noMoreFound {
		noMoreFound = true
		for i, match := range matches {
			if len(match) == 1 {
				for j, other := range matches {
					if i != j && p.hasField(other, match[0].name) {
						matches[j] = p.removeField(other, match[0].name)
						noMoreFound = false
					}
				}
			}
		}
	}

	result := 1
	for i, match := range matches {
		if strings.HasPrefix(match[0].name, "departure") {
			result *= log.yourTicket[i]
		}
	}

	fmt.Println(result)
}
