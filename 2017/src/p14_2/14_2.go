package p14_2

import (
	"fmt"
)

type Group struct {
	squares map[*Square]bool
}

func (g *Group) MergeGroups(other *Group) {
	for otherSquare := range other.squares {
		g.squares[otherSquare] = true
		otherSquare.group = g
	}
}

type Square struct {
	group *Group
}

func (s *Square) Print() {
	fmt.Printf("%v (%v)\n", s.group, len(s.group.squares))
}

func reverse(array []int, pos int, length int) {
	for i := 0; i < length/2; i++ {
		a := (pos + i) % len(array)
		b := (pos + length - i - 1) % len(array)

		array[a], array[b] = array[b], array[a]
	}
}

func getSparseHash(numbers []int) []int {

	numbers = append(numbers, []int{17, 31, 73, 47, 23}...)

	sparseHash := make([]int, 256)
	for i := 0; i < len(sparseHash); i++ {
		sparseHash[i] = i
	}

	current := 0
	skipSize := 0

	for round := 0; round < 64; round++ {
		for _, number := range numbers {
			reverse(sparseHash, current, number)

			current += number%len(sparseHash) + skipSize
			skipSize++
		}
	}

	return sparseHash
}

func getDenseHash(sparseHash []int) string {
	denseHash := ""

	for i := 0; i < 16; i++ {

		value := sparseHash[i*16]
		for j := 1; j < 16; j++ {
			value ^= sparseHash[i*16+j]
		}

		denseHash += fmt.Sprintf("%08b", value)
	}

	return denseHash
}

func getKnotHash(value string) string {
	numbers := []int{}
	for _, char := range value {
		numbers = append(numbers, int(char))
	}

	sparseHash := getSparseHash(numbers)

	return getDenseHash(sparseHash)
}

func generateGrid(input string) [][]*Square {
	grid := make([][]*Square, 128)
	for row := 0; row < 128; row++ {
		grid[row] = make([]*Square, 128)

		hash := getKnotHash(fmt.Sprintf("%v-%v", input, row))

		for column, char := range hash {
			if char == '1' {
				square := new(Square)
				square.group = new(Group)
				square.group.squares = make(map[*Square]bool)
				square.group.squares[square] = true
				grid[row][column] = square
			}
		}
	}

	return grid
}

func Solve() {
	input := "vbqugkhl"
	// input := "flqrgnkx"

	grid := generateGrid(input)

	for row := 0; row < 128; row++ {
		for column := 0; column < 128; column++ {

			if grid[row][column] == nil {
				continue
			}

			if row > 0 {
				neighbor := grid[row-1][column]

				if neighbor != nil {
					neighbor.group.MergeGroups(grid[row][column].group)
				}
			}

			if column > 0 {
				neighbor := grid[row][column-1]

				if neighbor != nil {
					neighbor.group.MergeGroups(grid[row][column].group)
				}
			}
		}
	}

	groupIds := make(map[*Group]bool)
	for row := 0; row < 128; row++ {
		for column := 0; column < 128; column++ {
			if grid[row][column] != nil {
				groupIds[grid[row][column].group] = true
			}
		}
	}

	fmt.Println(len(groupIds))
}
