package p22_1

import (
	"au"
	"fmt"
)

const Up = 0
const Right = 1
const Down = 2
const Left = 3

type Carrier struct {
	x, y int
	dir  int
}

func (c *Carrier) TurnLeft() {
	if c.dir > 0 {
		c.dir--
	} else {
		c.dir = 3
	}
}

func (c *Carrier) TurnRight() {
	c.dir = (c.dir + 1) % 4
}

func (c *Carrier) MoveForward() {
	if c.dir == Up {
		c.y--
	} else if c.dir == Right {
		c.x++
	} else if c.dir == Down {
		c.y++
	} else if c.dir == Left {
		c.x--
	}
}

func (c *Carrier) Work(grid *Grid) bool {
	addedInfection := false

	if grid.IsInfected(c.x, c.y) {
		c.TurnRight()
	} else {
		c.TurnLeft()
	}

	if grid.IsInfected(c.x, c.y) {
		grid.Clean(c.x, c.y)
	} else {
		grid.Infect(c.x, c.y)
		addedInfection = true
	}

	c.MoveForward()

	return addedInfection
}

type Grid struct {
	infectedState map[string]bool
	minX          int
	maxX          int
	minY          int
	maxY          int
}

func (g *Grid) getCoordKey(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func (g *Grid) updateMinMax(x, y int) {
	g.minX = au.MinInt(g.minX, x)
	g.maxX = au.MaxInt(g.maxX, x)

	g.minY = au.MinInt(g.minY, y)
	g.maxY = au.MaxInt(g.maxY, y)
}

func (g *Grid) IsInfected(x, y int) bool {
	ok, infected := g.infectedState[g.getCoordKey(x, y)]
	return ok && infected
}

func (g *Grid) Clean(x, y int) {
	delete(g.infectedState, g.getCoordKey(x, y))
	g.updateMinMax(x, y)
}

func (g *Grid) Infect(x, y int) {
	g.infectedState[g.getCoordKey(x, y)] = true
	g.updateMinMax(x, y)
}

func (g *Grid) Print() {
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if g.IsInfected(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func load(input []string) *Grid {
	grid := new(Grid)
	grid.infectedState = make(map[string]bool)

	for y, row := range input {
		for x, state := range row {
			if state == '#' {
				grid.Infect(x, y)
			} else {
				grid.Clean(x, y)
			}
		}
	}

	return grid
}

func Solve() {
	input := au.ReadInputAsStringArray("22")

	grid := load(input)

	startX := len(input) / 2
	startY := len(input) / 2
	carrier := Carrier{startX, startY, Up}

	burstInfectionCount := 0
	for i := 0; i < 10000; i++ {
		addedInfection := carrier.Work(grid)
		if addedInfection {
			burstInfectionCount++
		}
	}

	grid.Print()

	fmt.Println(burstInfectionCount)
}
