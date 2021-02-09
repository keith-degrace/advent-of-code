package p03_1

import (
	"au"
	"fmt"
)

type Entry struct {
	a int
	b int
	c int
}

func readEntries() ([]Entry) {
	entries := []Entry{};

	for _, input := range au.ReadInputAsStringArray("03") {
		a := au.ToNumber(input[0:5]);
		b := au.ToNumber(input[5:10]);
		c := au.ToNumber(input[10:15]);

		entries = append(entries, Entry{a, b, c})
	}

	return entries;
}

func isValid(entry Entry) bool {
	return entry.a + entry.b > entry.c &&
				 entry.a + entry.c > entry.b &&
				 entry.b + entry.c > entry.a;
}

func Solve() {
	entries := readEntries();
	// entries := []Entry {Entry{5, 10, 25}}

	validCount := 0;

	for _, entry := range entries {
		if (isValid(entry)) {
			validCount++;
		}
	}

	fmt.Println(validCount)
}
