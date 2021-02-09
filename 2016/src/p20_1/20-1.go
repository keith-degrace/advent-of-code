package p20_1

import (
	"au"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Range struct {
	lo int
	hi int
}

type Ranges []Range

func (s Ranges) Len() int {
	return len(s)
}

func (s Ranges) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Ranges) Less(i, j int) bool {
	if s[i].lo < s[j].lo {
		return true
	} else if s[i].lo > s[j].lo {
		return false
	} 
	
	return s[i].hi < s[j].hi
}


func testInputs() []string {
	return []string {
		"5-8",
		"0-2",
		"4-7",
	}
}

func parseInputs(inputs []string) []Range {
	ranges := make(Ranges, 0)

	for _,input := range inputs {
		parts := strings.Split(input, "-")

		lo := au.ToNumber(parts[0])
		hi := au.ToNumber(parts[1])

		ranges = append(ranges, Range { lo, hi })
	}

	sort.Sort(ranges)

	return ranges
}

func Solve() {
	fmt.Println("Starting\n")
	startTime := time.Now()

	inputs := au.ReadInputAsStringArray("20")
	// inputs  = testInputs()

	ranges := parseInputs(inputs)

	highestBlocked := ranges[0].hi

	for i := 1; i < len(ranges); i++ {
		if ranges[i].lo > (highestBlocked + 1) {
			fmt.Println(highestBlocked + 1)
			break;
		} else if ranges[i].hi > highestBlocked {
			highestBlocked = ranges[i].hi
		}
	}
	
	fmt.Println("\nCompleted in", time.Now().Sub(startTime))
}