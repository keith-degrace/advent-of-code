package p13_2

import (
	"au"
	"fmt"
	"math"
	"sort"
)

func isWall(x int, y int, code int) bool {
	if x < 0 || y < 0 {
		return true
	}

	sum := (x*x + 3*x + 2*x*y + y + y*y) + code

	count := 0

	for sum != 0 {
		if sum%2 == 1 {
			count++
		}

		sum /= 2
	}

	return count%2 != 0
}

func print(code int, width int, height int) {
	screen := au.NewScreen(width, height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if isWall(x, y, code) {
				screen.SetPixel(x, y, '#')
			}
		}
	}

	screen.Print()
}

func getCoordKey(x int, y int) string {
	return fmt.Sprintf("%vx%v", x, y)
}

type Node struct {
	id     string
	x      int
	y      int
	parent *Node
	gScore float64
	fScore float64
	hScore float64
}

func newNode(x int, y int) Node {
	return Node{
		id:     getCoordKey(x, y),
		x:      x,
		y:      y,
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
	nodes map[string]Node
}

func newOpenSet() OpenSet {
	return OpenSet{
		make(map[string]Node),
	}
}

func (this *OpenSet) push(node Node) {
	this.nodes[node.id] = node
}

func (this *OpenSet) popSmallestF() Node {
	nodes := []Node{}
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

type ClosedSet map[string]bool

func addToClosedSet(closedSet ClosedSet, node Node) {
	closedSet[node.id] = true
}

func isInClosedSet(closedSet ClosedSet, x int, y int) bool {
	_, ok := closedSet[getCoordKey(x, y)]
	return ok
}

func getNeighbors(node Node, code int) [][]int {
	offsets := [][]int{
		[]int{0, -1},
		[]int{1, 0},
		[]int{0, 1},
		[]int{-1, 0},
	}

	neighbors := [][]int{}

	for _, offset := range offsets {
		xx := node.x + offset[0]
		yy := node.y + offset[1]

		if isWall(xx, yy, code) {
			continue
		}

		neighbors = append(neighbors, []int{xx, yy})
	}

	return neighbors
}

func getDistance(node1 Node, node2 Node) float64 {
	dx := float64(node2.x - node2.x)
	dy := float64(node2.y - node2.y)

	return math.Pow(dx, 2) + math.Pow(dy, 2)
}

func buildPath(endNode Node) [][]int {
	path := [][]int{}

	for current := &endNode; current != nil; current = current.parent {
		path = append(path, []int{current.x, current.y})
	}

	path = path[:len(path)-1]

	for i := len(path)/2 - 1; i >= 0; i-- {
		opp := len(path) - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}

func getShortestPath(x1 int, y1 int, x2 int, y2 int, code int) [][]int {
	openSet := newOpenSet()
	closedSet := make(ClosedSet)

	startNode := newNode(x1, y1)

	openSet.push(startNode)

	for !openSet.isEmpty() {
		currentNode := openSet.popSmallestF()

		addToClosedSet(closedSet, currentNode)

		if currentNode.x == x2 && currentNode.y == y2 {
			return buildPath(currentNode)
		}

		for _, neighbor := range getNeighbors(currentNode, code) {
			if isInClosedSet(closedSet, neighbor[0], neighbor[1]) {
				continue
			}

			neighborNode := newNode(neighbor[0], neighbor[1])
			neighborNode.parent = &currentNode
			neighborNode.gScore = currentNode.gScore + 1
			neighborNode.hScore = getDistance(currentNode, neighborNode)
			neighborNode.fScore = neighborNode.gScore + neighborNode.hScore
			openSet.push(neighborNode)
		}
	}

	return [][]int{}
}

func Solve() {
	input := 1364

	count := 1

	for x := 0; x < 51; x++ {
		for y := 0; y < 51; y++ {
			if isWall(x, y, input) {
				continue
			}

			path := getShortestPath(1, 1, x, y, input)
			if len(path) > 0 && len(path) <= 50 {
				count += 1
			}

		}
	}

	fmt.Println(count)
}
