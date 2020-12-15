package main

import (
	"strings"
	"testing"
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
