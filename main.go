package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alokmenghrajani/adventofcode2020/day01"
	"github.com/alokmenghrajani/adventofcode2020/day02"
	"github.com/alokmenghrajani/adventofcode2020/day03"
	"github.com/alokmenghrajani/adventofcode2020/day04"
	"github.com/alokmenghrajani/adventofcode2020/day05"
	"github.com/alokmenghrajani/adventofcode2020/day06"
	"github.com/alokmenghrajani/adventofcode2020/day07"
	"github.com/alokmenghrajani/adventofcode2020/day08"
	"github.com/alokmenghrajani/adventofcode2020/day09"
	"github.com/alokmenghrajani/adventofcode2020/day10"
	"github.com/alokmenghrajani/adventofcode2020/day11"
	"github.com/alokmenghrajani/adventofcode2020/utils"
)

// Usage: go run main.go <NN>
// assumes input is in day<NN>/input.txt
func main() {
	d := day()
	fmt.Printf("Running day %02d\n", d)

	switch d {
	case 1:
		fmt.Printf("part 1: %d\n", day01.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day01.Part2(utils.Readfile(d)))
	case 2:
		fmt.Printf("part 1: %d\n", day02.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day02.Part2(utils.Readfile(d)))
	case 3:
		fmt.Printf("part 1: %d\n", day03.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day03.Part2(utils.Readfile(d)))
	case 4:
		fmt.Printf("part 1: %d\n", day04.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day04.Part2(utils.Readfile(d)))
	case 5:
		fmt.Printf("part 1: %d\n", day05.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day05.Part2(utils.Readfile(d)))
	case 6:
		fmt.Printf("part 1: %d\n", day06.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day06.Part2(utils.Readfile(d)))
	case 7:
		fmt.Printf("part 1: %d\n", day07.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day07.Part2(utils.Readfile(d)))
	case 8:
		fmt.Printf("part 1: %d\n", day08.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day08.Part2(utils.Readfile(d)))
	case 9:
		fmt.Printf("part 1: %d\n", day09.Part1(25, utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day09.Part2(25, utils.Readfile(d)))
	case 10:
		fmt.Printf("part 1: %d\n", day10.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day10.Part2(utils.Readfile(d)))
	case 11:
		fmt.Printf("part 1: %d\n", day11.Part1(utils.Readfile(d)))
		fmt.Printf("part 2: %d\n", day11.Part2(utils.Readfile(d)))
	default:
		panic(errors.New(fmt.Sprintf("no such day: %d", d)))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		return 11
	}
	day := utils.Atoi(os.Args[1], -1)
	return day
}
