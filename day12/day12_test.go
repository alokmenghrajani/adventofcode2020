package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`F10
N3
F7
R90
F11`)
	require.Equal(t, 25, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`F10
N3
F7
R90
F11`)
	require.Equal(t, 286, r)
}
