package main

import (
	"au"
	"fmt"
)

type Puzzle05_1 struct {
}

func (p Puzzle05_1) getSeatRow(seatSpecifier string) int {
	min := 0
	max := 127

	for i := 0; i < 7; i++ {

		if seatSpecifier[i] == 'F' {
			max = max - (max-min+1)/2
		} else {
			min = min + (max-min+1)/2
		}
	}

	return min
}

func (p Puzzle05_1) getSeatColumn(seatSpecifier string) int {
	min := 0
	max := 7

	for i := 7; i < 10; i++ {

		if seatSpecifier[i] == 'L' {
			max = max - (max-min+1)/2
		} else {
			min = min + (max-min+1)/2
		}
	}

	return min
}

func (p Puzzle05_1) run() {
	input := au.ReadInputAsStringArray("05")

	max := 0

	for _, seatSpecifier := range input {
		row := p.getSeatRow(seatSpecifier)
		column := p.getSeatColumn(seatSpecifier)

		value := row*8 + column

		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}
