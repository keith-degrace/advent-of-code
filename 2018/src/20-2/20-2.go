package main

import (
	"au"
	"fmt"
)

type Segment struct {
	value string
	options [][]Segment
}

func getClosingBraceIndex(input string, startingIndex int) int {
	braceLevel := 1

	for i := startingIndex; i < len(input); i++ {
		if input[i] == '(' {
			braceLevel++
		} else if input[i] == ')' {
			if braceLevel == 1 {
				return i
			} else {
				braceLevel--
			}
		}
	}

	return -1
}

func getOptions(input string) []string {
	options := make([]string, 0)

	currentOption := ""
	for i := 0; i < len(input); i++ {
		if input[i] == '|' {
			// Finish off the current value
			options = append(options, currentOption)
			currentOption = ""
			continue
		}

		// Skip over sub-option sections
		if input[i] == '(' {
			closingBraceIndex := getClosingBraceIndex(input, i + 1)
			currentOption += input[i:closingBraceIndex + 1]
			i = closingBraceIndex
			continue
		}

		currentOption += string(input[i])
	}

	if len(currentOption) > 0 {
		options = append(options, currentOption)
	}

	return options
}

func parseInput(input string) []Segment {
	segments := make([]Segment, 0)

	currentValue := ""
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			// Finish off the current value
			segments = append(segments, Segment{value:currentValue})
			currentValue = ""
			
			// Add a segment for the optional section
			closingBraceIndex := getClosingBraceIndex(input, i + 1)
			subString := input[i+1:closingBraceIndex]

			optionalSegment := Segment{}
			for _,option := range getOptions(subString) {
				optionalSegment.options = append(optionalSegment.options, parseInput(option))
			}

			segments = append(segments, optionalSegment)

			i = closingBraceIndex
		} else {
			currentValue += string(input[i])
		}
	}

	if len(currentValue) > 0 {
		segments = append(segments, Segment{value:currentValue})
	}

	return segments
}

type Room struct {
	coord au.Coord
	northDoor bool
	westDoor bool
	southDoor bool
	eastDoor bool
}

func getNextDoorCoord(coord au.Coord, direction rune) au.Coord {
	if direction == 'N' {
		return coord.ShiftY(-1)
	} else if direction == 'S' {
		return coord.ShiftY(+1)
	} else if direction == 'W' {
		return coord.ShiftX(-1)
	} else if direction == 'E' {
		return coord.ShiftX(+1)
	}

	panic("Wee")
}

func createOrUpdateNextRoom(rooms *map[string] Room, coord au.Coord, direction rune) au.Coord {
	currentRoom := (*rooms)[coord.Key()]

	nextCoord := getNextDoorCoord(coord, direction)
	nextRoom, ok := (*rooms)[nextCoord.Key()]
	if !ok {
		nextRoom = Room{coord: nextCoord}
	}

	if direction == 'N' {
		currentRoom.northDoor = true
		nextRoom.southDoor = true
	} else if direction == 'S' {
		currentRoom.southDoor = true
		nextRoom.northDoor = true
	} else if direction == 'W' {
		currentRoom.westDoor = true
		nextRoom.eastDoor = true
	} else if direction == 'E' {
		currentRoom.eastDoor = true
		nextRoom.westDoor = true
	}

	(*rooms)[nextCoord.Key()] = nextRoom
	(*rooms)[coord.Key()] = currentRoom

	return nextCoord
}

func createRooms(rooms *map[string] Room, segments []Segment, coord au.Coord) {
	for _,segment := range segments {
		for _,direction := range segment.value {
			coord = createOrUpdateNextRoom(rooms, coord, direction)
		}

		for _,option := range segment.options {
			createRooms(rooms, option, coord)
		}
	}
}

func getBounds(rooms *map[string] Room, x int, y int) (int, int, int, int) {
	minX := 999999
	minY := 999999
	maxX := 0
	maxY := 0

	for _,room := range *rooms {
		minX = au.MinInt(minX, room.coord.X())
		minY = au.MinInt(minY, room.coord.Y())
		maxX = au.MaxInt(maxX, room.coord.X())
		maxY = au.MaxInt(maxY, room.coord.Y())
	}

	return minX, minY, maxX, maxY
}

func printRoom(screen *au.DynamicScreen, room *Room) {
	screenX := room.coord.X() * 2 + 1
	screenY := room.coord.Y() * 2 + 1

	if room.coord.EqualsXY(0, 0) {
		screen.SetPixel(screenX, screenY, 'X')
	} else {
		screen.SetPixel(screenX, screenY, ' ')
	}

	screen.SetPixel(screenX - 1, screenY - 1, '#')
	screen.SetPixel(screenX - 1, screenY + 1, '#')
	screen.SetPixel(screenX + 1, screenY + 1, '#')
	screen.SetPixel(screenX + 1, screenY - 1, '#')

	if room.northDoor {
		screen.SetPixel(screenX, screenY - 1, ' ')
	} else {
		screen.SetPixel(screenX, screenY - 1, '#')
	}

	if room.southDoor {
		screen.SetPixel(screenX, screenY + 1, ' ')
	} else {
		screen.SetPixel(screenX, screenY + 1, '#')
	}

	if room.westDoor {
		screen.SetPixel(screenX - 1, screenY, ' ')
	} else {
		screen.SetPixel(screenX - 1, screenY, '#')
	}

	if room.eastDoor {
		screen.SetPixel(screenX + 1, screenY, ' ')
	} else {
		screen.SetPixel(screenX + 1, screenY, '#')
	}
}

func printMap(rooms *map[string] Room, path []au.Coord) {
	screen := au.NewDynamicScreen()

	for _,room := range *rooms {
		printRoom(screen, &room)
	}

	for _,coord := range path {
		x := coord.X() * 2 + 1
		y := coord.Y() * 2 + 1
	
		screen.SetPixel(x, y, '.')
	}

	screen.Print()
}

func getNeighbors(rooms *map[string] Room) func (coord au.Coord) []au.Coord {
	return func (coord au.Coord) []au.Coord {
		children := []au.Coord{}

		room := (*rooms)[coord.Key()]

	if room.northDoor {
			children = append(children, coord.ShiftY(-1))
	}

	if room.southDoor {
			children = append(children, coord.ShiftY(+1))
	}

	if room.westDoor {
			children = append(children, coord.ShiftX(-1))
	}

	if room.eastDoor {
			children = append(children, coord.ShiftX(+1))
	}

	return children
}
		}


func getRoomCount(rooms *map[string] Room) int {
	count := 0

	origin := au.NewCoord(0, 0)

	for _,room := range *rooms {
		path := au.GetShortestPath(origin, room.coord, getNeighbors(rooms))
		if len(path) >= 1000 {
			count++
		}
	}

	return count
}

func main() {
	input := au.ReadInputAsString("20")
	input = input[1:len(input)-1]

	segments := parseInput(input)

	rooms := make(map[string] Room)

	firstRoom := Room{}
	rooms[firstRoom.coord.Key()] = firstRoom

	createRooms(&rooms, segments, firstRoom.coord)

	fmt.Println(getRoomCount(&rooms))
}
