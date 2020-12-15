package main

import (
	"io"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(15, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day15(inputReader, 30000000)
	}
	return day15(inputReader, 2020)
})

func day15(inputReader io.Reader, n int) int {
	starter := aocutil.ReadAllCommaSeparatedInts(inputReader)

	// Keep track of previous turn (1-indexed) where each number was seen.
	memory := make([]int, n)

	for turn := 1; turn < len(starter); turn++ {
		memory[starter[turn-1]] = turn
	}

	prevNumber := starter[len(starter)-1]

	for turn := len(starter) + 1; turn <= n; turn++ {
		prevTurn := turn - 1
		var number int

		if prevPrevTurn := memory[prevNumber]; prevPrevTurn > 0 {
			number = prevTurn - prevPrevTurn
		}

		memory[prevNumber] = prevTurn
		prevNumber = number
	}

	return prevNumber
}
