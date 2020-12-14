package main

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay07Part1(t *testing.T) {
	runDayTests(t, 7, []dayTest{
		{
			input: `
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`,
			want: 4,
		},
	})
}

func TestDay07Part2(t *testing.T) {
	runDayTests(t, 7, []dayTest{
		{
			part2: true,
			input: `
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`,
			want: 32,
		},
		{
			part2: true,
			input: `
shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`,
			want: 126,
		},
	})
}

func BenchmarkDay07BaselineIO(b *testing.B) {
	day := 7
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		io.Copy(ioutil.Discard, input)
		input.Close()
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	day := 7
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](false, input)
		input.Close()
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	day := 7
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		days[day](true, input)
		input.Close()
	}
}
