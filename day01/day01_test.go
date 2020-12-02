package day01

import (
	"github.com/alecthomas/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1([]string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	})
	assert.Equal(t, 514579, r)
}

func TestPart1_edgecase(t *testing.T) {
	r := Part1([]string{
		"1010",
		"2019",
		"1",
	})
	assert.Equal(t, 2019, r)
}

func TestPart2(t *testing.T) {
	r := Part2([]string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	})
	assert.Equal(t, 241861950, r)
}

func TestPart2_edgecase(t *testing.T) {
	r := Part2([]string{
		"1009",
		"2",
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	})
	assert.Equal(t, 241861950, r)
}
