package p22_2

import (
	"au"
	"fmt"
	"regexp"
	"strconv"
)

func testInput() []string {
	return []string{
		"Filesystem            Size  Used  Avail  Use%",
		"/dev/grid/node-x0-y0   10T    8T     2T   80%",
		"/dev/grid/node-x0-y1   11T    6T     5T   54%",
		"/dev/grid/node-x0-y2   32T   28T     4T   87%",
		"/dev/grid/node-x1-y0    9T    7T     2T   77%",
		"/dev/grid/node-x1-y1    8T    0T     8T    0%",
		"/dev/grid/node-x1-y2   11T    7T     4T   63%",
		"/dev/grid/node-x2-y0   10T    6T     4T   60%",
		"/dev/grid/node-x2-y1    9T    8T     1T   88%",
		"/dev/grid/node-x2-y2    9T    6T     3T   66%",
	}
}

type Node struct {
	coord au.Coord
	size  int
	avail int
}

func (n Node) Used() int {
	return n.size - n.avail
}

type Grid struct {
	width     int
	height    int
	nodes     map[au.Coord]Node
	transfers int
}

func NewGrid() Grid {
	return Grid{0, 0, make(map[au.Coord]Node), 0}
}

func (g *Grid) Clone() Grid {
	copy := NewGrid()

	for coord, node := range g.nodes {
		copy.SetNode(coord, node)
	}

	return copy
}

func (g *Grid) SetNode(coord au.Coord, node Node) {
	g.nodes[coord] = node
	g.width = au.MaxInt(coord.X()+1, g.width)
	g.height = au.MaxInt(coord.Y()+1, g.height)
}

func (g *Grid) GetNode(coord au.Coord) Node {
	return g.nodes[coord]
}

func (g *Grid) GetNeighbors(coord au.Coord) []Node {
	var neighbors []Node

	// Top
	if coord.Y() > 0 {
		neighbors = append(neighbors, g.GetNode(coord.ShiftY(-1)))
	}

	// Left
	if coord.X() > 0 {
		neighbors = append(neighbors, g.GetNode(coord.ShiftX(-1)))
	}

	// Bottom
	if coord.Y() < g.height-1 {
		neighbors = append(neighbors, g.GetNode(coord.ShiftY(1)))
	}

	// Right
	if coord.X() < g.width-1 {
		neighbors = append(neighbors, g.GetNode(coord.ShiftX(1)))
	}

	return neighbors
}

func (g *Grid) Transfer(from au.Coord, to au.Coord) {
	fromNode := g.GetNode(from)
	toNode := g.GetNode(to)

	//fmt.Printf("Moving %v from (%v) to (%v)\n", fromNode.Used(), from, to)

	toNode.avail -= fromNode.Used()
	g.SetNode(to, toNode)

	fromNode.avail = fromNode.size
	g.SetNode(from, fromNode)

	au.Assert(toNode.Used() <= toNode.size)
	au.Assert(fromNode.Used() <= fromNode.size)

	g.transfers++
}

func (g Grid) Print() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			node := g.GetNode(au.NewCoord(x, y))
			fmt.Printf("%8s", fmt.Sprintf("%v/%v", node.size-node.avail, node.size))
		}

		fmt.Println()
	}

	fmt.Println()
}

func parse(input []string) Grid {
	grid := NewGrid()

	re := regexp.MustCompile("x([0-9]+)-y([0-9]+) *([0-9]+)T *([0-9]+)T *([0-9]+)T *([0-9]+)%")
	for _, line := range input {
		m := re.FindStringSubmatch(line)
		if m != nil {
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])
			size, _ := strconv.Atoi(m[3])
			avail, _ := strconv.Atoi(m[5])

			coord := au.NewCoord(x, y)

			grid.SetNode(coord, Node{coord, size, avail})
		}
	}

	return grid
}

func Solve() {
	input := au.ReadInputAsStringArray("22")
	//input = testInput()

	grid := parse(input)

	dataLocation := au.NewCoord(grid.width-1, 0)
	dataDestination := au.NewCoord(0, 0)
	emptyNode := au.NewCoord(3, 20)

	//
	dataPath := au.GetShortestPath(dataLocation, dataDestination, func(coord au.Coord) []au.Coord {
		var neighbors []au.Coord

		node := grid.GetNode(coord)

		for _, neighborNode := range grid.GetNeighbors(coord) {
			if node.Used() <= neighborNode.size {
				neighbors = append(neighbors, neighborNode.coord)
			}
		}

		return neighbors
	})

	for _, dataPathEntry := range dataPath {

		path := au.GetShortestPath(emptyNode, dataPathEntry, func(coord au.Coord) []au.Coord {
			var neighbors []au.Coord

			node := grid.GetNode(coord)

			for _, neighborNode := range grid.GetNeighbors(coord) {
				if dataLocation.Equals(neighborNode.coord) {
					continue
				}

				if node.Used() <= neighborNode.size {
					neighbors = append(neighbors, neighborNode.coord)
				}
			}

			return neighbors
		})

		// Move the empty node to the end of our path.
		grid.Transfer(path[0], emptyNode)
		for i := 1; i < len(path); i++ {
			grid.Transfer(path[i], path[i-1])
		}
		emptyNode = path[len(path)-1]

		// Move the data to the new empty node
		grid.Transfer(dataLocation, emptyNode)
		dataLocation, emptyNode = emptyNode, dataLocation
	}

	fmt.Println(grid.transfers)
}
