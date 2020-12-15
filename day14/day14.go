package day14

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

func Part1(input string) int {
	mask := ""
	memory := map[int]int{}
	for _, line := range strings.Split(input, "\n") {
		re1 := regexp.MustCompile(`mask = (.*)$`)
		match := re1.FindStringSubmatch(line)
		if len(match) > 0 {
			mask = match[1]
			continue
		}

		re2 := regexp.MustCompile(`mem\[(\d+)\] = (\d+)$`)
		match = re2.FindStringSubmatch(line)
		if len(match) > 0 {
			ms := compute2(mask, utils.MustAtoi(match[1]), 0)
			for _, v2 := range ms {
				memory[v2] = utils.MustAtoi(match[2])
			}
		} else {
			panic("here")
		}
	}

	sum := 0
	for _, v := range memory {
		sum += v
	}
	return sum
}

func compute2(mask string, v int, offset int) []int {
	if offset == len(mask) {
		return []int{v}
	}

	if mask[offset] == '1' {
		m := 1 << (35 - offset)
		v = v | m
		return compute2(mask, v, offset+1)
	}
	if mask[offset] == '0' {
		return compute2(mask, v, offset+1)
	}
	if mask[offset] == 'X' {
		r := compute2(mask, v, offset+1)
		t := len(r)
		for i := 0; i < t; i++ {
			v2 := r[i]
			m := 1 << (35 - offset)
			v2 = v2 ^ m
			r = append(r, v2)
		}
		return r
	}
	panic("unreach")
}

func compute(mask string, v int) int {
	for i := 0; i < len(mask); i++ {
		if mask[i] == 'X' {
			continue
		}
		if mask[i] == '0' {
			m := ^(1 << (35 - i))
			v = v & m
		}
		if mask[i] == '1' {
			m := 1 << (35 - i)
			v = v | m
		}
	}
	return v
}

//		// mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
//		//mem[8] = 11
//		//mem[7] = 101
//		//mem[8] = 0
//		fmt.Println(line)
//	}
//	return 0
//}
//
////func Part2(input string) int {
////	for _, line := range strings.Split(input, "\n") {
////		fmt.Printf("%s\n", line)
////	}
////	return 0
////}
