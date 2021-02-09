package main

import (
	"au"
	"fmt"
)

type Puzzle11_1 struct {
}

func (p Puzzle11_1) load(input []string) *au.Screen {

	screen := au.NewScreen(len(input[0]), len(input))

	for y, line := range input {
		for x := 0; x < len(line); x++ {
			screen.SetPixel(x, y, line[x])
		}
	}

	return screen
}

func (p Puzzle11_1) isEmpty(screen *au.Screen, x int, y int) bool {
	if x < 0 || x >= screen.Width {
		return true
	}

	if y < 0 || y >= screen.Height {
		return true
	}

	return screen.GetPixel(x, y) != '#'
}

func (p Puzzle11_1) run() {
	input := au.ReadInputAsStringArray("11")

	screen := p.load(input)
	nextScreen := screen.Clone()

	for {
		for y := 0; y < screen.Height; y++ {
			for x := 0; x < screen.Width; x++ {

				state := screen.GetPixel(x, y)

				if state == 'L' {
					if p.isEmpty(screen, x-1, y-1) &&
						p.isEmpty(screen, x-1, y) &&
						p.isEmpty(screen, x-1, y+1) &&
						p.isEmpty(screen, x+1, y-1) &&
						p.isEmpty(screen, x+1, y) &&
						p.isEmpty(screen, x+1, y+1) &&
						p.isEmpty(screen, x, y-1) &&
						p.isEmpty(screen, x, y+1) {
						nextScreen.SetPixel(x, y, '#')
					}
				} else if state == '#' {
					count := 0

					if !p.isEmpty(screen, x-1, y-1) {
						count++
					}
					if !p.isEmpty(screen, x-1, y) {
						count++
					}
					if !p.isEmpty(screen, x-1, y+1) {
						count++
					}
					if !p.isEmpty(screen, x+1, y-1) {
						count++
					}
					if count < 4 && !p.isEmpty(screen, x+1, y) {
						count++
					}
					if count < 4 && !p.isEmpty(screen, x+1, y+1) {
						count++
					}
					if count < 4 && !p.isEmpty(screen, x, y-1) {
						count++
					}
					if count < 4 && !p.isEmpty(screen, x, y+1) {
						count++
					}

					if count >= 4 {
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
