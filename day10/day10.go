package day10

import (
	"fmt"
	"sort"

	"github.com/alokmenghrajani/adventofcode2020/utils/inputs"
)

func Part1(input string) int {
	numbers := inputs.ToInts(input)
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })

	// append max
	max := numbers[len(numbers)-1]
	numbers = append(numbers, max+3)

	// prepend zero
	numbers = append([]int{0}, numbers...)

	ones := 0
	threes := 0
	for i := 1; i < len(numbers); i++ {
		delta := numbers[i] - numbers[i-1]
		if delta == 1 {
			ones++
		} else if delta == 3 {
			threes++
		}
	}
	return ones * threes
}

func Part2(input string) int {
	numbers := inputs.ToInts(input)
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })

	// append max
	max := numbers[len(numbers)-1]
	numbers = append(numbers, max+3)

	memoize = map[string]int{}
	x := solve(numbers, 0, 0)
	return x
}

var memoize map[string]int

func solve(numbers []int, current int, offset int) int {
	key := fmt.Sprintf("%d-%d", current, offset)
	r, ok := memoize[key]
	if ok {
		return r
	}

	if offset == len(numbers) {
		if current == numbers[offset-1] {
			// we are done when we reach the end and we picked our device
			return 1
		}
		return 0
	}

	delta := numbers[offset] - current
	if delta < 4 {
		r := solve(numbers, numbers[offset], offset+1) + solve(numbers, current, offset+1)
		memoize[key] = r
		return r
	} else {
		return 0
	}
}
