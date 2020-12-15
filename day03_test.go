package main

import (
	"testing"
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
