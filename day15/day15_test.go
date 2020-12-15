package day15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := compute(`0,3,6`, 10)
	require.Equal(t, 0, r)

	r = compute(`0,3,6`, 2020)
	require.Equal(t, 436, r)

	r = compute(`1,3,2`, 2020)
	require.Equal(t, 1, r)

	r = compute(`2,1,3`, 2020)
	require.Equal(t, 10, r)

	r = compute(`1,2,3`, 2020)
	require.Equal(t, 27, r)

	r = compute(`2,3,1`, 2020)
	require.Equal(t, 78, r)

	r = compute(`3,2,1`, 2020)
	require.Equal(t, 438, r)

	r = compute(`3,1,2`, 2020)
	require.Equal(t, 1836, r)
}

func TestPart2(t *testing.T) {
	r := compute(`0,3,6`, 30000000)
	require.Equal(t, 175594, r)

	r = compute(`1,3,2`, 30000000)
	require.Equal(t, 2578, r)

	r = compute(`2,1,3`, 30000000)
	require.Equal(t, 3544142, r)

	r = compute(`1,2,3`, 30000000)
	require.Equal(t, 261214, r)

	r = compute(`2,3,1`, 30000000)
	require.Equal(t, 6895259, r)

	r = compute(`3,2,1`, 30000000)
	require.Equal(t, 18, r)

	r = compute(`3,1,2`, 30000000)
	require.Equal(t, 362, r)
}
