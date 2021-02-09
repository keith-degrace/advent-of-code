package p24_1

import (
	"au"
	"fmt"
	"sort"
	"strings"
)

type Component struct {
	left  int
	right int
}

func areCompatible(first Component, second Component, third *Component) bool {
	// If first and second are connected on second's left side, then third must be connected to second's right side (if a third is provided)
	if first.left == second.left || first.right == second.left {
		return third == nil || second.right == third.left || second.right == third.right
	}

	// If first and second are connected on second's right side, then third must be connected to second's left side (if a third is provided)
	if first.left == second.right || first.right == second.right {
		return third == nil || second.left == third.left || second.left == third.right
	}

	return false
}

func getStrength(arrangement []Component) int {
	strength := 0

	for _, component := range arrangement {
		strength += component.left + component.right
	}

	return strength
}

func getOthers(components []Component, index int) []Component {
	others := []Component{}

	for i, component := range components {
		if i == index {
			continue
		}

		others = append(others, component)
	}

	return others
}

func permutate(arrangement []Component, rest []Component, apply func([]Component)) {

	subArrangement := make([]Component, len(arrangement)+1)
	copy(subArrangement, arrangement)

	for i, component := range rest {
		if len(arrangement) > 1 {
			if !areCompatible(arrangement[len(arrangement)-2], arrangement[len(arrangement)-1], &component) {
				continue
			}
		} else if len(arrangement) > 0 {
			if !areCompatible(arrangement[len(arrangement)-1], component, nil) {
				continue
			}
		}

		subArrangement[len(subArrangement)-1] = component
		subRest := getOthers(rest, i)

		if component.left == 0 || component.right == 0 {
			apply(subArrangement)
		} else {
			permutate(subArrangement, subRest, apply)
		}
	}
}

func Solve() {
	input := au.ReadInputAsStringArray("24")

	components := []Component{}
	for _, line := range input {
		parts := strings.Split(line, "/")

		component := Component{}
		component.left = au.ToNumber(parts[0])
		component.right = au.ToNumber(parts[1])

		components = append(components, component)
	}

	sort.Slice(components, func(i, j int) bool {
		maxI := au.MaxInt(components[i].left, components[i].right)
		maxJ := au.MaxInt(components[j].left, components[j].right)
		if maxI != maxJ {
			return maxI > maxJ
		}

		minI := au.MinInt(components[i].left, components[i].right)
		minJ := au.MinInt(components[j].left, components[j].right)

		return minI > minJ
	})

	strongest := 0

	permutate([]Component{}, components, func(arrangement []Component) {
		strength := getStrength(arrangement)
		if strength > strongest {
			strongest = strength
		}
	})

	fmt.Println(strongest)
}
