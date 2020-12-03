package day02

import (
	"github.com/alokmenghrajani/adventofcode2020/utils"
	"regexp"
	"strings"
)

func Part1(input []string) int {
	type line struct {
		Min int
		Max int
		C   string
		S   string
	}
	re := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)

	valid := 0

	for _, l := range input {
		var match line
		utils.ParseToStruct(re, l, &match)

		t := strings.Count(match.S, match.C)
		if t >= match.Min && t <= match.Max {
			valid++
		}
	}
	return valid
}

func Part2(input []string) int {
	type line struct {
		Offset1 int
		Offset2 int
		C       byte
		S       string
	}
	re := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)

	valid := 0
	for _, l := range input {
		var match line
		utils.ParseToStruct(re, l, &match)

		// golang doesn't have any xor, but != is a fine replacement.
		if (match.S[match.Offset1-1] == match.C) != (match.S[match.Offset2-1] == match.C) {
			valid++
		}
	}
	return valid
}
