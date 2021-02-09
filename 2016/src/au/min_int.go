package au

import (
	"math"
)

func MinInt(value1 int, value2 int) int {
	return int(math.Min(float64(value1), float64(value2)))
}
