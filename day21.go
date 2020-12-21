package main

import (
	"io"
	"sort"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(21, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day21Part2(inputReader)
	}
	return day21Part1(inputReader)
})

func day21Part1(inputReader io.Reader) interface{} {
	allergenToIngredients := make(map[string]map[string]bool)
	ingredientCounts := make(map[string]int)

	day21Parse(inputReader, func(ingredients map[string]bool, allergens []string) {
		for ingredient := range ingredients {
			ingredientCounts[ingredient]++
		}

		day21ReduceAllergenTheories(allergenToIngredients, ingredients, allergens)
	})

	for _, ingredients := range allergenToIngredients {
		for ingredient := range ingredients {
			delete(ingredientCounts, ingredient)
		}
	}

	result := 0

	for _, count := range ingredientCounts {
		result += count
	}

	return result
}

func day21Part2(inputReader io.Reader) interface{} {
	allergenToIngredients := make(map[string]map[string]bool)

	day21Parse(inputReader, func(ingredients map[string]bool, allergens []string) {
		day21ReduceAllergenTheories(allergenToIngredients, ingredients, allergens)
	})

	allergenToSingleIngredient := day21ResolveAllergensToSingleIngredient(allergenToIngredients)

	var allergens []string

	for allergen := range allergenToSingleIngredient {
		allergens = append(allergens, allergen)
	}

	sort.Strings(allergens)

	var ingredients []string

	for _, allergen := range allergens {
		ingredient := allergenToSingleIngredient[allergen]
		ingredients = append(ingredients, ingredient)
	}

	return strings.Join(ingredients, ",")
}

func day21Parse(inputReader io.Reader, visitFn func(ingredients map[string]bool, allergens []string)) {
	aocutil.VisitStrings(inputReader, func(v string) {
		split := strings.SplitN(v, " (contains ", 2)
		ingredients := strings.Fields(split[0])
		allergens := strings.Split(split[1][:len(split[1])-1], ", ")
		ingredientsSet := make(map[string]bool, len(ingredients))
		for _, ingredient := range ingredients {
			ingredientsSet[ingredient] = true
		}
		visitFn(ingredientsSet, allergens)
	})
}

func day21ReduceAllergenTheories(allergenToIngredients map[string]map[string]bool, ingredients map[string]bool, allergens []string) {
	for _, allergen := range allergens {
		if possibleIngredients, ok := allergenToIngredients[allergen]; !ok {
			ingredientsCopy := make(map[string]bool, len(ingredients))
			for ingredient := range ingredients {
				ingredientsCopy[ingredient] = true
			}
			allergenToIngredients[allergen] = ingredientsCopy
		} else {
			for ingredient := range possibleIngredients {
				if !ingredients[ingredient] {
					delete(allergenToIngredients[allergen], ingredient)
				}
			}
		}
	}
}

func day21ResolveAllergensToSingleIngredient(allergenToIngredients map[string]map[string]bool) map[string]string {
	allResolved := func() bool {
		for _, ingredients := range allergenToIngredients {
			if len(ingredients) > 1 {
				return false
			}
		}
		return true
	}

	any := func(set map[string]bool) string {
		for elem := range set {
			return elem
		}
		panic("no elements")
	}

	for !allResolved() {
		for allergen, ingredients := range allergenToIngredients {
			if len(ingredients) == 1 {
				ingredient := any(ingredients)
				for allergen2, ingredients2 := range allergenToIngredients {
					if allergen2 == allergen {
						continue
					}
					delete(ingredients2, ingredient)
				}
			}
		}
	}

	result := make(map[string]string)
	for allergen, ingredients := range allergenToIngredients {
		result[allergen] = any(ingredients)
	}

	return result
}
