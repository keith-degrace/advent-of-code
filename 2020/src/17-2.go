package main

import (
	"au"
	"fmt"
	"math"
)

type Puzzle17_2 struct {
}

type Coord17_2 struct {
	x, y, z, w int
}

func (c *Coord17_2) getNeighbors() []Coord17_2 {
	neighbors := []Coord17_2{}

	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			for dz := -1; dz < 2; dz++ {
				for dw := -1; dw < 2; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}

					neighbors = append(neighbors, Coord17_2{c.x + dx, c.y + dy, c.z + dz, c.w + dw})
				}
			}
		}
	}

	return neighbors
}

type Grid17_2 struct {
	data map[Coord17_2]bool
}

func NewGrid17_2() Grid17_2 {
	grid := Grid17_2{}
	grid.data = make(map[Coord17_2]bool)
	return grid
}

func (g *Grid17_2) GetMin() Coord17_2 {
	min := Coord17_2{math.MaxInt16, math.MaxInt16, math.MaxInt16, math.MaxInt16}

	for coord := range g.data {
		min.x = au.MinInt(min.x, coord.x)
		min.y = au.MinInt(min.y, coord.y)
		min.z = au.MinInt(min.z, coord.z)
		min.w = au.MinInt(min.w, coord.w)
	}

	return min
}

func (g *Grid17_2) GetMax() Coord17_2 {
	max := Coord17_2{math.MinInt16, math.MinInt16, math.MinInt16, math.MinInt16}

	for coord := range g.data {
		max.x = au.MaxInt(max.x, coord.x)
		max.y = au.MaxInt(max.y, coord.y)
		max.z = au.MaxInt(max.z, coord.z)
		max.w = au.MaxInt(max.w, coord.w)
	}

	return max
}

func (g *Grid17_2) IsActive(coord Coord17_2) bool {
	return g.data[coord] == true
}

func (g *Grid17_2) Set(coord Coord17_2, active bool) {
	g.data[coord] = active
}

func (p *Puzzle17_2) getActiveNeighborCount(grid Grid17_2, coord Coord17_2) int {
	count := 0

	for _, neighbor := range coord.getNeighbors() {
		if grid.IsActive(neighbor) {
			count++
		}
	}

	return count
}

func (p *Puzzle17_2) cycle(grid Grid17_2) Grid17_2 {
	newGrid := NewGrid17_2()

	min := grid.GetMin()
	max := grid.GetMax()

	for x := min.x - 2; x < max.x+2; x++ {
		for y := min.y - 2; y < max.y+2; y++ {
			for z := min.z - 2; z < max.z+2; z++ {
				for w := min.w - 2; w < max.w+2; w++ {
					coord := Coord17_2{x, y, z, w}
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
	}

	return newGrid
}

func (p Puzzle17_2) run() {
	input := au.ReadInputAsStringArray("17")

	grid := NewGrid17_2()

	for x, line := range input {
		for y := range line {
			grid.Set(Coord17_2{x, y, 0, 0}, line[y] == '#')
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
