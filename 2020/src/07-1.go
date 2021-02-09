package main

import (
	"au"
	"fmt"
	"regexp"
)

type Puzzle07_1 struct {
}

func (p Puzzle07_1) getHolders(tree map[string][]string, bag string) map[string]bool {
	outermostBags := make(map[string]bool)

	holders, ok := tree[bag]
	if ok {
		for _, holder := range holders {
			outermostBags[holder] = true

			for bag := range p.getHolders(tree, holder) {
				outermostBags[bag] = true
			}
		}
	}

	return outermostBags
}

func (p Puzzle07_1) run() {
	input := au.ReadInputAsStringArray("07")

	tree := make(map[string][]string)

	r1 := regexp.MustCompile("^(.+) bags contain (.*).$")
	r2 := regexp.MustCompile("[0-9]+ (.+?) bag")

	for _, rule := range input {
		m1 := r1.FindStringSubmatch(rule)

		bag := m1[1]

		for _, m2 := range r2.FindAllStringSubmatch(m1[2], -1) {
			containedBag := m2[1]

			tree[containedBag] = append(tree[containedBag], bag)
		}
	}

	fmt.Println(len(p.getHolders(tree, "shiny gold")))
}
