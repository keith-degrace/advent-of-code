package au

import (
	"fmt"
)

type Coord3D struct {
	x   int
	y   int
	z   int
	key string
}

func NewCoord3D(x int, y int, z int) Coord3D {
	return Coord3D{x, y, z, GetCoord3DKey(x, y, z)}
}

func (this *Coord3D) Shift(dx int, dy int, dz int) Coord3D {
	return NewCoord3D(this.x+dx, this.y+dy, this.z+dz)
}

func (this *Coord3D) ShiftX(dx int) Coord3D {
	return this.Shift(dx, 0, 0)
}

func (this *Coord3D) ShiftY(dy int) Coord3D {
	return this.Shift(0, dy, 0)
}

func (this *Coord3D) ShiftZ(dz int) Coord3D {
	return this.Shift(0, 0, dz)
}

func (this *Coord3D) X() int {
	return this.x
}

func (this *Coord3D) Y() int {
	return this.y
}

func (this *Coord3D) Z() int {
	return this.z
}

func (this *Coord3D) Key() string {
	if len(this.key) == 0 {
		this.key = GetCoord3DKey(this.X(), this.Y(), this.Z())
	}
	return this.key
}

func (this *Coord3D) Equals(other Coord3D) bool {
	return this.EqualsXYZ(other.x, other.y, other.z)
}

func (this *Coord3D) EqualsXYZ(x int, y int, z int) bool {
	return this.x == x && this.y == y && this.z == z
}

func (this Coord3D) String() string {
	return fmt.Sprintf("%v,%v,%v", this.x, this.y, this.z)
}

func GetCoord3DKey(x int, y int, z int) string {
	return fmt.Sprintf("%vx%v%v", x, y, z)
}
