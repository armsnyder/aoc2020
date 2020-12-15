package main

import (
	"testing"
)

func TestDay01Part1(t *testing.T) {
	runDayTests(t, 1, []dayTest{
		{
			input: `
1721
979
366
299
675
1456
`,
			want: 514579,
		},
	})
}

func TestDay01Part2(t *testing.T) {
	runDayTests(t, 1, []dayTest{
		{
			part2: true,
			input: `
1721
979
366
299
675
1456
`,
			want: 241861950,
		},
	})
}
