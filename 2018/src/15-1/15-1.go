package main

import (
	"au"
	"fmt"
	"math"
	"sort"
)

func testInputs1() []string {
	return []string {
		"#######",
		"#.G...#",
		"#...EG#",
		"#.#.#G#",
		"#..G#E#",
		"#.....#",
		"#######",
	}
}

func testInputs2() []string {
	return []string {
		"#########",
		"#G..G..G#",
		"#.......#",
		"#.......#",
		"#G..E..G#",
		"#.......#",
		"#.......#",
		"#G..G..G#",
		"#########",
	}
}

func testInputs3() []string {
	return []string {
		"#######",
		"#G..#E#",
		"#E#E.E#",
		"#G.##.#",
		"#...#E#",
		"#...E.#",
		"#######",
	}
}

func testInputs4() []string {
	return []string {
		"#######",
		"#E..EG#",
		"#.#G.E#",
		"#E.##E#",
		"#G..#.#",
		"#..E#.#",
		"#######",
	}
}

func testInputs5() []string {
	return []string {
		"#######",
		"#E.G#.#",
		"#.#G..#",
		"#G.#.G#",
		"#G..#.#",
		"#...E.#",
		"#######",
	}
}

func testInputs6() []string {
	return []string {
		"#######",
		"#.E...#",
		"#.#..G#",
		"#.###.#",
		"#E#G#G#",
		"#...#G#",
		"#######",
	}
}

func testInputs7() []string {
	return []string {
		"#########",
		"#G......#",
		"#.E.#...#",
		"#..##..G#",
		"#...##..#",
		"#...#...#",
		"#.G...G.#",
		"#.....G.#",
		"#########",
	}
}

type CaveMap struct {
	data map[string] bool
	width int
	height int
}

func (this *CaveMap) isWall(x int, y int) bool {
	_, ok := this.data[getCoordKey(x, y)]
	return ok
}

type Coord struct {
	x int
	y int
}

func (this *Coord) equals(other Coord) bool {
	return this.x == other.x && this.y == other.y
}

func getCoordKey(x int, y int) string {
	return fmt.Sprintf("%vx%v", x, y)
}

type Player struct {
	id int
	kind byte
	position Coord
	hitPoints int
	attackPower int
}

type byPlayer []Player

func (s byPlayer) Len() int {
	return len(s)
}
func (s byPlayer) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byPlayer) Less(i, j int) bool {
	if s[i].position.y < s[j].position.y {
		return true
	} else if s[i].position.y > s[j].position.y {
		return false
	} 
	
	return s[i].position.x < s[j].position.x
}

func parseInputs(inputs []string) (CaveMap, []Player) {
	caveMap := CaveMap{}
	caveMap.data = map[string] bool{}

	players := []Player{}

	nextPlayerId := 0

	for y,input := range inputs {
		for x,char := range input {
			if char == '#' {
				caveMap.data[getCoordKey(x, y)] = true
				caveMap.width = au.MaxInt(caveMap.width, x + 1)
				caveMap.height = au.MaxInt(caveMap.height, y + 1)
			} else if char == 'G' {
				players = append(players, Player{nextPlayerId, 'G', Coord { x, y }, 200, 3 })
				nextPlayerId++
			} else if char == 'E' {
				players = append(players, Player{nextPlayerId, 'E', Coord { x, y }, 200, 3 })
				nextPlayerId++
			}
		}
	}

	return caveMap, players
}

func printCaveMap(caveMap CaveMap, players []Player) {
	screen := au.NewScreen(caveMap.width, caveMap.height)

	for x := 0; x < caveMap.width; x++ {
		for y := 0; y < caveMap.height; y++ {
			if caveMap.isWall(x, y) {
				screen.SetPixel(x, y, '#')
			}
		}
	}

	for _,player := range players {
		screen.SetPixel(player.position.x, player.position.y, player.kind)
	}

	screen.Print()
}

func printCaveMapWithPath(caveMap CaveMap, players []Player, path []Coord) {
	screen := au.NewScreen(caveMap.width, caveMap.height)

	for x := 0; x < caveMap.width; x++ {
		for y := 0; y < caveMap.height; y++ {
			if caveMap.isWall(x, y) {
				screen.SetPixel(x, y, '#')
			}
		}
	}

	for _,player := range players {
		screen.SetPixel(player.position.x, player.position.y, player.kind)
	}

	for index,coord := range path {
		screen.SetPixel(coord.x, coord.y, byte('0' + (index % 10)))
	}

	screen.Print()
}

func printPlayerHP(players []Player) {
	sort.Sort(byPlayer(players))

	for _, player := range players {
		fmt.Printf("%v(%v), ", string(player.kind), player.hitPoints)
	}

	fmt.Println()
}

type Node struct {
	id string
	coord Coord
	parent *Node
	gScore float64
	fScore float64
	hScore float64
}

func newNode(coord Coord) Node {
	return Node {
		id: getCoordKey(coord.x, coord.y),
		coord: coord,
		gScore: 0,
		fScore: 0,
	}
}

type byNode []Node 

func (s byNode) Len() int {
	return len(s)
}
func (s byNode) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byNode) Less(i, j int) bool {
	if s[i].fScore < s[j].fScore {
		return true
	} else if s[i].fScore > s[j].fScore {
		return false
	} 

	if s[i].coord.y > s[j].coord.y {
		return true
	} else if s[i].coord.y < s[j].coord.y {
		return false
	} 
	
	return s[i].coord.x > s[j].coord.x
}

type OpenSet struct {
	nodes map [string] Node
}

func newOpenSet() OpenSet {
	return OpenSet {
		make(map [string] Node),
	}
}

func (this *OpenSet) push(node Node) {
	this.nodes[node.id] = node
}

func (this *OpenSet) popSmallestF() Node {
	nodes := []Node {}
	for _, node := range this.nodes {
		nodes = append(nodes, node)
	}

	sort.Sort(byNode(nodes))

	delete(this.nodes, nodes[0].id)

	return nodes[0]
}

func (this *OpenSet) isEmpty() bool {
	return len(this.nodes) == 0
}

type ClosedSet map [string] bool

func addToClosedSet(closedSet ClosedSet, coord Coord) {
	closedSet[getCoordKey(coord.x, coord.y)] = true
}

func isInClosedSet(closedSet ClosedSet, coord Coord) bool {
	_, ok := closedSet[getCoordKey(coord.x, coord.y)]
	return ok
}

func hasPlayer(players []Player, coord Coord, movingPlayer Player, targetPlayer Player) bool {
	for _,player := range players {
		if player.id == movingPlayer.id ||player.id == targetPlayer.id {
			continue
		}

		if player.position.equals(coord) {
			return true
		}
	}

	return false
}

func getNeighbors(node Node, caveMap CaveMap, players []Player, movingPlayer Player, targetPlayer Player) []Coord {
	offsets := [][] int {
		[]int { 0, -1},
		[]int { 1,  0},
		[]int { 0,  1},
		[]int {-1,  0},
	}

	children := []Coord{}

	for _, offset := range offsets {
		coord := Coord {
			node.coord.x + offset[0],
			node.coord.y + offset[1],
		}

		if caveMap.isWall(coord.x, coord.y) {
			continue
		}

		if hasPlayer(players, coord, movingPlayer, targetPlayer) {
			continue
		}

		children = append(children, coord)
	}

	return children
}

func getDistance(coord1 Coord, coord2 Coord) float64 {
	dx := float64(coord2.x - coord1.x)
	dy := float64(coord2.y - coord1.y)

	return math.Pow(dx, 2) + math.Pow(dy, 2)
}

func buildPath(endNode Node) []Coord {
	path := []Coord {}

	for current := &endNode; current != nil; current = current.parent {
		path = append(path, current.coord)
	}

	path = path[:len(path) - 1]

	for i := len(path)/2-1; i >= 0; i-- {
		opp := len(path)-1-i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}

func getShortestPath(caveMap CaveMap, players []Player, movingPlayer Player, targetPlayer Player) []Coord {
	openSet := newOpenSet()
	closedSet := make(ClosedSet)

	startNode := newNode(movingPlayer.position)

	openSet.push(startNode)

	for !openSet.isEmpty() {
		currentNode := openSet.popSmallestF()

		addToClosedSet(closedSet, currentNode.coord)

		if currentNode.coord.equals(targetPlayer.position) {
			return buildPath(currentNode)
		}
		
		for _,neighborCoord := range getNeighbors(currentNode, caveMap, players, movingPlayer, targetPlayer) {
			if isInClosedSet(closedSet, neighborCoord) {
				continue
			}

			neighborNode := newNode(neighborCoord)
			neighborNode.parent = &currentNode
			neighborNode.gScore = currentNode.gScore + 1
			neighborNode.hScore = getDistance(currentNode.coord, neighborNode.coord)
			neighborNode.fScore = neighborNode.gScore + neighborNode.hScore
			openSet.push(neighborNode)
		}
	}

	return []Coord{}
}

func getClosestEnemy(caveMap CaveMap, players []Player, player Player) ([]int, []Coord) {
	closestEnemies := []int{}
	closestPath := []Coord{}

	for candidatePlayerIndex,candidatePlayer := range players {
		if candidatePlayer.kind == player.kind {
			continue
		}
	
		path := getShortestPath(caveMap, players, player, candidatePlayer)
		if len(path) <= 0 {
			continue
		}

		if len(closestPath) == 0 || len(path) < len(closestPath) {
			closestEnemies = []int { candidatePlayerIndex }
			closestPath = path
		} else if len(path) == len(closestPath) {
			closestEnemies = append(closestEnemies, candidatePlayerIndex)
		}
	}

	return closestEnemies, closestPath
}

func getPlayerWithLowestHitPoints(players []Player, closestEnemiesIndices []int) int {
	lowestHitPointIndex := closestEnemiesIndices[0]

	for _,closestEnemiesIndex:= range closestEnemiesIndices {
		if players[closestEnemiesIndex].hitPoints < players[lowestHitPointIndex].hitPoints {
			lowestHitPointIndex = closestEnemiesIndex
		}
	}

	return lowestHitPointIndex
}

func iterate(caveMap CaveMap, players []Player) []Player {
	sort.Sort(byPlayer(players))
	
	for i := 0; i < len(players); i++ {
		closestEnemiesIndices, closestEnemyPath := getClosestEnemy(caveMap, players, players[i])
		
		if len(closestEnemyPath) > 1 {
			players[i].position.x = closestEnemyPath[0].x
			players[i].position.y = closestEnemyPath[0].y
		}

		if len(closestEnemyPath) == 1 || len(closestEnemyPath) == 2 {
			enemyToAttackIndex := getPlayerWithLowestHitPoints(players, closestEnemiesIndices)

			players[enemyToAttackIndex].hitPoints -= players[i].attackPower
			if players[enemyToAttackIndex].hitPoints <= 0 {
				players = append(players[:enemyToAttackIndex], players[enemyToAttackIndex+1:]...)
				if enemyToAttackIndex < i {
					i--
				}
			}
		}
	}

	return players;
}

func hasWinner(players []Player) bool {
	elfCount := 0
	goblingCount := 0

	for _,player := range players {
		if player.hitPoints <= 0 {
			continue
		}

		if player.kind == 'E' {
			elfCount++
		} else {
			goblingCount++
		}
	}

	return elfCount == 0 || goblingCount == 0
}

func getHitPointSum(players []Player) int {
	sum := 0

	for _,player := range players {
		sum += player.hitPoints;
	}

	return sum
}

func main() {
	inputs := au.ReadInputAsStringArray("15")
	// inputs = testInputs7()

	caveMap, players := parseInputs(inputs)

	round := 1
	for {
		players = iterate(caveMap, players)
		if hasWinner(players) {
			break;
		}
		
		round++
	}

	printCaveMap(caveMap, players)
	printPlayerHP(players)
	fmt.Println("Round", round)
	fmt.Println("HP", getHitPointSum(players))
	fmt.Println((round - 1) * getHitPointSum(players))
}