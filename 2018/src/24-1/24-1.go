package main

import (
	"au"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func testInputs() []string {
	return []string {
		"Immune System:",
		"17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2",
		"989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3",
		"",
		"Infection:",
		"801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1",
		"4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4",
	}
}

type Group struct {
	id string
	kind string
	units int
	hp int
	attackPower int
	attackType string
	initiative int
	immunities map[string]bool
	weaknesses map[string]bool
}

func (this *Group) getEffectivePower() int {
	return this.units * this.attackPower
}

func (this *Group) isImmuneTo(attackingGroup Group) bool {
	_, ok := this.immunities[attackingGroup.attackType]
	return ok
}

func (this *Group) isWeakAgainst(attackingGroup Group) bool {
	_, ok := this.weaknesses[attackingGroup.attackType]
	return ok
}

func (this *Group) print() {
	fmt.Println(this, this.getEffectivePower())
}

type Groups map [string] Group

func (this *Groups) get(id string) Group {
	group, ok := (*this)[id]
	if !ok {
		panic(fmt.Sprintf("Group %v does not exist", id))
	}

	return group
}

func (this *Groups) toArray() []Group {
	array := []Group {}

	for _,group := range *this {
		array = append(array, group)
	}

	return array
}

func (this *Groups) print() {
	for _,group := range *this {
		group.print()
	}
	fmt.Println()
}

func parseImmunities(input string) map[string]bool {
	immunities := make(map[string]bool)

	re := regexp.MustCompile("immune to ([^;\\)]*)")

	matches := re.FindStringSubmatch(input)
	if len(matches) > 0 {
		for _,token := range strings.Split(matches[1], ",") {
			immunities[strings.TrimSpace(token)] = true
		}
	}

	return immunities
}

func parseWeaknesses(input string) map[string]bool {
	weaknesses := make(map[string]bool)

	re := regexp.MustCompile("weak to ([^;\\)]*)")

	matches := re.FindStringSubmatch(input)
	if len(matches) > 0 {
		for _,token := range strings.Split(matches[1], ",") {
			weaknesses[strings.TrimSpace(token)] = true
		}
	}

	return weaknesses
}

func parseInputs(inputs []string) Groups {
	groups := Groups{}

	kind := "Immune System Group"

	re := regexp.MustCompile("(.*) units each with (.*) hit points (.*)with an attack that does (.*) (.*) damage at initiative (.*)")

	idCounter := 1

	for i := 1; i < len(inputs); i++ {
		matches := re.FindStringSubmatch(inputs[i])
		if len(matches) > 0 {
			group := Group {
				id: fmt.Sprintf("%v %v", kind, idCounter),
				kind: kind,
				units: au.ToNumber(matches[1]),
				hp: au.ToNumber(matches[2]),
				attackPower: au.ToNumber(matches[4]),
				attackType: matches[5],
				initiative: au.ToNumber(matches[6]),
				immunities: parseImmunities(matches[3]),
				weaknesses: parseWeaknesses(matches[3]),
			}

			idCounter++
			groups[group.id] = group
		} else if inputs[i] == "Infection:" {
			kind = "Infection Group"
			idCounter = 1
		} else if len(inputs[i]) > 0 {
			panic(fmt.Sprintf("Parse Error: %v", inputs[i]))
		}
	}

	return groups
}

func getIds(groups []Group) []string {
	ids := []string {}

	for _,group := range groups {
		ids = append(ids, group.id)
	}

	return ids
}

func getDamage(attackingGroup Group, defendingGroup Group) int {
	if attackingGroup.units <= 0 {
		return 0
	}

	if defendingGroup.isImmuneTo(attackingGroup) {
		return 0
	}

	if defendingGroup.isWeakAgainst(attackingGroup) {
		return attackingGroup.getEffectivePower() * 2
	}

	return attackingGroup.getEffectivePower()
}

type bySelection []Group

func (s bySelection) Len() int {
	return len(s)
}

func (s bySelection) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s bySelection) Less(i, j int) bool {
	effectivePower1 := s[i].getEffectivePower()
	effectivePower2 := s[j].getEffectivePower()

	if effectivePower1 > effectivePower2 {
		return true
	} else if effectivePower1 < effectivePower2 {
		return false
	}

	return s[i].initiative > s[j].initiative
}

func getSelectionOrder(groups Groups) []Group {
	orderedGroups := groups.toArray()
	sort.Sort(bySelection(orderedGroups))
	return orderedGroups
}

type byTarget struct {
	groups []Group
	attackingGroup Group
}

func (s byTarget) Len() int {
	return len(s.groups)
}

func (s byTarget) Swap(i, j int) {
	s.groups[i], s.groups[j] = s.groups[j], s.groups[i]
}

func (s byTarget) Less(i, j int) bool {
	damage1 := getDamage(s.attackingGroup, s.groups[i])
	damage2 := getDamage(s.attackingGroup, s.groups[j])

	if damage1 > damage2 {
		return true
	} else if damage1 < damage2 {
		return false
	}

	effectivePower1 := s.groups[i].getEffectivePower()
	effectivePower2 := s.groups[j].getEffectivePower()

	if effectivePower1 > effectivePower2 {
		return true
	} else if effectivePower1 < effectivePower2 {
		return false
	}

	return s.groups[i].initiative > s.groups[j].initiative
}

func getTarget(attackingGroup Group, groups []Group) (*Group, []Group) {
	orderedGroups := make([]Group, 0)
	for _,group := range groups {
		if group.kind != attackingGroup.kind && getDamage(attackingGroup, group) > 0 {
			orderedGroups = append(orderedGroups, group)
		}
	}

	if len(orderedGroups) == 0 {
		return nil, groups
	}

	sort.Sort(byTarget{ orderedGroups, attackingGroup })
	target := &orderedGroups[0]

	remaining := make([]Group, 0)
	for _,group := range groups {
		if group.id == target.id {
			continue
		}

		remaining = append(remaining, group)
	}

	return target, remaining
}

type byAttack []Group

func (s byAttack) Len() int {
	return len(s)
}

func (s byAttack) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byAttack) Less(i, j int) bool {
	return s[i].initiative > s[j].initiative
}

func getAttackOrder(groups Groups) []string {
	orderedGroups := groups.toArray()
	sort.Sort(byAttack(orderedGroups))
	return getIds(orderedGroups)
}

func iterate(groups Groups) Groups {
	selections := map[string]string {}
	availableGroups := groups.toArray()

	for _,selector := range getSelectionOrder(groups) {
		target, remaining := getTarget(selector, availableGroups)
		if target != nil {
			selections[selector.id] = target.id
			availableGroups = remaining
		}
	}

	for _,attackerId := range getAttackOrder(groups) {
		targetId, ok := selections[attackerId]
		if !ok {
			continue
		}

		attacker := groups.get(attackerId)
		target := groups.get(targetId)

		damage := getDamage(attacker, target)
		
		target.units -= damage / target.hp
		groups[targetId] = target
	}

	remainingGroups := make(Groups)

	for _, group := range groups {
		if group.units > 0 {
			remainingGroups[group.id] = group
		}
	}

	return remainingGroups
}

func hasWinner(groups Groups) bool {
	hasImmuneSystemGroup := false
	hasInfectionGroup := false

	for _,group := range groups {
		if group.kind == "Immune System Group" && group.units > 0 {
			hasImmuneSystemGroup = true
		}

		if group.kind == "Infection Group" && group.units > 0 {
			hasInfectionGroup = true
		}
	}

	return hasInfectionGroup != hasImmuneSystemGroup
}

func getWinnerUnitCount(groups Groups) int {
	count := 0

	for _,group := range groups {
		if group.units > 0 {
			count += group.units
		}
	}

	return count
}

func main() {
	inputs := au.ReadInputAsStringArray("24")
	// inputs  = testInputs()

	groups := parseInputs(inputs)

	for !hasWinner(groups) {
		groups = iterate(groups)
	}
	
	fmt.Println(getWinnerUnitCount(groups))
}