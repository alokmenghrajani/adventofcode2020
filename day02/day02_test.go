package day02

import (
	"github.com/alecthomas/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1([]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	})
	assert.Equal(t, 2, r)
}

func TestPart2(t *testing.T) {
	r := Part2([]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	})
	assert.Equal(t, 1, r)
}
