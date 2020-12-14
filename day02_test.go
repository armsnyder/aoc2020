package main

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
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

func BenchmarkDay02BaselineIO(b *testing.B) {
	day := 2
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		io.Copy(ioutil.Discard, input)
		input.Close()
	}
}

func BenchmarkDay02Part1(b *testing.B) {
	day := 2
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	day := 2
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
