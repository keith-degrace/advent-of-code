package main

import (
	"au"
	"fmt"
	"sort"
)

func testInputs() []string {
	return []string {
		"/>-<\\",
		"|   |" ,
		"| /<+-\\",
		"| | | v",
		"\\>+</ |",
		"  |   ^",
		"  \\<->/",
	}
}

type Rails []string

type Cart struct {
	x int
	y int
	direction byte
	nextTurn byte
	removed bool
}

func getValue(inputs []string, x int, y int) byte {
	if y < 0 || y > len(inputs) - 1 {
		return ' '
	}
	
	if x < 0 || x > len(inputs[y]) - 1 {
		return ' '
	}

	return inputs[y][x]
}

func getRailValue(inputs []string, x int, y int) string {
	above := getValue(inputs, x, y-1)
	below := getValue(inputs, x, y+1)
	left := getValue(inputs, x-1, y)
	right := getValue(inputs, x+1, y)

	if (above == '|' || above == '/' || above == '\\' || above == '+') &&
	 	 (below == '|' || below == '/' || below == '\\' || below == '+') &&
	 	 (left  == '-' || left  == '/' || left  == '\\' || left  == '+') &&
	 	 (right == '-' || right == '/' || right == '\\' || right == '+') {
			return "+"
	}

	if (above == '|' || above == '/' || above == '\\' || above == '+') &&
	 	 (below == '|' || below == '/' || below == '\\' || below == '+') {
			return "|"
	}

	if (left  == '-' || left  == '/' || left  == '\\' || left  == '+') &&
	 	 (right == '-' || right == '/' || right == '\\' || right == '+') {
			return "-"
	}

	fmt.Println("Error!")
	return " "
}

func parseInputs(inputs []string) (Rails, []Cart) {
	rails := make(Rails, 0)
	carts := make([]Cart, 0)

	for y,input := range inputs {
		row := ""
		for x,v := range input {
			if v == '<' || v == '>' || v == '^' || v == 'v' {
				row += getRailValue(inputs, x, y)
				carts = append(carts, Cart{x, y, byte(v), 'L', false})
			} else {
				row += string(v)
			}
		}

		rails = append(rails, row)
	}

	return rails, carts
}

func handleIntersection(carts []Cart, index int) {
	// Turn it in its new direction
	if carts[index].nextTurn == 'L' {
		if carts[index].direction == '>' {
			carts[index].direction = '^'
		} else if carts[index].direction == '^' {
			carts[index].direction = '<'
		} else if carts[index].direction == '<' {
			carts[index].direction = 'v'
		} else {
			carts[index].direction = '>'
		}
	} else if carts[index].nextTurn == 'R' {
		if carts[index].direction == '>' {
			carts[index].direction = 'v'
		} else if carts[index].direction == 'v' {
			carts[index].direction = '<'
		} else if carts[index].direction == '<' {
			carts[index].direction = '^'
		} else {
			carts[index].direction = '>'
		}
	}

	// Update the next intersection turn.
	if carts[index].nextTurn == 'L' {
		carts[index].nextTurn = 'S'
	} else if carts[index].nextTurn == 'R' {
		carts[index].nextTurn = 'L'
	} else {
		carts[index].nextTurn = 'R'
	}
}

type byCart []Cart

func (s byCart) Len() int {
	return len(s)
}
func (s byCart) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byCart) Less(i, j int) bool {
	if s[i].y < s[j].y {
		return true
	} else if s[i].y > s[j].y {
		return false
	} 
	
	return s[i].x < s[j].x
}

func cleanUpCrashes(carts []Cart) {
	for {
		crashFound := false

		for i := 0; i < len(carts); i++ {
			if carts[i].removed {
				continue
			}

			for j := i + 1; j < len(carts); j++ {
				if carts[j].removed {
					continue
				}

				if carts[i].x == carts[j].x && carts[i].y == carts[j].y {
					carts[i].removed = true
					carts[j].removed = true
				}			
			}
		}

		if !crashFound {
			break;
		}
	}
}

func getActiveCartCount(carts []Cart) int {
	count := 0

	for _,cart := range carts {
		if !cart.removed {
			count++
		}
	}

	return count
}

func getLastActiveCart(carts []Cart) Cart {
	for _,cart := range carts {
		if !cart.removed {
			return cart
		}
	}

	return Cart{}
}

func iterate(rails Rails, carts []Cart) bool {
	for i := 0; i < len(carts); i++ {
		if carts[i].removed {
			continue
		}

		if carts[i].direction == '>' {
			nextRail := getValue(rails, carts[i].x + 1, carts[i].y)

			carts[i].x++

			if nextRail == '/' {
				carts[i].direction = '^'
			} else if nextRail == '\\' {
				carts[i].direction = 'v'
			} else if nextRail == '+' {
				handleIntersection(carts, i)
			}
		} else if carts[i].direction == 'v' {
			nextRail := getValue(rails, carts[i].x, carts[i].y + 1)

			carts[i].y++

			if nextRail == '/' {
				carts[i].direction = '<'
			} else if nextRail == '\\' {
				carts[i].direction = '>'
			} else if nextRail == '+' {
				handleIntersection(carts, i)
			}
		} else if carts[i].direction == '<' {
			nextRail := getValue(rails, carts[i].x - 1, carts[i].y)

			carts[i].x--

			if nextRail == '/' {
				carts[i].direction = 'v'
			} else if nextRail == '\\' {
				carts[i].direction = '^'
			} else if nextRail == '+' {
				handleIntersection(carts, i)
			}
		} else if carts[i].direction == '^' {
			nextRail := getValue(rails, carts[i].x, carts[i].y - 1)

			carts[i].y--

			if nextRail == '/' {
				carts[i].direction = '>'
			} else if nextRail == '\\' {
				carts[i].direction = '<'
			} else if nextRail == '+' {
				handleIntersection(carts, i)
			}
		}

		cleanUpCrashes(carts);
	}

	if getActiveCartCount(carts) == 1 {
		lastActiveCart := getLastActiveCart(carts)
		fmt.Printf("%v,%v\n", lastActiveCart.x, lastActiveCart.y)
		return false
	}

	return true
}

func printState(rails Rails, carts []Cart) {
	width := 0
	for _,row := range rails {
		width = au.MaxInt(width, len(row))
	}

	screen := au.NewScreen(width, len(rails))

	for y := range rails {
		for x := range rails[y] {
			screen.SetPixel(x, y, rails[y][x])
		}
	}

	for _,cart := range carts {
		if !cart.removed {
			screen.SetPixel(cart.x, cart.y, cart.direction)
		}
	}

	screen.Print()
}

func main() {
	inputs := au.ReadInputAsStringArray("13")
	// inputs = testInputs()

	rails,carts := parseInputs(inputs)
	// printState(rails, carts)

	for {
		sort.Sort(byCart(carts))

		if !iterate(rails, carts) {
			break;
		}
		
		// printState(rails, carts)
	}
}
