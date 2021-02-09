package p17_2

import (
	"au"
	"fmt"
	"math"
	"sort"
	"time"
)

type Node struct {
	Coord au.Coord
	Parent *Node
	ScoreG float64
	ScoreF float64
	ScoreH float64
	Path string
}

func NewNode(coord au.Coord) Node {
	return Node {
		Coord: coord,
		ScoreG: 0,
		ScoreF: 0,
	}
}

func (this *Node) GetKey() string {
	return this.Coord.Key() + this.Path
}

type ByNode []Node 

func (s ByNode) Len() int {
	return len(s)
}

func (s ByNode) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByNode) Less(i, j int) bool {
	return s[i].ScoreF > s[j].ScoreF
}

type OpenSet map [string] Node

func (this *OpenSet) Push(node Node) {
	(*this)[node.GetKey()] = node
}

func (this *OpenSet) Pop() Node {
	nodes := []Node {}
	for _, node := range *this {
		nodes = append(nodes, node)
	}

	sort.Sort(ByNode(nodes))

	delete(*this, nodes[0].GetKey())

	return nodes[0]
}

func (this *OpenSet) IsEmpty() bool {
	return len(*this) == 0
}

type ClosedSet map [string] bool

func (this *ClosedSet) Add(node *Node) {
	(*this)[node.GetKey()] = true
}

func (this *ClosedSet) Has(node *Node) bool {
	_, ok := (*this)[node.GetKey()]
	return ok
}

func GetDistance(coord1 *au.Coord, coord2 *au.Coord) float64 {
	dx := float64(coord2.X() - coord1.X())
	dy := float64(coord2.Y() - coord1.Y())

	return math.Pow(dx, 2) + math.Pow(dy, 2)
}

func BuildCoordPath(endNode *Node) []au.Coord {
	path := []au.Coord {}

	for current := endNode; current != nil; current = current.Parent {
		path = append(path, current.Coord)
	}

	for i := len(path)/2-1; i >= 0; i-- {
		opp := len(path)-1-i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}

func convertPath(path []au.Coord) string {
	stringPath := ""

	for i := 1; i < len(path); i++ {
		if path[i].X() > path[i-1].X() {
			stringPath += "R"
		} else if path[i].X() < path[i-1].X() {
			stringPath += "L"
		} else if path[i].Y() > path[i-1].Y() {
			stringPath += "D"
		} else if path[i].Y() < path[i-1].Y() {
			stringPath += "U"
		}
	}

	return stringPath
}

func isDoorOpen(char byte) bool {
	return char == 'b' || char == 'c' || char == 'd' || char == 'e' || char == 'f'
}

func getNeighbors(node *Node, passcode string) []au.Coord {
	neighbors := []au.Coord{}

	hash := au.GetMD5(passcode + node.Path)[:4]

	// Up
	if node.Coord.Y() > 0 && isDoorOpen(hash[0]) {
		neighbors = append(neighbors, node.Coord.ShiftY(-1))
	}

	// Down
	if node.Coord.Y() < 3 && isDoorOpen(hash[1]) {
		neighbors = append(neighbors, node.Coord.ShiftY(+1))
	}

	// Left
	if node.Coord.X() > 0 && isDoorOpen(hash[2]) {
		neighbors = append(neighbors, node.Coord.ShiftX(-1))
	}

	// Right
	if node.Coord.X() < 3 && isDoorOpen(hash[3]) {
		neighbors = append(neighbors, node.Coord.ShiftX(+1))
	}

	return neighbors
}

func BuildPath(endNode *Node) string {
	return convertPath(BuildCoordPath(endNode))
}

func GetLongestPathLength(from au.Coord, to au.Coord, passcode string) int {
	longestPathLength := 0

	openSet := make(OpenSet)
	closedSet := make(ClosedSet)

	startNode := NewNode(from)

	openSet.Push(startNode)

	for !openSet.IsEmpty() {
		currentNode := openSet.Pop()

		closedSet.Add(&currentNode)

		if currentNode.Coord.Equals(to) {
			if len(currentNode.Path) > longestPathLength {
				longestPathLength = len(currentNode.Path)
			}
		} else {
		
			for _,neighborCoord := range getNeighbors(&currentNode, passcode) {
				neighborNode := NewNode(neighborCoord)
				neighborNode.Parent = &currentNode
				neighborNode.ScoreG = currentNode.ScoreG + 1
				neighborNode.ScoreH = GetDistance(&currentNode.Coord, &neighborNode.Coord)
				neighborNode.ScoreF = neighborNode.ScoreG + neighborNode.ScoreH
				neighborNode.Path = BuildPath(&neighborNode)

				if closedSet.Has(&neighborNode) {
					continue
				}

				openSet.Push(neighborNode)
			}
		}
	}

	return longestPathLength
}

func Solve() {
	fmt.Println("Starting\n")
	startTime := time.Now()

	au.AssertIntsEqual(GetLongestPathLength(au.NewCoord(0, 0), au.NewCoord(3, 3), "ihgpwlah"), 370)
	au.AssertIntsEqual(GetLongestPathLength(au.NewCoord(0, 0), au.NewCoord(3, 3), "kglvqrro"), 492)
	au.AssertIntsEqual(GetLongestPathLength(au.NewCoord(0, 0), au.NewCoord(3, 3), "ulqzkmiv"), 830)
	
	fmt.Println(GetLongestPathLength(au.NewCoord(0, 0), au.NewCoord(3, 3), "dmypynyp"))

	fmt.Println("\nCompleted in", time.Now().Sub(startTime))
}