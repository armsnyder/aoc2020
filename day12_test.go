package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func Test_day12(t *testing.T) {
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

func BenchmarkDay12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[12](true, aocutil.GetInput(12))
	}
}
