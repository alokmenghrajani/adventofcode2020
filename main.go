package main

import (
	"errors"
	"fmt"
	"github.com/alokmenghrajani/adventofcode2020/day01"
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
		day01.Part1(utils.Readfile(d))
		day01.Part2(utils.Readfile(d))
	default:
		panic(errors.New(fmt.Sprintf("no such day: %d", d)))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		return 1
	}
	day, err := strconv.Atoi(os.Args[1])
	utils.PanicOnErr(err)
	return day
}
