package day20

import (
	fmt "fmt"
	"math"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
	"github.com/alokmenghrajani/adventofcode2020/utils/grids"
	"github.com/alokmenghrajani/adventofcode2020/utils/inputs"
)

var tiles map[int]*grids.Grid // id => tile data
var solution *grids.Grid      // track where each tile goes, by its id
var available map[int]bool    // track which tiles haven't yet been placed.

func Part1(input string) int {
	solve(input)

	// return corners
	minX, maxX := solution.SizeX()
	minY, maxY := solution.SizeY()
	return solution.Get(minX, minY).(int) *
		solution.Get(minX, maxY).(int) *
		solution.Get(maxX, minY).(int) *
		solution.Get(maxX, maxY).(int)
}

func solve(input string) {
	parse(input)

	targetSize := int(math.Sqrt(float64(len(tiles))))

	// place all remaining tiles
	done := false
outer:
	for !done {
		done = true

		// for each spot in solution, try to add a neighboring cell
		minX, maxX := solution.SizeX()
		if maxX-minX+1 < targetSize {
			minX--
			maxX++
		}
		minY, maxY := solution.SizeY()
		if maxY-minY+1 < targetSize {
			minY--
			maxY++
		}

		for i := minX; i <= maxX; i++ {
			for j := minY; j <= maxY; j++ {
				for attempt, _ := range available {
					if place(i, j, attempt) {
						done = false
						delete(available, attempt)
						continue outer
					}
				}
			}
		}
	}

	if len(available) != 0 {
		solution.PrintN()
		panic("no solution?")
	}
	solution.PrintN()
}

func Part2(input string) int {
	solve(input)

	// build the image
	image := grids.NewGrid(' ')
	minX, maxX := solution.SizeX()
	minY, maxY := solution.SizeY()

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			place2(image, x-minX, y-minY, tiles[solution.Get(x, y).(int)])
		}
	}

	monsterX, monsterY := monster()
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			search2(image, monsterX, monsterY)
			image = rotate2(image)
		}
		image = flip2(image)
	}

	// count
	minX, maxX = image.SizeX()
	minY, maxY = image.SizeY()
	r := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if image.GetRune(x, y) == '#' {
				r++
			}
		}
	}
	return r
}

func monster() ([]int, []int) {
	m := `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `
	var rx []int
	var ry []int
	for y, line := range strings.Split(m, "\n") {
		for x, char := range line {
			if char == '#' {
				rx = append(rx, x)
				ry = append(ry, y)
			}
		}
	}
	return rx, ry
}

func search2(image *grids.Grid, monsterX, monsterY []int) {
	minX, maxX := image.SizeX()
	minY, maxY := image.SizeY()
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			checkMonster(image, monsterX, monsterY, x, y)
		}
	}
}

func checkMonster(image *grids.Grid, monsterX, monsterY []int, x, y int) {
	for i := range monsterX {
		if image.GetRune(x+monsterX[i], y+monsterY[i]) != '#' {
			return
		}
	}
	for i := range monsterX {
		image.Set(x+monsterX[i], y+monsterY[i], int32(fmt.Sprintf("%d", i%10)[0]))
	}
	fmt.Printf("found a monster at %d, %d\n", x, y)
}

func rotate2(src *grids.Grid) *grids.Grid {
	r := grids.NewGrid(' ')
	minX, maxX := src.SizeX()
	minY, maxY := src.SizeY()
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			r.Set(y, maxX-x, src.Get(x, y))
		}
	}
	return r
}

func flip2(src *grids.Grid) *grids.Grid {
	r := grids.NewGrid(' ')
	minX, maxX := src.SizeX()
	minY, maxY := src.SizeY()
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			r.Set(maxX-x, y, src.Get(x, y))
		}
	}
	return r
}

func place2(image *grids.Grid, x, y int, tile *grids.Grid) {
	for i := 1; i < 9; i++ {
		for j := 1; j < 9; j++ {
			image.Set(x*8+i-1, y*8+j-1, tile.Get(i, j))
		}
	}
}

func parse(input string) {
	tiles = map[int]*grids.Grid{} // id => tile data
	solution = grids.NewGrid(-1)  // track where each tile goes, by its id
	available = map[int]bool{}    // track which tiles haven't yet been placed.

	// parse input
	for i, tilesString := range strings.Split(input, "\n\n") {
		t := strings.Split(tilesString, "\n")
		firstRow := t[0]
		tileId := utils.MustAtoi(firstRow[5 : len(firstRow)-1])

		// convert tile to a grid
		grid := inputs.ToGrid(strings.Join(t[1:], "\n"), ' ')
		tiles[tileId] = grid

		if i == 0 {
			// the first tile gets placed in the center
			solution.Set(0, 0, tileId)
		} else {
			available[tileId] = true
		}
	}

}

func place(x, y int, attempt int) bool {
	if solution.Get(x, y) != -1 {
		// cell already taken
		return false
	}
	n := 0
	if solution.Get(x+1, y) != -1 {
		n++
	}
	if solution.Get(x-1, y) != -1 {
		n++
	}
	if solution.Get(x, y+1) != -1 {
		n++
	}
	if solution.Get(x, y-1) != -1 {
		n++
	}
	if n == 0 {
		return false
	}

	for dir := 0; dir <= 7; dir++ {
		s := rotate(tiles[attempt], dir)
		if solution.Get(x+1, y) != -1 {
			// we have a cell to our right
			if !checkRight(s, tiles[solution.Get(x+1, y).(int)]) {
				continue
			}
		}
		if solution.Get(x-1, y) != -1 {
			// we have a cell to our left
			if !checkRight(tiles[solution.Get(x-1, y).(int)], s) {
				continue
			}
		}
		if solution.Get(x, y+1) != -1 {
			// we have a cell below us
			if !checkBottom(s, tiles[solution.Get(x, y+1).(int)]) {
				continue
			}
		}
		if solution.Get(x, y-1) != -1 {
			// we have a cell above us
			if !checkBottom(tiles[solution.Get(x, y-1).(int)], s) {
				continue
			}
		}

		// we found something!
		solution.Set(x, y, attempt)
		tiles[attempt] = s
		return true
	}
	return false
}

func rotate(src *grids.Grid, dir int) *grids.Grid {
	if dir == 0 {
		return src
	}
	r := grids.NewGrid(' ')
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			switch dir {
			case 1:
				r.Set(i, 9-j, src.Get(i, j))
			case 2:
				r.Set(9-i, j, src.Get(i, j))
			case 3:
				r.Set(9-i, 9-j, src.Get(i, j))
			case 4:
				r.Set(j, i, src.Get(i, j))
			case 5:
				r.Set(j, 9-i, src.Get(i, j))
			case 6:
				r.Set(9-j, i, src.Get(i, j))
			case 7:
				r.Set(9-j, 9-i, src.Get(i, j))
			default:
				panic("meh")
			}
		}
	}
	return r
}

// check that [a][b] is a valid configuration
func checkRight(a, b *grids.Grid) bool {
	for j := 0; j < 10; j++ {
		if a.Get(9, j) != b.Get(0, j) {
			return false
		}
	}
	return true
}

// check that [a] is a valid configuration
//            [b]
func checkBottom(a, b *grids.Grid) bool {
	for i := 0; i < 10; i++ {
		if a.Get(i, 9) != b.Get(i, 0) {
			return false
		}
	}
	return true
}
