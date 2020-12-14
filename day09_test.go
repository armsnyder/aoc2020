package main

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay09Part1(t *testing.T) {
	inputReader := strings.NewReader(`
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`)
	got := day09Part1(inputReader, 5)
	want := 127
	if got != want {
		t.Errorf("day09Part1() = %v, want %v", got, want)
	}
}

func TestDay09Part2(t *testing.T) {
	inputReader := strings.NewReader(`
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`)
	got := day09Part2(inputReader, 5)
	want := 62
	if got != want {
		t.Errorf("day09Part1() = %v, want %v", got, want)
	}
}

func BenchmarkDay09BaselineIO(b *testing.B) {
	day := 9
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		io.Copy(ioutil.Discard, input)
		input.Close()
	}
}

func BenchmarkDay09Part1(b *testing.B) {
	day := 9
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay09Part2(b *testing.B) {
	day := 9
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
