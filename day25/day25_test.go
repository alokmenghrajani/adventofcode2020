package day25

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`5764801
17807724`)
	require.Equal(t, 14897079, r)
}
