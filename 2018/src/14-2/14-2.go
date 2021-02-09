package main

import (
	// "au"
	"container/list"
	"fmt"
)

type Scoreboard struct {
	recipes *list.List
	current1 *list.Element
	current2 *list.Element
}

func NewScoreboard() Scoreboard {
	scoreboard := Scoreboard{}

	scoreboard.recipes = list.New()
	scoreboard.current1 = scoreboard.recipes.PushBack(3)
	scoreboard.current2 = scoreboard.recipes.PushBack(7)

	return scoreboard;
}

func (this *Scoreboard) print() {
	for e := this.recipes.Front(); e != nil; e = e.Next() {
		value := ""
		if e == this.current1 {
			value = fmt.Sprintf("(%v)", e.Value)
		} else if e == this.current2 {
			value = fmt.Sprintf("[%v]", e.Value)
		} else {
			value = fmt.Sprintf(" %v ", e.Value)
		}

		fmt.Printf("%2v", value)
	}

	fmt.Println()
}

func addNewRecipes(scoreboard *Scoreboard) {
	score1 := scoreboard.current1.Value.(int)
	score2 := scoreboard.current2.Value.(int)

	sum := score1 + score2

	if sum > 19 { 
		panic(fmt.Sprintf("%v,%v,%v", score1, score2, sum))
	}

	if sum < 10 {
		scoreboard.recipes.PushBack(sum)
	} else {
		scoreboard.recipes.PushBack(1)
		scoreboard.recipes.PushBack(sum % 10)
	}
}

func getNewCurrent(scoreboard *Scoreboard, current *list.Element) *list.Element {
	advanceBy := current.Value.(int)

	for i := 0; i <= advanceBy; i++ {
		next := current.Next()
		if next != nil {
			current = next
		} else {
			current = scoreboard.recipes.Front()
		}
	}

	return current
}

func updateCurrents(scoreboard *Scoreboard) {
	scoreboard.current1 = getNewCurrent(scoreboard, scoreboard.current1)
	scoreboard.current2 = getNewCurrent(scoreboard, scoreboard.current2)
}

func iterate(scoreboard *Scoreboard) {
	addNewRecipes(scoreboard)
	updateCurrents(scoreboard)
}

func getCandidate(scoreboard *Scoreboard, len int, offset int) string {
	if scoreboard.recipes.Len() < (len + offset) {
		return ""
	}

	candidate := ""

	current := scoreboard.recipes.Back()
	for i := 0; i < (len + offset) - 1; i++ {
		current = current.Prev()
	}

	for i := 0; i < len; i++ {
		candidate += fmt.Sprintf("%v", current.Value.(int))
		current = current.Next()
	}

	return candidate
}

func main() {
	input := "633601"
	// input = "59414"

	scoreboard := NewScoreboard()

	for {
		candidate1 := getCandidate(&scoreboard, len(input), 0)
		if candidate1 == input {
			fmt.Println(scoreboard.recipes.Len() - len(input))
			break
		} 

		candidate2 := getCandidate(&scoreboard, len(input), 1)
		if candidate2 == input {
			fmt.Println(scoreboard.recipes.Len() - len(input) - 1)
			break
		} 

		iterate(&scoreboard)
	}
}