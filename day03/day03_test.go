package day03

import (
	"github.com/alecthomas/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)
	assert.Equal(t, uint(7), r)
}

func TestPart2(t *testing.T) {
	r := Part2(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)
	assert.Equal(t, uint(336), r)
}
