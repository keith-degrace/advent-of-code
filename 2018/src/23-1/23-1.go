package main

import (
	"au"
	"fmt"
	"math"
	"regexp"
)

func testInputs() []string {
	return []string {
		"pos=<0,0,0>, r=4",
		"pos=<1,0,0>, r=1",
		"pos=<4,0,0>, r=3",
		"pos=<0,2,0>, r=1",
		"pos=<0,5,0>, r=3",
		"pos=<0,0,3>, r=1",
		"pos=<1,1,1>, r=1",
		"pos=<1,1,2>, r=1",
		"pos=<1,3,1>, r=1",
	}
}

type Nanobot struct {
	x int
	y int
	z int
	r int
}

func parseInputs(inputs []string) []Nanobot {
	nanobots := []Nanobot {}

	re := regexp.MustCompile("pos=<(.+),(.+),(.+)>, r=(.+)")

	for _,input := range inputs {
		matches := re.FindStringSubmatch(input)

		nanobots = append(nanobots, Nanobot {
			x: au.ToNumber(matches[1]),
			y: au.ToNumber(matches[2]),
			z: au.ToNumber(matches[3]),
			r: au.ToNumber(matches[4]),
		})

	}

	return nanobots
}

func getStrongest(nanobots []Nanobot) Nanobot {
	strongest := nanobots[0]

	for i := 1; i < len(nanobots); i++ {
		if nanobots[i].r > strongest.r {
			strongest = nanobots[i]
		}
	}

	return strongest
}

func getDistance(nanobot1 Nanobot, nanobot2 Nanobot) int {
	dx := math.Abs(float64(nanobot1.x - nanobot2.x))
	dy := math.Abs(float64(nanobot1.y - nanobot2.y))
	dz := math.Abs(float64(nanobot1.z - nanobot2.z))

	return int(dx + dy + dz)
}

func getNanobotsInRange(nanobot Nanobot, nanobots []Nanobot) int {
	count := 0

	for i := 0; i < len(nanobots); i++ {
		distance := getDistance(nanobot, nanobots[i])
		if distance <= nanobot.r {
			count++
		}
	}

	return count
}

func main() {
	inputs := au.ReadInputAsStringArray("23")
	//inputs  = testInputs()

	nanobots := parseInputs(inputs)
	strongest := getStrongest(nanobots)

	fmt.Println(getNanobotsInRange(strongest, nanobots))
}