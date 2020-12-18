package day17

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

var minX, maxX, minY, maxY, minZ, maxZ int
var minZZ, maxZZ int

func Part1(input string) int {
	g := map[string]bool{}
	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if c == '#' {
				g[makeKey(x, y, 0)] = true
				minX = utils.IntMin(minX, x)
				maxX = utils.IntMax(maxX, x)
				minY = utils.IntMin(minY, y)
				maxY = utils.IntMax(maxY, y)
			}
		}
	}

	for i := 0; i < 6; i++ {
		g = compute(g)
	}

	return len(g)
}

func makeKey(x, y, z int) string {
	return fmt.Sprintf("%d:%d:%d", x, y, z)
}

func compute(g map[string]bool) map[string]bool {
	rmap := map[string]bool{}

	cMinX := minX
	cMaxX := maxX
	cMinY := minY
	cMaxY := maxY
	cMinZ := minZ
	cMaxZ := maxZ

	for x := cMinX - 1; x <= cMaxX+1; x++ {
		for y := cMinY - 1; y <= cMaxY+1; y++ {
			for z := cMinZ - 1; z <= cMaxZ+1; z++ {
				// count neighbors
				r := 0
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						for k := -1; k <= 1; k++ {
							if i == 0 && j == 0 && k == 0 {
								continue
							}
							if g[makeKey(x+i, y+j, z+k)] {
								r++
							}
						}
					}
				}

				if g[makeKey(x, y, z)] {
					if r == 2 || r == 3 {
						rmap[makeKey(x, y, z)] = true
						minX = utils.IntMin(minX, x)
						maxX = utils.IntMax(maxX, x)
						minY = utils.IntMin(minY, y)
						maxY = utils.IntMax(maxY, y)
						minZ = utils.IntMin(minZ, z)
						maxZ = utils.IntMax(maxZ, z)
					}
				} else {
					if r == 3 {
						rmap[makeKey(x, y, z)] = true
						minX = utils.IntMin(minX, x)
						maxX = utils.IntMax(maxX, x)
						minY = utils.IntMin(minY, y)
						maxY = utils.IntMax(maxY, y)
						minZ = utils.IntMin(minZ, z)
						maxZ = utils.IntMax(maxZ, z)
					}
				}
			}
		}
	}

	return rmap
}

func Part2(input string) int {
	g := map[string]bool{}
	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if c == '#' {
				g[makeKey2(x, y, 0, 0)] = true
				minX = utils.IntMin(minX, x)
				maxX = utils.IntMax(maxX, x)
				minY = utils.IntMin(minY, y)
				maxY = utils.IntMax(maxY, y)
			}
		}
	}

	for i := 0; i < 6; i++ {
		g = compute2(g)
	}

	return len(g)
}

func compute2(g map[string]bool) map[string]bool {
	rmap := map[string]bool{}

	cMinX := minX
	cMaxX := maxX
	cMinY := minY
	cMaxY := maxY
	cMinZ := minZ
	cMaxZ := maxZ
	cMinZZ := minZZ
	cMaxZZ := maxZZ

	for x := cMinX - 1; x <= cMaxX+1; x++ {
		for y := cMinY - 1; y <= cMaxY+1; y++ {
			for z := cMinZ - 1; z <= cMaxZ+1; z++ {
				for zz := cMinZZ - 1; zz <= cMaxZZ+1; zz++ {
					// count neighbors
					r := 0
					for i := -1; i <= 1; i++ {
						for j := -1; j <= 1; j++ {
							for k := -1; k <= 1; k++ {
								for kk := -1; kk <= 1; kk++ {
									if i == 0 && j == 0 && k == 0 && kk == 0 {
										continue
									}
									if g[makeKey2(x+i, y+j, z+k, zz+kk)] {
										r++
									}
								}
							}
						}
					}

					if g[makeKey2(x, y, z, zz)] {
						if r == 2 || r == 3 {
							rmap[makeKey2(x, y, z, zz)] = true
							minX = utils.IntMin(minX, x)
							maxX = utils.IntMax(maxX, x)
							minY = utils.IntMin(minY, y)
							maxY = utils.IntMax(maxY, y)
							minZ = utils.IntMin(minZ, z)
							maxZ = utils.IntMax(maxZ, z)
							minZZ = utils.IntMin(minZZ, zz)
							maxZZ = utils.IntMax(maxZZ, zz)
						}
					} else {
						if r == 3 {
							rmap[makeKey2(x, y, z, zz)] = true
							minX = utils.IntMin(minX, x)
							maxX = utils.IntMax(maxX, x)
							minY = utils.IntMin(minY, y)
							maxY = utils.IntMax(maxY, y)
							minZ = utils.IntMin(minZ, z)
							maxZ = utils.IntMax(maxZ, z)
							minZZ = utils.IntMin(minZZ, zz)
							maxZZ = utils.IntMax(maxZZ, zz)
						}
					}
				}
			}
		}
	}

	return rmap
}

func makeKey2(x, y, z, zz int) string {
	return fmt.Sprintf("%d:%d:%d:%d", x, y, z, zz)
}
