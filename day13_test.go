package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay13Part1(t *testing.T) {
	runDayTests(t, 13, []dayTest{
		{
			input: `
939
7,13,x,x,59,x,31,19
`,
			want: 295,
		},
	})
}

func TestDay13Part2(t *testing.T) {
	runDayTests(t, 13, []dayTest{
		{
			part2: true,
			input: `
xxx
7,13,x,x,59,x,31,19
`,
			want: int64(1068781),
		},
		{
			part2: true,
			input: `
xxx
17,x,13,19
`,
			want: int64(3417),
		},
		{
			part2: true,
			input: `
xxx
67,7,59,61
`,
			want: int64(754018),
		},
		{
			part2: true,
			input: `
xxx
67,x,7,59,61
`,
			want: int64(779210),
		},
		{
			part2: true,
			input: `
xxx
67,7,x,59,61
`,
			want: int64(1261476),
		},
		{
			part2: true,
			input: `
xxx
1789,37,47,1889
`,
			want: int64(1202161486),
		},
	})
}

func BenchmarkDay13Part1(b *testing.B) {
	day := 13
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay13Part2(b *testing.B) {
	day := 13
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
