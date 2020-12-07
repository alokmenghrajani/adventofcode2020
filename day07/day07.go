package day07

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

type bag struct {
	amount int
	name   string
}

var bags map[string][]bag

func Part1(input string) int {
	bags = map[string][]bag{}
	parse(input)

	// find everything which contains
	// "shiny gold" bag
	total := 0
	for k, _ := range bags {
		if contains(k, "shiny gold") {
			total++
		}
	}
	return total
}

func parse(input string) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// light lavender bags contain 5 clear teal bags, 2 wavy gold bags, 5 drab maroon bags, 2 posh cyan bags.
		// dim violet bags contain no other bags.
		re := regexp.MustCompile(`^(.*?) bags? contain`)
		match := re.FindStringSubmatch(line)
		k := match[1]

		re = regexp.MustCompile(`(\d+) (.+?) bags?`)
		matches := re.FindAllStringSubmatch(line, -1)
		for i := 0; i < len(matches); i++ {
			if bags[k] == nil {
				bags[k] = []bag{}
			}
			bags[k] = append(bags[k], bag{
				amount: utils.MustAtoi(matches[i][1]),
				name:   matches[i][2],
			})
		}
	}
}

func contains(key string, target string) bool {
	for _, v := range bags[key] {
		if v.name == target {
			return true
		}
	}
	for _, v := range bags[key] {
		if contains(v.name, target) {
			return true
		}
	}
	return false
}

func Part2(input string) int {
	bags = map[string][]bag{}
	parse(input)

	// weight returns the number of bags in a given bag, including itself. So subtract one.
	return weight("shiny gold") - 1
}

func weight(key string) int {
	if bags[key] == nil {
		// empty bag counts as 1
		return 1
	}

	// filled bags is 1 + amount * each bag's weight
	r := 1
	for _, v := range bags[key] {
		r += v.amount * weight(v.name)
	}
	return r
}
