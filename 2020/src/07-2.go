package main

import (
	"au"
	"fmt"
	"regexp"
	"strconv"
)

type Puzzle07_2 struct {
}

func (p Puzzle07_2) getContainedBagCount(tree map[string]map[string]int, bag string) int {
	count := 0

	containedBags, ok := tree[bag]
	if !ok {
		return 0
	}

	for containedBag, containedBagCount := range containedBags {
		count += containedBagCount + containedBagCount*p.getContainedBagCount(tree, containedBag)
	}

	return count
}

func (p Puzzle07_2) run() {
	input := au.ReadInputAsStringArray("07")

	tree := make(map[string]map[string]int)

	r1 := regexp.MustCompile("^(.+) bags contain (.*).$")
	r2 := regexp.MustCompile("([0-9]+) (.+?) bag")

	for _, rule := range input {
		m1 := r1.FindStringSubmatch(rule)

		bag := m1[1]

		for _, m2 := range r2.FindAllStringSubmatch(m1[2], -1) {
			count, _ := strconv.Atoi(m2[1])
			containedBag := m2[2]

			_, ok := tree[bag]
			if !ok {
				tree[bag] = make(map[string]int)
			}

			tree[bag][containedBag] = count
		}
	}

	fmt.Println(p.getContainedBagCount(tree, "shiny gold"))
}
