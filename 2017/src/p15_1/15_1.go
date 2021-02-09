package p15_1

import (
	"fmt"
)

func Solve() {
	a, b := 883, 879
	// a, b := 65, 8921

	count := 0
	for i := 0; i < 40000000; i++ {
		a = (a * 16807) % 2147483647
		b = (b * 48271) % 2147483647

		if a&0xFFFF == b&0xFFFF {
			count++
		}
	}

	fmt.Println(count)
}
