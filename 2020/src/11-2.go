package main

import (
	"au"
	"fmt"
)

type Puzzle11_2 struct {
}

func (p Puzzle11_2) load(input []string) *au.Screen {

	screen := au.NewScreen(len(input[0]), len(input))

	for y, line := range input {
		for x := 0; x < len(line); x++ {
			screen.SetPixel(x, y, line[x])
		}
	}

	return screen
}

func (p Puzzle11_2) isEmpty(screen *au.Screen, x int, y int, xSlope int, ySlope int) bool {

	for {
		x += xSlope
		y += ySlope

		if x < 0 || y < 0 || x >= screen.Width || y >= screen.Height {
			break
		}

		if screen.GetPixel(x, y) == 'L' {
			return true
		}

		if screen.GetPixel(x, y) == '#' {
			return false
		}
	}

	return true
}

func (p Puzzle11_2) run() {
	input := au.ReadInputAsStringArray("11")

	screen := p.load(input)
	nextScreen := screen.Clone()

	for {
		for y := 0; y < screen.Height; y++ {
			for x := 0; x < screen.Width; x++ {

				state := screen.GetPixel(x, y)

				if state == 'L' {
					if p.isEmpty(screen, x, y, -1, -1) &&
						p.isEmpty(screen, x, y, -1, 0) &&
						p.isEmpty(screen, x, y, -1, 1) &&
						p.isEmpty(screen, x, y, 1, -1) &&
						p.isEmpty(screen, x, y, 1, 0) &&
						p.isEmpty(screen, x, y, 1, 1) &&
						p.isEmpty(screen, x, y, 0, -1) &&
						p.isEmpty(screen, x, y, 0, 1) {
						nextScreen.SetPixel(x, y, '#')
					}
				} else if state == '#' {
					count := 0

					if !p.isEmpty(screen, x, y, -1, -1) {
						count++
					}
					if !p.isEmpty(screen, x, y, -1, 0) {
						count++
					}
					if !p.isEmpty(screen, x, y, -1, 1) {
						count++
					}
					if !p.isEmpty(screen, x, y, 1, -1) {
						count++
					}
					if !p.isEmpty(screen, x, y, 1, 0) {
						count++
					}
					if count < 5 && !p.isEmpty(screen, x, y, 1, 1) {
						count++
					}
					if count < 5 && !p.isEmpty(screen, x, y, 0, -1) {
						count++
					}
					if count < 5 && !p.isEmpty(screen, x, y, 0, 1) {
						count++
					}

					if count >= 5 {
						nextScreen.SetPixel(x, y, 'L')
					}
				}
			}
		}

		if screen.Equals(nextScreen) {
			screen = nextScreen
			break
		}

		screen = nextScreen
		nextScreen = screen.Clone()
	}

	screen.Print()

	fmt.Println(screen.GetCount('#'))
}
