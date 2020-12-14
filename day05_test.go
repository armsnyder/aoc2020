package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay05Part1(t *testing.T) {
	runDayTests(t, 5, []dayTest{
		{
			name:  "single",
			input: "FBFBBFFRLR",
			want:  357,
		},
		{
			name: "many",
			input: `
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL
`,
			want: 820,
		},
	})
}

func TestDay05Part2(t *testing.T) {
	runDayTests(t, 5, []dayTest{
		{
			part2: true,
			input: `
FBFBBFFRRR
FBFBBFFRLL
FBFBBFFRLR
`,
			want: 358,
		},
	})
}

func BenchmarkDay05Part1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[5](false, aocutil.GetInput(5))
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[5](true, aocutil.GetInput(5))
	}
}
