package main

import (
	"testing"
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
