package main

import (
	"testing"
)

func TestDay23Part1(t *testing.T) {
	runDayTests(t, 23, []dayTest{
		{
			input: "389125467",
			want:  "67384529",
		},
	})
}

func TestDay23Part2(t *testing.T) {
	runDayTests(t, 23, []dayTest{
		{
			part2: true,
			input: "389125467",
			want:  149245887792,
		},
	})
}
