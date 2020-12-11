package inputs

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils/grids"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

func ToInts(input string) []int {
	var r []int
	for _, line := range strings.Split(input, "\n") {
		r = append(r, utils.MustAtoi(line))
	}
	return r
}

func ToGrid(input string) *grids.Grid {
	grid := grids.NewGrid()

	for y, line := range strings.Split(input, "\n") {
		for x, rune := range line {
			grid.Set(x, y, rune)
		}
	}

	return grid
}
