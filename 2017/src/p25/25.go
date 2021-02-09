package p25

import (
	"au"
	"fmt"
	"math"
	"strings"
)

type Blueprint struct {
	beginStateId  string
	checksumAfter int
	states        map[string]State
}

type Condition struct {
	writeValue    int
	moveDirection string
	nextState     string
}

type State struct {
	id         string
	conditions [2]Condition
}

func load(input []string) Blueprint {

	blueprint := Blueprint{}
	blueprint.beginStateId = string(input[0][len(input[0])-2])
	blueprint.checksumAfter = au.ToNumber(strings.Split(input[1], " ")[5])
	blueprint.states = make(map[string]State)

	currentState := State{}
	currentCondition := 0

	for _, line := range input[3:] {
		if strings.HasPrefix(line, "In state ") {
			currentState.id = string(line[9])
		}

		if strings.HasPrefix(line, "  If the current value is") {
			currentCondition = au.ToNumber(string(line[len(line)-2]))
		}

		if strings.HasPrefix(line, "    - Write the value") {
			currentState.conditions[currentCondition].writeValue = au.ToNumber(string(line[len(line)-2]))
		}

		if strings.HasPrefix(line, "    - Move one slot to the ") {
			currentState.conditions[currentCondition].moveDirection = line[27 : len(line)-1]
		}

		if strings.HasPrefix(line, "    - Continue with state ") {
			currentState.conditions[currentCondition].nextState = string(line[len(line)-2])
		}

		if len(line) == 0 {
			blueprint.states[currentState.id] = currentState
			currentState = State{}
		}
	}

	blueprint.states[currentState.id] = currentState

	return blueprint
}

func printTape(tape map[int]int) {
	min := math.MaxInt16
	max := math.MinInt16

	for k, _ := range tape {
		min = au.MinInt(min, k)
		max = au.MaxInt(max, k)
	}

	for i := min; i <= max; i++ {
		value, _ := tape[i]
		fmt.Printf("%v ", value)
	}
	fmt.Println()
}

func Solve() {
	input := au.ReadInputAsStringArray("25")

	blueprint := load(input)

	tape := make(map[int]int)
	tape[0] = 0

	currentStateId := blueprint.beginStateId
	currentPosition := 0

	for i := 0; i < blueprint.checksumAfter; i++ {
		currentTapeValue, _ := tape[currentPosition]

		currentState, _ := blueprint.states[currentStateId]
		currentCondition := currentState.conditions[currentTapeValue]

		tape[currentPosition] = currentCondition.writeValue

		if currentCondition.moveDirection == "left" {
			currentPosition--
		} else {
			currentPosition++
		}

		// printTape(tape)

		currentStateId = currentCondition.nextState
	}

	checksum := 0
	for _, v := range tape {
		if v == 1 {
			checksum++
		}
	}
	fmt.Println(checksum)
}
