package day18

import (
	"strconv"
	"strings"

	"github.com/yhirose/go-peg"
)

func Part1(input string) int {
	parser, _ := peg.NewParser(`
    EXPR         ←  ATOM (BINOP ATOM)*
    ATOM         ←  NUMBER / '(' EXPR ')'
    BINOP        ←  < [+*] >
    NUMBER       ←  < [0-9] >
    %whitespace  ←  [ ]*
    ---
    %expr  = EXPR  
    %binop = L + * 
`)

	g := parser.Grammar
	g["EXPR"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		val := v.ToInt(0)
		if v.Len() > 1 {
			ope := v.ToStr(1)
			rhs := v.ToInt(2)
			switch ope {
			case "+":
				val += rhs
			case "*":
				val *= rhs
			}
		}
		return val, nil
	}
	g["BINOP"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return v.Token(), nil
	}
	g["NUMBER"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return strconv.Atoi(v.Token())
	}

	r := 0
	for _, line := range strings.Split(input, "\n") {
		val, _ := parser.ParseAndGetValue(line, nil)
		r += val.(int)
	}
	return r
}

func Part2(input string) int {
	parser, _ := peg.NewParser(`
    EXPR         ←  ATOM (BINOP ATOM)*
    ATOM         ←  NUMBER / '(' EXPR ')'
    BINOP        ←  < [+*] >
    NUMBER       ←  < [0-9] >
    %whitespace  ←  [ \t]*
    ---
    %expr  = EXPR
    %binop = L * 
    %binop = L + 
`)

	g := parser.Grammar
	g["EXPR"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		val := v.ToInt(0)
		if v.Len() > 1 {
			ope := v.ToStr(1)
			rhs := v.ToInt(2)
			switch ope {
			case "+":
				val += rhs
			case "*":
				val *= rhs
			}
		}
		return val, nil
	}
	g["BINOP"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return v.Token(), nil
	}
	g["NUMBER"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
		return strconv.Atoi(v.Token())
	}

	r := 0
	for _, line := range strings.Split(input, "\n") {
		val, _ := parser.ParseAndGetValue(line, nil)
		r += val.(int)
	}
	return r
}
