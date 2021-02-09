package p09_1

import (
	"au"
	"fmt"
)

func Solve() {
	input := au.ReadInputAsString("09")

	openGroupCount := 0
	inGarbage := false

	score := 0

	for i := 0; i < len(input); i++ {
		// Ignore anything after an exclamation point
		if input[i] == '!' {
			i++
			continue
		}

		if inGarbage {
			if input[i] == '>' {
				inGarbage = false
			}
		} else {
			if input[i] == '<' {
				inGarbage = true
			} else if input[i] == '{' {
				openGroupCount++
				score += openGroupCount
			} else if input[i] == '}' {
				openGroupCount--
			}
		}
	}

	fmt.Println(score)
}
