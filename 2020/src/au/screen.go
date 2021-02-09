package au

import (
	"fmt"
	"strings"
)

type Screen struct {
	Width  int
	Height int
	pixels []byte
}

func NewScreen(Width int, Height int) *Screen {
	screen := new(Screen)
	screen.Width = Width
	screen.Height = Height
	screen.pixels = make([]byte, Width*Height)
	return screen
}

func (this *Screen) Clone() *Screen {
	clone := NewScreen(this.Width, this.Height)
	for index, value := range this.pixels {
		clone.pixels[index] = value
	}
	return clone
}

func (this *Screen) validate(x int, y int) {
	if x < 0 || x >= this.Width {
		panic(fmt.Sprintf("GetPixel/SetPixel(%v, %v) - %v, %v - Value out of range", x, y, this.Width, this.Height))
	}

	if y < 0 || y >= this.Height {
		panic(fmt.Sprintf("GetPixel/SetPixel(%v, %v) - %v, %v - Value out of range", x, y, this.Width, this.Height))
	}
}

func (this *Screen) SetPixel(x int, y int, value byte) {
	this.validate(x, y)
	this.pixels[y*this.Width+x] = value
}

func (this *Screen) GetPixel(x int, y int) byte {
	this.validate(x, y)
	return this.pixels[y*this.Width+x]
}

func (this *Screen) GetCount(value byte) int {
	count := 0

	for _, aValue := range this.pixels {
		if aValue == value {
			count++
		}
	}

	return count
}

func (this *Screen) Equals(other *Screen) bool {
	if this.Width != other.Width {
		return false
	}

	if this.Height != other.Height {
		return false
	}

	for index, value := range this.pixels {
		if other.pixels[index] != value {
			return false
		}
	}

	return true
}

func (this *Screen) Print() int {
	count := 0

	fmt.Printf("%v%v%v\n", string("╔"), strings.Repeat("═", this.Width), string("╗"))

	for y := 0; y < this.Height; y++ {
		fmt.Print("║")
		for x := 0; x < this.Width; x++ {
			fmt.Print(string(this.GetPixel(x, y)))
		}
		fmt.Println("║")
	}

	fmt.Printf("%v%v%v\n", string("╚"), strings.Repeat("═", this.Width), string("╝"))

	return count
}
