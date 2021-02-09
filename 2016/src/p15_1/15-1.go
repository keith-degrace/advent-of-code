package p15_1

import (
	"au"
	"fmt"
	"regexp"
)

type Disc struct {
	positions int
	startingPosition int
}

func parseInputs(inputs []string) []Disc {
	discs := []Disc {}

	re := regexp.MustCompile("Disc #[0-9]+ has ([0-9]+) positions; at time=[0-9]+, it is at position ([0-9]+).")

	for _,input := range inputs {
		matches := re.FindStringSubmatch(input)

		disc := Disc {
			au.ToNumber(matches[1]),
			au.ToNumber(matches[2]),
		}

		discs = append(discs, disc)
	}

	return discs
}

func isDiscAtZero(disc Disc, time int) bool {
	position := (disc.startingPosition + time) % disc.positions
	return position == 0
}

func Solve() {
	inputs := au.ReadInputAsStringArray("15")

	discs := parseInputs(inputs)

	for time := 0; ; time++ {
		found := true
		for i := 0; i < len(discs); i++ {
			if !isDiscAtZero(discs[i], time + i) {
				found = false
				break
			}
		}

		if found {
			fmt.Println(time - 1)
			break
		}
	}
}

