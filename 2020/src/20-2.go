package main

import (
	"au"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Puzzle20_2 struct {
}

func (p Puzzle20_2) parse(input []string) []*Tile20_2 {
	tiles := []*Tile20_2{}

	var currentTile *Tile20_2

	tileHeaderRegex := regexp.MustCompile("Tile ([0-9]*):")

	for _, line := range input {

		m := tileHeaderRegex.FindStringSubmatch(line)
		if len(m) > 0 {
			currentTile = new(Tile20_2)
			currentTile.id = m[1]
			tiles = append(tiles, currentTile)
		} else if len(line) != 0 {
			currentTile.data = append(currentTile.data, strings.Split(line, ""))
		}
	}

	return tiles
}

func (p Puzzle20_2) getOpposite(edge int) int {
	switch edge {
	case Top:
		return Bottom
	case Bottom:
		return Top
	case Left:
		return Right
	case Right:
		return Left
	}

	au.Assert(false) // Should never happen
	return 0
}

func rotate(data [][]string) [][]string {
	newData := make([][]string, len(data))

	for i := 0; i < len(data); i++ {
		newData[i] = make([]string, len(data[i]))
	}

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data); x++ {
			destX := len(data) - y - 1
			destY := x

			newData[destX][destY] = data[x][y]
		}
	}

	return newData
}

func flip(data [][]string) [][]string {
	newData := make([][]string, len(data))

	for i := 0; i < len(data); i++ {
		newData[i] = make([]string, len(data[i]))
	}

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data); x++ {
			destX := x
			destY := len(data) - y - 1

			newData[destX][destY] = data[x][y]
		}
	}

	return newData
}

type Tile20_2 struct {
	id   string
	data [][]string
}

func (t *Tile20_2) print() {
	for _, row := range t.data {
		fmt.Println(row)
	}
}

func (t *Tile20_2) rotate() {
	t.data = rotate(t.data)
}

func (t *Tile20_2) flip() {
	t.data = flip(t.data)
}

func (t *Tile20_2) getEdge(which int) string {
	var edge string

	switch which {
	case Top:
		edge = strings.Join(t.data[0], "")
		break

	case Bottom:
		edge = strings.Join(t.data[len(t.data)-1], "")
		break

	case Left:
		for _, row := range t.data {
			edge += string(row[0])
		}
		break

	case Right:
		for _, row := range t.data {
			edge += string(row[len(row)-1])
		}
		break
	}

	return edge
}

func (t *Tile20_2) hasEdge(edge string) bool {
	for i := 1; i <= 4; i++ {
		candidate := t.getEdge(i)
		if edge == candidate || edge == au.ReverseString(candidate) {
			return true
		}
	}

	return false
}

func (p Puzzle20_2) getTopLeftTile(tiles map[string]*Tile20_2) *Tile20_2 {
	// Find one of the four tiles that only has two neighbors

	for _, candidate := range tiles {

		neighborEdges := make(map[int]bool)

		// Look at the four edges and see if any tiles match them.
		for edge := 1; edge <= 4; edge++ {

			candidateEdge := candidate.getEdge(edge)

			for _, otherTile := range tiles {
				if candidate.id == otherTile.id {
					continue
				}

				if otherTile.hasEdge(candidateEdge) {
					neighborEdges[edge] = true
					break
				}
			}
		}

		if len(neighborEdges) == 2 {
			// Orient it so that the top left corner is actual in the top left.

			_, hasTop := neighborEdges[Top]
			_, hasBottom := neighborEdges[Bottom]
			_, hasLeft := neighborEdges[Left]
			_, hasRight := neighborEdges[Right]

			if hasLeft && hasBottom {
				candidate.rotate()
			}

			if hasLeft && hasTop {
				candidate.rotate()
				candidate.rotate()
			}

			if hasTop && hasRight {
				candidate.rotate()
				candidate.rotate()
				candidate.rotate()
			}

			return candidate
		}
	}

	au.Assert(false)
	return nil
}

func (p Puzzle20_2) getTileWithTopLeftNeighbors(tiles map[string]*Tile20_2, topTile, leftTile *Tile20_2) *Tile20_2 {
	if topTile == nil && leftTile == nil {
		return p.getTopLeftTile(tiles)
	}

	for _, candidate := range tiles {
		if topTile != nil && !candidate.hasEdge(topTile.getEdge(Bottom)) {
			continue
		}

		if leftTile != nil && !candidate.hasEdge(leftTile.getEdge(Right)) {
			continue
		}

		// We have a match, not make sure it's oriented properly (in a dumb way)
		for rotate := 1; rotate <= 4; rotate++ {
			for flip := 1; flip <= 2; flip++ {
				topMatch := topTile == nil || candidate.getEdge(Top) == topTile.getEdge(Bottom)
				leftMatch := leftTile == nil || candidate.getEdge(Left) == leftTile.getEdge(Right)

				if topMatch && leftMatch {
					return candidate
				}

				candidate.flip()
			}

			candidate.rotate()
		}
	}

	au.Assert(false)
	return nil
}

func (p Puzzle20_2) buildgrid(tiles []*Tile20_2) [][]*Tile20_2 {
	gridDimension := int(math.Sqrt(float64(len(tiles))))

	grid := make([][]*Tile20_2, gridDimension)
	for i := 0; i < gridDimension; i++ {
		grid[i] = make([]*Tile20_2, gridDimension)
	}

	// Start from the top left corner
	remainingTiles := make(map[string]*Tile20_2)
	for _, tile := range tiles {
		remainingTiles[tile.id] = tile
	}

	for x := 0; x < gridDimension; x++ {
		for y := 0; y < gridDimension; y++ {
			var leftNeighbor *Tile20_2
			if x > 0 {
				leftNeighbor = grid[x-1][y]
			}

			var topNeighbor *Tile20_2
			if y > 0 {
				topNeighbor = grid[x][y-1]
			}

			grid[x][y] = p.getTileWithTopLeftNeighbors(remainingTiles, topNeighbor, leftNeighbor)
			au.Assert(grid[x][y] != nil)

			delete(remainingTiles, grid[x][y].id)
		}
	}

	return grid
}

func stitchImage(grid [][]*Tile20_2) [][]string {
	image := [][]string{}

	for gridY := 0; gridY < len(grid); gridY++ {
		for dataY := 1; dataY < 9; dataY++ {

			imageRow := ""

			for gridX := 0; gridX < len(grid[gridY]); gridX++ {
				tile := grid[gridX][gridY]

				for dataX := 1; dataX < 9; dataX++ {
					imageRow += tile.data[dataY][dataX]
				}
			}

			image = append(image, strings.Split(imageRow, ""))
		}
	}

	return image
}

func (p Puzzle20_2) printImage(image [][]string) {
	for _, imageRow := range image {
		fmt.Println(imageRow)
	}
}

func (p Puzzle20_2) hasMonsterAt(image [][]string, x, y int) bool {
	monster := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	for dy, line := range monster {
		imageY := y + dy - 1
		if imageY < 0 || imageY >= len(image) {
			return false
		}

		for dx, char := range line {
			imageX := x + dx
			if imageX >= len(image[imageY]) {
				return false
			}

			if char == '#' {
				if image[imageY][imageX] != "#" {
					return false
				}
			}
		}
	}

	return true
}

func (p Puzzle20_2) getRoughness(image [][]string) int {
	totalHash := 0
	for _, imageRow := range image {
		for _, char := range imageRow {
			if char == "#" {
				totalHash++
			}
		}
	}

	totalMonsters := 0
	for doRotate := 1; doRotate <= 4; doRotate++ {
		for doFlip := 1; doFlip <= 2; doFlip++ {

			for y := 0; y < len(image); y++ {
				for x := 0; x < len(image[y]); x++ {
					if p.hasMonsterAt(image, x, y) {
						totalMonsters++
					}
				}
			}

			if totalMonsters > 0 {
				return totalHash - (totalMonsters * 15)
			}

			image = flip(image)
		}

		image = rotate(image)
	}

	au.Assert(false)
	return 0
}

func (p Puzzle20_2) run() {
	input := au.ReadInputAsStringArray("20")

	tiles := p.parse(input)

	grid := p.buildgrid(tiles)

	image := stitchImage(grid)

	fmt.Println(p.getRoughness(image))
}
