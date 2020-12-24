package day22

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`)
	require.Equal(t, 306, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`)
	require.Equal(t, 291, r)
}
