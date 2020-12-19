package day19

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils/inputs"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

type rule struct {
	left   []int
	right  []int
	string byte
}

func (r rule) String() string {
	if len(r.right) > 0 {
		return fmt.Sprintf("%v | %v", r.left, r.right)
	} else if r.string != 0 {
		return fmt.Sprintf("\"%s\"", string(r.string))
	} else {
		return fmt.Sprintf("%v", r.left)
	}
}

var rules map[int]rule

func Part1(input string) int {
	parts := strings.Split(input, "\n\n")
	parseRules(parts[0], false)

	r := 0
	for _, line := range strings.Split(parts[1], "\n") {
		t := check(line, rules[0])
		if t[len(line)] {
			r++
		}
	}
	return r
}

func Part2(input string) int {
	parts := strings.Split(input, "\n\n")
	parseRules(parts[0], true)

	r := 0
	for _, line := range strings.Split(parts[1], "\n") {
		t := check(line, rules[0])
		if t[len(line)] {
			r++
		}
	}
	return r
}

func parseRules(input string, isPart2 bool) {
	rules = map[int]rule{}
	for _, line := range strings.Split(input, "\n") {
		re1 := regexp.MustCompile(`^(\d+): "(.)"$`)
		match := re1.FindStringSubmatch(line)
		if len(match) > 0 {
			n := utils.MustAtoi(match[1])
			r := rule{left: []int{}, right: []int{}, string: match[2][0]}
			rules[n] = r
			continue
		}
		re2 := regexp.MustCompile(`^(\d+): ([0-9 ]+)$`)
		match = re2.FindStringSubmatch(line)
		if len(match) > 0 {
			n := utils.MustAtoi(match[1])
			t := inputs.ToInts(match[2], " ")
			r := rule{left: t, right: []int{}, string: 0}
			rules[n] = r
			continue
		}
		re3 := regexp.MustCompile(`^(\d+): ([0-9 ]+) \| ([0-9 ]+)$`)
		match = re3.FindStringSubmatch(line)
		if len(match) > 0 {
			n := utils.MustAtoi(match[1])
			t1 := inputs.ToInts(match[2], " ")
			t2 := inputs.ToInts(match[3], " ")
			r := rule{left: t1, right: t2, string: 0}
			rules[n] = r
			continue
		}
	}

	if isPart2 {
		rules[8] = rule{left: []int{42}, right: []int{42, 8}}
		rules[11] = rule{left: []int{42, 31}, right: []int{42, 11, 31}}
	}
}

func check(input string, current rule) map[int]bool {
	if len(current.right) > 0 {
		// a fork in the road
		left := rule{left: current.left, right: []int{}, string: 0}
		r1 := check(input, left)

		right := rule{left: current.right, right: []int{}, string: 0}
		r2 := check(input, right)

		for k, _ := range r2 {
			r1[k] = true
		}
		return r1
	}

	if len(current.left) > 1 {
		// a sequence
		head := rule{left: []int{current.left[0]}, right: []int{}, string: 0}
		r1 := check(input, head)

		tail := rule{left: current.left[1:], right: []int{}, string: 0}
		r := map[int]bool{}
		for matched1, _ := range r1 {
			r2 := check(input[matched1:], tail)
			for matched2, _ := range r2 {
				r[matched1+matched2] = true
			}
		}
		return r
	}

	if len(current.left) == 1 {
		// resolve rule
		next := rules[current.left[0]]
		return check(input, next)
	}

	// we are done
	if len(input) > 0 && input[0] == current.string {
		return map[int]bool{1: true}
	}
	return map[int]bool{}
}
