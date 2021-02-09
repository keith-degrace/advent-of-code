package main

import (
	"au"
	"fmt"
	"regexp"
	"strconv"
)

type Puzzle14_2 struct {
}

func (p Puzzle14_2) permutate(mask string, index int, value int, setValue func(string, int)) {
	if index == len(mask)-1 {
		if mask[index] == 'X' {
			setValue(mask[:index]+"0"+mask[index+1:], value)
			setValue(mask[:index]+"1"+mask[index+1:], value)
		} else {
			setValue(mask, value)
		}
	} else {
		if mask[index] == 'X' {
			p.permutate(mask[:index]+"0"+mask[index+1:], index+1, value, setValue)
			p.permutate(mask[:index]+"1"+mask[index+1:], index+1, value, setValue)
		} else {
			p.permutate(mask, index+1, value, setValue)
		}
	}
}

func (p Puzzle14_2) run() {
	input := au.ReadInputAsStringArray("14")

	maskRegex := regexp.MustCompile("^mask = ([X01]+)$")
	memRegex := regexp.MustCompile("^mem\\[([0-9]+)\\] = ([0-9]+)$")

	mask := ""

	memory := make(map[string]int)

	for _, line := range input {

		maskMatches := maskRegex.FindStringSubmatch(line)
		if maskMatches != nil {
			mask = maskMatches[1]
		} else {
			memMatches := memRegex.FindStringSubmatch(line)

			memIndex, _ := strconv.Atoi(memMatches[1])
			memValue, _ := strconv.Atoi(memMatches[2])

			maskedMemIndex := fmt.Sprintf("%036b", memIndex)
			for i := 0; i < 36; i++ {
				if mask[i] != '0' {
					maskedMemIndex = maskedMemIndex[:i] + string(mask[i]) + maskedMemIndex[i+1:]
				}
			}

			p.permutate(maskedMemIndex, 0, memValue, func(memIndex string, value int) {
				memory[memIndex] = value
			})
		}
	}

	sum := 0
	for _, value := range memory {
		sum += value
	}

	fmt.Println(sum)
}
