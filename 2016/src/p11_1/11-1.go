package p11_2

import (
	"au"
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"strconv"
)

func testInputs() []string {
	return []string{
		"The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.",
		"The second floor contains a hydrogen generator.",
		"The third floor contains a lithium generator.",
		"The fourth floor contains nothing relevant.",
	}
}

func testInputs2() []string {
	return []string{
		"The first floor contains a aa-compatible microchip and a bb-compatible microchip.",
		"The second floor contains a aa generator.",
		"The third floor contains a bb generator.",
	}
}

func testInputs3() []string {
	return []string{
		"The first floor contains a bb-compatible microchip.",
		"The second floor contains a bb generator.",
		"The third floor contains a aa-compatible microchip and a aa generator.",
	}
}

type Item struct {
	id string
	kind string
}

type Floor []Item

type Config struct {
	floors []Floor
	elevator int
}

func parseFloor(input string) Floor {
	floor := make(Floor, 0)

	generatorRegExp := regexp.MustCompile("[a-z]* generator")
	for _, item := range generatorRegExp.FindAllString(input, -1) {
		id := strings.ToUpper(item[:2])
		floor = append(floor, Item{id, "G"})
	}

	microchipRegExp := regexp.MustCompile("[a-z]*-compatible microchip")
	for _, item := range microchipRegExp.FindAllString(input, -1) {
		id := strings.ToUpper(item[:2])
		floor = append(floor, Item{id, "M"})
	}

	return floor;
}

func parseInputs(inputs []string) Config {
	config := Config{ make([]Floor, len(inputs)), 0 }

	for index, input := range inputs {
		config.floors[index] = parseFloor(input)
	}

	return config;
}

func printConfig(config Config) {
	for i := len(config.floors)-1; i >= 0; i-- {
		elevatorString := "."
		if i == config.elevator {
			elevatorString = "E"
		}

		fmt.Printf("F%v %v %v\n", i+1, elevatorString, config.floors[i])
	}
	fmt.Println(getConfigId(config))
	fmt.Println("")
}

func copyFloor(floor Floor) Floor {
	floorCopy := make(Floor, len(floor))
	copy(floorCopy, floor)
	return floorCopy;
}

func hasChip(floor Floor) bool {
	for _, item := range floor {
		if item.kind == "M" {
			return true
		}
	}

	return false
}

func hasGenerator(floor Floor) bool {
	for _, item := range floor {
		if item.kind == "G" {
			return true
		}
	}

	return false
}

func hasDisconnectedChip(floor Floor) bool {
	generatorMap := make(map[string] bool)
	for _, item := range floor {
		if item.kind == "G" {
			generatorMap[item.id] = true
		}
	}

	for _, item := range floor {
		if item.kind == "M" {
			_, ok := generatorMap[item.id]
			if !ok {
				return true
			}
		}
	}

	return false
}

func isFloorStable(floor Floor) bool {
	// If there are 0 or 1 item, it is safe.
	if len(floor) < 2 {
		return true
	}

	// Floor is always stable if there are no generators.
	if !hasGenerator(floor) {
		return true
	}

	// Floor is always stable if there are no chips.
	if !hasChip(floor) {
		return true
	}

	// Finally, we know we have at least one generator so if there are any disconnected
	// chips, we are in trouble.
	return !hasDisconnectedChip(floor)
}

func completed(config Config) bool {
	// If all floors except the top one are empty, we are done.
	for i := 0; i < len(config.floors) - 1; i++ {
		if len(config.floors[i]) != 0 {
			return false;
		}
	}

	return true
}

func getFloorId(floor Floor) string {
	if len(floor) == 0 {
		return ""
	}

	items := make([]string, 0)

	for _,item := range floor {
		items = append(items, fmt.Sprintf("%v_%v", item.kind, item.id))
	}
	
	sort.Strings(items)

	return strings.Join(items,",")
}

func getConfigId(config Config) string {
	var buffer bytes.Buffer

	buffer.WriteString("E")
	buffer.WriteString(strconv.Itoa(config.elevator))
	buffer.WriteString(" ")

	for index,floor := range config.floors {
		buffer.WriteString("F")
		buffer.WriteString(strconv.Itoa(index))
		buffer.WriteString("[")
		buffer.WriteString(getFloorId(floor))
		buffer.WriteString("] ")
	}

	return buffer.String()
}

func moveItems(config Config, fromFloor int, toFloor int, indicesToMove []int) (Config, error) {
	// Add the item to the target floor.  Is the floor still stable?
	newToFloor := copyFloor(config.floors[toFloor])
	for _,indexToMove := range indicesToMove {
		newToFloor = append(newToFloor, config.floors[fromFloor][indexToMove])
	}
	if !isFloorStable(newToFloor) {
		return Config{}, errors.New("Unstable")
	}

	// Remove item from this floor.  Is the floor still stable?
	newFromFloor := copyFloor(config.floors[fromFloor])
	for _,indexToMove := range indicesToMove {
		newFromFloor = append(newFromFloor[:indexToMove], newFromFloor[indexToMove+1:]...)
	}
	if !isFloorStable(newFromFloor) {
		return Config{}, errors.New("Unstable")
	}
	
	newFloors := make([]Floor, len(config.floors))
	for i := 0; i < len(config.floors); i++ {
		if i == fromFloor {
			newFloors[i] = newFromFloor
		} else if i == toFloor {
			newFloors[i] = newToFloor
		} else {
			newFloors[i] = config.floors[i]
		}
	}

	return Config { newFloors, toFloor }, nil
}

func getChildren(config Config) []Config {
	children := make([]Config, 0)

	// Priority 1: Try moving up two items up
	if config.elevator < len(config.floors) - 1 {
		// Evaluate pair moves
		for i := 0; i < len(config.floors[config.elevator]); i++ {
			for j := i + 1; j < len(config.floors[config.elevator]); j++ {
				child, err := moveItems(config, config.elevator, config.elevator+1, []int{j, i})
				if err == nil {
					children = append(children, child)
				}
			}
		}
	}

	// Priority 2: Try moving one item up
	if config.elevator < len(config.floors) - 1 {
		// Evaluate single moves
		for i := 0; i < len(config.floors[config.elevator]); i++ {
			child, err := moveItems(config, config.elevator, config.elevator+1, []int{i})
			if err == nil {
				children = append(children, child)
			}
		}
	}

	// Priority 3: Try moving just one item down
	if config.elevator > 0 {
		// Evaluate single moves
		for i := 0; i < len(config.floors[config.elevator]); i++ {
			child, err := moveItems(config, config.elevator, config.elevator-1, []int{i})
			if err == nil {
				children = append(children, child)
			}
		}
	}

	// Priority 4: Move two items down
	if config.elevator > 0 {
		// Evaluate pair moves
		for i := 0; i < len(config.floors[config.elevator]); i++ {
			for j := i + 1; j < len(config.floors[config.elevator]); j++ {
				child, err := moveItems(config, config.elevator, config.elevator-1, []int{j, i})
				if err == nil {
					children = append(children, child)
				}
			}
		}
	}
	
	return children
}

type OpenSet struct {
	list *list.List
	set map [string] bool
}

func newOpenSet() OpenSet {
	return OpenSet {
		list.New(),
		make(map [string] bool),
	}
}

func (this *OpenSet) push(config Config) {
	this.list.PushFront(config)

	configId := getConfigId(config)
	this.set[configId] = true
}

func (this *OpenSet) pop() Config {
	subtreeRootElement := this.list.Back()
	this.list.Remove(subtreeRootElement)
	config := subtreeRootElement.Value.(Config)

	configId := getConfigId(config)
	delete(this.set, configId)

	return config 
}

func (this *OpenSet) isEmpty() bool {
	return this.list.Len() == 0
}

func (this *OpenSet) has(config Config) bool {
	configId := getConfigId(config)
	_, ok := this.set[configId]
	return ok
}

type ClosedSet map [string] bool

func addToClosedSet(closedSet ClosedSet, config Config) {
	configId := getConfigId(config)
	closedSet[configId] = true
}

func isInClosedSet(closedSet ClosedSet, config Config) bool {
	configId := getConfigId(config)
	_, ok := closedSet[configId]
	return ok
}

func getStepCount(config Config, meta map[string] Config) int {
	count := 0

	var ok bool
	current := config

	for {
		current, ok = meta[getConfigId(current)]
		if !ok {
			break;
		}

		count++
	}

	return count - 1
}

func process(config Config) int {
	openSet := newOpenSet()
	closedSet := make(ClosedSet)

	meta := map[string] Config{}

	root := config
	meta[getConfigId(root)] = Config{}
	openSet.push(root)

	bestStepCount := 999999999

	for !openSet.isEmpty() {
		subtreeRoot := openSet.pop()

		if completed(subtreeRoot) {
			bestStepCount = au.MinInt(bestStepCount, getStepCount(subtreeRoot, meta))
		}

		for _,child := range getChildren(subtreeRoot) {
			if isInClosedSet(closedSet, child) {
				continue
			}

			if !openSet.has(child) {
				meta[getConfigId(child)] = subtreeRoot
				openSet.push(child)
			}
		}

		addToClosedSet(closedSet, subtreeRoot)
	}

	return bestStepCount
}

func Solve() {
	inputs := au.ReadInputAsStringArray("11")
	// inputs := testInputs()

	config := parseInputs(inputs)

	result := process(config)
	fmt.Println(result)
}

