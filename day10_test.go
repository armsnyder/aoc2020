package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func Test_day10(t *testing.T) {
	runDayTests(t, 10, []dayTest{
		{
			input: `
16
10
15
5
1
11
7
19
6
12
4
`,
			want: 35,
		},
		{
			input: `
28
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
3
`,
			want: 220,
		},
		{
			part2: true,
			input: `
16
10
15
5
1
11
7
19
6
12
4
`,
			want: 8,
		},
		{
			part2: true,
			input: `
28
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
3
`,
			want: 19208,
		},
	})
}

func BenchmarkDay10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[10](true, aocutil.GetInput(10))
	}
}
