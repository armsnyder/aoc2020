package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func Test_day11(t *testing.T) {
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

func BenchmarkDay11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[11](true, aocutil.GetInput(11))
	}
}
