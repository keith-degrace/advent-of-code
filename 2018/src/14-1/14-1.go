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

func getNextTenScore(scoreboard *Scoreboard, input int) string {
	current := scoreboard.recipes.Front()
	for i := 0; i < input - 1; i++ {
		current = current.Next()
	}

	nextTenScore := ""
	for i := 0; i < 10; i++ {
		current = current.Next()
		nextTenScore += fmt.Sprintf("%v", current.Value.(int))
	}

	return nextTenScore
}

func main() {
	input := 633601
	// input := 2018

	scoreboard := NewScoreboard()

	for {
		iterate(&scoreboard)

		if scoreboard.recipes.Len() >= (input + 10) {
			break;
		}
	}

	fmt.Println(getNextTenScore(&scoreboard, input))
}