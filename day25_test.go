package main

import (
	"testing"
)

func TestDay25Part1(t *testing.T) {
	runDayTests(t, 25, []dayTest{
		{
			input: `
17807724
5764801
`,
			want: 14897079,
		},
		{
			input: `
5764801
17807724
`,
			want: 14897079,
		},
	})
}
