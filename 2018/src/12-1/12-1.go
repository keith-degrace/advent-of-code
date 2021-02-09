package main

import (
	"au"
	"fmt"
	"regexp"
)

func testInputs() []string {
	return []string {
		"initial state: #..#.#..##......###...###",
		"",
		"...## => #",
		"..#.. => #",
		".#... => #",
		".#.#. => #",
		".#.## => #",
		".##.. => #",
		".#### => #",
		"#.#.# => #",
		"#.### => #",
		"##.#. => #",
		"##.## => #",
		"###.. => #",
		"###.# => #",
		"####. => #",
	}
}

type Pattern struct {
	values []byte
	next byte
}

func parseInputs(inputs []string) (map[int] byte, []Pattern) {
	initialState := make(map[int] byte)
	for index,char := range inputs[0][15:] {
		initialState[index] = byte(char)
	}

	re := regexp.MustCompile("(.*) => (.)")

	patterns := make([]Pattern, 0)
	for _,input := range inputs {
		matches := re.FindStringSubmatch(input)
		if len(matches) == 0 {
			continue
		} 

		pattern := Pattern {
			values: []byte(matches[1]),
			next: matches[2][0],
		}

		patterns = append(patterns, pattern)
	}

	return initialState, patterns
}

func getPot(state map[int] byte, index int) byte {
	value,ok := state[index]
	if ok {
		return value
	} else {
		return '.'
	}
}

func getPots(state map[int] byte, index int) []byte {
	return []byte{
		getPot(state, index - 2),
		getPot(state, index - 1),
		getPot(state, index),
		getPot(state, index + 1),
		getPot(state, index + 2),
	}
}

func getMinIndex(state map[int] byte) int {
	min := 999999999

	for k,v := range state {
		if v != '.' {
			min = au.MinInt(min, k)
		}
	}

	return min
}

func getMaxIndex(state map[int] byte) int {
	max := 0

	for k,v := range state {
		if v != '.' {
			max = au.MaxInt(max, k)
		}
	}

	return max
}

func generate(state map[int] byte, patterns []Pattern) map[int] byte {
	newState := make(map[int] byte)
	for k,v := range state {
		newState[k] = v
	}

	minIndex := getMinIndex(state) - 10
	maxIndex := getMaxIndex(state) + 10

	for i := minIndex; i <= maxIndex; i++ {
		pots := getPots(state, i)

		found := false
		for _,pattern := range patterns {
			if string(pots) == string(pattern.values) {
				newState[i] = pattern.next
				found = true
				break
			}
		}

		if !found {
			newState[i] = '.'
		}
	}

	return newState
}

func getPlantSum(state map[int] byte) int {
	count := 0

	for k,v := range state {
		if v == '#' {
			count += k
		}
	}

	return count
}

func printState(state map[int] byte) {
	minIndex := getMinIndex(state)
	maxIndex := getMaxIndex(state)

	for i := minIndex; i <= maxIndex; i++ {
		fmt.Print(string(getPot(state, i)))
	}

	fmt.Println()
}

func printStates(states []map[int] byte) {
	minIndex := getMinIndex(states[0]) - 10
	maxIndex := getMaxIndex(states[0]) + 10

	for _,state := range states {
		minIndex = au.MinInt(minIndex, getMinIndex(state))
		maxIndex = au.MaxInt(maxIndex, getMaxIndex(state))
	}

	for _,state := range states {
		for i := minIndex; i <= maxIndex; i++ {
			fmt.Print(string(getPot(state, i)))
		}
		
		fmt.Println()
	}
}

func main() {
	inputs := au.ReadInputAsStringArray("12")
	// inputs := testInputs()

	state,patterns := parseInputs(inputs)

	states := []map[int] byte { state }

	for i := 0; i < 20; i++ {
		state = generate(state, patterns)
		states = append(states, state)
	}

	printStates(states)

	fmt.Println(getPlantSum(state))
}
