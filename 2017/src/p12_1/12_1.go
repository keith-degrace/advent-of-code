package p12_1

import (
	"au"
	"fmt"
	"regexp"
	"strings"
)

type Program struct {
	connected []*Program
}

func isConnectedToZero(programMap map[string][]string, program string, visiting map[string]bool) bool {
	if program == "0" {
		return true
	}

	visiting[program] = false

	for _, connectedProgram := range programMap[program] {
		if _, ok := visiting[connectedProgram]; ok {
			continue
		}

		if connectedProgram == "0" || isConnectedToZero(programMap, connectedProgram, visiting) {
			return true
		}
	}

	delete(visiting, program)

	return false
}

func Solve() {
	input := au.ReadInputAsStringArray("12")

	programMap := make(map[string][]string)

	lineRegex := regexp.MustCompile("([0-9]+) <-> (.*)")
	for _, line := range input {
		m := lineRegex.FindStringSubmatch(line)

		program := m[1]
		connected := strings.Split(m[2], ", ")

		programMap[program] = connected
	}

	count := 0

	for program := range programMap {
		visited := make(map[string]bool)

		if isConnectedToZero(programMap, program, visited) {
			count++
		}
	}

	fmt.Println(count)
}
