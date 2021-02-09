package p12_2

import (
	"au"
	"fmt"
	"regexp"
	"strings"
)

func Solve() {
	input := au.ReadInputAsStringArray("12")

	groups := []map[string]bool{}

	lineRegex := regexp.MustCompile("([0-9]+) <-> (.*)")
	for _, line := range input {
		m := lineRegex.FindStringSubmatch(line)

		program := m[1]
		connectedPrograms := strings.Split(m[2], ", ")

		group := make(map[string]bool)
		group[program] = true

		for _, connectedProgram := range connectedPrograms {
			group[connectedProgram] = true
		}

		groups = append(groups, group)
	}

	// Just keep merging groups until there are no more intersecting ones
	for {
		intersecting := false

		for i := 0; i < len(groups)-1; i++ {

			for j := i + 1; j < len(groups); j++ {

				// Go through all of group i and look for a conflict with group j
				for program1 := range groups[i] {

					if _, ok := groups[j][program1]; ok {

						// Merge group j into group i
						for k, v := range groups[j] {
							groups[i][k] = v
						}

						// Empty out group j
						groups[j] = make(map[string]bool)

						intersecting = true
						break
					}
				}
			}
		}

		if !intersecting {
			break
		}
	}

	// Count all non-empty groups.
	count := 0
	for _, group := range groups {
		if len(group) > 0 {
			count++
		}
	}

	fmt.Println(count)
}
