package main

import (
	"au"
	"fmt"
	"regexp"
)

func testInputs() []string {
	return []string {
		"x=495, y=2..7",
		"y=7, x=495..501",
		"x=501, y=3..7",
		"x=498, y=2..4",
		"x=506, y=1..2",
		"x=498, y=10..13",
		"x=504, y=10..13",
		"y=13, x=498..504",
	}
}

type Vein struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func parseInputs(inputs []string) []Vein {
	veins := []Vein{}

	verticalRe := regexp.MustCompile("x=([0-9]+), y=([0-9]+)..([0-9]+)")
	horizontalRe := regexp.MustCompile("y=([0-9]+), x=([0-9]+)..([0-9]+)")

	for _,input := range inputs {
		verticalMatches := verticalRe.FindStringSubmatch(input)
		if len(verticalMatches) > 0 {
			vein := Vein {
				au.ToNumber(verticalMatches[1]),
				au.ToNumber(verticalMatches[2]),
				au.ToNumber(verticalMatches[1]),
				au.ToNumber(verticalMatches[3]),
			}
			veins = append(veins, vein)
		}

		horizontalMatches := horizontalRe.FindStringSubmatch(input)
		if len(horizontalMatches) > 0 {
			vein := Vein {
				au.ToNumber(horizontalMatches[2]),
				au.ToNumber(horizontalMatches[1]),
				au.ToNumber(horizontalMatches[3]),
				au.ToNumber(horizontalMatches[1]),
			}
			veins = append(veins, vein)
		}
	}

	return veins
}

func getBounds(veins []Vein) (int, int, int, int) {
	minX := 999999999
	minY := 999999999
	maxX := 0
	maxY := 0

	for _,vein := range veins {
		minX = au.MinInt(minX, vein.x1)
		minX = au.MinInt(minX, vein.x2)
		minY = au.MinInt(minY, vein.y1)
		minY = au.MinInt(minY, vein.y2)

		maxX = au.MaxInt(maxX, vein.x1)
		maxX = au.MaxInt(maxX, vein.x2)
		maxY = au.MaxInt(maxY, vein.y1)
		maxY = au.MaxInt(maxY, vein.y2)
	}

	return minX, minY, maxX, maxY
}

func createScreen(veins []Vein) *au.Screen {
	minX, minY, maxX, maxY := getBounds(veins)

	width := maxX - minX + 3
	height := maxY - minY + 2

	offsetX := -minX + 1
	offsetY := -minY + 1

	screen := au.NewScreen(width, height)

	screen.SetPixel(500 + offsetX, 0, '+')
	screen.SetPixel(500 + offsetX, 1, '|')

	for _, vein := range veins {
		for x := vein.x1; x <= vein.x2; x++ {
			screen.SetPixel(x + offsetX, vein.y1 + offsetY, '#')
		}

		for y := vein.y1; y <= vein.y2; y++ {
			screen.SetPixel(vein.x1 + offsetX, y + offsetY, '#')
		}
	}

	return screen
}

func canContainWater(screen *au.Screen, x int, y int) bool {
	foundLeftBlock := false
	for left := x - 1; left >= 0; left-- {
		if screen.GetPixel(left, y) == '#' {
			foundLeftBlock = true
		}
	}

	if !foundLeftBlock {
		return false
	}

	for right := x + 1; right < screen.Width; right++ {
		if screen.GetPixel(right, y) == '#' {
			return true
		}
	}

	return false
}

func expandWater(screen *au.Screen, x int, y int) bool {
	updated := false

	leftContained := false
	left := x
	for left > 0 {
		if screen.GetPixel(left - 1, y) == '#' {
			leftContained = true
			break
		}

		below := screen.GetPixel(left, y + 1)
		if below != '#' && below != '~' {
			leftContained = false
			break;
		}

		left--
	}

	rightContained := false
	right := x
	for right < screen.Width - 1 {
		if screen.GetPixel(right + 1, y) == '#' {
			rightContained = true
			break
		}

		below := screen.GetPixel(right, y + 1)
		if below != '#' && below != '~' {
			rightContained = false
			break;
		}

		right++
	}

	if leftContained && rightContained {
		for xx := left; xx <= right; xx++ {
			if screen.GetPixel(xx, y) != '~' {
				screen.SetPixel(xx, y, '~')
				updated = true
			}
		}
	} else {
		for xx := left; xx <= right; xx++ {
			if screen.GetPixel(xx, y) != '|' {
				screen.SetPixel(xx, y, '|')
				updated = true
			}
		}
	}

	return updated
}

func iterate(screen *au.Screen) bool {
	done := true

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height - 1; y++ {
			kind := screen.GetPixel(x, y)

			if kind == '|' {
				below := screen.GetPixel(x, y + 1)

				if below == 0 {
					screen.SetPixel(x, y + 1, '|')
					done = false
				}

				if below == '#' || below == '~' {
					if expandWater(screen, x, y) {
						done = false
					}
				}
			}
		}
	}

	// screen.Print()

	return done
}

func getRestCount(screen *au.Screen) int {
	count := 0

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			kind := screen.GetPixel(x, y)
			if kind == '~' {
				count++
			}
		}
	}

	return count
}

func main() {
	inputs := au.ReadInputAsStringArray("17")
	// inputs = testInputs()

	veins := parseInputs(inputs)

	screen := createScreen(veins)

	iteration := 1
	for {
		if iterate(screen) {
			break
		}

		iteration++
	}

	fmt.Println(getRestCount(screen))
}