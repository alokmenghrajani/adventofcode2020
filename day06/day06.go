package day06

import (
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n\n")
	total := 0
	for _, line := range lines {
		answers := strings.Split(line, "\n")
		n := map[rune]int{}
		for _, answer := range answers {
			for _, c := range answer {
				n[c]++
			}
		}
		total += len(n)
	}
	return total
}

func Part2(input string) int {
	lines := strings.Split(input, "\n\n")
	total := 0
	for _, line := range lines {
		answers := strings.Split(line, "\n")
		n := map[rune]int{}
		for _, answer := range answers {
			for _, c := range answer {
				n[c]++
			}
		}
		for _, v := range n {
			if v == len(answers) {
				total++
			}
		}
	}
	return total
}
