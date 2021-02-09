package main

import (
	"au"
	"fmt"
	"regexp"
	"sort"
	"time"
)

func testInputs() []string {
	return []string{
		"pos=<10,12,12>, r=2",
		"pos=<12,14,12>, r=2",
		"pos=<16,12,12>, r=4",
		"pos=<14,14,14>, r=6",
		"pos=<50,50,50>, r=200",
		"pos=<10,10,10>, r=5",
	}
}

func parseInputs(inputs []string) []*Nanobot {
	nanobots := []*Nanobot{}

	re := regexp.MustCompile("pos=<(.+),(.+),(.+)>, r=(.+)")

	for _, input := range inputs {
		matches := re.FindStringSubmatch(input)

		nanobots = append(nanobots, NewNanobot(
			au.ToNumber(matches[1]),
			au.ToNumber(matches[2]),
			au.ToNumber(matches[3]),
			au.ToNumber(matches[4])))
	}

	return nanobots
}

func getDistance(x1, y1, z1, x2, y2, z2 int) int {
	return au.AbsInt(x1-x2) + au.AbsInt(y1-y2) + au.AbsInt(z1-z2)
}

// Axis-Aligned Bounding Box
type BoundingBox struct {
	min au.Coord3D
	max au.Coord3D
	mid au.Coord3D
}

func NewBoundingBox(minX, minY, minZ, maxX, maxY, maxZ int) *BoundingBox {
	box := new(BoundingBox)

	box.min = au.NewCoord3D(minX, minY, minZ)
	box.max = au.NewCoord3D(maxX, maxY, maxZ)
	box.mid = au.NewCoord3D(minX+box.Width()/2, minY+box.Height()/2, minZ+box.Depth()/2)

	return box
}

func (a *BoundingBox) Clone() *BoundingBox {
	copy := new(BoundingBox)

	copy.min = a.min
	copy.max = a.max

	return copy
}

func (b *BoundingBox) Width() int {
	return b.max.X() - b.min.X()
}

func (b *BoundingBox) Height() int {
	return b.max.Y() - b.min.Y()
}

func (b *BoundingBox) Depth() int {
	return b.max.Z() - b.min.Z()
}

func (b *BoundingBox) Intersects(other *BoundingBox) bool {
	return (b.min.X() <= other.max.X() && b.max.X() >= other.min.X()) &&
		(b.min.Y() <= other.max.Y() && b.max.Y() >= other.min.Y()) &&
		(b.min.Z() <= other.max.Z() && b.max.Z() >= other.min.Z())
}

func (b *BoundingBox) Contains(point au.Coord3D) bool {
	return point.X() >= b.min.X() && point.X() <= b.max.X() &&
		point.Y() >= b.min.Y() && point.Y() <= b.max.Y() &&
		point.Z() >= b.min.Z() && point.Z() <= b.max.Z()
}

func (b *BoundingBox) Partition() []*BoundingBox {
	return []*BoundingBox{
		NewBoundingBox(b.min.X(), b.min.Y(), b.min.Z(), b.mid.X(), b.mid.Y(), b.mid.Z()),
		NewBoundingBox(b.mid.X(), b.min.Y(), b.min.Z(), b.max.X(), b.mid.Y(), b.mid.Z()),
		NewBoundingBox(b.min.X(), b.mid.Y(), b.min.Z(), b.mid.X(), b.max.Y(), b.mid.Z()),
		NewBoundingBox(b.mid.X(), b.mid.Y(), b.min.Z(), b.max.X(), b.max.Y(), b.mid.Z()),
		NewBoundingBox(b.min.X(), b.min.Y(), b.mid.Z(), b.mid.X(), b.mid.Y(), b.max.Z()),
		NewBoundingBox(b.mid.X(), b.min.Y(), b.mid.Z(), b.max.X(), b.mid.Y(), b.max.Z()),
		NewBoundingBox(b.min.X(), b.mid.Y(), b.mid.Z(), b.mid.X(), b.max.Y(), b.max.Z()),
		NewBoundingBox(b.mid.X(), b.mid.Y(), b.mid.Z(), b.max.X(), b.max.Y(), b.max.Z()),
	}
}

type Nanobot struct {
	coord       au.Coord3D
	radius      int
	corners     []au.Coord3D
	boundingBox *BoundingBox
}

func NewNanobot(x int, y int, z int, radius int) *Nanobot {
	nanobot := new(Nanobot)
	nanobot.coord = au.NewCoord3D(x, y, z)
	nanobot.radius = radius

	// Calculate the corners
	minX := nanobot.coord.X() - nanobot.radius
	maxX := nanobot.coord.X() + nanobot.radius

	minY := nanobot.coord.Y() - nanobot.radius
	maxY := nanobot.coord.Y() + nanobot.radius

	minZ := nanobot.coord.Z() - nanobot.radius
	maxZ := nanobot.coord.Z() + nanobot.radius

	nanobot.corners = []au.Coord3D{
		au.NewCoord3D(minX, y, z),
		au.NewCoord3D(maxX, y, z),
		au.NewCoord3D(x, minY, z),
		au.NewCoord3D(x, maxY, z),
		au.NewCoord3D(x, y, minZ),
		au.NewCoord3D(x, y, maxZ),
	}

	// Calculate the axis aligned bounding box
	nanobot.boundingBox = new(BoundingBox)
	nanobot.boundingBox.min = au.NewCoord3D(
		nanobot.coord.X()-nanobot.radius,
		nanobot.coord.Y()-nanobot.radius,
		nanobot.coord.Z()-nanobot.radius)

	nanobot.boundingBox.max = au.NewCoord3D(
		nanobot.coord.X()+nanobot.radius,
		nanobot.coord.Y()+nanobot.radius,
		nanobot.coord.Z()+nanobot.radius)

	return nanobot
}

func (n *Nanobot) IsInRange(x int, y int, z int) bool {
	return getDistance(n.coord.X(), n.coord.Y(), n.coord.Z(), x, y, z) <= n.radius
}

func (n *Nanobot) IntersectsBox(box *BoundingBox) bool {
	// Because we are using manhattan distance, the radius is not a sphere, instead
	// it's a octahedron (i.e. 3D diamond).  We can check for overlap by checking if
	// at least one corner of one octahedron is in the other

	// Let's do a cheap bounding box test first.
	if !n.boundingBox.Intersects(box) {
		return false
	}

	// Check if any of the corners of the nanobot range are in the bounding box
	for _, corner := range n.corners {
		if box.Contains(corner) {
			return true
		}
	}

	// Check if any of the corners of the box are within range of the nanobot
	return n.IsInRange(box.mid.X(), box.mid.Y(), box.mid.Z()) ||
		n.IsInRange(box.min.X(), box.min.Y(), box.min.Z()) ||
		n.IsInRange(box.min.X(), box.max.Y(), box.min.Z()) ||
		n.IsInRange(box.max.X(), box.min.Y(), box.min.Z()) ||
		n.IsInRange(box.max.X(), box.max.Y(), box.min.Z()) ||
		n.IsInRange(box.min.X(), box.min.Y(), box.max.Z()) ||
		n.IsInRange(box.min.X(), box.max.Y(), box.max.Z()) ||
		n.IsInRange(box.max.X(), box.min.Y(), box.max.Z()) ||
		n.IsInRange(box.max.X(), box.max.Y(), box.max.Z())
}

func getNanobotBoundingBox(nanobots []*Nanobot) *BoundingBox {
	minX := nanobots[0].boundingBox.min.X()
	minY := nanobots[0].boundingBox.min.Y()
	minZ := nanobots[0].boundingBox.min.Z()
	maxX := nanobots[0].boundingBox.max.X()
	maxY := nanobots[0].boundingBox.max.Y()
	maxZ := nanobots[0].boundingBox.max.Z()

	for i := 1; i < len(nanobots); i++ {
		minX = au.MinInt(minX, nanobots[i].boundingBox.min.X())
		minY = au.MinInt(minY, nanobots[i].boundingBox.min.Y())
		minZ = au.MinInt(minZ, nanobots[i].boundingBox.min.Z())
		maxX = au.MaxInt(maxX, nanobots[i].boundingBox.max.X())
		maxY = au.MaxInt(maxY, nanobots[i].boundingBox.max.Y())
		maxZ = au.MaxInt(maxZ, nanobots[i].boundingBox.max.Z())
	}

	boundingBox := new(BoundingBox)
	boundingBox.min = au.NewCoord3D(minX, minY, minZ)
	boundingBox.max = au.NewCoord3D(maxX, maxY, maxZ)

	return boundingBox
}

func search(nanobots []*Nanobot) {
	type StackEntry struct {
		box      *BoundingBox
		nanobots []*Nanobot
	}

	firstEntry := StackEntry{getNanobotBoundingBox(nanobots), nanobots}

	stack := []StackEntry{firstEntry}

	for len(stack) > 0 {
		// Pop
		current := stack[0]
		stack = stack[1:]

		// If our best candidate has reached a size of 1, then we've found it!
		if current.box.Width() == 1 && current.box.Height() == 1 && current.box.Depth() == 1 {
			fmt.Printf("Distance  : %v\n", getDistance(0, 0, 0, current.box.mid.X(), current.box.mid.Y(), current.box.mid.Z()))
			fmt.Println()
			return
		}

		// Partition the current box into 8 sub boxes.
		for _, subBox := range current.box.Partition() {
			subBoxNanobots := []*Nanobot{}
			for _, nanobot := range nanobots {
				if nanobot.IntersectsBox(subBox) {
					subBoxNanobots = append(subBoxNanobots, nanobot)
				}
			}

			stack = append(stack, StackEntry{subBox, subBoxNanobots})
		}

		// This is key... prioritize the based on the box with the highest number of nanobots, followed by the one closest to origin.
		sort.Slice(stack, func(i, j int) bool {
			if len(stack[i].nanobots) > len(stack[j].nanobots) {
				return true
			} else if len(stack[i].nanobots) < len(stack[j].nanobots) {
				return false
			}

			distanceI := getDistance(0, 0, 0, stack[i].box.mid.X(), stack[i].box.mid.Y(), stack[i].box.mid.Z())
			distanceJ := getDistance(0, 0, 0, stack[j].box.mid.X(), stack[j].box.mid.Y(), stack[j].box.mid.Z())

			return distanceI > distanceJ
		})
	}
}

func main() {
	fmt.Printf("Starting at %v\n\n", time.Now())
	startTime := time.Now()

	inputs := au.ReadInputAsStringArray("23")
	nanobots := parseInputs(inputs)

	search(nanobots)

	fmt.Printf("\nCompleted at %v (in %v)\n", time.Now(), time.Now().Sub(startTime))
}
