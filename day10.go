package main

import (
	"io"
	"sort"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(10, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day10Part2(inputReader)
	}

	return day10Part1(inputReader)
})

func day10Part1(inputReader io.Reader) int {
	jolts := day10SortedJolts(inputReader)
	joltDiffs1 := 0
	joltDiffs3 := 1

	for i := 1; i < len(jolts); i++ {
		diff := jolts[i] - jolts[i-1]

		switch diff {
		case 1:
			joltDiffs1++
		case 3:
			joltDiffs3++
		}
	}

	return joltDiffs1 * joltDiffs3
}

func day10Part2(inputReader io.Reader) int {
	jolts := day10SortedJolts(inputReader)
	waysFromIndex := make([]int, len(jolts))
	waysFromIndex[len(jolts)-1] = 1

	for i := len(jolts) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0 && jolts[i]-jolts[j] <= 3; j-- {
			waysFromIndex[j] += waysFromIndex[i]
		}
	}

	return waysFromIndex[0]
}

func day10SortedJolts(inputReader io.Reader) []int {
	jolts := aocutil.ReadAllInts(inputReader)
	jolts = append(jolts, 0)
	sort.Ints(jolts)
	return jolts
}
