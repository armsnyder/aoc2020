package main

import (
	"testing"
)

func TestDay21Part1(t *testing.T) {
	runDayTests(t, 21, []dayTest{
		{
			input: `
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`,
			want: 5,
		},
	})
}

func TestDay21Part2(t *testing.T) {
	runDayTests(t, 21, []dayTest{
		{
			part2: true,
			input: `
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`,
			want: "mxmxvkd,sqjhc,fvjkl",
		},
	})
}
