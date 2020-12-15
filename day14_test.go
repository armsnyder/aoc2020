package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
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

func BenchmarkDay14Part1(b *testing.B) {
	day := 14
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay14Part2(b *testing.B) {
	day := 14
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
