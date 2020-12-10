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
	joltDiffs3 := 0

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
	memo := make([]int, len(jolts))

	return day10CountCombinations(jolts, 0, memo)
}

func day10SortedJolts(inputReader io.Reader) []int {
	jolts := aocutil.ReadAllInts(inputReader)

	sort.Ints(jolts)

	jolts = append([]int{0}, jolts...)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	return jolts
}

func day10CountCombinations(jolts []int, start int, memo []int) (total int) {
	if memo[start] > 0 {
		return memo[start]
	}

	defer func() {
		memo[start] = total
	}()

	if start == len(jolts)-1 {
		return 1
	}

	for i := start + 1; i < len(jolts) && jolts[i]-jolts[start] <= 3; i++ {
		total += day10CountCombinations(jolts, i, memo)
	}

	return total
}
