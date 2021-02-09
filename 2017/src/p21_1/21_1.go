package p21_1

import (
	"au"
	"fmt"
	"strings"
)

type EnhancementRule struct {
	input  *Grid
	output *Grid
}

func (e *EnhancementRule) Matches(grid *Grid, x, y int) bool {
	g := e.input

	for flip := 0; flip <= 2; flip++ {
		for rotate := 0; rotate <= 4; rotate++ {
			if g.Matches(grid, x, y) {
				return true
			}

			g = g.Rotated()
		}

		g = g.Flipped()
	}

	return false
}

func (e *EnhancementRule) Apply(grid *Grid, x, y int) {
	for dx := 0; dx < e.output.Size(); dx++ {
		for dy := 0; dy < e.output.Size(); dy++ {
			grid.data[x+dx][y+dy] = e.output.data[dx][dy]
		}
	}
}

type Grid struct {
	data [][]string
}

func NewGrid(size int) *Grid {
	grid := new(Grid)

	grid.data = make([][]string, size)
	for i := 0; i < size; i++ {
		grid.data[i] = make([]string, size)
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			grid.data[x][y] = "O"
		}
	}

	return grid
}

func LoadGrid(input string) *Grid {
	rows := strings.Split(input, "/")

	gridSize := len(rows)
	grid := NewGrid(gridSize)

	for y, row := range rows {
		for x := range row {
			grid.data[x][y] = string(row[x])
		}
	}

	return grid
}

func (g *Grid) Matches(other *Grid, offsetX, offsetY int) bool {
	for x := 0; x < g.Size(); x++ {
		for y := 0; y < g.Size(); y++ {
			if other.data[x+offsetX][y+offsetY] != g.data[x][y] {
				return false
			}
		}
	}

	return true
}

func (g *Grid) Flipped() *Grid {
	flipped := NewGrid(g.Size())

	for x := 0; x < g.Size(); x++ {
		for y := 0; y < g.Size(); y++ {
			flipped.data[x][g.Size()-y-1] = g.data[x][y]
		}
	}

	return flipped
}

func (g *Grid) Rotated() *Grid {
	rotated := NewGrid(g.Size())

	for x := 0; x < g.Size(); x++ {
		for y := 0; y < g.Size(); y++ {
			rotated.data[x][y] = g.data[y][g.Size()-x-1]
		}
	}

	return rotated
}

func (g *Grid) Size() int {
	return len(g.data)
}

func (g *Grid) Print() {
	for y := 0; y < g.Size(); y++ {
		for x := 0; x < g.Size(); x++ {
			fmt.Printf("%v", g.data[x][y])
		}
		fmt.Println()
	}
}

func load(input []string) []EnhancementRule {
	rules := []EnhancementRule{}

	for _, line := range input {
		parts := strings.Split(line, " => ")
		inputData := parts[0]
		outoutData := parts[1]

		rule := EnhancementRule{}
		rule.input = LoadGrid(inputData)
		rule.output = LoadGrid(outoutData)

		rules = append(rules, rule)
	}

	return rules
}

func enhance(inputGrid *Grid, rules []EnhancementRule) *Grid {
	var inputSquareSize int
	var outputSquareSize int

	if inputGrid.Size()%2 == 0 {
		inputSquareSize = 2
		outputSquareSize = 3
	} else {
		inputSquareSize = 3
		outputSquareSize = 4
	}

	squareCount := inputGrid.Size() / inputSquareSize

	outputGridSize := squareCount * outputSquareSize
	outputGrid := NewGrid(outputGridSize)

	for x := 0; x < squareCount; x++ {
		inputX := x * inputSquareSize
		outputX := x * outputSquareSize

		for y := 0; y < squareCount; y++ {
			inputY := y * inputSquareSize
			outputY := y * outputSquareSize

			for _, rule := range rules {
				if rule.output.Size() == outputSquareSize {
					if rule.Matches(inputGrid, inputX, inputY) {
						rule.Apply(outputGrid, outputX, outputY)
					}
				}
			}
		}
	}

	return outputGrid
}

func Solve() {
	input := au.ReadInputAsStringArray("21")

	rules := load(input)

	grid := LoadGrid(".#./..#/###")

	for i := 0; i < 5; i++ {
		grid = enhance(grid, rules)
	}

	onCount := 0
	for y := 0; y < grid.Size(); y++ {
		for x := 0; x < grid.Size(); x++ {
			if grid.data[x][y] == "#" {
				onCount++
			}
		}
	}
	fmt.Println(onCount)
}
