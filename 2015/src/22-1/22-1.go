package main

import (
	"au"
	"container/list"
	"fmt"
)

type Player struct {
	hp int
	mana int
	manaSpent int
	shieldCounter int
	rechargeCounter int
}

type Boss struct {
	hp int
	damage int
	poisonCounter int
}

type Game struct {
	player Player
	boss Boss
	turn byte
}

func magicMissile(player Player, boss Boss) (Player, Boss) {
	player.mana -= 53
	player.manaSpent += 53
	boss.hp -= 4
	return player, boss
}

func drain(player Player, boss Boss) (Player, Boss) {
	player.hp += 2
	player.mana -= 73
	player.manaSpent += 73
	boss.hp -= 2
	return player, boss
}

func shield(player Player) Player {
	player.mana -= 113
	player.manaSpent += 113
	player.shieldCounter = 6
	return player
}

func poison(player Player, boss Boss) (Player, Boss) {
	player.mana -= 173
	player.manaSpent += 173
	boss.poisonCounter = 6
	return player, boss
}

func recharge(player Player) Player {
	player.mana -= 229
	player.manaSpent += 229
	player.rechargeCounter = 5
	return player
}

func applyEffects(player Player, boss Boss) (Player, Boss) {
	if player.shieldCounter > 0 {
		player.shieldCounter--
	}

	if player.rechargeCounter > 0 {
		player.rechargeCounter--
		player.mana += 101
	}

	if boss.poisonCounter > 0 {
		boss.poisonCounter--
		boss.hp -= 3
	}

	return player, boss
}

func bossAttack(player Player, boss Boss) (Player, Boss) {
	if player.shieldCounter > 0 {
		player.hp -= boss.damage - 7
	} else {
		player.hp -= boss.damage
	}

	return player, boss
}

func getPlayerId(player Player) string {
	return fmt.Sprintf("%v,%v,%v,%v", player.hp, player.mana, player.shieldCounter, player.rechargeCounter)
}

func getBossId(boss Boss) string {
	return fmt.Sprintf("%v,%v,%v", boss.hp, boss.damage, boss.poisonCounter)
}

func getGameId(game Game) string {
	return fmt.Sprintf("[%v][%v]", getPlayerId(game.player), getBossId(game.boss))
}

type OpenSet struct {
	list *list.List
	set map [string] bool
}

func newOpenSet() OpenSet {
	return OpenSet {
		list.New(),
		make(map [string] bool),
	}
}

func (this *OpenSet) push(game Game) {
	this.list.PushFront(game)

	gameId := getGameId(game)
	this.set[gameId] = true
}

func (this *OpenSet) pop() Game {
	subtreeRootElement := this.list.Back()
	this.list.Remove(subtreeRootElement)
	game := subtreeRootElement.Value.(Game)

	gameId := getGameId(game)
	delete(this.set, gameId)

	return game 
}

func (this *OpenSet) isEmpty() bool {
	return this.list.Len() == 0
}

func (this *OpenSet) has(game Game) bool {
	gameId := getGameId(game)
	_, ok := this.set[gameId]
	return ok
}

type ClosedSet map [string] bool

func addToClosedSet(closedSet ClosedSet, game Game) {
	gameId := getGameId(game)
	closedSet[gameId] = true
}

func isInClosedSet(closedSet ClosedSet, game Game) bool {
	gameId := getGameId(game)
	_, ok := closedSet[gameId]
	return ok
}

func getChildren(game Game) []Game {
	children := make([]Game, 0)

	if game.turn == 'P' {

		if game.player.mana >= 53 {
			player, boss := applyEffects(game.player, game.boss)
			if boss.hp > 0 {
				player, boss = magicMissile(player, boss)
			}

			children = append(children, Game {player, boss, 'B'})
		}

		if game.player.mana >= 73 {
			player, boss := applyEffects(game.player, game.boss)
			if boss.hp > 0 {
				player, boss = drain(player, boss)
			}

			children = append(children, Game {player, boss, 'B'})
		}

		if game.player.mana >= 113 {
			player, boss := applyEffects(game.player, game.boss)
			if boss.hp > 0 {
				player = shield(player)
			}

			children = append(children, Game {player, boss, 'B'})
		}

		if game.player.mana >= 173 {
			player, boss := applyEffects(game.player, game.boss)
			if boss.hp > 0 {
				player, boss = poison(player, boss)
			}

			children = append(children, Game {player, boss, 'B'})
		}

		if game.player.mana >= 229 {
			player, boss := applyEffects(game.player, game.boss)
			if boss.hp > 0 {
				player = recharge(player)
			}

			children = append(children, Game {player, boss, 'B'})
		}

	} else {
		player, boss := applyEffects(game.player, game.boss)
		if boss.hp > 0 {
			player, boss = bossAttack(player, boss)
		}

		children = append(children, Game {player, boss, 'P'})
	}

	return children
}

func getMinManaToWin(game Game) int {
	openSet := newOpenSet()
	closedSet := make(ClosedSet)

	meta := map[string] Game{}

	root := game
	meta[getGameId(root)] = Game{}
	openSet.push(root)

	minManaSpent := 999999999

	for !openSet.isEmpty() {
		subtreeRoot := openSet.pop()

		if subtreeRoot.boss.hp <= 0 {
			minManaSpent = au.MinInt(minManaSpent, subtreeRoot.player.manaSpent)
		}

		for _,child := range getChildren(subtreeRoot) {
			if child.player.hp <= 0 {
				continue
			}
			if child.player.manaSpent >= minManaSpent {
				continue
			}

			if isInClosedSet(closedSet, child) {
				continue
			}

			if !openSet.has(child) {
				meta[getGameId(child)] = subtreeRoot
				openSet.push(child)
			}
		}

		addToClosedSet(closedSet, subtreeRoot)
	}

	return minManaSpent
}

func main() {
	boss := Boss{ hp: 71, damage: 10 }
	player := Player{ hp: 50, mana: 500 }

	game := Game{ player, boss, 'P' }

	minManaSpent := getMinManaToWin(game)
	fmt.Println(minManaSpent)
}
