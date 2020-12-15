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
	memory := make(map[int]int)

	for turn := 0; turn < len(starter)-1; turn++ {
		memory[starter[turn]] = turn
	}

	prevNumber := starter[len(starter)-1]

	for turn := len(starter); turn < n; turn++ {
		prevTurn := turn - 1
		var number int

		if prevPrevTurn, ok := memory[prevNumber]; ok {
			number = prevTurn - prevPrevTurn
		}

		memory[prevNumber] = prevTurn
		prevNumber = number
	}

	return prevNumber
}
