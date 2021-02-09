package main

import (
	"au"
	"regexp"
)

func testInputs() []string {
	return []string {
		"position=< 9,  1> velocity=< 0,  2>",
		"position=< 7,  0> velocity=<-1,  0>",
		"position=< 3, -2> velocity=<-1,  1>",
		"position=< 6, 10> velocity=<-2, -1>",
		"position=< 2, -4> velocity=< 2,  2>",
		"position=<-6, 10> velocity=< 2, -2>",
		"position=< 1,  8> velocity=< 1, -1>",
		"position=< 1,  7> velocity=< 1,  0>",
		"position=<-3, 11> velocity=< 1, -2>",
		"position=< 7,  6> velocity=<-1, -1>",
		"position=<-2,  3> velocity=< 1,  0>",
		"position=<-4,  3> velocity=< 2,  0>",
		"position=<10, -3> velocity=<-1,  1>",
		"position=< 5, 11> velocity=< 1, -2>",
		"position=< 4,  7> velocity=< 0, -1>",
		"position=< 8, -2> velocity=< 0,  1>",
		"position=<15,  0> velocity=<-2,  0>",
		"position=< 1,  6> velocity=< 1,  0>",
		"position=< 8,  9> velocity=< 0, -1>",
		"position=< 3,  3> velocity=<-1,  1>",
		"position=< 0,  5> velocity=< 0, -1>",
		"position=<-2,  2> velocity=< 2,  0>",
		"position=< 5, -2> velocity=< 1,  2>",
		"position=< 1,  4> velocity=< 2,  1>",
		"position=<-2,  7> velocity=< 2, -2>",
		"position=< 3,  6> velocity=<-1, -1>",
		"position=< 5,  0> velocity=< 1,  0>",
		"position=<-6,  0> velocity=< 2,  0>",
		"position=< 5,  9> velocity=< 1, -2>",
		"position=<14,  7> velocity=<-2,  0>",
		"position=<-3,  6> velocity=< 2, -1>",
	}
}

type Point struct {
	x int
	y int
	speedX int
	speedY int
}

type Simulation struct {
	points []Point

	minX int
	minY int
	maxX int
	maxY int
}

func parseInputs(inputs []string) Simulation {
	simulation := Simulation{
		points: make([]Point, 0),
	}

	re := regexp.MustCompile("position=<(.*),(.*)> velocity=<(.*),(.*)>")

	simulation.minX = 10000000
	simulation.minY = 10000000
	simulation.maxX = 0
	simulation.maxY = 0

	for _,input := range inputs {
		matches := re.FindStringSubmatch(input)

		point := Point{
			x: au.ToNumber(matches[1]),
			y: au.ToNumber(matches[2]),
			speedX: au.ToNumber(matches[3]),
			speedY: au.ToNumber(matches[4]),
		}

		simulation.points = append(simulation.points, point)

		simulation.minX = au.MinInt(simulation.minX, point.x)
		simulation.minY = au.MinInt(simulation.minY, point.y)
		simulation.maxX = au.MaxInt(simulation.maxX, point.x)
		simulation.maxY = au.MaxInt(simulation.maxY, point.y)
	}

	return simulation;
}

func (this *Simulation) simulate() {
	this.minX = 10000000
	this.minY = 10000000
	this.maxX = 0
	this.maxY = 0

	for i := 0; i < len(this.points); i++ {
		this.points[i].x += this.points[i].speedX
		this.points[i].y += this.points[i].speedY

		this.minX = au.MinInt(this.minX, this.points[i].x)
		this.minY = au.MinInt(this.minY, this.points[i].y)
		this.maxX = au.MaxInt(this.maxX, this.points[i].x)
		this.maxY = au.MaxInt(this.maxY, this.points[i].y)
	}
}


func (this *Simulation) getHeight() int {
	return this.maxY - this.minY + 1
}

func (this *Simulation) print() {
	width := this.maxX - this.minX + 1
	height := this.maxY - this.minY + 1

	screen := au.NewScreen(width, height)

	for _,point := range this.points {
		shiftedX := point.x - this.minX
		shiftedY := point.y - this.minY
		
		screen.SetPixel(shiftedX, shiftedY, '#')
	}

	screen.Print();
}

func main() {
	inputs := au.ReadInputAsStringArray("10")
	// inputs := testInputs()

	simulation := parseInputs(inputs)
	// print(points)

	for simulation.getHeight() != 10 {
		simulation.simulate()
	}

	simulation.print()
}
