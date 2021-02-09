package main

import (
	"au"
	"fmt"
	"math"
)

type Puzzle17_1 struct {
}

type Coord17_1 struct {
	x, y, z int
}

func (c *Coord17_1) getNeighbors() []Coord17_1 {
	neighbors := []Coord17_1{}

	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			for dz := -1; dz < 2; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}

				neighbors = append(neighbors, Coord17_1{c.x + dx, c.y + dy, c.z + dz})
			}
		}
	}

	return neighbors
}

type Grid17_1 struct {
	data map[Coord17_1]bool
}

func NewGrid17_1() Grid17_1 {
	grid := Grid17_1{}
	grid.data = make(map[Coord17_1]bool)
	return grid
}

func (g *Grid17_1) GetMin() Coord17_1 {
	min := Coord17_1{math.MaxInt16, math.MaxInt16, math.MaxInt16}

	for coord := range g.data {
		min.x = au.MinInt(min.x, coord.x)
		min.y = au.MinInt(min.y, coord.y)
		min.z = au.MinInt(min.z, coord.z)
	}

	return min
}

func (g *Grid17_1) GetMax() Coord17_1 {
	max := Coord17_1{math.MinInt16, math.MinInt16, math.MinInt16}

	for coord := range g.data {
		max.x = au.MaxInt(max.x, coord.x)
		max.y = au.MaxInt(max.y, coord.y)
		max.z = au.MaxInt(max.z, coord.z)
	}

	return max
}

func (g *Grid17_1) IsActive(coord Coord17_1) bool {
	return g.data[coord] == true
}

func (g *Grid17_1) Set(coord Coord17_1, active bool) {
	g.data[coord] = active
}

func (p *Puzzle17_1) getActiveNeighborCount(grid Grid17_1, coord Coord17_1) int {
	count := 0

	for _, neighbor := range coord.getNeighbors() {
		if grid.IsActive(neighbor) {
			count++
		}
	}

	return count
}

func (p *Puzzle17_1) cycle(grid Grid17_1) Grid17_1 {
	newGrid := NewGrid17_1()

	min := grid.GetMin()
	max := grid.GetMax()

	for x := min.x - 2; x < max.x+2; x++ {
		for y := min.y - 2; y < max.y+2; y++ {
			for z := min.z - 2; z < max.z+2; z++ {
				coord := Coord17_1{x, y, z}
				activeNeighborCount := p.getActiveNeighborCount(grid, coord)
				active := grid.IsActive(coord)

				if active == true {
					newGrid.Set(coord, activeNeighborCount == 2 || activeNeighborCount == 3)
				} else {
					newGrid.Set(coord, activeNeighborCount == 3)
				}
			}
		}
	}

	return newGrid
}

func (p Puzzle17_1) run() {
	input := au.ReadInputAsStringArray("17")

	grid := NewGrid17_1()

	for x, line := range input {
		for y := range line {
			grid.Set(Coord17_1{x, y, 0}, line[y] == '#')
		}
	}

	for i := 0; i < 6; i++ {
		grid = p.cycle(grid)
	}

	activeCount := 0
	for _, active := range grid.data {
		if active == true {
			activeCount++
		}
	}

	fmt.Println(activeCount)
}
