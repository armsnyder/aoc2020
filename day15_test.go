package main

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay15Part1(t *testing.T) {
	runDayTests(t, 15, []dayTest{
		{
			input: `0,3,6`,
			want:  436,
		},
		{
			input: `1,3,2`,
			want:  1,
		},
		{
			input: `2,1,3`,
			want:  10,
		},
		{
			input: `1,2,3`,
			want:  27,
		},
		{
			input: `2,3,1`,
			want:  78,
		},
		{
			input: `3,2,1`,
			want:  438,
		},
		{
			input: `3,1,2`,
			want:  1836,
		},
	})
}

func TestDay15Part2(t *testing.T) {
	runDayTests(t, 15, []dayTest{
		{
			part2: true,
			input: `0,3,6`,
			want:  175594,
		},
	})
}

func BenchmarkDay15BaselineIO(b *testing.B) {
	day := 15
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		io.Copy(ioutil.Discard, input)
		input.Close()
	}
}

func BenchmarkDay15Part1(b *testing.B) {
	day := 15
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay15Part2(b *testing.B) {
	day := 15
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
