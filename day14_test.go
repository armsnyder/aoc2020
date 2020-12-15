package main

import (
	"testing"
)

func TestDay14Part1(t *testing.T) {
	runDayTests(t, 14, []dayTest{
		{
			input: `
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`,
			want: int64(165),
		},
	})
}

func TestDay14Part2(t *testing.T) {
	runDayTests(t, 14, []dayTest{
		{
			part2: true,
			input: `
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`,
			want: int64(208),
		},
	})
}
