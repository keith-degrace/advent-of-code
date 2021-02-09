package main

import (
	"au"
	"fmt"
	"strconv"
	"strings"
)

type Puzzle13_1 struct {
}

func (p Puzzle13_1) getBusses(input string) []int {

	var busses []int

	for _, entry := range strings.Split(input, ",") {
		if entry != "x" {
			busId, _ := strconv.Atoi(entry)
			busses = append(busses, busId)
		}
	}

	return busses
}

func (p Puzzle13_1) run() {
	input := au.ReadInputAsStringArray("13")

	timestamp, _ := strconv.Atoi(input[0])
	busses := p.getBusses(input[1])

	smallestWait := 10000000
	smallestResult := 0

	for _, bus := range busses {

		wait := bus - timestamp%bus

		if wait < smallestWait {
			smallestWait = wait
			smallestResult = wait * bus
		}
	}

	fmt.Println(smallestResult)
}
