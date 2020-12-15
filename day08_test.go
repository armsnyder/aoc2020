package main

import (
	"testing"
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
