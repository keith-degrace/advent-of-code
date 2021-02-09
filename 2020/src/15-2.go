package main

import (
	"au"
	"fmt"
	"strings"
)

type Puzzle15_2 struct {
}

type Log struct {
	log        map[int][]int
	lastNumber int
}

func NewLog() Log {
	return Log{make(map[int][]int), -1}
}

func (s *Log) Add(number int, turn int) {
	newList := []int{}

	list, ok := s.log[number]
	if ok {
		newList = append(newList, list[len(list)-1])
	}

	newList = append(newList, turn)

	s.log[number] = newList

	s.lastNumber = number
}

func (s *Log) GetMagicLastCallNumber() int {
	list, ok := s.log[s.lastNumber]
	if !ok {
		return -1
	}

	if len(list) < 2 {
		return -1
	}

	return list[1] - list[0]
}

func (p Puzzle15_2) run() {
	input := "16,1,0,18,12,14,19"
	// input := "0,3,6"

	log := NewLog()
	turn := 0

	for _, entry := range strings.Split(input, ",") {
		turn++
		log.Add(au.ToNumber(entry), turn)
	}

	for turn < 30000000 {
		turn++
		var newNumber int

		lastCallNumber := log.GetMagicLastCallNumber()

		if lastCallNumber == -1 {
			newNumber = 0
		} else {
			newNumber = lastCallNumber
		}

		log.Add(newNumber, turn)
	}

	fmt.Printf("%vth number spoken is %v\n\n", turn, log.lastNumber)
}
