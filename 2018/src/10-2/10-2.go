package main

import (
	"au"
	"fmt"
	"regexp"
)

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

	simulation := parseInputs(inputs)

	seconds := 0
	for simulation.getHeight() != 10 {
		simulation.simulate()
		seconds++
	}

	fmt.Println(seconds)
}
