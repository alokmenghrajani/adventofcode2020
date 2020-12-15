package day14

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)
	require.Equal(t, 165, r)
}

//func TestPart2(t *testing.T) {
//	r := Part2(``)
//	require.Equal(t, 0, r)
//}
