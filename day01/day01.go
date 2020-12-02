package day01

import (
	"github.com/alokmenghrajani/adventofcode2020/utils"
)

// Inefficiently loop over entire dataset twice
func Part1(input []string) int {
	numbers := utils.StringsToInts(input)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				return numbers[i] * numbers[j]
			}
		}
	}
	panic("no solution")
}

// Inefficiently loop over entire dataset thrice
func Part2(input []string) int {
	numbers := utils.StringsToInts(input)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					return numbers[i] * numbers[j] * numbers[k]
				}
			}
		}
	}

	panic("no solution")
}
