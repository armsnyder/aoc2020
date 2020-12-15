package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay11Part1(t *testing.T) {
	runDayTests(t, 11, []dayTest{
		{
			input: `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`,
			want: 37,
		},
	})
}

func TestDay11Part2(t *testing.T) {
	runDayTests(t, 11, []dayTest{
		{
			part2: true,
			input: `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`,
			want: 26,
		},
	})
}

func BenchmarkDay11Part1(b *testing.B) {
	day := 11
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	day := 11
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
