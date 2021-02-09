package au

import (
	"fmt"
	"strings"
)

type Screen struct {
	Width int
	Height int
	pixels []byte
}

func NewScreen(Width int, Height int) *Screen {
	screen := new(Screen)
	screen.Width = Width
	screen.Height = Height
	screen.pixels = make([]byte, Width * Height)
	return screen
}

func (this *Screen) SetPixel(x int, y int, value byte) {
	this.pixels[y * this.Width + x] = value
}

func (this *Screen) GetPixel(x int, y int) byte {
	return this.pixels[y * this.Width + x]
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
