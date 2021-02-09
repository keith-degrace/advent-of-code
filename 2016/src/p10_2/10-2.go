package p10_2

import (
	"au"
	"fmt"
	"regexp"
)

func testInputs() []string {
	return []string{
		"value 5 goes to bot 2",
		"bot 2 gives low to bot 1 and high to bot 0",
		"value 3 goes to bot 1",
		"bot 1 gives low to output 1 and high to bot 0",
		"bot 0 gives low to output 2 and high to output 0",
		"value 2 goes to bot 2",
	}
}

func initialAssignments(inputs []string) map [int] []int {
	bots := map [int] []int {}

	re := regexp.MustCompile("value ([0-9]*) goes to bot ([0-9]*)")

	for _, input := range inputs {
		matches := re.FindStringSubmatch(input)
		if len(matches) > 0 {
			value := au.ToNumber(matches[1])
			bot := au.ToNumber(matches[2])

			bots[bot] = append(bots[bot], value)
		}
	}

	return bots
}

func getLow(values []int) int {
	if len(values) == 1 || values[0] < values[1] {
		return values[0]
	} else {
		return values[1]
	}
}

func getHigh(values []int) int {
	if len(values) == 1 || values[1] < values[0] {
		return values[0]
	} else {
		return values[1]
	}
}

func applyActions(inputs []string, bots map [int] []int) {
	outputs := map [int] []int {}

	re := regexp.MustCompile("bot ([0-9]*) gives low to ([a-z]*) ([0-9]*) and high to ([a-z]*) ([0-9]*)")

	for (len(inputs) > 0) {
		rejects := []string{}

		for _, input := range inputs {
			matches := re.FindStringSubmatch(input)
			if len(matches) > 0 {
				giverBot := au.ToNumber(matches[1])
				lowReceiverType := matches[2]
				lowReceiver := au.ToNumber(matches[3])
				highReceiverType := matches[4]
				highReceiver := au.ToNumber(matches[5])

				if (len(bots[giverBot]) < 2) {
					rejects = append(rejects, input);
					continue;
				}
	
				loChip := getLow(bots[giverBot])
				hiChip := getHigh(bots[giverBot])

				if (lowReceiverType == "bot") {
					bots[lowReceiver] = append(bots[lowReceiver], loChip)
				} else {
					outputs[lowReceiver] = append(outputs[lowReceiver], loChip)
				}

				if (highReceiverType == "bot") {
					bots[highReceiver] = append(bots[highReceiver], hiChip)
				} else {
					outputs[highReceiver] = append(outputs[highReceiver], hiChip)
				}

				bots[giverBot] = []int {}
			}
		}

		inputs = rejects
	}

	fmt.Println(outputs[0][0] * outputs[1][0] * outputs[2][0])
}

func Solve() {
	inputs := au.ReadInputAsStringArray("10")
	// inputs := testInputs()

	bots := initialAssignments(inputs)

	applyActions(inputs, bots)
}
