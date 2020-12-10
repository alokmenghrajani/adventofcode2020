package day10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	r := Part1(`16
10
15
5
1
11
7
19
6
12
4`)
	require.Equal(t, 7*5, r)

	r = Part1(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)
	require.Equal(t, 22*10, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`16
10
15
5
1
11
7
19
6
12
4`)
	require.Equal(t, 8, r)

	r = Part2(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)
	require.Equal(t, 19208, r)
}
