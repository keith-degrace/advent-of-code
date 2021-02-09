package main

import (
	"au"
	"fmt"
)

type Puzzle05_2 struct {
}

func (p Puzzle05_2) getSeatRow(seatSpecifier string) int {
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

func (p Puzzle05_2) getSeatColumn(seatSpecifier string) int {
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

func (p Puzzle05_2) run() {
	input := au.ReadInputAsStringArray("05")

	screen := au.NewScreen(128, 8)

	for _, seatSpecifier := range input {
		row := p.getSeatRow(seatSpecifier)
		column := p.getSeatColumn(seatSpecifier)

		screen.SetPixel(row, column, 'O')
	}

	for row := 1; row < 127; row++ {
		for column := 1; column < 7; column++ {

			if screen.GetPixel(row, column) != 'O' &&
				screen.GetPixel(row-1, column) == 'O' &&
				screen.GetPixel(row+1, column) == 'O' &&
				screen.GetPixel(row, column-1) == 'O' &&
				screen.GetPixel(row, column+1) == 'O' {
				fmt.Println(row*8 + column)
				break
			}

		}
	}
}
