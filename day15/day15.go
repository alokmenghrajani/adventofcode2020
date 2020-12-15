package day15

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

type state struct {
	offset      int
	prevValue   *int
	latestValue int
}

func Part1(input string) int {
	return compute(input, 2020)
}

func Part2(input string) int {
	return compute(input, 30000000)
}

func compute(input string, loop int) int {
	numbers := strings.Split(input, ",")
	seen := map[int]int{}
	var s state
	for i, v := range numbers {
		s = add(seen, i+1, utils.MustAtoi(v))
	}

	for s.offset < loop {
		if s.prevValue == nil {
			s = add(seen, s.offset+1, 0)
		} else {
			s = add(seen, s.offset+1, s.offset-*s.prevValue)
		}
	}
	return s.latestValue
}

func add(seen map[int]int, offset, value int) state {
	t, ok := seen[value]
	seen[value] = offset
	if ok {
		return state{offset: offset, prevValue: &t, latestValue: value}
	} else {
		return state{offset: offset, prevValue: nil, latestValue: value}
	}
}
