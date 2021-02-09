package p13_1

import (
	"au"
	"fmt"
	"strings"
)

type Layer struct {
	id    int
	depth int
	pos   int
	dir   int
}

func Solve() {
	input := au.ReadInputAsStringArray("13")

	layerCount := 0

	layers := make(map[int]*Layer)
	for _, line := range input {
		tokens := strings.Split(line, ": ")

		id := au.ToNumber(tokens[0])
		depth := au.ToNumber(tokens[1])

		layers[id] = new(Layer)
		layers[id].id = id
		layers[id].depth = depth
		layers[id].pos = 0
		layers[id].dir = 1

		layerCount = au.MaxInt(layerCount, id+1)
	}

	severity := 0

	for current := 0; current < layerCount; current++ {

		if layer, ok := layers[current]; ok {
			if layer.pos == 0 {
				severity += layer.id * layer.depth
			}
		}

		for _, layer := range layers {
			layer.pos += layer.dir
			if layer.pos < 0 {
				layer.pos = 1
				layer.dir *= -1
			} else if layer.pos >= layer.depth {
				layer.pos = layer.depth - 2
				layer.dir *= -1
			}
		}
	}

	fmt.Println(severity)
}
