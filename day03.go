package main

import (
	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(3, func(part2 bool, rawInput []byte) interface{} {
	input := aocutil.Strings(rawInput)

	var slopes [][2]int
	if part2 {
		slopes = [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	} else {
		slopes = [][2]int{{3, 1}}
	}

	result := 1

	for _, slope := range slopes {
		result *= day03CountTreesOnSlope(input, slope)
	}

	return result
})

func day03CountTreesOnSlope(input []string, slope [2]int) int {
	var row, col, total int

	for row < len(input) {
		index := col % len(input[row])

		if input[row][index] == '#' {
			total++
		}

		row += slope[1]
		col += slope[0]
	}

	return total
}
