package inputs

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

func ToInts(input string) []int {
	var r []int
	for _, line := range strings.Split(input, "\n") {
		r = append(r, utils.MustAtoi(line))
	}
	return r
}

func ToGrid(input string) [][]byte {
	rows := strings.Split(input, "\n")

	r := make([][]byte, len(rows))
	for i := range r {
		r[i] = []byte(rows[i])
	}

	return r
}
