package main

import (
	"fmt"
	"au"
)

func main() {
	frequencySeen := make(map[int]bool)
	frequencySeen[0] = true;

	frequency := 0
	for {
		for _, frequencyChange := range au.ReadInputAsNumberArray("01") {
			frequency += frequencyChange

			_, ok := frequencySeen[frequency]
			if ok {
				fmt.Println(frequency)
				return
			}

			frequencySeen[frequency] = true;
		}
	}
}
