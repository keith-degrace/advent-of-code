package main

import (
	"au"
	"fmt"
)

type Puzzle24_2 struct {
}

func (p Puzzle24_2) parse(input []string) [][]string {
	pathes := [][]string{}

	for _, line := range input {
		path := []string{}

		current := 0
		for current < len(line) {
			currentpath := string(line[current])

			if (currentpath == "n" || currentpath == "s") && current < len(line)-1 {
				nextChar := line[current+1]

				if nextChar == 'e' || nextChar == 'w' {
					currentpath += string(line[current+1])
				}
			}

			path = append(path, currentpath)
			current += len(currentpath)
		}

		pathes = append(pathes, path)
	}

	return pathes
}

type Coord24_2 struct {
	x, y int
}

func (p Puzzle24_2) getNeighborSteps() []string {
	return []string{"e", "ne", "se", "w", "nw", "se"}
}

func (p Puzzle24_2) getNeighborCoord(coord Coord24_2, step string) Coord24_2 {
	if step == "e" {
		return Coord24_2{coord.x + 2, coord.y}
	} else if step == "ne" {
		return Coord24_2{coord.x + 1, coord.y - 2}
	} else if step == "se" {
		return Coord24_2{coord.x + 1, coord.y + 2}
	} else if step == "w" {
		return Coord24_2{coord.x - 2, coord.y}
	} else if step == "nw" {
		return Coord24_2{coord.x - 1, coord.y - 2}
	} else if step == "sw" {
		return Coord24_2{coord.x - 1, coord.y + 2}
	}

	fmt.Println(step)
	au.Assert(false)
	return coord
}

func (p Puzzle24_2) getNeighborCoords(coord Coord24_2) []Coord24_2 {
	neighbors := []Coord24_2{}

	for _, step := range p.getNeighborSteps() {
		neighbors = append(neighbors, p.getNeighborCoord(coord, step))
	}

	return neighbors
}

func (p Puzzle24_2) getInitialConfiguration(pathes [][]string) map[Coord24_2]bool {
	tileState := make(map[Coord24_2]bool)

	for _, path := range pathes {
		coord := Coord24_2{0, 0}

		// Apply the path
		for _, step := range path {
			coord = p.getNeighborCoord(coord, step)
		}

		currentState, found := tileState[coord]
		if !found {
			p.setTile(tileState, coord, false)
		} else {
			p.setTile(tileState, coord, !currentState)
		}
	}

	return tileState
}

func (p Puzzle24_2) setTile(tileState map[Coord24_2]bool, coord Coord24_2, value bool) {
	tileState[coord] = value

	// Make sure all neighbor tiles exists to simplify things later.
	for _, neighborCoord := range p.getNeighborCoords(coord) {
		if _, ok := tileState[neighborCoord]; !ok {
			tileState[neighborCoord] = true
		}
	}
}

func (p Puzzle24_2) isBlack(tileState map[Coord24_2]bool, coord Coord24_2) bool {
	isWhite, ok := tileState[coord]
	return ok && !isWhite
}

func (p Puzzle24_2) getBlackNeighborCount(tileState map[Coord24_2]bool, coord Coord24_2) int {
	count := 0

	// East
	if p.isBlack(tileState, p.getNeighborCoord(coord, "e")) {
		count++
	}

	// North East
	if p.isBlack(tileState, p.getNeighborCoord(coord, "ne")) {
		count++
	}

	// South East
	if p.isBlack(tileState, p.getNeighborCoord(coord, "se")) {
		count++
	}

	// West
	if p.isBlack(tileState, p.getNeighborCoord(coord, "w")) {
		count++
	}

	// North West
	if p.isBlack(tileState, p.getNeighborCoord(coord, "nw")) {
		count++
	}

	// South West
	if p.isBlack(tileState, p.getNeighborCoord(coord, "sw")) {
		count++
	}

	return count
}

func (p Puzzle24_2) run() {
	input := au.ReadInputAsStringArray("24")

	pathes := p.parse(input)

	tileState := p.getInitialConfiguration(pathes)

	for i := 0; i < 100; i++ {
		newTileState := make(map[Coord24_2]bool)
		for coord, isWhite := range tileState {

			blackNeighborCount := p.getBlackNeighborCount(tileState, coord)

			if !isWhite && (blackNeighborCount == 0 || blackNeighborCount > 2) {
				p.setTile(newTileState, coord, true)
			} else if isWhite && blackNeighborCount == 2 {
				p.setTile(newTileState, coord, false)
			} else {
				p.setTile(newTileState, coord, isWhite)
			}
		}

		tileState = newTileState
	}

	// Count black tiles
	blackCount := 0
	for _, state := range tileState {
		if !state {
			blackCount++
		}
	}

	fmt.Println(blackCount)
}
