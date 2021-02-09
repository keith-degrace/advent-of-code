package main

import (
	"au"
	"fmt"
	"strings"
	"container/list"
)

func getMarbleIndex(marbles *list.List, marble int) int {
	index := 0;

	for element := marbles.Front(); element != nil; element = element.Next() {
		if element.Value == marble {
			return index
		}

		index++;
	}

	return -1;
}

func getElementAtOffset(marbles *list.List, element *list.Element, offset int) *list.Element {
	if (offset > 0) {
		for offset != 0 {
			element = element.Next()
			if element == nil {
				element = marbles.Front()
			}
			offset--
		}
	} else {
		for offset != 0 {
			element = element.Prev()
			if element == nil {
				element = marbles.Back()
			}
			offset++
		}
	}

	return element
}

type Game struct {
	marbles *list.List
	scores map[int] int
	currentMarble *list.Element
}

func (this *Game) play(marbleToPlay int, player int) {
	if this.marbles.Len() == 1 {
		this.currentMarble = this.marbles.PushBack(marbleToPlay)
		return
	}

	if marbleToPlay != 0 && marbleToPlay % 23 == 0 {
		marble7 := getElementAtOffset(this.marbles, this.currentMarble, -7)
		this.marbles.Remove(marble7)

		this.scores[player] += marbleToPlay + marble7.Value.(int)

		this.currentMarble = getElementAtOffset(this.marbles, this.currentMarble, -6)
	} else {
		marble2 := getElementAtOffset(this.marbles, this.currentMarble, 1)
		this.currentMarble = this.marbles.InsertAfter(marbleToPlay, marble2)
	}
}

func (this *Game) getHighScore() int {
	highScore := 0

	for _,score := range this.scores {
		highScore = au.MaxInt(highScore, score)
	}

	return highScore
}

func (this *Game) print(player int) {
	fmt.Printf("[%2d] ", player + 1)

	for element := this.marbles.Front(); element != nil; element = element.Next() {
		marble := element.Value

		if marble == this.currentMarble.Value {
			fmt.Printf("(%3d) ", marble)
		} else {
			fmt.Printf(" %3d  ", marble)
		}
	}

	fmt.Printf("\n")
}

func main() {
	input := au.ReadInputAsString("09")
	// input := "9 players; last marble is worth 25 points"
	// input := "10 players; last marble is worth 1618 points"
	// input := "13 players; last marble is worth 7999 points"
	// input := "17 players; last marble is worth 1104 points"
	// input := "21 players; last marble is worth 6111 points"
	// input := "30 players; last marble is worth 5807 points"

	playerCount := au.ToNumber(strings.Split(input, " ")[0])
	lastMarble := au.ToNumber(strings.Split(input, " ")[6]) * 100

	game := Game{}
	game.marbles = list.New()
	game.scores = map[int]int{}

	game.currentMarble = game.marbles.PushBack(0)
	// game.print(-1)

	marbleToPlay := 0
	player := -1

	for i := 0; i < lastMarble; i++ {
		player = (player + 1) % playerCount;
		marbleToPlay++

		game.play(marbleToPlay, player)
		// game.print(player)
		// fmt.Println(game.getHighScore(), game.lastMarbleScore, game.scores)
	}

	fmt.Println(game.getHighScore())
}
