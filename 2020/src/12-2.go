package main

import (
	"au"
	"fmt"
	"math"
	"strconv"
)

type Puzzle12_2 struct {
}

func (p Puzzle12_2) move(x int, y int, f byte, value int) (int, int) {
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

func (p Puzzle12_2) rotate(x int, y int, angle int) (int, int) {

	times := int(math.Abs(float64(angle / 90)))

	for i := 0; i < times; i++ {
		if angle < 0 {
			x, y = -y, x
		} else {
			x, y = y, -x
		}
	}

	return x, y
}

func (p Puzzle12_2) run() {
	input := au.ReadInputAsStringArray("12")

	waypointX := 10
	waypointY := 1

	boatX := 0
	boatY := 0

	for _, line := range input {

		action := line[0]
		value, _ := strconv.Atoi(line[1:])

		if action == 'L' {
			waypointX, waypointY = p.rotate(waypointX, waypointY, -value)
		} else if action == 'R' {
			waypointX, waypointY = p.rotate(waypointX, waypointY, value)
		} else if action == 'F' {
			boatX += waypointX * value
			boatY += waypointY * value
		} else {
			waypointX, waypointY = p.move(waypointX, waypointY, action, value)
		}
	}

	fmt.Println(math.Abs(float64(boatX)) + math.Abs(float64(boatY)))
}
