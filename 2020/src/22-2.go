package main

import (
	"au"
	"container/list"
	"fmt"
	"strings"
)

type Puzzle22_2 struct {
}

type Player22_2 struct {
	id    int
	cards *list.List
}

func NewPlayer22_2(id int) *Player22_2 {
	player := new(Player22_2)
	player.id = id
	player.cards = list.New()
	return player
}

func (p Player22_2) getScore() int {
	score := 0

	i := 1
	for card := p.cards.Back(); card != nil; card = card.Prev() {
		score += card.Value.(int) * i
		i++
	}

	return score
}

func (p Player22_2) serialize() string {
	value := fmt.Sprintf("[%v] ", p.id)

	for card := p.cards.Front(); card != nil; card = card.Next() {
		value += fmt.Sprintf("%v ", card.Value.(int))
	}

	return value
}

func (p Player22_2) print() {
	fmt.Println(p.serialize())
}

type History22_2 struct {
	rounds map[string]bool
}

func NewHistory22_2() History22_2 {
	return History22_2{make(map[string]bool)}
}

func (h *History22_2) AddRound(players []*Player22_2) {
	h.rounds[players[0].serialize()+players[1].serialize()] = true
}

func (h *History22_2) HasSeen(players []*Player22_2) bool {
	_, ok := h.rounds[players[0].serialize()+players[1].serialize()]
	return ok
}

func (p Puzzle22_2) parse(input []string) []*Player22_2 {
	players := make([]*Player22_2, 2)

	currentPlayer := -1
	for _, line := range input {
		if strings.HasPrefix(line, "Player") {
			currentPlayer++
			players[currentPlayer] = NewPlayer22_2(currentPlayer)
		} else if len(line) > 0 {
			players[currentPlayer].cards.PushBack(au.ToNumber(line))
		}
	}

	return players
}

func (p Puzzle22_2) playRecursiveGame(players []*Player22_2, topCard0, topCard1 int) int {
	recursivePlayers := make([]*Player22_2, 2)

	recursivePlayers[0] = NewPlayer22_2(0)
	{
		card0 := players[0].cards.Front()
		for i := 0; i < topCard0; i++ {
			recursivePlayers[0].cards.PushBack(card0.Value.(int))
			card0 = card0.Next()
		}
	}

	recursivePlayers[1] = NewPlayer22_2(1)
	{
		card1 := players[1].cards.Front()
		for i := 0; i < topCard1; i++ {
			recursivePlayers[1].cards.PushBack(card1.Value.(int))
			card1 = card1.Next()
		}
	}

	winner, _ := p.playGame(recursivePlayers)

	return winner
}

func (p Puzzle22_2) playGame(players []*Player22_2) (int, int) {
	history := NewHistory22_2()

	for round := 1; ; round++ {
		if history.HasSeen(players) {
			return 0, players[0].getScore()
		}

		history.AddRound(players)

		// Draw top cards
		topCard0 := players[0].cards.Front()
		players[0].cards.Remove(topCard0)

		topCard1 := players[1].cards.Front()
		players[1].cards.Remove(topCard1)

		// If both players have at least as many cards remaining in their deck as the value of the card they
		// just drew, the winner of the round is determined by playing a new game of Recursive Combat
		roundWinner := -1
		if players[0].cards.Len() >= topCard0.Value.(int) && players[1].cards.Len() >= topCard1.Value.(int) {
			roundWinner = p.playRecursiveGame(players, topCard0.Value.(int), topCard1.Value.(int))
		} else {

			if topCard0.Value.(int) > topCard1.Value.(int) {
				roundWinner = 0
			} else if topCard0.Value.(int) < topCard1.Value.(int) {
				roundWinner = 1
			}
		}

		// Distribute card to winner
		if roundWinner == 0 {
			players[0].cards.PushBack(topCard0.Value.(int))
			players[0].cards.PushBack(topCard1.Value.(int))
		} else if roundWinner == 1 {
			players[1].cards.PushBack(topCard1.Value.(int))
			players[1].cards.PushBack(topCard0.Value.(int))
		} else {
			players[0].cards.PushBack(topCard0.Value.(int))
			players[1].cards.PushBack(topCard1.Value.(int))
		}

		// See if we have a winner.
		if players[0].cards.Len() == 0 {
			return 1, players[1].getScore()
		}

		if players[1].cards.Len() == 0 {
			return 0, players[0].getScore()
		}
	}

}

func (p Puzzle22_2) run() {
	input := au.ReadInputAsStringArray("22")

	players := p.parse(input)

	_, score := p.playGame(players)
	fmt.Println(score)
}
