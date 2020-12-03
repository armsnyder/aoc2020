package main

import (
	"testing"
)

func Test_day02(t *testing.T) {
	runDayTests(t, 2, []dayTest{
		{
			input: []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"},
			want:  2,
		},
		{
			part2: true,
			input: []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"},
			want:  1,
		},
	})
}
