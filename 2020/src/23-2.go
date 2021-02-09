package main

import (
	"au"
	"container/ring"
	"fmt"
	"strings"
)

type Puzzle23_2 struct {
}

func (p Puzzle23_2) run() {
	input := au.ToNumbers(strings.Split("219748365", ""))
	// input := au.ToNumbers(strings.Split("389125467", ""))

	largestValue := 0

	cups := ring.New(1000000)

	ringMap := make(map[int]*ring.Ring)

	for _, entry := range input {
		cups.Value = entry
		ringMap[entry] = cups
		largestValue = au.MaxInt(largestValue, entry)
		cups = cups.Next()
	}

	for i := 0; i < 1000000-len(input); i++ {
		value := largestValue + 1 + i
		cups.Value = value
		ringMap[value] = cups
		cups = cups.Next()
	}

	largestValue = cups.Prev().Value.(int)

	current := cups

	for i := 0; i < 10000000; i++ {
		removed := current.Unlink(3)
		removed1 := removed.Value.(int)
		removed2 := removed.Next().Value.(int)
		removed3 := removed.Next().Next().Value.(int)

		destinationTargetValue := current.Value.(int) - 1

		for {
			if destinationTargetValue < 1 {
				destinationTargetValue = largestValue
			}

			if destinationTargetValue != removed1 && destinationTargetValue != removed2 && destinationTargetValue != removed3 {
				destination := ringMap[destinationTargetValue]
				destination.Link(removed)
				break
			}

			destinationTargetValue--
		}

		current = current.Next()
	}

	one := ringMap[1]

	cup1 := one.Next().Value.(int)
	cup2 := one.Next().Next().Value.(int)

	fmt.Println(cup1 * cup2)
}
