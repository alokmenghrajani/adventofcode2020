package day22

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

func Part1(input string) int {
	var player1 []int
	var player2 []int

	cards := strings.Split(input, "\n\n")
	for _, line := range strings.Split(cards[0], "\n")[1:] {
		player1 = append(player1, utils.MustAtoi(line))
	}
	for _, line := range strings.Split(cards[1], "\n")[1:] {
		player2 = append(player2, utils.MustAtoi(line))
	}

	winner := combatPart1(player1, player2)

	r := 0
	f := 1
	for i := len(winner) - 1; i >= 0; i-- {
		r = r + winner[i]*f
		f++
	}
	return r
}

func combatPart1(player1, player2 []int) []int {
	for len(player1) > 0 && len(player2) > 0 {
		top1 := player1[0]
		player1 = player1[1:]

		top2 := player2[0]
		player2 = player2[1:]

		if top1 > top2 {
			player1 = append(player1, top1)
			player1 = append(player1, top2)
		} else {
			player2 = append(player2, top2)
			player2 = append(player2, top1)
		}
	}
	if len(player1) > 0 {
		return player1
	}
	if len(player2) > 0 {
		return player2
	}
	panic("invalid state")
}

func Part2(input string) int {
	var player1 []int
	var player2 []int

	cards := strings.Split(input, "\n\n")
	for _, line := range strings.Split(cards[0], "\n")[1:] {
		player1 = append(player1, utils.MustAtoi(line))
	}
	for _, line := range strings.Split(cards[1], "\n")[1:] {
		player2 = append(player2, utils.MustAtoi(line))
	}

	_, winner := combatPart2(player1, player2)

	r := 0
	f := 1
	for i := len(winner) - 1; i >= 0; i-- {
		r = r + winner[i]*f
		f++
	}
	return r
}

func intsToStrings(a []int) []string {
	var r []string
	for _, n := range a {
		r = append(r, fmt.Sprintf("%d", n))
	}
	return r
}

func combatPart2(player1, player2 []int) (bool, []int) {
	memoize := map[string]bool{}
	for len(player1) > 0 && len(player2) > 0 {
		t := strings.Join(intsToStrings(player1), ",")
		if memoize[t] {
			return true, player1
		}
		memoize[t] = true

		t = strings.Join(intsToStrings(player2), ",")
		if memoize[t] {
			return true, player1
		}
		memoize[t] = true

		top1 := player1[0]
		player1 = player1[1:]

		top2 := player2[0]
		player2 = player2[1:]

		var r bool
		if len(player1) >= top1 && len(player2) >= top2 {
			copyPlayer1 := make([]int, top1)
			copy(copyPlayer1, player1)

			copyPlayer2 := make([]int, top2)
			copy(copyPlayer2, player2)

			r, _ = combatPart2(copyPlayer1, copyPlayer2)
		} else {
			r = top1 > top2
		}
		if r {
			player1 = append(player1, top1)
			player1 = append(player1, top2)
		} else {
			player2 = append(player2, top2)
			player2 = append(player2, top1)
		}
	}
	if len(player1) > 0 {
		return true, player1
	}
	if len(player2) > 0 {
		return false, player2
	}
	panic("invalid state")
}
