package p04_2

import (
	"au"
	"fmt"
	"regexp"
)

func testInputs() []string {
	return []string{
		"qzmt-zixmtkozy-ivhz-343[aaaaa]",
	}
}

type Room struct {
	encryptedName  string
	sectorId       int
	storedChecksum string
}

func decryptLetter(letter rune, sectorId int) string {
	if letter == '-' {
		return " "
	}

	asciiCode := int(letter)
	decryptedAsciiCode := (asciiCode+sectorId-97)%26 + 97

	return string(rune(decryptedAsciiCode))
}

func (this *Room) getDecryptedName() string {
	decryptedName := ""

	for _, encryptedLetter := range this.encryptedName {
		decryptedName += decryptLetter(encryptedLetter, this.sectorId)
	}

	return decryptedName
}

func parseInputs(inputs []string) []Room {
	rooms := []Room{}

	re := regexp.MustCompile("^([a-z-]*)-([0-9]*)\\[(.*)\\]$")

	for _, input := range inputs {
		matches := re.FindStringSubmatch(input)

		encryptedName := matches[1]
		sectorId := au.ToNumber(matches[2])
		storedChecksum := matches[3]

		rooms = append(rooms, Room{encryptedName, sectorId, storedChecksum})
	}

	return rooms
}

func Solve() {
	inputs := au.ReadInputAsStringArray("04")
	// inputs := testInputs()

	rooms := parseInputs(inputs)

	for _, room := range rooms {
		if room.getDecryptedName() == "northpole object storage" {
			fmt.Println(room.sectorId)
		}
	}
}
