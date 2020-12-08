package day08

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)
	assert.Equal(t, 5, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)
	assert.Equal(t, 8, r)
}
