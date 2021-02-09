package p02_2

import (
	"au"
	"fmt"
)

func applyInstructions(currentButton string, instruction string) (string) {

	moveTable := map[string] string {
		"1D": "3",
		
		"2R": "3",
		"2D": "6",

		"3U": "1",
		"3D": "7",
		"3L": "2",
		"3R": "4",

		"4L": "3",
		"4D": "8",

		"5R": "6",

		"6U": "2",
		"6D": "A",
		"6L": "5",
		"6R": "7",

		"7U": "3",
		"7D": "B",
		"7L": "6",
		"7R": "8",

		"8U": "4",
		"8D": "C",
		"8L": "7",
		"8R": "9",

		"9L": "8",

		"AR": "B",
		"AU": "6",

		"BU": "7",
		"BD": "D",
		"BL": "A",
		"BR": "C",

		"CU": "8",
		"CL": "B",

		"DU": "B",
	}


	for _, move := range instruction {
		moveKey := currentButton + string(move);

		newButton, ok := moveTable[moveKey]
		if ok {
			currentButton = newButton
		}
	}

	return currentButton
}

func Solve() {
	instructions := au.ReadInputAsStringArray("02");
	// instructions := []string { "ULL", "RRDDD", "LURDL", "UUUUD" };

	currentButton := "5";

	code := []string{};

	for _, instruction := range instructions {
		currentButton = applyInstructions(currentButton, instruction)
		code = append(code, currentButton);
	}

	fmt.Println(code);
}
