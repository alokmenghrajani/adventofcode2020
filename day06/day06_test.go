package day06

import (
	"github.com/alecthomas/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1(`abcx
abcy
abcz`)
	assert.Equal(t, 6, r)

	r = Part1(`abc

a
b
c

ab
ac

a
a
a
a

b`)
	assert.Equal(t, 11, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`abc

a
b
c

ab
ac

a
a
a
a

b`)
	assert.Equal(t, 6, r)
}
