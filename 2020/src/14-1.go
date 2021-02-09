package main

import (
	"au"
	"fmt"
	"regexp"
	"strconv"
)

type Puzzle14_1 struct {
}

func (p Puzzle14_1) applyMask(value uint64, mask string) uint64 {

	for i := 0; i < len(mask); i++ {
		bitPosition := uint64(len(mask) - i - 1)

		if mask[i] == '0' {
			value = value & uint64(^(1 << bitPosition))
		} else if mask[i] == '1' {
			value = value | uint64(1<<bitPosition)
		}
	}

	return value
}

func (p Puzzle14_1) run() {
	input := au.ReadInputAsStringArray("14")

	maskRegex := regexp.MustCompile("^mask = ([X01]+)$")
	memRegex := regexp.MustCompile("^mem\\[([0-9]+)\\] = ([0-9]+)$")

	mask := ""
	memory := make(map[int]uint64)

	for _, line := range input {

		maskMatches := maskRegex.FindStringSubmatch(line)
		if maskMatches != nil {
			mask = maskMatches[1]
		} else {
			memMatches := memRegex.FindStringSubmatch(line)

			index, _ := strconv.Atoi(memMatches[1])
			value, _ := strconv.Atoi(memMatches[2])

			memory[index] = p.applyMask(uint64(value), mask)
		}
	}

	sum := uint64(0)
	for _, value := range memory {
		sum += value
	}

	fmt.Println(sum)
}
