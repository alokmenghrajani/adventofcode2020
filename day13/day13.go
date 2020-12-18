package day13

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	soonest := utils.MustAtoi(lines[0])
	scheduleString := strings.Split(lines[1], ",")
	var schedule []int
	for _, s := range scheduleString {
		if s == "x" {
			schedule = append(schedule, 0)
			continue
		}
		schedule = append(schedule, utils.MustAtoi(s))
	}

	// find id and at what time
	when := map[int]int{}
	for _, v := range schedule {
		if v == 0 {
			continue
		}
		t := int(math.Ceil(float64(soonest)/float64(v))) * v
		when[v] = t
	}

	k := utils.MapFindMin(when).(int)
	return k * (when[k] - soonest)
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	scheduleString := strings.Split(lines[1], ",")
	var crtA []*big.Int
	var crtN []*big.Int
	for offset, s := range scheduleString {
		if s == "x" {
			continue
		}
		v := utils.MustAtoi(s)
		crtA = append(crtA, big.NewInt(int64(v-offset%v)))
		crtN = append(crtN, big.NewInt(int64(v)))
	}

	r, err := crt(crtA, crtN)
	utils.PanicOnErr(err)
	return int(r.Int64())
}

var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
