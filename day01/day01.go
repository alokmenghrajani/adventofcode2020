package day01

import (
	"fmt"
	"github.com/alokmenghrajani/adventofcode2020/utils"
)

// Inefficiently loop over entire dataset twice
func Part1(input []string) {
	numbers := utils.StringsToInts(input)

	for _, i := range numbers {
		for _, j := range numbers {
			if i+j == 2020 {
				fmt.Printf("%d\n", i*j)
				return
			}
		}
	}
}

// Inefficiently loop over entire dataset thrice
func Part2(input []string) {
	numbers := utils.StringsToInts(input)

	for _, i := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if i+j+k == 2020 {
					fmt.Printf("%d\n", i*j*k)
					return
				}
			}
		}
	}
}
