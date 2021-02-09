package au

import (
	"fmt"
)

type Coord struct {
	x int
	y int
	key string
}

func NewCoord(x int, y int) Coord {
	return Coord { x, y, GetCoordKey(x, y) }
}

func (this *Coord) Shift(dx int, dy int) Coord {
	return NewCoord(this.x + dx, this.y + dy)
}

func (this *Coord) ShiftX(dx int) Coord {
	return this.Shift(dx, 0)
}

func (this *Coord) ShiftY(dy int) Coord {
	return this.Shift(0, dy)
}

func (this *Coord) X() int {
	return this.x
}

func (this *Coord) Y() int {
	return this.y
}

func (this *Coord) Key() string {
	if len(this.key) == 0 {
		this.key = GetCoordKey(this.X(), this.Y())
	}
	return this.key
}

func (this *Coord) Equals(other Coord) bool {
	return this.EqualsXY(other.x, other.y)
}

func (this *Coord) EqualsXY(x int, y int) bool {
	return this.x == x && this.y == y
}

func GetCoordKey(x int, y int) string {
	return fmt.Sprintf("%vx%v", x, y)
}
