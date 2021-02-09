package main

import (
	"au"
	"container/list"
	"fmt"
	"math"
	"strings"
)

func testInputs1() []string {
	return []string {
	"	0,0,0,0",
	"	3,0,0,0",
	"	0,3,0,0",
	"	0,0,3,0",
	"	0,0,0,3",
	"	0,0,0,6",
	"	9,0,0,0",
	" 12,0,0,0",
	}
}

func testInputs2() []string {
	return []string {
		"-1,2,2,0",
		"0,0,2,-2",
		"0,0,0,-2",
		"-1,2,0,0",
		"-2,-2,-2,2",
		"3,0,2,-1",
		"-1,3,2,2",
		"-1,0,-1,0",
		"0,2,1,-2",
		"3,0,0,0",
	}
}

func testInputs3() []string {
	return []string {
		"1,-1,0,1",
		"2,0,-1,0",
		"3,2,-1,0",
		"0,0,3,1",
		"0,0,-1,-1",
		"2,3,-2,0",
		"-2,2,0,0",
		"2,-2,0,-1",
		"1,-1,0,-1",
		"3,2,0,2",
	}
}

func testInputs4() []string {
	return []string {
		"1,-1,-1,-2",
		"-2,-2,0,1",
		"0,2,1,3",
		"-2,3,-2,1",
		"0,2,3,-2",
		"-1,-1,1,-2",
		"0,-2,-1,0",
		"-2,2,3,-1",
		"1,2,2,0",
		"-1,-2,0,-2",
	}
}

type Coord struct {
	x int
	y int
	z int
	t int

	group *list.List
}

type Constellation struct {
	min Coord
	max Coord
}

func parseInputs(inputs []string) []Coord {
	coords := []Coord{}

	for _,input := range inputs {
		tokens := strings.Split(strings.TrimSpace(input), ",")

		coord := Coord {
			x: au.ToNumber(tokens[0]),
			y: au.ToNumber(tokens[1]),
			z: au.ToNumber(tokens[2]),
			t: au.ToNumber(tokens[3]),
			group: list.New(),
		}

		coords = append(coords, coord)
	}

	return coords
}

func getDistance(coord1 Coord, coord2 Coord) int {
	dx := math.Abs(float64(coord2.x - coord1.x))
	dy := math.Abs(float64(coord2.y - coord1.y))
	dz := math.Abs(float64(coord2.z - coord1.z))
	dt := math.Abs(float64(coord2.t - coord1.t))

	return int(dx + dy + dz + dt)
}

func main() {
	inputs := au.ReadInputAsStringArray("25")
	// inputs  = testInputs3()
	
	coords := parseInputs(inputs)
	for i := 0; i < len(coords); i++ {
		coords[i].group.PushBack(&coords[i])
	}

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			if coords[i].group == coords[j].group {
				continue
			}

			if getDistance(coords[i], coords[j]) > 3 {
				continue
			}

			for current := coords[j].group.Front(); current != nil; current = current.Next() {
				coordToMove := current.Value.(*Coord)
				coordToMove.group = coords[i].group
				coords[i].group.PushBack(coordToMove)
			}
		}
	}

	groups := map [*list.List] bool{}
	for _,coord := range coords {
		groups[coord.group] = true
	}

	fmt.Println(len(groups))
}