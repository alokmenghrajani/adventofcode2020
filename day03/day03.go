package day03

import (
	"github.com/alokmenghrajani/adventofcode2020/utils"
	"strings"
)

func Part1(input string) uint {
	lines := strings.Split(input, "\n")
	grid := utils.Grid(lines)
	return collisions(grid, 3, 1)
}

func Part2(input string) uint {
	lines := strings.Split(input, "\n")
	grid := utils.Grid(lines)
	deltas := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	r := uint(1)
	for _, delta := range deltas {
		r *= collisions(grid, delta[0], delta[1])
	}
	return r
}

func collisions(grid [][]byte, dx, dy int) uint {
	x := 0
	y := 0
	c := uint(0)
	cols := len(grid[0])
	for {
		x += dx
		y += dy
		if y >= len(grid) {
			return c
		}
		if grid[y][x%cols] == '#' {
			c++
		}
	}
}
