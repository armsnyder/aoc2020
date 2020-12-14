package main

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay08Part1(t *testing.T) {
	runDayTests(t, 8, []dayTest{
		{
			input: `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`,
			want: 5,
		},
	})
}

func TestDay08Part2(t *testing.T) {
	runDayTests(t, 8, []dayTest{
		{
			part2: true,
			input: `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`,
			want: 8,
		},
	})
}

func BenchmarkDay08BaselineIO(b *testing.B) {
	day := 8
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		io.Copy(ioutil.Discard, input)
		input.Close()
	}
}

func BenchmarkDay08Part1(b *testing.B) {
	day := 8
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay08Part2(b *testing.B) {
	day := 8
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
