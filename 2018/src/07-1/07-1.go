package main

import (
	"au"
	"fmt"
	"regexp"
	"sort"
	"strings"
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

func parseInputs(inputs []string) map[string] []string {
	steps := map[string] []string{}

	re := regexp.MustCompile("Step (.) must be finished before step (.) can begin.")

	for _, input := range inputs {
		matches := re.FindStringSubmatch(input)

		dependency := matches[1]
		step := matches[2]

		steps[step] = append(steps[step], dependency)
	}

	return steps
}

func topoSort(steps map[string] []string) string {

	// L ← Empty list that will contain the sorted elements
	L := []string{}

	// S ← Set of all nodes with no incoming edge
	S := []string{}
	{
		allDependencies := map[string] bool{}
		for _, dependencies  := range steps {
			for _, dependency := range dependencies {
				allDependencies[dependency] = true
			}
		}

		for dependency := range allDependencies {
			_, ok := steps[dependency]
			if !ok {
				S = append(S, dependency)
			}
		}
	}

	// while S is non-empty do
	for len(S) > 0 {
		// We want tie breakers to be alphabetical.
		sort.Strings(S)

		// remove a node n from S
		var n string
		n, S = S[0], S[1:]

		// add n to tail of L
		L = append(L, n)

		// for each node m with an edge e from n to m do
		for m, dependencies := range steps {
			for index, dependency := range dependencies {
				if dependency == n {
					// remove edge e from the graph
					steps[m] = append(steps[m][:index], steps[m][index+1:]...)

					// if m has no other incoming edges then
					if len(steps[m]) == 0 {
						// insert m into S
						S = append(S, m)
					}
					break;
				}
			}
		}
	}

	return strings.Join(L, "")
}

func main() {
	inputs := au.ReadInputAsStringArray("07")
	//inputs := testInputs();

	steps := parseInputs(inputs)
	fmt.Println(topoSort(steps))
}
