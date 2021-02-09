package p03_2

import (
	"au"
	"fmt"
	"os"
)

func GetValue(coords map[au.Coord]int, coord au.Coord) int {
	value, ok := coords[coord]
	if !ok {
		return 0
	}

	return value
}

func GetNeighborSum(coords map[au.Coord]int, coord au.Coord) int {
	sum := 0
	sum += GetValue(coords, coord.ShiftX(-1))
	sum += GetValue(coords, coord.ShiftX(1))
	sum += GetValue(coords, coord.ShiftY(-1))
	sum += GetValue(coords, coord.ShiftY(1))
	sum += GetValue(coords, coord.Shift(-1, -1))
	sum += GetValue(coords, coord.Shift(-1, 1))
	sum += GetValue(coords, coord.Shift(1, -1))
	sum += GetValue(coords, coord.Shift(1, 1))
	return sum

}

func Solve() {
	input := 347991

	values := make(map[au.Coord]int)

	x := 0
	y := 0

	direction := 1

	values[au.NewCoord(0, 0)] = 1

	for width := 1; ; width++ {

		for dx := 0; dx < width; dx++ {
			x += direction

			coord := au.NewCoord(x, y)
			values[coord] = GetNeighborSum(values, coord)

			if values[coord] > input {
				fmt.Println(values[coord])
				os.Exit(0)
			}
		}

		direction = -1 * direction

		for dy := 0; dy < width; dy++ {
			y += direction

			coord := au.NewCoord(x, y)
			values[coord] = GetNeighborSum(values, coord)

			if values[coord] > input {
				fmt.Println(values[coord])
				os.Exit(0)
			}
		}
	}

}
