package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay01Part1(t *testing.T) {
	runDayTests(t, 1, []dayTest{
		{
			input: `
1721
979
366
299
675
1456
`,
			want: 514579,
		},
	})
}

func TestDay01Part2(t *testing.T) {
	runDayTests(t, 1, []dayTest{
		{
			part2: true,
			input: `
1721
979
366
299
675
1456
`,
			want: 241861950,
		},
	})
}

func BenchmarkDay01Part1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[1](false, aocutil.GetInput(1))
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[1](true, aocutil.GetInput(1))
	}
}
