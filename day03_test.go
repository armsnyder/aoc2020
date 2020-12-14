package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay03Part1(t *testing.T) {
	runDayTests(t, 3, []dayTest{
		{
			input: `
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`,
			want: 7,
		},
	})
}

func TestDay03Part2(t *testing.T) {
	runDayTests(t, 3, []dayTest{
		{
			part2: true,
			input: `
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`,
			want: 336,
		},
	})
}

func BenchmarkDay03Part1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[3](false, aocutil.GetInput(3))
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[3](true, aocutil.GetInput(3))
	}
}
