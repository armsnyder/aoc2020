package main

import (
	"testing"
)

func Test_day08(t *testing.T) {
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
