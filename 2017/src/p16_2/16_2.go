package p16_2

import (
	"au"
	"fmt"
	"strings"
)

func initPrograms() []string {
	programs := []string{}

	for i := 0; i < 16; i++ {
		programs = append(programs, string('a'+i))
	}

	return programs
}

func applyMoves(programs []string, pos int, moves []string) int {

	for _, move := range moves {
		// Spin
		if move[0] == 's' {
			amount := au.ToNumber(move[1:])
			for i := 0; i < amount; i++ {
				if pos > 0 {
					pos--
				} else {
					pos = len(programs) - 1
				}
			}
		}

		// Exchange
		if move[0] == 'x' {
			positions := strings.Split(move[1:], "/")

			position1 := au.ToNumber(positions[0])
			position2 := au.ToNumber(positions[1])

			index1 := (pos + position1) % len(programs)
			index2 := (pos + position2) % len(programs)

			programs[index1], programs[index2] = programs[index2], programs[index1]
		}

		// Partner
		if move[0] == 'p' {
			partner1 := string(move[1])
			partner2 := string(move[3])

			index1 := 0
			for i := 0; i < len(programs); i++ {
				if programs[i] == partner1 {
					index1 = i
					break
				}
			}

			for i := 0; i < len(programs); i++ {
				if programs[i] == partner2 {
					programs[i], programs[index1] = programs[index1], programs[i]
					break
				}
			}
		}
	}

	return pos
}

func Solve() {
	input := au.ReadInputAsString("16")

	moves := strings.Split(input, ",")

	programs := initPrograms()
	pos := 0

	possibleResults := []string{}
	for {
		pos = applyMoves(programs, pos, moves)

		result := ""
		for i := 0; i < len(programs); i++ {
			result += programs[(i+pos)%len(programs)]
		}

		if len(possibleResults) > 0 && result == possibleResults[0] {
			break
		}

		possibleResults = append(possibleResults, result)
	}

	// The looping pattern does not include the original configuration, so that's why it's a billion minus one.
	fmt.Println(possibleResults[(1000000000-1)%len(possibleResults)])
}
