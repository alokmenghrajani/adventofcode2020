package day11

import (
	"github.com/alokmenghrajani/adventofcode2020/utils/grids"

	"github.com/alokmenghrajani/adventofcode2020/utils/inputs"
)

func Part1(input string) int {
	grid := inputs.ToGrid(input)
	_, maxX := grid.SizeX()
	_, maxY := grid.SizeY()

	for done := false; !done; {
		done = true
		nextGrid := grids.NewGrid()

		for i := 0; i <= maxX; i++ {
			for j := 0; j <= maxY; j++ {
				v := grid.GetRune(i, j)
				if v == 'L' && countPart1(grid, i, j) == 0 {
					nextGrid.Set(i, j, '#')
					done = false
				} else if v == '#' && countPart1(grid, i, j) >= 4 {
					nextGrid.Set(i, j, 'L')
					done = false
				} else {
					nextGrid.Set(i, j, v)
				}
			}
		}
		grid = nextGrid
	}

	// count
	c := 0
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if grid.GetRune(i, j) == '#' {
				c++
			}
		}
	}
	return c
}

func countPart1(grid *grids.Grid, x, y int) int {
	r := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if grid.GetRune(x+i, y+j) == '#' {
				r++
			}
		}
	}
	return r
}

func Part2(input string) int {
	grid := inputs.ToGrid(input)
	_, maxX := grid.SizeX()
	_, maxY := grid.SizeY()

	for done := false; !done; {
		done = true
		nextGrid := grids.NewGrid()

		for i := 0; i <= maxX; i++ {
			for j := 0; j <= maxY; j++ {
				v := grid.GetRune(i, j)
				if v == 'L' && countPart2(grid, i, j) == 0 {
					nextGrid.Set(i, j, '#')
					done = false
				} else if v == '#' && countPart2(grid, i, j) >= 5 {
					nextGrid.Set(i, j, 'L')
					done = false
				} else {
					nextGrid.Set(i, j, v)
				}
			}
		}
		grid = nextGrid
	}

	// count
	c := 0
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if grid.GetRune(i, j) == '#' {
				c++
			}
		}
	}
	return c
}

func countPart2(grid *grids.Grid, x, y int) int {
	r := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for d := 1; ; d++ {
				tx := x + i*d
				ty := y + j*d
				v := grid.GetRune(tx, ty)
				if v == '#' {
					r++
					break
				} else if v != '.' {
					break
				}
			}
		}
	}
	return r
}

//
//	for _, line := range strings.Split(input, "\n") {
//		fmt.Println(line)
//	}
//	//	d := 1
//	//	for {
//	//		ty := y + i*d
//	//		tx := x + j*d
//	//		if ty < 0 || ty >= len(grid) || tx < 0 || tx >= len(grid[0]) {
//	//			break
//	//		}
//	//		if grid[ty][tx] == '#' {
//	//			r++
//	//			break
//	//		}
//	//		if grid[ty][tx] == 'L' {
//	//			break
//	//		}
//	//		d++
//	//	}
//	//}
//	//}
//	//return r
//	//numbers := inputs.ToInts(input)
//	//sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })
//
//	return 0
//}
