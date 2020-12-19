package day19

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEdgeCase(t *testing.T) {
	r := Part1(`0: 1 3
1: 2 | 2 3
2: "a"
3: "b"

ab
aab
abb
aba
b
a
aaba`)
	require.Equal(t, 2, r)
}

func TestPart1(t *testing.T) {
	parseRules(`0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"`, false)
	r := check("b", rules[3])
	require.True(t, r[1])

	r = check("a", rules[1])
	require.True(t, r[1])

	r = check("ab", rules[2])
	require.True(t, r[2])

	r = check("ba", rules[2])
	require.True(t, r[2])

	r = check("aab", rules[0])
	require.True(t, r[3])

	r = check("aba", rules[0])
	require.True(t, r[3])

	r2 := Part1(`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`)
	require.Equal(t, 2, r2)
}

func TestPart2(t *testing.T) {
	part2 := `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`
	r := Part1(part2)
	require.Equal(t, 3, r)

	r = Part2(part2)
	require.Equal(t, 12, r)
}
