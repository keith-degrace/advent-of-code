package main

import (
	"au"
	"fmt"
	"strings"
)

func getMarbleIndex(marbles []int, marble int) int {
	for i := 0; i < len(marbles); i++ {
		if marbles[i] == marble {
			return i
		}
	}

	return -1;
}

type Game struct {
	marbles []int
	scores map[int] int
}

func (this *Game) play(currentMarble int, marbleToPlay int, player int) int {
	if len(this.marbles) == 1 {
		this.marbles = append(this.marbles, marbleToPlay)
		return marbleToPlay
	}

	indexOfCurrentMarble := getMarbleIndex(this.marbles, currentMarble);

	if marbleToPlay != 0 && marbleToPlay % 23 == 0 {
		nextCurrentMarbleIndex := ((indexOfCurrentMarble - 6) + len(this.marbles)) % len(this.marbles)
		nextCurrentMarble := this.marbles[nextCurrentMarbleIndex]

		marbleToRemoveIndex := ((indexOfCurrentMarble - 7) + len(this.marbles)) % len(this.marbles)
		marbleToRemove := this.marbles[marbleToRemoveIndex]

		this.marbles = append(this.marbles[:marbleToRemoveIndex], this.marbles[marbleToRemoveIndex+1:]...)

		this.scores[player] += marbleToPlay + marbleToRemove

		return nextCurrentMarble
	} else {
		newMarbleIndex := (indexOfCurrentMarble + 1) % len(this.marbles) + 1

		this.marbles = append(this.marbles, 0)
		copy(this.marbles[newMarbleIndex+1:], this.marbles[newMarbleIndex:])
		this.marbles[newMarbleIndex] = marbleToPlay

		return marbleToPlay
	}
}

func (this *Game) getHighScore() int {
	highScore := 0

	for _,score := range this.scores {
		highScore = au.MaxInt(highScore, score)
	}

	return highScore
}

func (this *Game) print(current int, player int) {
	fmt.Printf("[%2d] ", player + 1)

	for _,marble := range this.marbles {
		if marble == current {
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
	lastMarble := au.ToNumber(strings.Split(input, " ")[6])

	game := Game{}

	marbleToPlay := 0
	currentMarble := 0
	game.marbles = []int{0}
	game.scores = map[int]int{}
	player := -1
	// game.print(currentMarble, player)

	for i := 0; i < lastMarble; i++ {
		player = (player + 1) % playerCount;
		marbleToPlay++

		currentMarble = game.play(currentMarble, marbleToPlay, player)
		// game.print(currentMarble, player)
		// fmt.Println(game.getHighScore(), game.lastMarbleScore, game.scores)
	}

	fmt.Println(game.getHighScore())
}
