package main

import (
	"testing"
)

func TestDay22Part1(t *testing.T) {
	runDayTests(t, 22, []dayTest{
		{
			input: `
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`,
			want: 306,
		},
	})
}

func TestDay22Part2(t *testing.T) {
	runDayTests(t, 22, []dayTest{
		{
			part2: true,
			input: `
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`,
			want: 291,
		},
	})
}
