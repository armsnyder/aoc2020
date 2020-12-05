package main

import (
	"io"
	"sort"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(5, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day05Part2(inputReader)
	}
	return day05Part1(inputReader)
})

func day05Part1(inputReader io.Reader) interface{} {
	max := -1
	aocutil.VisitStrings(inputReader, func(s string) {
		seatID := day05GetSeatID(s)
		if seatID > max {
			max = seatID
		}
	})
	return max
}

func day05Part2(inputReader io.Reader) interface{} {
	var seats []int
	aocutil.VisitStrings(inputReader, func(s string) {
		seats = append(seats, day05GetSeatID(s))
	})
	sort.Ints(seats)
	for i := 0; i < len(seats)-1; i++ {
		if seats[i+1]-seats[i] != 1 {
			return seats[i] + 1
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
