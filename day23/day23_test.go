package day23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`389125467`)
	require.Equal(t, "67384529", r)
}

func TestPart2(t *testing.T) {
	r := Part2(`389125467`)
	require.Equal(t, 149245887792, r)
}
