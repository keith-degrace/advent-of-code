package main

import (
	"au"
	"fmt"
	"math"
	"strings"
)

func testInputs() []string {
	return []string {
		"1, 1",
		"1, 6",
		"8, 3",
		"3, 4",
		"5, 5",
		"8, 9",
	}
}

type Coord struct {
	x int
	y int
}

func parseCoords(inputs []string) []Coord {
	coords := []Coord{}

	for _, input := range inputs {
		stringValues := strings.Split(input, ",")

		x := au.ToNumber(stringValues[0])
		y := au.ToNumber(stringValues[1])

		coords = append(coords, Coord {x, y})
	}

	return coords
}

func getTopLeft(coords []Coord) Coord {
	topLeft := Coord{1000000, 1000000}

	for _, coord := range coords {
		if coord.x < topLeft.x {
			topLeft.x = coord.x
		}

		if coord.y < topLeft.y {
			topLeft.y = coord.y
		}
	}

	return topLeft
}

func getBottomRight(coords []Coord) Coord {
	bottomRight := Coord{0, 0}

	for _, coord := range coords {
		if coord.x > bottomRight.x {
			bottomRight.x = coord.x
		}

		if coord.y > bottomRight.y {
			bottomRight.y = coord.y
		}
	}

	return bottomRight
}

func shiftCoords(coords []Coord) []Coord {
	topLeft := getTopLeft(coords)

	shiftedCoords := []Coord{}

	for _, coord := range coords {
		shiftedCoords = append(shiftedCoords, Coord{coord.x - topLeft.x, coord.y - topLeft.y})
	}

	return shiftedCoords
}

func createScreen(coords []Coord) *au.Screen {
	bottomRight := getBottomRight(coords)

	width := bottomRight.x + 1
	height := bottomRight.y + 1

	return au.NewScreen(width, height)
}

func renderCoords(screen *au.Screen, coords []Coord) {
	for index, coord := range coords {
		screen.SetPixel(coord.x, coord.y, byte(index + 65))
	}
}

func getDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x2 - x1)) + math.Abs(float64(y2 - y1)))
}

func getClosestCoord(x int, y int, coords []Coord) (int) {
	minDistance := getDistance(coords[0].x, coords[0].y, x, y)
	minCoord := 0
	count := 1

	for i := 1; i < len(coords); i++ {
		distance := getDistance(coords[i].x, coords[i].y, x, y)
		if distance == minDistance {
			count++
		}

		if distance < minDistance {
			minDistance = distance
			minCoord = i
			count = 1
		}
	}

	if count > 1 {
		return -1;
	}

	return minCoord
}

func renderClosestCoords(screen *au.Screen, coords []Coord) {
	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			minCoord := getClosestCoord(x, y, coords)
			if minCoord == -1 {
				screen.SetPixel(x, y, '.')
			} else {
				screen.SetPixel(x, y, byte(minCoord + 97))
			}
		}
	}
}

func getArea(screen *au.Screen, coordIndex int) int {
	count := 0

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			value := screen.GetPixel(x, y)

			if value == byte(coordIndex + 97) {
				if x == 0 || y == 0 || x == screen.Width - 1 || y == screen.Height - 1 {
					return -1;
				}

				count++
			}
		}
	}

	return count
}

func getLargestArea(screen *au.Screen, coords []Coord) int {
	largest := -1;

	for index := range coords {
		largest = au.MaxInt(largest, getArea(screen, index))
	}

	return largest
}

func main() {
	inputs := au.ReadInputAsStringArray("06")
	// inputs := testInputs();

	coords := parseCoords(inputs)
	coords = shiftCoords(coords)

	screen := createScreen(coords)

	renderCoords(screen, coords)
	renderClosestCoords(screen, coords)

	fmt.Println(getLargestArea(screen, coords))
}
