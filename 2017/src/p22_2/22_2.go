package p22_2

import (
	"au"
	"fmt"
)

const Clean = 0
const Weakened = 1
const Infected = 2
const Flagged = 3

type Grid struct {
	infectedState map[string]int
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

func (g *Grid) GetState(x, y int) int {
	state, ok := g.infectedState[g.getCoordKey(x, y)]
	if !ok {
		return Clean
	}

	return state
}

func (g *Grid) SetState(x, y, state int) {
	g.infectedState[g.getCoordKey(x, y)] = state
	g.updateMinMax(x, y)
}

func (g *Grid) Print() {
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			fmt.Printf("%v", g.GetState(x, y))
		}
		fmt.Println()
	}
}

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

func (c *Carrier) TurnAround() {
	c.dir = (c.dir + 2) % 4
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

	currentNodeState := grid.GetState(c.x, c.y)
	if currentNodeState == Clean {
		c.TurnLeft()
	} else if currentNodeState == Weakened {
		// No turn
	} else if currentNodeState == Infected {
		c.TurnRight()
	} else if currentNodeState == Flagged {
		c.TurnAround()
	}

	if currentNodeState == Clean {
		grid.SetState(c.x, c.y, Weakened)
	} else if currentNodeState == Weakened {
		grid.SetState(c.x, c.y, Infected)
		addedInfection = true
	} else if currentNodeState == Infected {
		grid.SetState(c.x, c.y, Flagged)
	} else if currentNodeState == Flagged {
		grid.SetState(c.x, c.y, Clean)
	}

	c.MoveForward()

	return addedInfection
}

func load(input []string) *Grid {
	grid := new(Grid)
	grid.infectedState = make(map[string]int)

	for y, row := range input {
		for x, state := range row {
			if state == '#' {
				grid.SetState(x, y, Infected)
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
	for i := 0; i < 10000000; i++ {
		addedInfection := carrier.Work(grid)
		if addedInfection {
			burstInfectionCount++
		}
	}

	// grid.Print()

	fmt.Println(burstInfectionCount)
}
