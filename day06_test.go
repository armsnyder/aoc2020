package main

import (
	"testing"
)

func Test_day06(t *testing.T) {
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
