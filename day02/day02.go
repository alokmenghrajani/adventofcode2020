package day02

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	re := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)
	valid := 0
	for _, l := range input {
		match := re.FindStringSubmatch(l)
		min, _ := strconv.Atoi(match[1])
		max, _ := strconv.Atoi(match[2])
		c := match[3]
		s := match[4]

		t := strings.Count(s, c)
		if t >= min && t <= max {
			valid++
		}
	}
	return valid
}

func Part2(input []string) int {
	re := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)
	valid := 0
	for _, l := range input {
		match := re.FindStringSubmatch(l)
		offset1, _ := strconv.Atoi(match[1])
		offset2, _ := strconv.Atoi(match[2])
		c := match[3][0]
		s := match[4]

		// golang doesn't have any xor, but != is a fine replacement.
		if (s[offset1-1] == c) != (s[offset2-1] == c) {
			valid++
		}
	}
	return valid
}
