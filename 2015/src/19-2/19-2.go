package main

import (
	"au"
	"container/list"
	"fmt"
	"regexp"
	"sort"
)

func testInputs() []string {
	return []string {
		"e => H",
		"e => O",
		"H => HO",
		"H => OH",
		"O => HH",
		"",
		"HOHOHO",
	}
}

type Transform struct {
	from string
	to string
}
type Transforms []Transform

func parseInputs(inputs []string) (Transforms, string) {
	transforms := make(Transforms, 0)

	re := regexp.MustCompile("(.*) => (.*)")

	for _,input := range inputs {
		matches := re.FindStringSubmatch(input)
		if len(matches) > 0 {
			to := matches[1]
			from := matches[2]

			transforms = append(transforms, Transform{from, to})
		}
	}

	targetMolecule := inputs[len(inputs) - 1]

	return transforms, targetMolecule
}

type OpenSet struct {
	list *list.List
	set map [string] bool
}

func newOpenSet() OpenSet {
	return OpenSet {
		list.New(),
		make(map [string] bool),
	}
}

func (this *OpenSet) push(molecule string) {
	this.list.PushFront(molecule)
	this.set[molecule] = true
}

func (this *OpenSet) pop() string {
	subtreeRootElement := this.list.Front()
	this.list.Remove(subtreeRootElement)
	molecule := subtreeRootElement.Value.(string)

	delete(this.set, molecule)

	return molecule 
}

func (this *OpenSet) isEmpty() bool {
	return this.list.Len() == 0
}

func (this *OpenSet) has(molecule string) bool {
	_, ok := this.set[molecule]
	return ok
}

type ClosedSet map [string] bool

func addToClosedSet(closedSet ClosedSet, molecule string) {
	closedSet[molecule] = true
}

func isInClosedSet(closedSet ClosedSet, molecule string) bool {
	_, ok := closedSet[molecule]
	return ok
}

type byAscendingLength []string

func (s byAscendingLength) Len() int {
	return len(s)
}
func (s byAscendingLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byAscendingLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func getPermutations(molecule string, transforms Transforms) []string {
	if len(molecule) == 0 {
		return make([]string, 0)
	}

	permutations := make([]string, 0)

	for _,transform := range transforms {
		if len(molecule) < len(transform.from) {
			continue
		}
	
		for index := 0; index < len(molecule) - len(transform.from) + 1; index++ {
			substringStartIndex := index
			substringEndIndex := index + len(transform.from)
	
			substring := molecule[substringStartIndex:substringEndIndex]
			if substring == transform.from {
				preString := molecule[:substringStartIndex]
				postString := molecule[index +  + len(transform.from):]
				permutations = append(permutations, preString + transform.to + postString)
			}
		}
	}

	sort.Sort(byAscendingLength(permutations))

	return permutations
}

func getStepCount(molecule string, meta map[string] string) int {
	count := 0

	var ok bool
	current := molecule

	for {
		current, ok = meta[current]
		if !ok {
			break;
		}

		count++
	}

	return count - 1
}

func process(startingMolecule string, targetMolecule string, transforms Transforms) int {
	openSet := newOpenSet()
	closedSet := make(ClosedSet)

	meta := map[string] string{}

	root := startingMolecule
	meta[root] = ""
	openSet.push(root)

	bestStepCount := 999999999
	shortestSoFar := len(startingMolecule)

	for !openSet.isEmpty() {
		subtreeRoot := openSet.pop()

		if len(subtreeRoot) > shortestSoFar {
			continue
		}

		for _,child := range getPermutations(subtreeRoot, transforms) {
			if child == targetMolecule {
				meta[child] = subtreeRoot
				bestStepCount = au.MinInt(bestStepCount, getStepCount(child, meta))
			}

			if isInClosedSet(closedSet, child) {
				continue
			}

			if !openSet.has(child) {
				meta[child] = subtreeRoot
				openSet.push(child)
				shortestSoFar = au.MinInt(shortestSoFar, len(child))
			}
		}

		addToClosedSet(closedSet, subtreeRoot)
	}

	return bestStepCount
}

func main() {
	inputs := au.ReadInputAsStringArray("19")
	// inputs = testInputs()

	transforms, targetMolecule := parseInputs(inputs)

	result := process(targetMolecule, "e", transforms)
	fmt.Println(result)
}
