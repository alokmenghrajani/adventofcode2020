package day05

import (
	"github.com/alokmenghrajani/adventofcode2020/utils"
	"strconv"
	"strings"
)

func Part1(input string) int {
	boardingPasses := strings.Split(input, "\n")
	max := 0
	for _, boardingPass := range boardingPasses {
		max = utils.IntMax(max, compute(boardingPass))
	}
	return max
}

func Part2(input string) int {
	boardingPasses := strings.Split(input, "\n")
	all := map[int]bool{}
	for _, boardingPass := range boardingPasses {
		t := compute(boardingPass)
		all[t] = true
	}

	max := 0
	for _, s := range boardingPasses {
		max = utils.IntMax(max, compute(s))
	}

	min := 0
	for _, s := range boardingPasses {
		min = utils.IntMin(min, compute(s))
	}

	for i := min + 1; i < max-1; i++ {
		if !all[i] && all[i-1] && all[i+1] {
			return i
		}
	}

	return utils.MaxInt
}

func compute(s string) int {
	s = strings.ReplaceAll(s, "B", "1")
	s = strings.ReplaceAll(s, "F", "0")
	s = strings.ReplaceAll(s, "R", "1")
	s = strings.ReplaceAll(s, "L", "0")
	n, err := strconv.ParseInt(s, 2, 0)
	utils.PanicOnErr(err)
	return int(n)
}
