package au

import (
	"math"
	"sort"
)

type GetNeighbors func(coord Coord) []Coord

type Node struct {
	Coord Coord
	Parent *Node
	ScoreG float64
	ScoreF float64
	ScoreH float64
}

func NewNode(coord Coord) Node {
	return Node {
		Coord: coord,
		ScoreG: 0,
		ScoreF: 0,
	}
}

type ByNode []Node 

func (s ByNode) Len() int {
	return len(s)
}

func (s ByNode) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByNode) Less(i, j int) bool {
	return s[i].ScoreF < s[j].ScoreF
}

type OpenSet map [string] Node

func (this *OpenSet) Push(node Node) {
	(*this)[node.Coord.Key()] = node
}

func (this *OpenSet) Pop() Node {
	nodes := []Node {}
	for _, node := range *this {
		nodes = append(nodes, node)
	}

	sort.Sort(ByNode(nodes))

	delete(*this, nodes[0].Coord.Key())

	return nodes[0]
}

func (this *OpenSet) IsEmpty() bool {
	return len(*this) == 0
}

type ClosedSet map [string] bool

func (this *ClosedSet) Add(coord *Coord) {
	(*this)[coord.Key()] = true
}

func (this *ClosedSet) Has(coord *Coord) bool {
	_, ok := (*this)[coord.Key()]
	return ok
}

func GetDistance(coord1 *Coord, coord2 *Coord) float64 {
	dx := float64(coord2.x - coord1.x)
	dy := float64(coord2.y - coord1.y)

	return math.Pow(dx, 2) + math.Pow(dy, 2)
}

func BuildPath(endNode *Node) []Coord {
	path := []Coord {}

	for current := endNode; current != nil; current = current.Parent {
		path = append(path, current.Coord)
	}

	path = path[:len(path) - 1]

	for i := len(path)/2-1; i >= 0; i-- {
		opp := len(path)-1-i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}

func GetShortestPath(from Coord, to Coord, getNeighbors GetNeighbors) []Coord {
	openSet := make(OpenSet)
	closedSet := make(ClosedSet)

	startNode := NewNode(from)

	openSet.Push(startNode)

	for !openSet.IsEmpty() {
		currentNode := openSet.Pop()

		closedSet.Add(&currentNode.Coord)

		if currentNode.Coord.Equals(to) {
			return BuildPath(&currentNode)
		}
		
		for _,neighborCoord := range getNeighbors(currentNode.Coord) {
			if closedSet.Has(&neighborCoord) {
				continue
			}

			neighborNode := NewNode(neighborCoord)
			neighborNode.Parent = &currentNode
			neighborNode.ScoreG = currentNode.ScoreG + 1
			neighborNode.ScoreH = GetDistance(&currentNode.Coord, &neighborNode.Coord)
			neighborNode.ScoreF = neighborNode.ScoreG + neighborNode.ScoreH
			openSet.Push(neighborNode)
		}
	}

	return []Coord{}
}
