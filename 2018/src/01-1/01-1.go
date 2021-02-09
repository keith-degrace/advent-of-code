package main

import (
	"fmt"
	"au"
)

func main() {
	frequency := 0

	for _, frequencyChange := range au.ReadInputAsNumberArray("01") {
		frequency += frequencyChange
	}

	fmt.Println(frequency)
}
