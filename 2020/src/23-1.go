package main

import (
	"au"
	"container/ring"
	"fmt"
	"strings"
)

type Puzzle23_1 struct {
}

func (p Puzzle23_1) run() {
	input := au.ToNumbers(strings.Split("219748365", ""))
	// input := au.ToNumbers(strings.Split("389125467", ""))

	cups := ring.New(len(input))
	for _, entry := range input {
		cups.Value = entry
		cups = cups.Next()
	}

	current := cups

	for i := 0; i < 100; i++ {
		removed := current.Unlink(3)

		destination := current.Next()
		destinationValue := current.Value.(int) - 1

		for {
			if destination.Value.(int) == destinationValue {
				break
			}

			destination = destination.Next()
			if destination == current {

				destinationValue--
				if destinationValue == -1 {
					for value := current.Next(); value != current; value = value.Next() {
						destinationValue = au.MaxInt(destinationValue, value.Value.(int))
					}
				}
			}
		}

		destination.Link(removed)

		current = current.Next()
	}

	one := current
	for ; one.Value.(int) != 1; one = one.Next() {
	}

	for current := one.Next(); current != one; current = current.Next() {
		fmt.Printf("%v", current.Value.(int))
	}
}
