package main

import (
	"au"
	"fmt"
	"strings"
)

func testInputs() []string {
	return []string {
		"depth: 510",
		"target: 10,10",
	}
}

func parseInputs(inputs []string) (int, au.Coord) {
	depth := au.ToNumber(inputs[0][6:])	
	targetX := au.ToNumber(strings.Split(inputs[1][7:], ",")[0])	
	targetY := au.ToNumber(strings.Split(inputs[1][8:], ",")[1])	
	return depth, au.NewCoord(targetX, targetY)
}

type Regions struct {
	caveDepth int
	target au.Coord

	geoIndexCache map[string] int
	erosionLevelCache map[string] int
	regionTypeCache map[string] int
}

func (this *Regions) calculateGeoIndex(coord au.Coord) int {
	if coord.X() == 0 && coord.Y() == 0 {
		return 0
	}

	if coord.X() == this.target.X() && coord.Y() == this.target.Y() {
		return 0
	}

	if coord.Y() == 0 {
		return coord.X() * 16807
	}

	if coord.X() == 0 {
		return coord.Y() * 48271
	}

	return this.getErosionLevel(coord.ShiftX(-1)) * this.getErosionLevel(coord.ShiftY(-1))
}

func (this *Regions) getGeoIndex(coord au.Coord) int {
	value, ok := this.geoIndexCache[coord.Key()]
	if !ok {
		value = this.calculateGeoIndex(coord)
		this.geoIndexCache[coord.Key()] = value
	}

	return value
}

func (this *Regions) getErosionLevel(coord au.Coord) int {
	value, ok := this.erosionLevelCache[coord.Key()]
	if !ok {
		value = (this.getGeoIndex(coord) + this.caveDepth) % 20183
		this.erosionLevelCache[coord.Key()] = value
	}

	return value
}

func (this *Regions) getRegionType(coord au.Coord) int {
	value, ok := this.regionTypeCache[coord.Key()]
	if !ok {
		value = this.getErosionLevel(coord) % 3
		this.regionTypeCache[coord.Key()] = value
	}

	return value
}

func (this *Regions) getRiskLevel() int {
	risk := 0

	for x := 0; x <= this.target.X(); x++ {
		for y := 0; y <= this.target.Y(); y++ {
			if x != 0 || y != 0 || x != this.target.X() || y != this.target.Y() {
				risk += this.getRegionType(au.NewCoord(x, y))
			}
		}
	}

	return risk
}

// func printCave(caveDepth int, target au.Coord) {
// 	screen := au.NewScreen(target.X() + 1, target.Y() + 1)

// 	for x := 0; x <= target.X(); x++ {
// 		for y := 0; y <= target.Y(); y++ {
// 			switch getRegionType(au.NewCoord(x, y), target, caveDepth) {
// 			case 0: screen.SetPixel(x, y, '.')
// 			case 1: screen.SetPixel(x, y, '=')
// 			case 2: screen.SetPixel(x, y, '|')
// 			}
// 		}
// 	}

// 	screen.Print()
// }

func main() {
	inputs := au.ReadInputAsStringArray("22")
	//inputs  = testInputs()

	caveDepth, target := parseInputs(inputs)

	regions := Regions {
		caveDepth: caveDepth,
		target: target,
		geoIndexCache: map[string] int{},
		erosionLevelCache: map[string] int{},
		regionTypeCache: map[string] int{},
	}

	//printCave(caveDepth, target)
	
	fmt.Println(regions.getRiskLevel())
}