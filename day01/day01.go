package day01

import (
	"github.com/alokmenghrajani/adventofcode2020/utils"
	"regexp"
	"strings"
)

type number struct {
	N int
}

// Inefficiently loop over entire dataset twice
func Part1(input string) int {
	lines := strings.Split(input, "\n")
	var numbers []number
	for _, line := range lines {
		var n number
		re := regexp.MustCompile(`\d+`)
		if utils.ParseToStruct(re, line, &n) {
			numbers = append(numbers, n)
		}
	}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i].N+numbers[j].N == 2020 {
				return numbers[i].N * numbers[j].N
			}
		}
	}
	panic("no solution")
}

// Inefficiently loop over entire dataset thrice
func Part2(input string) int {
	lines := strings.Split(input, "\n")
	var numbers []number
	for _, line := range lines {
		var n number
		re := regexp.MustCompile(`\d+`)
		if utils.ParseToStruct(re, line, &n) {
			numbers = append(numbers, n)
		}
	}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i].N+numbers[j].N+numbers[k].N == 2020 {
					return numbers[i].N * numbers[j].N * numbers[k].N
				}
			}
		}
	}

	panic("no solution")
}
