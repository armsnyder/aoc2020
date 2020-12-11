package main

import (
	"io"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(11, func(part2 bool, inputReader io.Reader) interface{} {
	seats := aocutil.ReadAllStrings(inputReader)

	lastCount := 0

	for {
		day11DoRound(seats, part2)

		count := day11CountOccupiedSeats(seats)

		if lastCount == count {
			break
		}

		lastCount = count
	}

	return lastCount
})

func day11CountOccupiedSeats(seats []string) int {
	total := 0

	for _, line := range seats {
		for _, ch := range line {
			if ch == '#' {
				total++
			}
		}
	}

	return total
}

func day11DoRound(seats []string, part2 bool) {
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

	var positionsToOccupy [][2]int
	var positionsToEmpty [][2]int

	leaveSeatThreshold := 4
	if part2 {
		leaveSeatThreshold = 5
	}

	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[0]); j++ {
			switch seats[i][j] {
			case 'L':
				ok := true

				for _, dir := range directions {
					if day11CheckOccupied(seats, i, j, dir, part2) {
						ok = false
						break
					}
				}

				if ok {
					positionsToOccupy = append(positionsToOccupy, [2]int{i, j})
				}

			case '#':
				total := 0

				for _, dir := range directions {
					if day11CheckOccupied(seats, i, j, dir, part2) {
						total++
					}
				}

				if total >= leaveSeatThreshold {
					positionsToEmpty = append(positionsToEmpty, [2]int{i, j})
				}
			}
		}
	}

	for _, pos := range positionsToOccupy {
		seats[pos[0]] = seats[pos[0]][:pos[1]] + "#" + seats[pos[0]][pos[1]+1:]
	}

	for _, pos := range positionsToEmpty {
		seats[pos[0]] = seats[pos[0]][:pos[1]] + "L" + seats[pos[0]][pos[1]+1:]
	}
}

func day11CheckOccupied(seats []string, i, j int, dir [2]int, sightMode bool) bool {
	i += dir[0]
	j += dir[1]

	for day11CheckIndex(seats, i, j) {
		switch seats[i][j] {
		case '#':
			return true
		case 'L':
			return false
		}

		if !sightMode {
			return false
		}

		i += dir[0]
		j += dir[1]
	}

	return false
}

func day11CheckIndex(seats []string, i, j int) bool {
	if i < 0 || i >= len(seats) {
		return false
	}

	if j < 0 || j >= len(seats[i]) {
		return false
	}

	return true
}
