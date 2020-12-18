package day18

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`1 + 2 * 3 + 4 * 5 + 6`)
	require.Equal(t, 71, r)

	r = Part1(`1 + (2 * 3) + (4 * (5 + 6))`)
	require.Equal(t, 51, r)

	r = Part1(`2 * 3 + (4 * 5)`)
	require.Equal(t, 26, r)

	r = Part1(`5 + (8 * 3 + 9 + 3 * 4 * 3)`)
	require.Equal(t, 437, r)

	r = Part1(`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`)
	require.Equal(t, 12240, r)

	r = Part1(`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`)
	require.Equal(t, 13632, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`1 + 2 * 3 + 4 * 5 + 6`)
	require.Equal(t, 231, r)

	r = Part2(`1 + (2 * 3) + (4 * (5 + 6))`)
	require.Equal(t, 51, r)

	r = Part2(`2 * 3 + (4 * 5)`)
	require.Equal(t, 46, r)

	r = Part2(`5 + (8 * 3 + 9 + 3 * 4 * 3)`)
	require.Equal(t, 1445, r)

	r = Part2(`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`)
	require.Equal(t, 669060, r)

	r = Part2(`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`)
	require.Equal(t, 23340, r)
}
