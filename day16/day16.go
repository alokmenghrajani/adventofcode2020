package day16

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

type rule struct {
	name                   string
	min1, max1, min2, max2 int
	offsets                map[int]bool
	final                  int
}

func Part1(input string) int {
	var rules []rule

	parts := strings.Split(input, "\n\n")
	for _, line := range strings.Split(parts[0], "\n") {
		re := regexp.MustCompile(`^(.*?): (\d+)-(\d+) or (\d+)-(\d+)`)
		match := re.FindStringSubmatch(line)
		r := rule{
			name:    match[1],
			min1:    utils.MustAtoi(match[2]),
			max1:    utils.MustAtoi(match[3]),
			min2:    utils.MustAtoi(match[4]),
			max2:    utils.MustAtoi(match[5]),
			offsets: map[int]bool{},
		}
		rules = append(rules, r)
	}

	res := 0
	t := strings.Split(parts[2], "\n")
	for _, line := range t[1:] {
		for _, n := range strings.Split(line, ",") {
			n2 := utils.MustAtoi(n)
			ok := false
			for _, r := range rules {
				if (n2 >= r.min1 && n2 <= r.max1) || (n2 >= r.min2 && n2 <= r.max2) {
					ok = true
					break
				}
			}
			if !ok {
				res += n2
			}
		}
	}
	return res
}

func Part2(input string) int {
	var rules []rule

	parts := strings.Split(input, "\n\n")
	for _, line := range strings.Split(parts[0], "\n") {
		re := regexp.MustCompile(`^(.*?): (\d+)-(\d+) or (\d+)-(\d+)`)
		match := re.FindStringSubmatch(line)
		r := rule{
			name:    match[1],
			min1:    utils.MustAtoi(match[2]),
			max1:    utils.MustAtoi(match[3]),
			min2:    utils.MustAtoi(match[4]),
			max2:    utils.MustAtoi(match[5]),
			offsets: map[int]bool{},
		}
		rules = append(rules, r)
	}

	t := strings.Split(parts[2], "\n")
	for _, line := range t[1:] {
		for offset, n := range strings.Split(line, ",") {
			n2 := utils.MustAtoi(n)
			ok := false
			for _, r := range rules {
				if (n2 >= r.min1 && n2 <= r.max1) || (n2 >= r.min2 && n2 <= r.max2) {
					ok = true
					break
				}
			}
			if ok {
				for rnumber, r := range rules {
					if (n2 >= r.min1 && n2 <= r.max1) || (n2 >= r.min2 && n2 <= r.max2) {
						_, tok := rules[rnumber].offsets[offset]
						if !tok {
							rules[rnumber].offsets[offset] = true
						}
					} else {
						rules[rnumber].offsets[offset] = false
					}
				}

			}
		}
	}

	// clean
	for i, _ := range rules {
		for k, v := range rules[i].offsets {
			if !v {
				delete(rules[i].offsets, k)
			}
		}
	}

	// find positions
	done := false
	for !done {
		done = true
		for i, _ := range rules {
			if len(rules[i].offsets) == 1 {
				var k int
				for k, _ = range rules[i].offsets {
					break
				}
				rules[i].final = k
				done = false
				for j, _ := range rules {
					delete(rules[j].offsets, k)
				}
			}
		}
	}

	tickets := strings.Split(parts[1], "\n")
	ticket := strings.Split(tickets[1], ",")
	res := 1
	for _, r := range rules {
		if strings.HasPrefix(r.name, "departure") {
			res = res * utils.MustAtoi(ticket[r.final])
		}
	}

	return res
}
