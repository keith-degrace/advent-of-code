package au

import (
	"math"
)

func MaxInt(value1 int, value2 int) int {
	return int(math.Max(float64(value1), float64(value2)))
}
