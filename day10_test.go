package main

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay10Part1(t *testing.T) {
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
	})
}

func TestDay10Part2(t *testing.T) {
	runDayTests(t, 10, []dayTest{
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

func BenchmarkDay10BaselineIO(b *testing.B) {
	day := 10
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		io.Copy(ioutil.Discard, input)
		input.Close()
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	day := 10
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	day := 10
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
