package main

import (
	"au"
	"fmt"
)

type Puzzle18_1 struct {
}

func (p *Puzzle18_1) evaluate(expression string) int {
	result := 0
	currentOperator := ""

	openParanthesesCount := 0
	paranthesesContent := ""

	for _, char := range expression {

		if openParanthesesCount == 0 {
			if char == '(' {
				openParanthesesCount = 1
				paranthesesContent = ""
			} else if char == '+' || char == '*' {
				currentOperator = string(char)
			} else if char != ' ' {
				number := au.ToNumber(string(char))

				if currentOperator == "+" {
					result += number
				} else if currentOperator == "*" {
					result *= number
				} else {
					result = number
				}
			}

		} else {
			switch char {
			case '(':
				openParanthesesCount++
				paranthesesContent += string(char)
				break
			case ')':
				openParanthesesCount--

				if openParanthesesCount == 0 {
					paranthesesValue := p.evaluate(paranthesesContent)

					if currentOperator == "+" {
						result += paranthesesValue
					} else if currentOperator == "*" {
						result *= paranthesesValue
					} else {
						result = paranthesesValue
					}
				} else {
					paranthesesContent += string(char)
				}
				break
			default:
				paranthesesContent += string(char)
			}
		}
	}

	return result
}

func (p Puzzle18_1) run() {
	input := au.ReadInputAsStringArray("18")

	sum := 0
	for _, expression := range input {
		sum += p.evaluate(expression)
	}

	fmt.Println(sum)
}
