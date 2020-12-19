package day09

import (
	"github.com/alokmenghrajani/adventofcode2020/utils"
	"github.com/alokmenghrajani/adventofcode2020/utils/inputs"
)

func Part1(nPrevious int, input string) int {
	n := inputs.ToInts(input, "\n")
	for i := nPrevious; i < len(n); i++ {
		ok := true
	outer:
		for j := i - nPrevious; j < i; j++ {
			for k := i - nPrevious; k < i; k++ {
				if k == j {
					continue
				}
				if n[j]+n[k] == n[i] {
					ok = false
					break outer
				}
			}
		}
		if ok {
			return n[i]
		}
	}
	panic("no solution?")
}

func Part2(nPrevious int, input string) int {
	target := Part1(nPrevious, input)

	n := inputs.ToInts(input, "\n")

	s := 2
	for {
		for i := 0; i < len(n)-s+1; i++ {
			a, b, c := sumMinMax(n[i : i+s])
			if a == target {
				return b + c
			}
		}
		s++
		if s > len(n) {
			panic("no solution?")
		}
	}
}

func sumMinMax(s []int) (int, int, int) {
	r := 0
	min := utils.MaxInt
	max := utils.MinInt
	for _, n := range s {
		r += n
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return r, min, max
}
