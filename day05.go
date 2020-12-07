package main

import (
	"sort"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(5, func(part2 bool, rawInput []byte) interface{} {
	if part2 {
		return day05Part2(rawInput)
	}
	return day05Part1(rawInput)
})

func day05Part1(rawInput []byte) interface{} {
	max := -1
	boardingPasses := aocutil.Strings(rawInput)
	for _, boardingPass := range boardingPasses {
		seatID := day05GetSeatID(boardingPass)
		if seatID > max {
			max = seatID
		}
	}
	return max
}

func day05Part2(rawInput []byte) interface{} {
	var seatIDs []int
	boardingPasses := aocutil.Strings(rawInput)
	for _, boardingPass := range boardingPasses {
		seatIDs = append(seatIDs, day05GetSeatID(boardingPass))
	}
	sort.Ints(seatIDs)
	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1]-seatIDs[i] != 1 {
			return seatIDs[i] + 1
		}
	}
	panic("no solution")
}

func day05GetSeatID(boardingPass string) int {
	seatID := 0
	for i, ch := range boardingPass {
		switch ch {
		case 'B', 'R':
			seatID |= 1 << (9 - i)
		}
	}
	return seatID
}
