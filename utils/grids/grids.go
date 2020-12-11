package grids

import (
	"fmt"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

// TODO:
// - make this an arbitrary dimension thing
type Grid struct {
	minX, maxX, minY, maxY int
	data                   map[string]interface{}
}

func NewGrid() *Grid {
	return &Grid{
		data: map[string]interface{}{},
	}
}

func (g *Grid) SizeX() (int, int) {
	return g.minX, g.maxX
}

func (g *Grid) SizeY() (int, int) {
	return g.minY, g.maxY
}

func (g *Grid) Get(x, y int) interface{} {
	k := key(x, y)
	v, ok := g.data[k]
	if !ok {
		return rune(0)
	}
	return v
}

func (g *Grid) GetRune(x, y int) rune {
	r := g.Get(x, y)
	return r.(rune)
}

func (g *Grid) Set(x, y int, v interface{}) {
	k := key(x, y)
	g.data[k] = v
	g.minX = utils.IntMin(g.minX, x)
	g.maxX = utils.IntMax(g.maxX, x)
	g.minY = utils.IntMin(g.minY, y)
	g.maxY = utils.IntMax(g.maxY, y)
}

func (g *Grid) Print() {
	for j := g.minY; j <= g.maxY; j++ {
		for i := g.minX; i <= g.maxX; i++ {
			fmt.Printf("%c", g.GetRune(i, j))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func key(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}
