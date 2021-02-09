package p24_2

import (
	"au"
	"fmt"
	"math"
)

func testInputs() []string {
	return []string{
		"###########",
		"#0.1.....2#",
		"#.#######.#",
		"#4.......3#",
		"###########",
	}
}

func load(inputs []string) (*au.DynamicScreen, []au.Coord) {
	screen := au.NewDynamicScreen()

	coordMap := make(map[int]au.Coord)

	for y, line := range inputs {
		for x, _ := range line {
			screen.SetPixel(x, y, line[x])

			if line[x] != '#' && line[x] != '.' {
				coordMap[au.ToNumber(string(line[x]))] = au.NewCoord(x, y)
			}
		}
	}

	var coords []au.Coord

	for i := 0; i < len(coordMap); i++ {
		coords = append(coords, coordMap[i])
	}

	return screen, coords
}

func getPathLengths(screen *au.DynamicScreen, coords []au.Coord) [][]int {
	pathLengths := make([][]int, len(coords))
	for i := range pathLengths {
		pathLengths[i] = make([]int, len(coords))
	}

	for i := 0; i <= len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {

			path := au.GetShortestPath(coords[i], coords[j], func(coord au.Coord) []au.Coord {
				var neighboors []au.Coord

				left := coord.ShiftX(-1)
				if screen.GetPixel(left.X(), left.Y()) != '#' {
					neighboors = append(neighboors, left)
				}

				right := coord.ShiftX(1)
				if screen.GetPixel(right.X(), right.Y()) != '#' {
					neighboors = append(neighboors, right)
				}

				above := coord.ShiftY(1)
				if screen.GetPixel(above.X(), above.Y()) != '#' {
					neighboors = append(neighboors, above)
				}

				below := coord.ShiftY(-1)
				if screen.GetPixel(below.X(), below.Y()) != '#' {
					neighboors = append(neighboors, below)
				}

				return neighboors
			})

			pathLengths[i][j] = len(path)
			pathLengths[j][i] = len(path)
		}
	}

	return pathLengths
}

func permutate(locations []int, size int, apply func(locations []int)) {
	if size == 1 {
		apply(locations)
	}

	for i := 0; i < size; i++ {
		permutate(locations, size-1, apply)

		if size%2 == 1 {
			locations[0], locations[size-1] = locations[size-1], locations[0]
		} else {
			locations[i], locations[size-1] = locations[size-1], locations[i]
		}
	}
}

func getShortestPath(pathLengths [][]int) int {
	var locations []int
	for i := 1; i < len(pathLengths); i++ {
		locations = append(locations, i)
	}

	minPathLength := math.MaxInt16

	permutate(locations, len(locations), func(locations []int) {

		candidatePathLength := pathLengths[0][locations[0]]

		for i := 1; i < len(locations); i++ {
			candidatePathLength += pathLengths[locations[i-1]][locations[i]]
		}

		candidatePathLength += pathLengths[0][locations[len(locations)-1]]

		minPathLength = au.MinInt(minPathLength, candidatePathLength)
	})

	return minPathLength
}

func Solve() {
	inputs := au.ReadInputAsStringArray("24")
	// inputs := testInputs()

	screen, coords := load(inputs)

	pathLengths := getPathLengths(screen, coords)

	fmt.Println(getShortestPath(pathLengths))
}
