package p17_1

import (
	"container/ring"
	"fmt"
)

func Solve() {
	input := 377
	// input := 3

	current := ring.New(1)
	current.Value = 0

	for i := 1; i <= 2017; i++ {
		current := current.Move(input + 1)

		newValue := ring.New(1)
		newValue.Value = i

		current.Link(newValue)
		current = newValue
	}

	fmt.Println(current.Next().Value.(int))
}
