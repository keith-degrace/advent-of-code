package main

import (
	"au"
	"fmt"
	"regexp"
)

type Pattern struct {
	values [5]bool
	next bool
}

func parseInputs(inputs []string) (map[int] bool, []Pattern) {
	initialState := make(map[int] bool)
	for index,char := range inputs[0][15:] {
		if char == '#' {
			initialState[index] = true
		}
	}

	re := regexp.MustCompile("(.*) => (.)")

	patterns := make([]Pattern, 0)
	for _,input := range inputs {
		matches := re.FindStringSubmatch(input)
		if len(matches) == 0 {
			continue
		} 

		pattern := Pattern {}
		for i,v := range matches[1] {
			pattern.values[i] = v == '#'
		}

		pattern.next = matches[2][0] == '#'
	

		patterns = append(patterns, pattern)
	}

	return initialState, patterns
}

func stabilized(state1 map[int] bool, state2 map[int] bool) bool {
	return toString(state1) == toString(state2)
}

func getPot(state map[int] bool, index int) bool {
	value,ok := state[index]
	return ok && value
}

func getPots(state map[int] bool, index int) [5]bool {
	return [5]bool{
		getPot(state, index - 2),
		getPot(state, index - 1),
		getPot(state, index),
		getPot(state, index + 1),
		getPot(state, index + 2),
	}
}

func getMinIndex(state map[int] bool) int {
	min := 999999999

	for k,v := range state {
		if v {
			min = au.MinInt(min, k)
		}
	}

	return min
}

func getMaxIndex(state map[int] bool) int {
	max := 0

	for k,v := range state {
		if v {
			max = au.MaxInt(max, k)
		}
	}

	return max
}

func generate(state map[int] bool, patterns []Pattern) map[int] bool {
	newState := make(map[int] bool)
	for k,v := range state {
		newState[k] = v
	}

	minIndex := getMinIndex(state) - 2
	maxIndex := getMaxIndex(state) + 2

	for i := minIndex; i <= maxIndex; i++ {
		pots := getPots(state, i)

		for _,pattern := range patterns {
			if pots == pattern.values {
				newState[i] = pattern.next
				break
			}
		}
	}

	return newState
}

func getPlantSum(state map[int] bool) int {
	count := 0

	for k,v := range state {
		if v {
			count += k
		}
	}

	return count
}

func toString(state map[int] bool) string {
	output := ""

	minIndex := getMinIndex(state)
	maxIndex := getMaxIndex(state)

	for i := minIndex; i <= maxIndex; i++ {
		if getPot(state, i) {
			output += "#"
		} else {
			output += "."
		}
	}
	
	return output
}

func main() {
	inputs := au.ReadInputAsStringArray("12")

	state, patterns := parseInputs(inputs)

	for i := 0; i < 50000000000; i++ {
		newState := generate(state, patterns)

		stabilized := stabilized(state, newState)
		if stabilized {
			fmt.Println("Stabilized after", i+1, "generations")

			currentPlantSum := getPlantSum(newState)
			lastPlantSum := getPlantSum(state)

			sumIncreasePerGeneration := currentPlantSum - lastPlantSum
			fmt.Println("Plant sum is increasing by", sumIncreasePerGeneration, "per generaton.")

			generationsLeft := 50000000000 - (i + 1)
			fmt.Println(generationsLeft, "generations left")

			fmt.Println("Final sum should be", currentPlantSum + generationsLeft * sumIncreasePerGeneration)

			break;
		}

		state = newState
	}
}
