package main

import (
	"errors"
	"fmt"
	"github.com/alokmenghrajani/adventofcode2020/day01"
	"github.com/alokmenghrajani/adventofcode2020/day02"
	"github.com/alokmenghrajani/adventofcode2020/day03"
	"github.com/alokmenghrajani/adventofcode2020/day04"
	"github.com/alokmenghrajani/adventofcode2020/utils"
	"os"
	"strconv"
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
	default:
		panic(errors.New(fmt.Sprintf("no such day: %d", d)))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		return 4
	}
	day, err := strconv.Atoi(os.Args[1])
	utils.PanicOnErr(err)
	return day
}
