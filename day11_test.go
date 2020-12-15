package main

import (
	"testing"
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
