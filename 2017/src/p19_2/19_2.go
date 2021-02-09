package p19_2

import (
	"au"
	"fmt"
)

const West = 1
const South = 2
const East = 3
const North = 4

type Grid struct {
	data    []string
	current []int
}

func (g *Grid) Width() int {
	return len(g.data[0])
}

func (g *Grid) Height() int {
	return len(g.data)
}

func (g *Grid) Get(x, y int) byte {
	if x < 0 || y < 0 || x > g.Width()-1 || y > g.Height()-1 {
		return ' '
	}

	return g.data[y][x]
}

func (g *Grid) Print() {
	for y, row := range g.data {
		for x, char := range row {
			if x == g.current[0] && y == g.current[1] {
				fmt.Printf("X")
			} else {
				fmt.Printf("%v", string(char))
			}
		}
		fmt.Println()
	}
}

func findStart(grid *Grid) []int {
	for x := 0; x < grid.Width(); x++ {
		if grid.Get(x, 0) == '|' {
			return []int{x, 0}
		}
	}

	au.Assert(false)
	return nil
}

func advance(coord []int, dir int) []int {
	switch dir {
	case East:
		return []int{coord[0] + 1, coord[1]}
	case South:
		return []int{coord[0], coord[1] + 1}
	case West:
		return []int{coord[0] - 1, coord[1]}
	case North:
		return []int{coord[0], coord[1] - 1}
	default:
		au.Assert(false)
		return nil
	}
}

func Solve() {
	input := au.ReadInputAsStringArray("19")

	grid := new(Grid)
	grid.data = input

	grid.current = findStart(grid)
	dir := South

	steps := 0

	for {
		steps++

		grid.current = advance(grid.current, dir)

		current := grid.Get(grid.current[0], grid.current[1])
		if current == ' ' {
			break
		}

		if (current >= 'A' && current <= 'Z') || current == '|' || current == '-' {
			// Keep going
		} else if current == '+' {
			nextCoord := advance(grid.current, dir)
			next := grid.Get(nextCoord[0], nextCoord[1])

			// If the next move is not possible, we need to turn
			if next == ' ' {
				if dir == South || dir == North {
					eastValue := grid.Get(grid.current[0]+1, grid.current[1])
					if eastValue != ' ' && eastValue != '|' {
						dir = East
					} else {
						dir = West
					}
				} else {
					northValue := grid.Get(grid.current[0], grid.current[1]-1)
					if northValue != ' ' && northValue != '-' {
						dir = North
					} else {
						dir = South
					}
				}
			}
		} else {
			au.Assert(false)
		}
	}

	fmt.Println(steps)
}
