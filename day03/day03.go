package day03

import (
	"github.com/alokmenghrajani/adventofcode2020/utils/grids"
	"github.com/alokmenghrajani/adventofcode2020/utils/inputs"
)

func Part1(input string) uint {
	grid := inputs.ToGrid(input)
	return collisions(grid, 3, 1)
}

func Part2(input string) uint {
	grid := inputs.ToGrid(input)
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

func collisions(grid *grids.Grid, dx, dy int) uint {
	x := 0
	y := 0
	c := uint(0)
	_, cols := grid.SizeX()
	_, rows := grid.SizeY()
	for y <= rows {
		x += dx
		y += dy
		if grid.Get(x%(cols+1), y) == '#' {
			c++
		}
	}
	return c
}
