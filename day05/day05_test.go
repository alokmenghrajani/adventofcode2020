package day05

import (
	"github.com/alecthomas/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	r := compute("FBFBBFFRLR")
	assert.Equal(t, 357, r)

	r = compute("BFFFBBFRRR")
	assert.Equal(t, 567, r)

	r = compute("FFFBBBFRRR")
	assert.Equal(t, 119, r)

	r = compute("BBFFBBFRLL")
	assert.Equal(t, 820, r)
}
