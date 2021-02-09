package main

import (
	"au"
	"fmt"
)

func testInputs() []string {
	return []string {
		".#.#...|#.",
		".....#|##|",
		".|..|...#.",
		"..|#.....#",
		"#.#|||#|#|",
		"...#.||...",
		".|....|...",
		"||...#|.#|",
		"|.||||..|.",
		"...#.|..|.",
	}
}

func parseInputs(inputs []string) *au.Screen {
	width := len(inputs[0])
	height := len(inputs)

	screen := au.NewScreen(width, height)

	for y, input := range inputs {
		for x, value := range input {
			screen.SetPixel(x, y, byte(value))
		}
	}

	return screen
}

func getAdjecentItems(screen *au.Screen, x int, y int) []byte {
	items := []byte{}

	for xx := x - 1; xx <= x + 1; xx++ {
		if xx < 0 || xx >= screen.Width {
			continue
		}

		for yy := y - 1; yy <= y + 1; yy++ {
			if yy < 0 || yy >= screen.Height {
				continue
			}

			if xx != x || yy != y {
				items = append(items, screen.GetPixel(xx, yy))
			}
		}
	}

	return items
}

func getAdjacentItemCount(screen *au.Screen, x int, y int, kind byte) int {
	count := 0

	for _,adjacentItem := range getAdjecentItems(screen, x, y) {
		if adjacentItem == kind {
			count++
		}
	}

	return count
}

func hasThreeOrMoreAdjacentTrees(screen *au.Screen, x int, y int) bool {
	return getAdjacentItemCount(screen, x, y, '|') >= 3
}

func hasThreeOrMoreAdjacentLumberyards(screen *au.Screen, x int, y int) bool {
	return getAdjacentItemCount(screen, x, y, '#') >= 3
}

func hasAdjacentTreesAndLumberyard(screen *au.Screen, x int, y int) bool {
	return getAdjacentItemCount(screen, x, y, '|') >= 1 && getAdjacentItemCount(screen, x, y, '#') >= 1
}

func iterate(screen *au.Screen) {
	reference := screen.Clone()

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			acre := reference.GetPixel(x, y)

			if acre == '.' {
				if hasThreeOrMoreAdjacentTrees(reference, x, y) {
					screen.SetPixel(x, y, '|')
				}
			} else if acre == '|' {
				if hasThreeOrMoreAdjacentLumberyards(reference, x, y) {
					screen.SetPixel(x, y, '#')
				}
			} else if acre == '#' {
				if !hasAdjacentTreesAndLumberyard(reference, x, y) {
					screen.SetPixel(x, y, '.')
				}

			}
		}
	}
}

func getTotalResourceValue(screen *au.Screen) int {
	woodAcreCount := 0
	lumberyardCount := 0

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			acre := screen.GetPixel(x, y)

			if acre == '|' {
				woodAcreCount++
			} else if acre == '#' {
				lumberyardCount++
			}
		}
	}

	return woodAcreCount * lumberyardCount
}

func main() {
	inputs := au.ReadInputAsStringArray("18")
	// inputs = testInputs()

	screen := parseInputs(inputs)
	//screen.Print()

	for i := 0; i < 10; i++ {
		iterate(screen)
	}

	//screen.Print()

	fmt.Println(getTotalResourceValue(screen))
}