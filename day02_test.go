package main

import (
	"testing"
)

func TestDay02Part1(t *testing.T) {
	runDayTests(t, 2, []dayTest{
		{
			input: `
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`,
			want: 2,
		},
	})
}

func TestDay02Part2(t *testing.T) {
	runDayTests(t, 2, []dayTest{
		{
			part2: true,
			input: `
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`,
			want: 1,
		},
	})
}
