package main

import (
	"au"
	"fmt"
)

type Puzzle24_1 struct {
}

func (p Puzzle24_1) parse(input []string) [][]string {
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

func (p Puzzle24_1) getCoordinate(path []string) (int, int) {
	x := 0
	y := 0

	for _, step := range path {
		if step == "e" {
			x += 2
		} else if step == "ne" {
			x += 1
			y -= 2
		} else if step == "se" {
			x += 1
			y += 2
		} else if step == "w" {
			x -= 2
		} else if step == "nw" {
			x -= 1
			y -= 2
		} else if step == "sw" {
			x -= 1
			y += 2
		}
	}

	return x, y
}

func (p Puzzle24_1) run() {
	input := au.ReadInputAsStringArray("24")

	pathes := p.parse(input)

	tileState := make(map[string]bool) /* true is white */

	for _, path := range pathes {
		x, y := p.getCoordinate(path)

		coordKey := fmt.Sprintf("%v,%v", x, y)

		currentState, found := tileState[coordKey]
		if !found {
			tileState[coordKey] = false
		} else {
			tileState[coordKey] = !currentState
		}
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
