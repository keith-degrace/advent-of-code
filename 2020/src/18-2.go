package main

import (
	"au"
	"fmt"
	"regexp"
)

type Puzzle18_2 struct {
}

func (p *Puzzle18_2) evaluate(expression string) int {
	// Evaluate all parantheses
	parantheseRegExp := regexp.MustCompile("\\(([^()]*)\\)")
	for {
		match := parantheseRegExp.FindStringSubmatchIndex(expression)
		if len(match) == 0 {
			break
		}

		subExpression := p.evaluate(expression[match[2]:match[3]])

		expression = expression[:match[0]] + fmt.Sprintf("%v", subExpression) + expression[match[1]:]
	}

	// Evaluate all additions
	addRegExp := regexp.MustCompile("([0-9]+) \\+ ([0-9]+)")
	for {
		match := addRegExp.FindStringSubmatchIndex(expression)
		if len(match) == 0 {
			break
		}

		left := au.ToNumber(expression[match[2]:match[3]])
		right := au.ToNumber(expression[match[4]:match[5]])

		expression = expression[:match[0]] + fmt.Sprintf("%v", left+right) + expression[match[1]:]
	}

	// Evaluate all multiplications
	mulRegExp := regexp.MustCompile("([0-9]+) \\* ([0-9]+)")
	for {
		match := mulRegExp.FindStringSubmatchIndex(expression)
		if len(match) == 0 {
			break
		}

		left := au.ToNumber(expression[match[2]:match[3]])
		right := au.ToNumber(expression[match[4]:match[5]])

		expression = expression[:match[0]] + fmt.Sprintf("%v", left*right) + expression[match[1]:]
	}

	// Should only be a single number left
	return au.ToNumber(expression)
}

func (p Puzzle18_2) run() {
	input := au.ReadInputAsStringArray("18")

	sum := 0
	for _, expression := range input {
		sum += p.evaluate(expression)
	}

	fmt.Println(sum)
}
