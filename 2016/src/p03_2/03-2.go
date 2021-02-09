package p03_2

import (
	"au"
	"fmt"
	"regexp"
)

func testInputs() ([]string) {
	return []string {
		"101 301 501",
		"102 302 502",
		"103 303 503",
		"201 401 601",
		"202 402 602",
		"203 403 603",
	}
}

type Entry struct {
	a int
	b int
	c int
}

func parseInputs(inputs []string) ([]Entry) {
	entries := []Entry{}

	chunkCount := len(inputs) / 3

	re := regexp.MustCompile("[ ]*([0-9]*)[ ]*([0-9]*)[ ]*([0-9]*)[ ]*")

	for i := 0; i < chunkCount; i++ {
		line1 := inputs[i * 3 + 0]
		line2 := inputs[i * 3 + 1]
		line3 := inputs[i * 3 + 2]

		line1Matches := re.FindStringSubmatch(line1);
		line2Matches := re.FindStringSubmatch(line2);
		line3Matches := re.FindStringSubmatch(line3);

		a := au.ToNumber(line1Matches[1])
		b := au.ToNumber(line2Matches[1])
		c := au.ToNumber(line3Matches[1])
		entries = append(entries, Entry{a, b, c})

		a = au.ToNumber(line1Matches[2])
		b = au.ToNumber(line2Matches[2])
		c = au.ToNumber(line3Matches[2])
		entries = append(entries, Entry{a, b, c})

		a = au.ToNumber(line1Matches[3])
		b = au.ToNumber(line2Matches[3])
		c = au.ToNumber(line3Matches[3])
		entries = append(entries, Entry{a, b, c})
	}

	return entries
}

func isValid(entry Entry) bool {
	return entry.a + entry.b > entry.c &&
				 entry.a + entry.c > entry.b &&
				 entry.b + entry.c > entry.a
}

func Solve() {
	inputs := au.ReadInputAsStringArray("03")
	// inputs := testInputs()
	
	entries := parseInputs(inputs)

	validCount := 0

	for _, entry := range entries {
		if (isValid(entry)) {
			validCount++
		}
	}

	fmt.Println(validCount)
}
