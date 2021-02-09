package main

import (
	"au"
	"fmt"
	"regexp"
	"strings"
)

type Puzzle19_1 struct {
}

func (p Puzzle19_1) parse(input []string) (map[int]Rule19_1, []string) {
	lineIndex := 0

	// Load all rules
	charRules := make(map[int]*CharRule19_1)
	sequenceRules := make(map[int]*SequenceRule19_1)
	sequenceRuleSetIds := make(map[int][][]int)

	for len(input[lineIndex]) != 0 {
		charRuleRegex := regexp.MustCompile("([0-9]+): \"(.)\"")

		charRuleMatch := charRuleRegex.FindStringSubmatch(input[lineIndex])
		if len(charRuleMatch) > 0 {
			rule := new(CharRule19_1)
			rule.id = au.ToNumber(charRuleMatch[1])
			rule.char = charRuleMatch[2][0]
			charRules[rule.id] = rule
		} else {
			sequenceRuleRegex := regexp.MustCompile("([0-9]+): (.*)")

			sequenceRuleMatch := sequenceRuleRegex.FindStringSubmatch(input[lineIndex])
			au.Assert(len(sequenceRuleMatch) > 0)

			ruleId := au.ToNumber(sequenceRuleMatch[1])

			sequenceString := sequenceRuleMatch[2]
			au.Assert(len(sequenceString) > 0)

			ruleSets := [][]int{[]int{}}
			for _, token := range strings.Split(sequenceString, " ") {
				switch token {
				case " ":
					continue

				case "|":
					ruleSets = append(ruleSets, []int{})

				default:
					ruleSets[len(ruleSets)-1] = append(ruleSets[len(ruleSets)-1], au.ToNumber(token))
				}
			}

			rule := new(SequenceRule19_1)
			rule.id = ruleId
			rule.cache = make(map[string]bool)
			sequenceRules[ruleId] = rule
			sequenceRuleSetIds[ruleId] = ruleSets
		}

		lineIndex++
	}

	valuesToTest := input[lineIndex+1:]

	ruleMap := make(map[int]Rule19_1)

	for _, charRule := range charRules {
		ruleMap[charRule.id] = charRule
	}

	for _, sequenceRule := range sequenceRules {
		sequenceRule.ruleSets = [][]Rule19_1{}

		for _, ruleSetIds := range sequenceRuleSetIds[sequenceRule.id] {

			ruleSet := []Rule19_1{}

			for _, ruleId := range ruleSetIds {
				rule, ok := charRules[ruleId]
				if ok {
					ruleSet = append(ruleSet, rule)
				} else {
					rule, ok := sequenceRules[ruleId]
					if ok {
						ruleSet = append(ruleSet, rule)
					}
				}
			}

			sequenceRule.ruleSets = append(sequenceRule.ruleSets, ruleSet)
		}

		ruleMap[sequenceRule.id] = sequenceRule
	}

	return ruleMap, valuesToTest
}

type Rule19_1 interface {
	getId() int
	matches(text string) bool
}

type CharRule19_1 struct {
	id   int
	char byte
}

func (r *CharRule19_1) getId() int {
	return r.id
}

func (r *CharRule19_1) matches(text string) bool {
	return len(text) == 1 && text[0] == r.char
}

type SequenceRule19_1 struct {
	id           int
	ruleSets     [][]Rule19_1
	minMatchSize int
	cache        map[string]bool
}

func (r *SequenceRule19_1) getId() int {
	return r.id
}

func (r *SequenceRule19_1) ruleSetMatches(ruleSet []Rule19_1, text string) bool {
	if len(text) < len(ruleSet) {
		return false
	}

	if len(ruleSet) == 1 {
		return ruleSet[0].matches(text)
	}

	for i := 1; i < len(text); i++ {
		if !ruleSet[0].matches(text[:i]) {
			continue
		}

		if r.ruleSetMatches(ruleSet[1:], text[i:]) {
			return true
		}
	}

	return false
}

func (r *SequenceRule19_1) matches(text string) bool {
	value, ok := r.cache[text]
	if ok {
		return value
	}

	for _, ruleSet := range r.ruleSets {
		if r.ruleSetMatches(ruleSet, text) {
			r.cache[text] = true
			return true
		}
	}

	r.cache[text] = false
	return false
}

func (p Puzzle19_1) run() {
	input := au.ReadInputAsStringArray("19")

	ruleMap, valuesToTest := p.parse(input)

	count := 0
	for _, valueToTest := range valuesToTest {
		match := ruleMap[0].matches(valueToTest)

		if match {
			count++
		}
	}

	fmt.Println(count)
}
