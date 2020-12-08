package day08

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

type ops struct {
	name string
	arg  int
	ran  bool
}

func Part1(input string) int {
	prog := parse(input)
	_, acc := run(prog)
	return acc
}

func Part2(input string) int {
	prog := parse(input)

	// try changing each jmp and nop
	for i, v := range prog {
		if v.name == "jmp" {
			prog[i].name = "nop"
			ok, acc := run(prog)
			if ok {
				return acc
			}
			prog[i].name = "jmp"
		}
		if v.name == "nop" {
			prog[i].name = "jmp"
			ok, acc := run(prog)
			if ok {
				return acc
			}
			prog[i].name = "nop"
		}
	}
	panic("No solution?")
}

func parse(input string) []ops {
	var r []ops
	for _, line := range strings.Split(input, "\n") {
		re := regexp.MustCompile(`(...) ([+-])(\d+)`)
		match := re.FindStringSubmatch(line)
		v := utils.MustAtoi(match[3])
		if match[2] == "-" {
			v = -v
		}
		op := ops{
			name: match[1],
			arg:  v,
			ran:  false,
		}
		r = append(r, op)
	}
	return r
}

func run(prog []ops) (bool, int) {
	// reset ran
	for i := range prog {
		prog[i].ran = false
	}

	// run code
	ip := 0
	acc := 0
	for {
		if ip == len(prog) {
			return true, acc
		}
		op := prog[ip]
		if op.ran {
			return false, acc
		}
		prog[ip].ran = true

		switch op.name {
		case "acc":
			acc += op.arg
			ip++
		case "jmp":
			ip += op.arg
		case "nop":
			ip++
		}
	}
}
