package p09_2

import (
	"au"
	"fmt"
)

func Solve() {
	input := au.ReadInputAsString("09")

	openGroupCount := 0
	inGarbage := false

	garbageCount := 0

	for i := 0; i < len(input); i++ {
		if inGarbage {
			if input[i] == '!' {
				i++
			} else if input[i] == '>' {
				inGarbage = false
			} else {
				garbageCount++
			}
		} else {
			if input[i] == '<' {
				inGarbage = true
			} else if input[i] == '{' {
				openGroupCount++
			} else if input[i] == '}' {
				openGroupCount--
			}
		}
	}

	fmt.Println(garbageCount)
}
