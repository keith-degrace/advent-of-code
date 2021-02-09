package p07_1

import (
	"au"
	"fmt"
	"regexp"
	"strings"
)

type Tower07_1 struct {
	name      string
	weight    int
	subTowers map[string]bool
}

func Solve() {
	input := au.ReadInputAsStringArray("07")

	towers := []Tower07_1{}

	re := regexp.MustCompile("(.+) \\(([0-9]+)\\)( -> (.*))?")
	for _, line := range input {
		m := re.FindStringSubmatch(line)

		tower := Tower07_1{}

		tower.name = m[1]
		tower.weight = au.ToNumber(m[2])

		if len(m) > 3 {
			tower.subTowers = make(map[string]bool)
			for _, subTowerName := range strings.Split(m[4], ", ") {
				tower.subTowers[subTowerName] = true
			}
		}

		towers = append(towers, tower)

		fmt.Println(tower)
	}

	for i := 0; i < len(towers); i++ {

		foundAsSubTower := false
		for j := 0; j < len(towers); j++ {
			if i == j {
				continue
			}

			_, ok := towers[j].subTowers[towers[i].name]
			if ok {
				foundAsSubTower = true
				break
			}

		}

		if !foundAsSubTower {
			fmt.Println(towers[i].name)
			break
		}
	}
}
