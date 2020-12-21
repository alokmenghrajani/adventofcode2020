package day21

import (
	"sort"
	"strings"
)

type ingredient string
type allergen string

type blah struct {
	ingredients map[ingredient]bool
	allergens   map[allergen]bool
}

func Part1(input string) int {
	allIngredients := map[ingredient]bool{}
	allAllergens := map[allergen]bool{}
	var blahs []blah
	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " (contains ")
		b := blah{
			ingredients: map[ingredient]bool{},
			allergens:   map[allergen]bool{},
		}
		ingredients := strings.Split(l[0], " ")
		for _, v := range ingredients {
			b.ingredients[ingredient(v)] = true
			allIngredients[ingredient(v)] = true
		}

		l[1] = l[1][0 : len(l[1])-1]
		allergens := strings.Split(l[1], ", ")
		for _, v := range allergens {
			b.allergens[allergen(v)] = true
			allAllergens[allergen(v)] = true
		}

		blahs = append(blahs, b)
	}

	safe := map[ingredient]bool{}
	for i, _ := range allIngredients {
		safe[i] = true
	}

	for a, _ := range allAllergens {
		potentialIngredients := map[ingredient]bool{}
		for i, _ := range allIngredients {
			potentialIngredients[i] = true
		}

		for _, b := range blahs {
			if !b.allergens[a] {
				continue
			}
			for i, _ := range potentialIngredients {
				if !b.ingredients[i] {
					delete(potentialIngredients, i)
				}
			}
		}

		for i, _ := range potentialIngredients {
			delete(safe, i)
		}
	}

	r := 0
	for i, _ := range safe {
		t := 0
		for _, b := range blahs {
			if b.ingredients[i] {
				t++
			}
		}
		r += t
	}
	return r
}

func Part2(input string) string {
	allIngredients := map[ingredient]bool{}
	allAllergens := map[allergen]bool{}
	var blahs []blah
	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " (contains ")
		b := blah{
			ingredients: map[ingredient]bool{},
			allergens:   map[allergen]bool{},
		}
		ingredients := strings.Split(l[0], " ")
		for _, v := range ingredients {
			b.ingredients[ingredient(v)] = true
			allIngredients[ingredient(v)] = true
		}

		l[1] = l[1][0 : len(l[1])-1]
		allergens := strings.Split(l[1], ", ")
		for _, v := range allergens {
			b.allergens[allergen(v)] = true
			allAllergens[allergen(v)] = true
		}

		blahs = append(blahs, b)
	}

	safe := map[ingredient]bool{}
	for i, _ := range allIngredients {
		safe[i] = true
	}

	for a, _ := range allAllergens {
		potentialIngredients := map[ingredient]bool{}
		for i, _ := range allIngredients {
			potentialIngredients[i] = true
		}

		for _, b := range blahs {
			if !b.allergens[a] {
				continue
			}
			for i, _ := range potentialIngredients {
				if !b.ingredients[i] {
					delete(potentialIngredients, i)
				}
			}
		}

		for i, _ := range potentialIngredients {
			delete(safe, i)
		}
	}

	r := 0
	for i, _ := range safe {
		t := 0
		for _, b := range blahs {
			if b.ingredients[i] {
				t++
			}
		}
		r += t
	}

	type pair struct {
		i ingredient
		a allergen
	}

	var solution []pair
	usedIngredients := map[ingredient]bool{}
	usedAllergens := map[allergen]bool{}
	done := false
	for !done {
		done = true
		for a, _ := range allAllergens {
			if usedAllergens[a] {
				continue
			}
			potentialIngredients := map[ingredient]bool{}
			for i, _ := range allIngredients {
				if safe[i] {
					continue
				}
				if usedIngredients[i] {
					continue
				}
				potentialIngredients[i] = true
			}

			for _, b := range blahs {
				if !b.allergens[a] {
					continue
				}
				for i, _ := range potentialIngredients {
					if !b.ingredients[i] {
						delete(potentialIngredients, i)
					}
				}
			}
			if len(potentialIngredients) == 1 {
				var k ingredient
				for k, _ = range potentialIngredients {
					break
				}
				usedAllergens[a] = true
				usedIngredients[k] = true
				solution = append(solution, pair{k, a})
				done = false
			}
		}
	}
	if len(solution) != len(allAllergens) {
		panic("no solution")
	}

	sort.Slice(solution, func(i, j int) bool {
		return solution[i].a < solution[j].a
	})

	var t []string
	for _, p := range solution {
		t = append(t, string(p.i))
	}
	return strings.Join(t, ",")
}
