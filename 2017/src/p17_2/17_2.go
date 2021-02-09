package p17_2

import (
	"container/ring"
	"fmt"
)

func Solve() {
	input := 377

	zero := ring.New(1)
	zero.Value = 0

	current := zero
	for i := 1; i <= 50000000; i++ {
		current = current.Move(input)

		newValue := ring.New(1)
		newValue.Value = i

		current.Link(newValue)
		current = newValue
	}

	// Takes like 5 minutes to complete
	fmt.Println(zero.Next().Value.(int))
}
