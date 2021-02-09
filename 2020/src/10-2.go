package main

import (
	"au"
	"fmt"
	"sort"
)

type Puzzle10_2 struct {
}

func (p Puzzle10_2) getPathCount(input []int, cache map[int]int, index int) int {
	value, ok := cache[index]
	if ok {
		return value
	}

	if index >= len(input)-1 {
		return 1
	}

	count := 0

	if index < len(input)-1 && input[index+1]-input[index] < 4 {
		count += p.getPathCount(input, cache, index+1)

		if index < len(input)-2 && input[index+2]-input[index] < 4 {
			count += p.getPathCount(input, cache, index+2)

			if index < len(input)-3 && input[index+3]-input[index] < 4 {
				count += p.getPathCount(input, cache, index+3)
			}
		}
	}

	cache[index] = count

	return count
}

func (p Puzzle10_2) run() {
	input := au.ReadInputAsNumberArray("10")

	sort.Ints(input)
	input = append([]int{0}, input...)

	cache := make(map[int]int)

	count := p.getPathCount(input, cache, 0)

	fmt.Println(count)
}
