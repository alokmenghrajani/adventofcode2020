package day12

import (
	"math"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

func Part1(input string) int {
	d := []int{1, 0}
	pos := []int{0, 0}
	for _, line := range strings.Split(input, "\n") {
		//fmt.Printf("%d\t%d\n", pos[0], pos[1])
		l := line[0]
		v := utils.MustAtoi(line[1:])
		switch l {
		case 'N':
			pos[1] += v
		case 'S':
			pos[1] -= v
		case 'E':
			pos[0] += v
		case 'W':
			pos[0] -= v
		case 'F':
			pos[0] += d[0] * v
			pos[1] += d[1] * v
		case 'L':
			for {
				t := []int{0, 0}
				t[0] = d[0]
				t[1] = d[1]

				d[0] = -t[1]
				d[1] = t[0]
				v = v - 90
				if v == 0 {
					break
				}
			}
		case 'R':
			for {
				t := []int{0, 0}
				t[0] = d[0]
				t[1] = d[1]

				d[0] = t[1]
				d[1] = -t[0]
				v = v - 90
				if v == 0 {
					break
				}
			}
		}
	}
	i := int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
	return i
}

func Part2(input string) int {
	//d := []int{1, 0}
	pos := []int{0, 0}
	w := []int{10, 1}
	for _, line := range strings.Split(input, "\n") {
		l := line[0]
		v := utils.MustAtoi(line[1:])
		switch l {
		case 'N':
			w[1] += v
		case 'S':
			w[1] -= v
		case 'E':
			w[0] += v
		case 'W':
			w[0] -= v
		case 'F':
			pos[0] += w[0] * v
			pos[1] += w[1] * v
		case 'L':
			for {
				t := []int{0, 0}
				t[0] = w[0]
				t[1] = w[1]

				w[0] = -t[1]
				w[1] = t[0]
				v = v - 90
				if v == 0 {
					break
				}
			}
		case 'R':
			for {
				t := []int{0, 0}
				t[0] = w[0]
				t[1] = w[1]

				w[0] = t[1]
				w[1] = -t[0]
				v = v - 90
				if v == 0 {
					break
				}
			}
		}
	}
	i := int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
	return i
}
