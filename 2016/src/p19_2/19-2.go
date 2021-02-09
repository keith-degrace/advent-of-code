package p19_2

import (
	"au"
	"container/list"
	"fmt"
	"time"
)

type Elf struct {
	index int
	gifts int
}

func getNext(list *list.List, element *list.Element, offset int) *list.Element {
	current := element.Next()

	for {
		if current == nil {
			current = list.Front()
		}

		offset--

		if offset == 0 {
			return current
		}

		current = current.Next()
	}
}

func getWinningElf(elfCount int) int {
	elves := list.New()
	for i := 0; i < elfCount; i++ {
		elves.PushBack(Elf{i, 1})
	}

	// Two cursors, the victim one moves either by 1 when the list length is even, or by 2 if the list legth is odd.

	current := elves.Front()
	victim := getNext(elves, current, elves.Len()/2)
	for elves.Len() > 1 {
		current.Value = Elf{
			index: current.Value.(Elf).index,
			gifts: current.Value.(Elf).gifts + victim.Value.(Elf).gifts,
		}

		nextVictim := getNext(elves, victim, (elves.Len()%2)+1)
		elves.Remove(victim)

		current = getNext(elves, current, 1)
		victim = nextVictim
	}

	return elves.Front().Value.(Elf).index + 1
}

func Solve() {
	fmt.Println("Starting\n")
	startTime := time.Now()

	au.AssertIntsEqual(getWinningElf(5), 2)

	fmt.Println(getWinningElf(3018458))

	fmt.Println("\nCompleted in", time.Now().Sub(startTime))
}
