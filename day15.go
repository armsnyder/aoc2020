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

func day15(inputReader io.Reader, n int32) int {
	starter := aocutil.ReadAllCommaSeparatedInts(inputReader)

	// Keep track of previous turn (1-indexed) where each number was seen.
	memory := make([]int32, n)

	for turn := int32(1); turn < int32(len(starter)); turn++ {
		memory[starter[turn-1]] = turn
	}

	prevNumber := int32(starter[len(starter)-1])

	for turn := int32(len(starter) + 1); turn <= n; turn++ {
		prevTurn := turn - 1
		var number int32

		if prevPrevTurn := memory[prevNumber]; prevPrevTurn > 0 {
			number = prevTurn - prevPrevTurn
		}

		memory[prevNumber] = prevTurn
		prevNumber = number
	}

	return int(prevNumber)
}
