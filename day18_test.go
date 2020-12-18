package main

import (
	"testing"
)

func TestDay18Part1(t *testing.T) {
	runDayTests(t, 18, []dayTest{
		{
			input: `1 + 2 * 3 + 4 * 5 + 6`,
			want:  71,
		},
		{
			input: `1 + (2 * 3) + (4 * (5 + 6))`,
			want:  51,
		},
		{
			input: `2 * 3 + (4 * 5)`,
			want:  26,
		},
		{
			input: `5 + (8 * 3 + 9 + 3 * 4 * 3)`,
			want:  437,
		},
		{
			input: `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`,
			want:  12240,
		},
		{
			input: `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`,
			want:  13632,
		},
		{
			input: `
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
`,
			want: 26335,
		},
	})
}

func TestDay18Part2(t *testing.T) {
	runDayTests(t, 18, []dayTest{
		{
			part2: true,
			input: `2 + 3`,
			want:  5,
		},
		{
			part2: true,
			input: `2 * 3`,
			want:  6,
		},
		{
			part2: true,
			input: `2 * 3 + 4`,
			want:  14,
		},
		{
			part2: true,
			input: `2 + 3 * 4`,
			want:  20,
		},
		{
			part2: true,
			input: `(2 * 3) + 4`,
			want:  10,
		},
		{
			part2: true,
			input: `2 * (3 + 4)`,
			want:  14,
		},
		{
			part2: true,
			input: `1 + 2 * 3 + 4 * 5 + 6`,
			want:  231,
		},
		{
			part2: true,
			input: `1 + (2 * 3) + (4 * (5 + 6))`,
			want:  51,
		},
		{
			part2: true,
			input: `2 * 3 + (4 * 5)`,
			want:  46,
		},
		{
			part2: true,
			input: `5 + (8 * 3 + 9 + 3 * 4 * 3)`,
			want:  1445,
		},
		{
			part2: true,
			input: `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`,
			want:  669060,
		},
		{
			part2: true,
			input: `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`,
			want:  23340,
		},
	})
}
