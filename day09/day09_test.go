package day09

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(5, `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)
	assert.Equal(t, 127, r)
}

func TestPart2(t *testing.T) {
	r := Part2(5, `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)
	assert.Equal(t, 62, r)
}
