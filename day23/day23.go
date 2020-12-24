package day23

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

type cup struct {
	id   int
	next *cup
}

func Part1(input string) string {
	var first *cup
	var prev *cup
	cups := map[int]*cup{}
	max := 0
	for _, n := range strings.Split(input, "") {
		curr := &cup{
			id: utils.MustAtoi(n),
		}
		max = utils.IntMax(max, curr.id)
		cups[curr.id] = curr
		if first == nil {
			first = curr
		} else {
			prev.next = curr
		}
		prev = curr
	}
	prev.next = first

	curr := first
	for t := 0; t < 100; t++ {
		// remove 3 cups
		cup1 := curr.next
		cup2 := cup1.next
		cup3 := cup2.next
		curr.next = cup3.next

		n := curr.id
		for {
			n--
			if n == 0 {
				n = max
			}
			if n != cup1.id && n != cup2.id && n != cup3.id {
				break
			}
		}

		dest := cups[n]
		cup3.next = dest.next
		dest.next = cup1
		curr = curr.next
	}

	r := ""
	curr = cups[1].next
	for curr != cups[1] {
		r = fmt.Sprintf("%s%d", r, curr.id)
		curr = curr.next
	}

	return r
}

func Part2(input string) int {
	var first *cup
	var prev *cup
	cups := map[int]*cup{}
	max := 0
	for _, n := range strings.Split(input, "") {
		curr := &cup{
			id: utils.MustAtoi(n),
		}
		max = utils.IntMax(max, curr.id)
		cups[curr.id] = curr
		if first == nil {
			first = curr
		} else {
			prev.next = curr
		}
		prev = curr
	}
	for len(cups) < 1_000_000 {
		curr := &cup{
			id: max + 1,
		}
		max = utils.IntMax(max, curr.id)
		cups[curr.id] = curr
		prev.next = curr
		prev = curr
	}
	prev.next = first

	curr := first
	for t := 0; t < 10_000_000; t++ {
		// remove 3 cups
		cup1 := curr.next
		cup2 := cup1.next
		cup3 := cup2.next
		curr.next = cup3.next

		n := curr.id
		for {
			n--
			if n == 0 {
				n = max
			}
			if n != cup1.id && n != cup2.id && n != cup3.id {
				break
			}
		}

		dest := cups[n]
		cup3.next = dest.next
		dest.next = cup1
		curr = curr.next
	}
	return cups[1].next.id * cups[1].next.next.id
}
