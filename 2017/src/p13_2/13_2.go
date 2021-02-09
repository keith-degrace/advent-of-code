package p13_2

import (
	"au"
	"fmt"
	"strings"
)

func Solve() {
	input := au.ReadInputAsStringArray("13")

	layers := make(map[int]int)
	for _, line := range input {
		tokens := strings.Split(line, ": ")

		id := au.ToNumber(tokens[0])
		depth := au.ToNumber(tokens[1])

		layers[id] = depth
	}

	delay := 0
	for {

		found := true

		// The moving back and forward thing is a distraction.  We can just do a circular
		// thing with a module instead.  So it's just a matter of finding a spot where
		// the module is non-zero across the layers.
		for id, depth := range layers {
			if (id+delay)%((depth-1)*2) == 0 {
				found = false
				break
			}
		}

		if found {
			break
		}

		delay++
	}

	fmt.Println(delay)
}
