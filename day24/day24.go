package day24

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils/grids"
)

func Part1(input string) int {
	grid := grids.NewGrid(false)
	for _, line := range strings.Split(input, "\n") {
		x := 0
		y := 0
		for i := 0; i < len(line); {
			if line[i] == 'e' {
				x++
				i++
			} else if line[i] == 'w' {
				x--
				i++
			} else if line[i:i+2] == "se" {
				x++
				y--
				i += 2
			} else if line[i:i+2] == "sw" {
				y--
				i += 2
			} else if line[i:i+2] == "ne" {
				y++
				i += 2
			} else if line[i:i+2] == "nw" {
				x--
				y++
				i += 2
			}
		}
		t := grid.Get(x, y).(bool)
		t = !t
		grid.Set(x, y, t)
	}

	return count(grid)
}

func count(grid *grids.Grid) int {
	// count true
	minX, maxX := grid.SizeX()
	minY, maxY := grid.SizeY()
	r := 0
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if grid.Get(i, j).(bool) {
				r++
			}
		}
	}
	return r
}

func Part2(input string) int {
	grid := grids.NewGrid(false)
	for _, line := range strings.Split(input, "\n") {
		x := 0
		y := 0
		for i := 0; i < len(line); {
			if line[i] == 'e' {
				x++
				i++
			} else if line[i] == 'w' {
				x--
				i++
			} else if line[i:i+2] == "se" {
				x++
				y--
				i += 2
			} else if line[i:i+2] == "sw" {
				y--
				i += 2
			} else if line[i:i+2] == "ne" {
				y++
				i += 2
			} else if line[i:i+2] == "nw" {
				x--
				y++
				i += 2
			}
		}
		t := grid.Get(x, y).(bool)
		t = !t
		grid.Set(x, y, t)
	}

	for d := 1; d <= 100; d++ {
		newGrid := grids.NewGrid(false)
		minX, maxX := grid.SizeX()
		minY, maxY := grid.SizeY()
		for i := minX - 1; i <= maxX+1; i++ {
			for j := minY - 1; j <= maxY+1; j++ {
				r := 0
				if grid.Get(i+1, j).(bool) {
					r++
				}
				if grid.Get(i-1, j).(bool) {
					r++
				}
				if grid.Get(i, j+1).(bool) {
					r++
				}
				if grid.Get(i, j-1).(bool) {
					r++
				}
				if grid.Get(i-1, j+1).(bool) {
					r++
				}
				if grid.Get(i+1, j-1).(bool) {
					r++
				}

				v := grid.Get(i, j).(bool)
				if v && (r == 0 || r > 2) {
					newGrid.Set(i, j, !v)
				} else if !v && r == 2 {
					newGrid.Set(i, j, !v)
				} else {
					newGrid.Set(i, j, v)
				}
			}
		}
		grid = newGrid
	}

	return count(grid)
}
