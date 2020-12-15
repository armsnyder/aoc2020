package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay06Part1(t *testing.T) {
	runDayTests(t, 6, []dayTest{
		{
			input: `
abc

a
b
c

ab
ac

a
a
a
a

b
`,
			want: 11,
		},
	})
}

func TestDay06Part2(t *testing.T) {
	runDayTests(t, 6, []dayTest{
		{
			part2: true,
			input: `
abc

a
b
c

ab
ac

a
a
a
a

b
`,
			want: 6,
		},
	})
}

func BenchmarkDay06Part1(b *testing.B) {
	day := 6
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay06Part2(b *testing.B) {
	day := 6
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
