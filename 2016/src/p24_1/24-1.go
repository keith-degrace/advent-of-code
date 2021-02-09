package p24_1

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

func remove(values []int, index int) []int {
	rest := append([]int(nil), values...)

	copy(rest[index:], rest[index+1:])

	return rest[:len(rest)-1]
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

func getShortestPath(locations []int, pathLengths [][]int) int {
	if len(locations) == 2 {
		return pathLengths[locations[0]][locations[1]]
	}

	from := locations[0]

	minPathLength := math.MaxInt16

	for i := 1; i < len(locations); i++ {
		to := locations[i]

		var subLocations []int
		subLocations = append(subLocations, to)
		subLocations = append(subLocations, remove(locations[1:], i-1)...)
		subPathLength := getShortestPath(subLocations, pathLengths)

		pathLength := pathLengths[from][to] + subPathLength

		minPathLength = au.MinInt(minPathLength, pathLength)
	}

	return minPathLength
}

func Solve() {
	inputs := au.ReadInputAsStringArray("24")
	// inputs := testInputs()

	screen, coords := load(inputs)

	pathLengths := getPathLengths(screen, coords)

	var locations []int
	for i := 0; i < len(pathLengths); i++ {
		locations = append(locations, i)
	}

	fmt.Println(getShortestPath(locations, pathLengths))
}
