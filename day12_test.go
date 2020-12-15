package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay12Part1(t *testing.T) {
	runDayTests(t, 12, []dayTest{
		{
			input: `
F10
N3
F7
R90
F11
`,
			want: 25,
		},
	})
}

func TestDay12Part2(t *testing.T) {
	runDayTests(t, 12, []dayTest{
		{
			part2: true,
			input: `
F10
N3
F7
R90
F11
`,
			want: 286,
		},
	})
}

func BenchmarkDay12Part1(b *testing.B) {
	day := 12
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	day := 12
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
