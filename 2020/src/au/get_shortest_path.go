package au

import (
	"math"
	"sort"
)

type GetNeighbors func(coord Coord) []Coord

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
		id: coord.Key(),
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
	return s[i].fScore < s[j].fScore
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

func addToClosedSet(closedSet *ClosedSet, coord *Coord) {
	(*closedSet)[coord.Key()] = true
}

func isInClosedSet(closedSet *ClosedSet, coord *Coord) bool {
	_, ok := (*closedSet)[coord.Key()]
	return ok
}

func getDistance(coord1 *Coord, coord2 *Coord) float64 {
	dx := float64(coord2.x - coord1.x)
	dy := float64(coord2.y - coord1.y)

	return math.Pow(dx, 2) + math.Pow(dy, 2)
}

func buildPath(endNode *Node) []Coord {
	path := []Coord {}

	for current := endNode; current != nil; current = current.parent {
		path = append(path, current.coord)
	}

	path = path[:len(path) - 1]

	for i := len(path)/2-1; i >= 0; i-- {
		opp := len(path)-1-i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}

func GetShortestPath(from Coord, to Coord, getNeighbors GetNeighbors) []Coord {
	openSet := newOpenSet()
	closedSet := make(ClosedSet)

	startNode := newNode(from)

	openSet.push(startNode)

	for !openSet.isEmpty() {
		currentNode := openSet.popSmallestF()

		addToClosedSet(&closedSet, &currentNode.coord)

		if currentNode.coord.Equals(to) {
			return buildPath(&currentNode)
		}
		
		for _,neighborCoord := range getNeighbors(currentNode.coord) {
			if isInClosedSet(&closedSet, &neighborCoord) {
				continue
			}

			neighborNode := newNode(neighborCoord)
			neighborNode.parent = &currentNode
			neighborNode.gScore = currentNode.gScore + 1
			neighborNode.hScore = getDistance(&neighborNode.coord, &to)
			neighborNode.fScore = neighborNode.gScore + neighborNode.hScore
			openSet.push(neighborNode)
		}
	}

	return []Coord{}
}
