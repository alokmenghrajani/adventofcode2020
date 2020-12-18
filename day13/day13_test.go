package day13

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`939
7,13,x,x,59,x,31,19`)
	require.Equal(t, 295, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`0
7,13,x,x,59,x,31,19`)
	require.Equal(t, 1068781, r)
}
