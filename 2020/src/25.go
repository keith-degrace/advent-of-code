package main

import "fmt"

type Puzzle25_1 struct {
}

func (p Puzzle25_1) transform(value int, subjectNumber int, loopSize int) int {
	for i := 0; i < loopSize; i++ {
		value = (value * subjectNumber) % 20201227
	}

	return value
}

func (p Puzzle25_1) findLoopSize(publicKey int) int {
	value := 1
	for i := 0; ; i++ {
		value = p.transform(value, 7, 1)
		if publicKey == value {
			return i + 1
		}
	}
}

func (p Puzzle25_1) run() {
	cardPublicKey, doorPublicKey := 15113849, 4206373
	// cardPublicKey, doorPublicKey := 5764801, 17807724

	doorLoopSize := p.findLoopSize(doorPublicKey)

	fmt.Println(p.transform(1, cardPublicKey, doorLoopSize))
}
