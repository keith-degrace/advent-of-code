package main

import (
	"au"
	"fmt"
	"regexp"
	"strings"
)

type Puzzle20_1 struct {
}

func (p Puzzle20_1) parse(input []string) []*Tile20_1 {
	tiles := []*Tile20_1{}

	var currentTile *Tile20_1

	tileHeaderRegex := regexp.MustCompile("Tile ([0-9]*):")

	for _, line := range input {

		m := tileHeaderRegex.FindStringSubmatch(line)
		if len(m) > 0 {
			currentTile = new(Tile20_1)
			currentTile.id = m[1]
			tiles = append(tiles, currentTile)
		} else if len(line) != 0 {
			currentTile.data = append(currentTile.data, strings.Split(line, ""))
		}
	}

	return tiles
}

const Top = 1
const Bottom = 2
const Left = 3
const Right = 4

func (p Puzzle20_1) getOpposite(edge int) int {
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

type Tile20_1 struct {
	id   string
	data [][]string
}

func (t *Tile20_1) print() {
	for _, row := range t.data {
		fmt.Println(row)
	}
}

func (t *Tile20_1) rotate() {
	newData := make([][]string, 10)

	for i := 0; i < 10; i++ {
		newData[i] = make([]string, 10)
	}

	for y := 0; y < len(t.data); y++ {
		for x := 0; x < len(t.data); x++ {
			destX := 10 - y - 1
			destY := x

			newData[destX][destY] = t.data[x][y]
		}
	}

	t.data = newData
}

func (t *Tile20_1) flip() {
	newData := make([][]string, 10)

	for i := 0; i < 10; i++ {
		newData[i] = make([]string, 10)
	}

	for y := 0; y < len(t.data); y++ {
		for x := 0; x < len(t.data); x++ {
			destX := x
			destY := 10 - y - 1

			newData[destX][destY] = t.data[x][y]
		}
	}

	t.data = newData
}

func (t *Tile20_1) getEdge(which int) string {
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

func (p Puzzle20_1) getNeighbor(edge int, tile *Tile20_1, tiles []*Tile20_1) *Tile20_1 {
	for _, otherTile := range tiles {
		if tile.id == otherTile.id {
			continue
		}

		for i := 0; i < 2; i++ {
			for i := 0; i < 4; i++ {
				if tile.getEdge(edge) == otherTile.getEdge(p.getOpposite(edge)) {
					return otherTile
				}

				otherTile.rotate()
			}
			otherTile.flip()
		}
	}

	return nil
}

func (p Puzzle20_1) run() {
	input := au.ReadInputAsStringArray("20")

	tiles := p.parse(input)

	result := 1

	for _, candidate := range tiles {

		neighborCount := 0

		for edge := 1; edge <= 4; edge++ {
			if p.getNeighbor(edge, candidate, tiles) != nil {
				neighborCount++
			}
		}

		if neighborCount == 2 {
			result *= au.ToNumber(candidate.id)
		}
	}

	fmt.Println(result)
}
