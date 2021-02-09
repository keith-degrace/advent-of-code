package main

import (
	"au"
	"fmt"
	"regexp"
)

func testInputs() []string {
	return []string {
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin.",
	}
}

type Step struct {
	id string
	dependencies []string
	dependents []string
}

func parseInputs(inputs []string) map[string] Step {
	re := regexp.MustCompile("Step (.) must be finished before step (.) can begin.")

	allSteps := map[string] bool{}
	dependencyMap := map[string] []string{}
	dependentMap := map[string] []string{}

	for _, input := range inputs {
		matches := re.FindStringSubmatch(input)

		dependencyId := matches[1]
		stepId := matches[2]

		dependencyMap[stepId] = append(dependencyMap[stepId], dependencyId)
		allSteps[stepId] = true;

		dependentMap[dependencyId] = append(dependentMap[dependencyId], stepId)
		allSteps[dependencyId] = true;
	}

	steps := map[string] Step{}
	for step := range allSteps {
		steps[step] = Step {
			id: step,
			dependencies: dependencyMap[step],
			dependents: dependentMap[step],
		}
	}

	return steps
}

func topoSort(steps map[string] Step) []string {
	stepsToProcess := map[string] []string{}
	for _,step := range steps {
		stepsToProcess[step.id] = step.dependencies
	}

	// L ← Empty list that will contain the sorted elements
	L := []string{}

	// S ← Set of all nodes with no incoming edge
	S := []string{}
	{
		for _,step := range steps {
			if len(step.dependencies) == 0 {
				S = append(S, step.id)
			}
		}
	}

	// while S is non-empty do
	for len(S) > 0 {
		// remove a node n from S
		var n string
		n, S = S[0], S[1:]

		// add n to tail of L
		L = append(L, n)

		// for each node m with an edge e from n to m do
		for m, dependencies := range stepsToProcess {
			for index, dependency := range dependencies {
				if dependency == n {
					// remove edge e from the graph
					stepsToProcess[m] = append(stepsToProcess[m][:index], stepsToProcess[m][index+1:]...)

					// if m has no other incoming edges then
					if len(stepsToProcess[m]) == 0 {
						// insert m into S
						S = append(S, m)
					}
					break;
				}
			}
		}
	}

	return L
}

func getTime(step string) int {
	return int(step[0]) - 4
}

func allDependenciesDone(step Step, stepsDone map [string] bool) bool {
	for _,dependency := range step.dependencies {
		_, ok := stepsDone[dependency]
		if !ok {
			return false
		}
	}

	return true
}

func popNextStep(stepsToDo []string, allSteps map[string] Step, stepsDone map [string] bool) (string, []string) {
	if len(stepsToDo) > 0 {
		for i := 0; i < len(stepsToDo); i++ {
			stepToDo := allSteps[stepsToDo[i]]

			if (allDependenciesDone(stepToDo, stepsDone)) {
				return stepToDo.id, append(stepsToDo[:i], stepsToDo[i+1:]...)
			}
		}
	}

	return "", stepsToDo
}

func work(allSteps map[string] Step) int {
	type Worker struct {
		timeLeft int
		step string
	}

	stepsToDo := topoSort(allSteps)

	stepsDone := map [string] bool{};
	stepCount := len(stepsToDo)
		
	workers := make([]Worker, 5)

	time := 0

	for {

		// Advance work
		for i := 0; i<len(workers); i++ {
			if workers[i].timeLeft > 0 {
				workers[i].timeLeft--

				if (workers[i].timeLeft == 0) {
					stepsDone[workers[i].step] = true
					if (len(stepsDone) == stepCount) {
						return time;
					}
				}
			}
		}

		// Put idle workers to work.
		for i := 0; i<len(workers); i++ {
			if workers[i].timeLeft == 0 {
				var nextStepId string
				nextStepId, stepsToDo = popNextStep(stepsToDo, allSteps, stepsDone)
				if nextStepId != "" {
					workers[i].step = nextStepId
					workers[i].timeLeft = getTime(nextStepId)
				}
			}
    }
		
		time++
	}

	// Should never get here.
	return -1;
}

func main() {
	inputs := au.ReadInputAsStringArray("07")
	// inputs = testInputs();

	steps := parseInputs(inputs)

	time := work(steps)

	fmt.Println(time)
}
