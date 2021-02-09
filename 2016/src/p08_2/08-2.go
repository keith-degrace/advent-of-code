package p08_2

import (
	"au"
	"regexp"
)

func testInputs() []string {
	return []string {
		"rect 3x2",
		"rotate column x=1 by 1",
		"rotate row y=0 by 4",
		"rotate column x=1 by 1",
	};
}

func rect(screen *au.Screen, width int, height int) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			screen.SetPixel(x, y, '#')
		}
	}
}

func rotateColumn(screen *au.Screen, column int, offset int) {
	values := []byte{};

	for y := 0; y < screen.Height; y++ {
		values = append(values, screen.GetPixel(column, y))
	}

	for y := 0; y < screen.Height; y++ {
		screen.SetPixel(column, (y + offset) % screen.Height, values[y])
	}
}

func rotateRow(screen *au.Screen, row int, offset int) {
	values := []byte{};

	for x := 0; x < screen.Width; x++ {
		values = append(values, screen.GetPixel(x, row))
	}

	for x := 0; x < screen.Width; x++ {
		screen.SetPixel((x + offset) % screen.Width, row, values[x])
	}
}

func applyOperation(screen *au.Screen, operation string) {
	rectRegExp := regexp.MustCompile("^rect ([0-9]+)x([0-9]+)")
	rectMatches := rectRegExp.FindStringSubmatch(operation)
	if (len(rectMatches) > 0) {
		width := au.ToNumber(rectMatches[1]);
		height := au.ToNumber(rectMatches[2]);

		rect(screen, width, height)
		return;
	}

	rotateColumnRegExp := regexp.MustCompile("^rotate column x=([0-9]+) by ([0-9]+)")
	rotateColumnMatches := rotateColumnRegExp.FindStringSubmatch(operation)
	if (len(rotateColumnMatches) > 0) {
		column := au.ToNumber(rotateColumnMatches[1]);
		offset := au.ToNumber(rotateColumnMatches[2]);

		rotateColumn(screen, column, offset)
		return;
	}

	rotateRowRegExp := regexp.MustCompile("^rotate row y=([0-9]+) by ([0-9]+)")
	rotateRowMatches := rotateRowRegExp.FindStringSubmatch(operation)
	if (len(rotateRowMatches) > 0) {
		row := au.ToNumber(rotateRowMatches[1]);
		offset := au.ToNumber(rotateRowMatches[2]);

		rotateRow(screen, row, offset)
		return;
	}
}

func applyOperations(screen *au.Screen, operations []string) {
	for _, operation := range operations {
		applyOperation(screen, operation)
	}
}

func Solve() {
	operations := au.ReadInputAsStringArray("08")
	// operations := testInputs()

	screen := au.NewScreen(50, 6)
	applyOperations(screen, operations)
	screen.Print()
}
