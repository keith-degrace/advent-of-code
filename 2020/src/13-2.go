package main

import (
	"au"
	"fmt"
	"strconv"
	"strings"
)

type Puzzle13_2 struct {
}

type Bus13_2 struct {
	id    uint64
	index uint64
}

func (p Puzzle13_2) getBusses(input string) []Bus13_2 {

	var busses []Bus13_2

	for index, entry := range strings.Split(input, ",") {
		if entry != "x" {
			id, _ := strconv.Atoi(entry)
			busses = append(busses, Bus13_2{(uint64)(id), (uint64)(index)})
		}
	}

	return busses
}

func (p Puzzle13_2) run() {
	input := au.ReadInputAsStringArray("13")

	busses := p.getBusses(input[1])

	timestamp := (uint64)(0)
	currentIncrement := busses[0].id

	nextIndexIncrementCapture := 1
	lastIndexCaptureMatchTimestamp := (uint64)(0)

	for {
		found := true

		for i := 1; i < len(busses); i++ {
			if (timestamp+busses[i].index)%busses[i].id != 0 {
				found = false
				break
			} else if i == nextIndexIncrementCapture {
				if lastIndexCaptureMatchTimestamp != 0 {
					nextIndexIncrementCapture++
					currentIncrement = timestamp - lastIndexCaptureMatchTimestamp
					lastIndexCaptureMatchTimestamp = 0
				} else {
					lastIndexCaptureMatchTimestamp = timestamp
				}
			}
		}

		if found {
			break
		}

		timestamp += currentIncrement
	}

	fmt.Println(timestamp - busses[0].index)
}
