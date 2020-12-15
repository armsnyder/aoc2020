package main

import (
	"testing"
)

func TestDay06Part1(t *testing.T) {
	runDayTests(t, 6, []dayTest{
		{
			input: `
abc

a
b
c

ab
ac

a
a
a
a

b
`,
			want: 11,
		},
	})
}

func TestDay06Part2(t *testing.T) {
	runDayTests(t, 6, []dayTest{
		{
			part2: true,
			input: `
abc

a
b
c

ab
ac

a
a
a
a

b
`,
			want: 6,
		},
	})
}
