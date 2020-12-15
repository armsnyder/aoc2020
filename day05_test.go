package main

import (
	"testing"
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
