package day21

import (
	"fmt"
	"sort"
	"strings"

	"github.com/crillab/gophersat/bf"
)

type ingredient string
type allergen string

type food struct {
	ingredients map[ingredient]bool
	allergens   map[allergen]bool
}

var allIngredients map[ingredient]int
var allAllergens map[allergen]bool
var foods []food

func Part1(input string) int {
	solution := parseAndSolve(input)
	r := 0
	for i, _ := range allIngredients {
		if solution[fmt.Sprintf("%s-safe", i)] {
			r += allIngredients[i]
		}
	}
	return r
}

func Part2(input string) string {
	solution := parseAndSolve(input)

	// sort all the allergens
	var keys []string
	for a, _ := range allAllergens {
		keys = append(keys, string(a))
	}
	sort.Strings(keys)

	// emit r
	var r []string
	for _, a := range keys {
		for i, _ := range allIngredients {
			if solution[fmt.Sprintf("%s-%s", i, a)] {
				r = append(r, string(i))
			}
		}
	}
	return strings.Join(r, ",")
}

// Use a SAT solver to figure out which ingredient maps to which allergen. SAT solvers work with boolean variables, so
// we have to do some conversion. Let's say we have 4 ingredients (I1 thru I4) and 3 allergens (A1 thru A3), we can
// create the following constraints where In-Am are new boolean variables:
//
// each ingredient maps to zero or more allergens:
//   Unique(I1-A1, I1-A2, I1-A3, I1-safe)
//   Unique(I2-A1, I2-A2, I2-A3, I2-safe)
//   ...
//
// each allergen maps to one ingredient:
//   Unique(I1-A1, I2-A1, I3-A1, I4-A1)
//   Unique(I1-A2, I2-A2, I3-A2, I4-A2)
//   ...
//
// A food entry, such as: "I1 I2 I3 (contains A1 A2)" then gets mapped to:
//   Unique(I1-A1, I2-A1, I3-A1)
//   Unique(I1-A2, I2-A2, I3-A2)
func parseAndSolve(input string) map[string]bool {
	allIngredients = map[ingredient]int{}
	allAllergens = map[allergen]bool{}
	foods = []food{}

	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " (contains ")
		f := food{
			ingredients: map[ingredient]bool{},
			allergens:   map[allergen]bool{},
		}
		ingredients := strings.Split(l[0], " ")
		for _, v := range ingredients {
			f.ingredients[ingredient(v)] = true
			allIngredients[ingredient(v)]++
		}

		l[1] = strings.TrimSuffix(l[1], ")")
		allergens := strings.Split(l[1], ", ")
		for _, v := range allergens {
			f.allergens[allergen(v)] = true
			allAllergens[allergen(v)] = true
		}

		foods = append(foods, f)
	}

	model := bf.True

	for i, _ := range allIngredients {
		var vars []string
		for a, _ := range allAllergens {
			vars = append(vars, fmt.Sprintf("%s-%s", i, a))
		}
		// an ingredient is either safe or is tied to exactly one allergen
		vars = append(vars, fmt.Sprintf("%s-safe", i))
		model = bf.And(model, bf.Unique(vars...))
	}

	for a, _ := range allAllergens {
		var vars []string
		for i, _ := range allIngredients {
			vars = append(vars, fmt.Sprintf("%s-%s", i, a))
		}
		// an allergen is tied to exactly one ingredient
		model = bf.And(model, bf.Unique(vars...))
	}

	// convert foods into constraints
	for _, f := range foods {
		for a, _ := range f.allergens {
			var vars []string
			for i, _ := range f.ingredients {
				vars = append(vars, fmt.Sprintf("%s-%s", i, a))
			}
			model = bf.And(model, bf.Unique(vars...))
		}
	}

	solution := bf.Solve(model)
	if solution == nil {
		panic("unsat?")
	}
	return solution
}
