package p04_1

import (
	"au"
	"fmt"
	"regexp"
	"sort"
)

func testInputs() ([]string) {
	return []string {
		"aaaaa-bbb-z-y-x-123[abxyz]",
		"a-b-c-d-e-f-g-h-987[abcde]",
		"not-a-real-room-404[oarel]",
		"totally-real-room-200[decoy]",
	}
}

type Room struct {
	encryptedName string
	sectorId int
	storedChecksum string
}

func (this *Room) calculateChecksum() string {
	statMap := map[rune] int {}
	for _, letter := range this.encryptedName {
		if letter != '-' {
			statMap[letter]++
		}
	}

	type Stat struct {
		letter string
		count int
	}

	stats := []Stat {}
	for letter,count := range statMap {
		stats = append(stats, Stat{string(letter), count})
	}

	sort.Slice(stats, func(i, j int) bool {
		if (stats[i].count > stats[j].count) {
			return true
		} else if (stats[i].count < stats[j].count) {
			return false
		}

		return stats[i].letter < stats[j].letter
	})

	stats = stats[0:5]

	checksum := ""
	for _, stat := range stats {
		checksum += stat.letter
	}

	return checksum
}

func (this *Room) isReal() bool {
	return this.storedChecksum == this.calculateChecksum()
}

func parseInputs(inputs []string) []Room {
	rooms := []Room{}

	re := regexp.MustCompile("^([a-z-]*)-([0-9]*)\\[(.*)\\]$")

	for _, input := range inputs {
		matches := re.FindStringSubmatch(input)

		encryptedName := matches[1]
		sectorId := au.ToNumber(matches[2])
		storedChecksum := matches[3]
	
		rooms = append(rooms, Room {encryptedName, sectorId, storedChecksum})
	}

	return rooms
}

func getRealRooms(rooms []Room) ([]Room) {
	realRooms := []Room {}

	for _, room := range rooms {
		if (room.isReal()) {
			realRooms = append(realRooms, room)
		}
	}

	return realRooms
}

func getSectorIdSum(rooms []Room) (int) {
	sectorIdSum := 0

	for _, room := range rooms {
		sectorIdSum += room.sectorId;
	}

	return sectorIdSum
}

func Solve() {
	inputs := au.ReadInputAsStringArray("04")
	// inputs := testInputs()

	rooms := parseInputs(inputs)
	realRooms := getRealRooms(rooms)
	sectorIdSum := getSectorIdSum(realRooms)

	fmt.Println(sectorIdSum)
}
