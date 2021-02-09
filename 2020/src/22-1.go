package main

import (
	"au"
	"container/list"
	"fmt"
	"strings"
)

type Puzzle22_1 struct {
}

type Player22_1 struct {
	id    int
	cards *list.List
}

func NewPlayer22_1(id int) *Player22_1 {
	player := new(Player22_1)
	player.id = id
	player.cards = list.New()
	return player
}

func (p Player22_1) getScore() int {
	score := 0

	i := 1
	for card := p.cards.Back(); card != nil; card = card.Prev() {
		score += card.Value.(int) * i
		i++
	}

	return score
}

func (p Player22_1) print() {
	fmt.Printf("Player %v: ", p.id)
	for card := p.cards.Front(); card != nil; card = card.Next() {
		fmt.Printf("%v ", card.Value.(int))
	}
	fmt.Println()
}

func (p Puzzle22_1) parse(input []string) []*Player22_1 {
	players := make([]*Player22_1, 2)

	currentPlayer := -1
	for _, line := range input {
		if strings.HasPrefix(line, "Player") {
			currentPlayer++
			players[currentPlayer] = NewPlayer22_1(currentPlayer)
		} else if len(line) > 0 {
			players[currentPlayer].cards.PushBack(au.ToNumber(line))
		}
	}

	return players
}

func (p Puzzle22_1) playRound(players []*Player22_1) {

	topCard0 := players[0].cards.Front()
	players[0].cards.Remove(topCard0)

	topCard1 := players[1].cards.Front()
	players[1].cards.Remove(topCard1)

	if topCard0.Value.(int) > topCard1.Value.(int) {
		players[0].cards.PushBack(topCard0.Value.(int))
		players[0].cards.PushBack(topCard1.Value.(int))
	} else if topCard0.Value.(int) < topCard1.Value.(int) {
		players[1].cards.PushBack(topCard1.Value.(int))
		players[1].cards.PushBack(topCard0.Value.(int))
	} else {
		players[0].cards.PushBack(topCard0.Value.(int))
		players[1].cards.PushBack(topCard1.Value.(int))
	}
}

func (p Puzzle22_1) run() {
	input := au.ReadInputAsStringArray("22")

	players := p.parse(input)

	for _, player := range players {
		player.print()
	}

	round := 1
	for {
		p.playRound(players)

		if players[0].cards.Len() == 0 {
			fmt.Printf("Player %v wins with score of %v\n", players[1].id, players[1].getScore())
			break
		}

		if players[1].cards.Len() == 0 {
			fmt.Printf("Player %v wins with score of %v\n", players[0].id, players[0].getScore())
			break
		}

		round++
	}
}
