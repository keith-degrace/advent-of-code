package au

import (
	"fmt"
	"strings"
)

type DynamicScreen struct {
	minX int
	minY int
	maxX int
	maxY int
	pixels map [string] byte
}

func NewDynamicScreen() *DynamicScreen {
	screen := new(DynamicScreen)
	screen.minX = 0
	screen.minY = 0
	screen.maxX = 0
	screen.maxY = 0
	screen.pixels = make(map [string] byte)
	return screen
}

func (this *DynamicScreen) Clone() *DynamicScreen {
	clone := new(DynamicScreen)
	
	clone.minX = this.minX
	clone.minY = this.minY
	clone.maxX = this.maxX
	clone.maxY = this.maxY
	
	for key,value := range this.pixels {
		clone.pixels[key] = value
	}

	return clone
}

func (this *DynamicScreen) Width() int {
	return this.maxX - this.minX + 1
}

func (this *DynamicScreen) Height() int {
	return this.maxY - this.minY + 1
}

func (this *DynamicScreen) SetPixel(x int, y int, value byte) {
	this.pixels[GetCoordKey(x, y)] = value

	this.minX = MinInt(this.minX, x)
	this.minY = MinInt(this.minY, y)
	this.maxX = MaxInt(this.maxX, x)
	this.maxY = MaxInt(this.maxY, y)
}

func (this *DynamicScreen) GetPixel(x int, y int) byte {
	value, ok := this.pixels[GetCoordKey(x, y)]
	if !ok {
		return 0
	}

	return value
}

func (this *DynamicScreen) Print() int {
	count := 0

	width := this.Width()
	height := this.Height()

	offsetX := this.minX
	offsetY := this.minY

	fmt.Printf("%v%v%v\n", string("╔"), strings.Repeat("═", width), string("╗"))

	for y := 0; y < height; y++ {
		fmt.Print("║")
		for x := 0; x < width; x++ {
			fmt.Print(string(this.GetPixel(x + offsetX, y + offsetY)))
		}
		fmt.Println("║")
	}

	fmt.Printf("%v%v%v\n", string("╚"), strings.Repeat("═", width), string("╝"))

	return count
}
