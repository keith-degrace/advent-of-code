package main

import (
	"au"
	"fmt"
	"math"
	"strconv"
)

type Puzzle12_1 struct {
}

func (p Puzzle12_1) move(x int, y int, f byte, value int) (int, int) {
	if f == 'N' {
		return x, y + value
	}

	if f == 'S' {
		return x, y - value
	}

	if f == 'E' {
		return x + value, y
	}

	if f == 'W' {
		return x - value, y
	}

	return x, y
}

func (p Puzzle12_1) rotate(f byte, angle int) byte {

	times := int(math.Abs(float64(angle / 90)))

	for i := 0; i < times; i++ {
		if angle < 0 {
			if f == 'E' {
				f = 'N'
			} else if f == 'N' {
				f = 'W'
			} else if f == 'W' {
				f = 'S'
			} else if f == 'S' {
				f = 'E'
			}
		} else {
			if f == 'E' {
				f = 'S'
			} else if f == 'S' {
				f = 'W'
			} else if f == 'W' {
				f = 'N'
			} else if f == 'N' {
				f = 'E'
			}
		}
	}

	return f
}

func (p Puzzle12_1) run() {
	input := au.ReadInputAsStringArray("12")

	var f byte = 'E'
	x := 0
	y := 0

	for _, line := range input {

		action := line[0]
		value, _ := strconv.Atoi(line[1:])

		if action == 'L' {
			f = p.rotate(f, -value)
		} else if action == 'R' {
			f = p.rotate(f, value)
		} else if action == 'F' {
			x, y = p.move(x, y, f, value)
		} else {
			x, y = p.move(x, y, action, value)
		}
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}
