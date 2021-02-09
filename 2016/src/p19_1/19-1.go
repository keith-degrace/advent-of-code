package p19_1

import (
	"au"
	"fmt"
	"container/list"
	"time"
)

type Elf struct {
	index int
	gifts int
}

func getNext(list *list.List, element *list.Element) *list.Element {
	next := element.Next()

	if next == nil {
		next = list.Front()
	}

	return next;
}

func getWinningElf(elfCount int) int {
	elves := list.New()
	for i := 0; i < elfCount; i++ {
		elves.PushBack(Elf{i, 1})
	}

	current := elves.Front()
	for elves.Len() > 1 {
		next := getNext(elves, current)

		current.Value = Elf {
			index: current.Value.(Elf).index,
			gifts: current.Value.(Elf).gifts + next.Value.(Elf).gifts,
		}

		elves.Remove(next)

		current = getNext(elves, current)
	}

	return elves.Front().Value.(Elf).index + 1;
}

func Solve() {
	fmt.Println("Starting\n")
	startTime := time.Now()

	au.AssertIntsEqual(getWinningElf(5), 3)
	
	fmt.Println(getWinningElf(3018458))

	fmt.Println("\nCompleted in", time.Now().Sub(startTime))
}