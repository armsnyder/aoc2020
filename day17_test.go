package main

import (
	"testing"
)

func TestDay17Part1(t *testing.T) {
	runDayTests(t, 17, []dayTest{
		{
			input: `
.#.
..#
###
`,
			want: 112,
		},
	})
}

func TestDay17Part2(t *testing.T) {
	runDayTests(t, 17, []dayTest{
		{
			part2: true,
			input: `
.#.
..#
###
`,
			want: 848,
		},
	})
}
