package main

import (
	"testing"
)

func Test_day05(t *testing.T) {
	runDayTests(t, 5, []dayTest{
		{
			name:  "part 1 single",
			input: "FBFBBFFRLR",
			want:  357,
		},
		{
			name: "part 1 many",
			input: `
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL
`,
			want: 820,
		},
		{
			name:  "part 2",
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
