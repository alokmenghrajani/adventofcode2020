package day13

import (
	"fmt"
	"math"
	"strings"

	"github.com/alokmenghrajani/adventofcode2020/utils"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	//soonest := utils.MustAtoi(lines[0])
	scheduleString := strings.Split(lines[1], ",")
	var schedule []int
	for _, s := range scheduleString {
		if s == "x" {
			schedule = append(schedule, 0)
			continue
		}
		schedule = append(schedule, utils.MustAtoi(s))
	}

	//// find id and at what time
	//when := map[int]int{}
	//for _, v := range schedule {
	//	t := int(math.Ceil(float64(soonest)/float64(v))) * v
	//	when[v] = t
	//}
	//

	last := schedule[len(schedule)-1]
	i := int(math.Ceil(float64(100000000000000)/float64(last))) * last
	j := 0
	for {
		ok := true
		j = i - (len(schedule) - 1)
		for offset, v := range schedule {
			if v == 0 {
				continue
			}
			t := int(math.Ceil(float64(j+offset)/float64(v))) * v
			if t != j+offset {
				ok = false
				break
			}
		}
		if ok {
			break
		}
		i += last
	}
	return j
}

func Part2(input string) int {
	for _, line := range strings.Split(input, "\n") {
		fmt.Printf("%s\n", line)
	}
	return 0
}
