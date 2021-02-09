package main

import (
	"au"
	"fmt"
	"math"
	"sort"
	"strings"
)

func testInputs() []string {
	return []string {
		"depth: 510",
		"target: 10,10",
	}
}

func parseInputs(inputs []string) (int, au.Coord) {
	depth := au.ToNumber(inputs[0][6:])	
	targetX := au.ToNumber(strings.Split(inputs[1][7:], ",")[0])	
	targetY := au.ToNumber(strings.Split(inputs[1][8:], ",")[1])	
	return depth, au.NewCoord(targetX, targetY)
}

const Rocky = "Rocky"
const Wet = "Wet"
const Narrow = "Narrow"

const Nothing = 0
const Torch = 1
const ClimbingGear = 2

type Regions struct {
	caveDepth int
	target au.Coord

	geoIndexCache map[string] int
	erosionLevelCache map[string] int
	regionTypeCache map[string] string
}

func (this *Regions) calculateGeoIndex(coord au.Coord) int {
	if coord.X() == 0 && coord.Y() == 0 {
		return 0
	}

	if coord.Equals(this.target) {
		return 0
	}

	if coord.Y() == 0 {
		return coord.X() * 16807
	}

	if coord.X() == 0 {
		return coord.Y() * 48271
	}

	return this.getErosionLevel(coord.ShiftX(-1)) * this.getErosionLevel(coord.ShiftY(-1))
}

func (this *Regions) getGeoIndex(coord au.Coord) int {
	value, ok := this.geoIndexCache[coord.Key()]
	if !ok {
		value = this.calculateGeoIndex(coord)
		this.geoIndexCache[coord.Key()] = value
	}

	return value
}

func (this *Regions) getErosionLevel(coord au.Coord) int {
	value, ok := this.erosionLevelCache[coord.Key()]
	if !ok {
		value = (this.getGeoIndex(coord) + this.caveDepth) % 20183
		this.erosionLevelCache[coord.Key()] = value
	}

	return value
}

func (this *Regions) getRegionType(coord au.Coord) string {
	value, ok := this.regionTypeCache[coord.Key()]
	if !ok {
		switch this.getErosionLevel(coord) % 3 {
			case 0: value = Rocky
			case 1: value = Wet
			case 2: value = Narrow
			default: panic("What?!?")
		}

		this.regionTypeCache[coord.Key()] = value
	}

	if len(value) == 0 {
		panic("What?!?")
	}
	return value
}

func printCave(regions *Regions, path []au.Coord) {
	screen := au.NewDynamicScreen()

	// for x := 0; x <= regions.target.X(); x++ {
	// 	for y := 0; y <= regions.target.Y(); y++ {
	// 		switch regions.getRegionType(au.NewCoord(x, y)) {
	// 		case Rocky: screen.SetPixel(x, y, '.')
	// 		case Wet: screen.SetPixel(x, y, '=')
	// 		case Narrow: screen.SetPixel(x, y, '|')
	// 		}
	// 	}
	// }

	for _,coord := range path {
		switch regions.getRegionType(coord) {
			case Rocky: screen.SetPixel(coord.X(), coord.Y(), '.')
			case Wet: screen.SetPixel(coord.X(), coord.Y(), '=')
			case Narrow: screen.SetPixel(coord.X(), coord.Y(), '|')
		}
	}

	screen.Print()
}

type NodeKey [3]int

func getNodeKey(coord *au.Coord, tool int) NodeKey {
	return NodeKey { coord.X(), coord.Y(), tool }
}

type Node struct {
	key NodeKey
	coord au.Coord
	regionType string
	tool int
	parent *Node
	gScore float64
	fScore float64
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

func newNode(coord au.Coord, tool int, regions *Regions) Node {
	return Node {
		key: getNodeKey(&coord, tool),
		coord: coord,
		regionType: regions.getRegionType(coord),
		tool: tool,
		gScore: 0,
		fScore: 0,
	}
}

type OpenSet struct {
	nodes []Node
}

func newOpenSet() OpenSet {
	return OpenSet {
		make([]Node, 0),
	}
}

func (this *OpenSet) push(node Node) {
	this.nodes = append(this.nodes, node)
	sort.Sort(byNode(this.nodes))
}

func (this *OpenSet) pop() Node {
	node := this.nodes[0]
	this.nodes = this.nodes[1:]
	return node
}

func (this *OpenSet) isEmpty() bool {
	return len(this.nodes) == 0
}

type ClosedSet map [NodeKey] bool

func addToClosedSet(closedSet *ClosedSet, key NodeKey) {
	(*closedSet)[key] = true
}

func isInClosedSet(closedSet *ClosedSet, key NodeKey) bool {
	_, ok := (*closedSet)[key]
	return ok
}

func getDistance(coord1 *au.Coord, coord2 *au.Coord) float64 {
	dx := math.Abs(float64(coord2.X() - coord1.X()))
	dy := math.Abs(float64(coord2.Y() - coord1.Y()))

	return dx + dy
}

func buildPath(endNode *Node, regions *Regions) []au.Coord {
	path := []au.Coord {}
	nodes := []*Node {}

	for current := endNode; current != nil; current = current.parent {
		path = append(path, current.coord)
		nodes = append(nodes, current)
	}

	path = path[:len(path) - 1]

	for i := len(path)/2-1; i >= 0; i-- {
		opp := len(path)-1-i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}

func getAllowedTools(coord au.Coord, regions *Regions) []int {
	switch regions.getRegionType(coord) {
	case Rocky:
		return []int {ClimbingGear, Torch}
	case Wet:
		return []int {ClimbingGear, Nothing}
	case Narrow:
		return []int {Torch, Nothing}
	}

	panic("impossible")
}

func isToolAllowed(coord au.Coord, regions *Regions, tool int) bool {
	for _,allowedTool := range getAllowedTools(coord, regions) {
		if allowedTool == tool {
			return true
		}
	}

	return false
}

func getNeighbors(coord au.Coord, regions *Regions) []au.Coord {
	neighbors := []au.Coord {}

	neighbors = append(neighbors, coord.ShiftX(+1))
	neighbors = append(neighbors, coord.ShiftY(+1))

	if coord.X() > 0 {
		neighbors = append(neighbors, coord.ShiftX(-1))
	}

	if coord.Y() > 0 {
		neighbors = append(neighbors, coord.ShiftY(-1))
	}

	return neighbors
}

func GetShortestPath(regions *Regions) ([]au.Coord, int) {
	openSet := newOpenSet()
	closedSet := make(ClosedSet)

	targetCoord := au.NewCoord(10, 10)
	targetTool := Torch

	gScores := map[NodeKey] float64{}

	startNode := newNode(au.NewCoord(0, 0), Torch, regions)
	gScores[startNode.key] = startNode.gScore

	openSet.push(startNode)

	for !openSet.isEmpty() {
		currentNode := openSet.pop()

		// if len(openSet.nodes) % 1000 == 0 {
		// 	fmt.Println(currentNode)
		// 	fmt.Println(len(openSet.nodes))
		// }

		addToClosedSet(&closedSet, currentNode.key)

		if currentNode.coord.Equals(regions.target) && currentNode.tool == targetTool {
			return buildPath(&currentNode, regions), int(currentNode.gScore)
			fmt.Println(currentNode.gScore)
			continue
		}

		for _,neighborCoord := range getNeighbors(currentNode.coord, regions) {
			if !isToolAllowed(neighborCoord, regions, currentNode.tool) {
				continue
			}

			nodeKey := getNodeKey(&neighborCoord, currentNode.tool)

			if isInClosedSet(&closedSet, nodeKey) {
				continue
			}

			gScore := currentNode.gScore + 1

			existingScore, ok := gScores[nodeKey]
			if ok && gScore >= existingScore {
				continue
			}

			gScores[nodeKey] = gScore

			newNode := newNode(neighborCoord, currentNode.tool, regions)
			newNode.parent = &currentNode
			newNode.gScore = gScore
			newNode.fScore = newNode.gScore + getDistance(&currentNode.coord, &targetCoord)

			openSet.push(newNode)
		}

		for _,newNodeTool := range getAllowedTools(currentNode.coord, regions) {
			if newNodeTool == currentNode.tool {
				continue
			}

			nodeKey := getNodeKey(&currentNode.coord, newNodeTool)

			if isInClosedSet(&closedSet, nodeKey) {
				continue
			}

			gScore := currentNode.gScore + 7

			existingScore, ok := gScores[nodeKey]
			if ok && gScore >= existingScore {
				continue
			}

			gScores[nodeKey] = gScore

			newNode := newNode(currentNode.coord, newNodeTool, regions)
			newNode.parent = &currentNode
			newNode.gScore = gScore
			newNode.fScore = newNode.gScore + getDistance(&currentNode.coord, &targetCoord)

			openSet.push(newNode)
		}
	}

	return []au.Coord{}, -1
}

func main() {
	inputs := au.ReadInputAsStringArray("22")
	// inputs  = testInputs()

	caveDepth, target := parseInputs(inputs)

	regions := Regions {
		caveDepth: caveDepth,
		target: target,
		geoIndexCache: map[string] int{},
		erosionLevelCache: map[string] int{},
		regionTypeCache: map[string] string{},
	}

	path, cost := GetShortestPath(&regions)
	printCave(&regions, path)
	fmt.Println(cost)
}