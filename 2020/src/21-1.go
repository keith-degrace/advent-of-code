package main

import (
	"au"
	"fmt"
	"strings"
)

type Puzzle21_1 struct {
}

type Food21_1 struct {
	ingredients []string
	allergens   []string
}

func (p Puzzle21_1) parse(input []string) []Food21_1 {
	foods := []Food21_1{}

	for _, line := range input {
		parts := strings.Split(line, " (contains ")

		ingredients := strings.Split(parts[0], " ")

		allergens := []string{}
		if len(parts) == 2 {
			allergens = strings.Split(parts[1][:len(parts[1])-1], ", ")
		}

		foods = append(foods, Food21_1{ingredients, allergens})
	}

	return foods
}

func (p Puzzle21_1) getAllAllergens(foods []Food21_1) map[string]map[string]bool {
	candidateIngredientMap := make(map[string]map[string]int)

	for _, food := range foods {
		for _, allergen := range food.allergens {
			if _, ok := candidateIngredientMap[allergen]; !ok {
				candidateIngredientMap[allergen] = make(map[string]int)
			}

			for _, ingredient := range food.ingredients {
				candidateIngredientMap[allergen][ingredient]++
			}
		}
	}

	refinedCandidateIngredientMap := make(map[string]map[string]bool)
	for allergen, candidateIngredients := range candidateIngredientMap {
		refinedCandidateIngredientMap[allergen] = make(map[string]bool)

		allergenFoodCount := 0
		for _, food := range foods {
			for _, foodAllergen := range food.allergens {
				if allergen == foodAllergen {
					allergenFoodCount++
					break
				}
			}
		}

		for candidateIngredient, count := range candidateIngredients {
			if count == allergenFoodCount {
				refinedCandidateIngredientMap[allergen][candidateIngredient] = true
			}

		}
	}

	return refinedCandidateIngredientMap
}

func (p Puzzle21_1) getDefiniteAllergen(allergenMap map[string]map[string]bool) string {
	for allergen, ingredients := range allergenMap {
		if len(ingredients) == 1 {
			return allergen
		}
	}

	return ""
}

func (p Puzzle21_1) run() {
	input := au.ReadInputAsStringArray("21")

	foods := p.parse(input)

	allergenMap := p.getAllAllergens(foods)

	finalMap := make(map[string]string)

	for {
		definiteAllergen := p.getDefiniteAllergen(allergenMap)
		if len(definiteAllergen) == 0 {
			break
		}

		for ingredient, _ := range allergenMap[definiteAllergen] {
			for allergen, ingredients := range allergenMap {
				if allergen == definiteAllergen {
					continue
				}

				delete(ingredients, ingredient)
			}

			delete(allergenMap, definiteAllergen)
			finalMap[ingredient] = definiteAllergen
		}
	}

	count := 0
	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			if _, ok := finalMap[ingredient]; !ok {
				count++
			}
		}
	}

	fmt.Println(count)
}
