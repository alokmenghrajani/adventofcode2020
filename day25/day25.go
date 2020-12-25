package day25

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

func Part1(input string) int {
	doorPubkey := utils.MustAtoi(strings.Split(input, "\n")[0])
	cardPubkey := utils.MustAtoi(strings.Split(input, "\n")[1])

	n1 := f(7, cardPubkey)
	return f2(doorPubkey, n1)
}

func f(subject int, find int) int {
	n := 1
	i := 0
	for {
		i++
		n = (n * subject) % 20201227
		if n == find {
			return i
		}
	}
}

func f2(subject int, times int) int {
	n := 1
	for i := 0; i < times; i++ {
		n = (n * subject) % 20201227
	}
	return n
}
