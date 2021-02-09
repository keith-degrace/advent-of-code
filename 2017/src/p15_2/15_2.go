package p15_2

import (
	"container/list"
	"fmt"
)

func Solve() {
	a, b := 883, 879
	// a, b := 65, 8921

	stackA := list.New()
	stackB := list.New()

	pairs := 0
	count := 0
	for pairs < 5000000 {
		a = (a * 16807) % 2147483647
		b = (b * 48271) % 2147483647

		if a%4 == 0 {
			stackA.PushBack(a)
		}

		if b%8 == 0 {
			stackB.PushBack(b)
		}

		for stackA.Len() > 0 && stackB.Len() > 0 {
			headA := stackA.Front()
			stackA.Remove(headA)
			valueA := headA.Value.(int)

			headB := stackB.Front()
			stackB.Remove(headB)
			valueB := headB.Value.(int)

			if valueA&0xFFFF == valueB&0xFFFF {
				count++
			}

			pairs++
		}
	}

	fmt.Println(count)
}
