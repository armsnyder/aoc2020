package main

import (
	"testing"
)

func Test_day01(t *testing.T) {
	runDayTests(t, 1, []dayTest{
		{
			input: []int{1721, 979, 366, 299, 675, 1456},
			want:  514579,
		},
		{
			part2: true,
			input: []int{1721, 979, 366, 299, 675, 1456},
			want:  241861950,
		},
	})
}
