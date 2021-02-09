package p07_2

import (
	"au"
	"fmt"
	"regexp"
	"strings"
)

type Tower07_2 struct {
	name      string
	weight    int
	subTowers map[string]*Tower07_2
}

func load(input []string) *Tower07_2 {

	towers := make(map[string]*Tower07_2)
	subTowerMap := make(map[string][]string)

	// Load all towers
	re := regexp.MustCompile("(.+) \\(([0-9]+)\\)( -> (.*))?")
	for _, line := range input {
		m := re.FindStringSubmatch(line)

		tower := new(Tower07_2)

		tower.name = m[1]
		tower.weight = au.ToNumber(m[2])

		if len(m) > 3 {
			subTowerMap[tower.name] = strings.Split(m[4], ", ")
		}

		towers[tower.name] = tower
	}

	// Link all towers with their sub towers
	for _, tower := range towers {
		tower.subTowers = make(map[string]*Tower07_2)

		for _, subTowerName := range subTowerMap[tower.name] {
			subTower, ok := towers[subTowerName]
			if ok {
				tower.subTowers[subTower.name] = subTower
			}
		}

		towers[tower.name] = tower
	}

	// Find the root tower
	var rootTower *Tower07_2
	for name1, tower1 := range towers {

		foundAsSubTower := false
		for name2, tower2 := range towers {
			if name1 == name2 {
				continue
			}

			_, ok := tower2.subTowers[name1]
			if ok {
				foundAsSubTower = true
				break
			}

		}

		if !foundAsSubTower {
			rootTower = tower1
			break
		}
	}

	return rootTower
}

func getTotalWeight(tower *Tower07_2) int {
	total := tower.weight

	for _, subTower := range tower.subTowers {
		total += getTotalWeight(subTower)
	}

	return total
}

func getOddTower(towers map[string]*Tower07_2) (*Tower07_2, int, int) {
	weightGroups := make(map[int][]*Tower07_2)

	for _, tower := range towers {
		weight := getTotalWeight(tower)
		weightGroups[weight] = append(weightGroups[weight], tower)
	}

	if len(weightGroups) == 1 {
		return nil, 0, 0
	}

	fmt.Println(weightGroups)

	weights := []int{}
	for weight := range weightGroups {
		weights = append(weights, weight)
	}

	fmt.Println(weights)

	if len(weightGroups[weights[0]]) == 1 {
		return weightGroups[weights[0]][0], weights[0], weights[1]
	} else {
		au.AssertIntsEqual(len(weightGroups[weights[1]]), 1)
		return weightGroups[weights[1]][0], weights[1], weights[0]
	}
}

func findBadWeight(tower *Tower07_2) bool {
	if len(tower.subTowers) == 0 {
		return false
	}

	oddTower, actualWeight, expectedWeight := getOddTower(tower.subTowers)
	if oddTower != nil {
		fmt.Printf("%v was expected to weigh %v but weighs %v\n", oddTower.name, expectedWeight, actualWeight)

		if !findBadWeight(oddTower) {
			// Did not find a bad weight in the sub towers so we are the problem.
			fmt.Println(oddTower.weight + (expectedWeight - actualWeight))
		}

		return true
	} else {
		// Sub towers look good, so we are the problem.  Report to the previous iteration and report the difference.
		return false
	}
}

func Solve() {
	input := au.ReadInputAsStringArray("07")

	rootTower := load(input)

	findBadWeight(rootTower)
}
