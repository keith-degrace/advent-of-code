package main

import (
	"au"
	"fmt"
	"strings"
)

type Puzzle15_1 struct {
}

func (p Puzzle15_1) LastCalled(numbers []int, value int) int {
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] == value {
			return i + 1
		}
	}

	return -1
}

func (p Puzzle15_1) run() {
	input := "16,1,0,18,12,14,19"
	// input := "0,3,6"

	numbers := []int{}
	for _, entry := range strings.Split(input, ",") {
		numbers = append(numbers, au.ToNumber(entry))
	}

	for len(numbers) < 2020 {
		var newNumber int

		turn := len(numbers) + 1
		lastNumber := numbers[len(numbers)-1]
		lastNumberCalledOnTurn := p.LastCalled(numbers[:len(numbers)-1], lastNumber)

		if lastNumberCalledOnTurn == -1 {
			newNumber = 0
		} else {
			newNumber = turn - 1 - lastNumberCalledOnTurn
		}

		numbers = append(numbers, newNumber)

		fmt.Printf("%vth number spoken is %v\n\n", turn, newNumber)
	}
}
