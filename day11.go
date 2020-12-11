package main

import (
	"io"
	"sort"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(11, func(part2 bool, inputReader io.Reader) interface{} {
	seats := aocutil.ReadAllStrings(inputReader)
	count := 0

	for {
		delta := day11DoRound(seats, part2)

		if delta == 0 {
			return count
		}

		count += delta
	}
})

func day11DoRound(seats []string, part2 bool) (occupiedSeatDelta int) {
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

	positionsToOccupy := make([][]int, len(seats))
	positionsToEmpty := make([][]int, len(seats))

	leaveSeatThreshold := 4
	if part2 {
		leaveSeatThreshold = 5
	}

	for i := 0; i < len(seats); i++ {
	outer:
		for j := 0; j < len(seats[0]); j++ {
			switch seats[i][j] {
			case 'L':
				for _, dir := range directions {
					if day11CheckOccupied(seats, i, j, dir, part2) {
						continue outer
					}
				}

				positionsToOccupy[i] = append(positionsToOccupy[i], j)

			case '#':
				total := 0

				for _, dir := range directions {
					if day11CheckOccupied(seats, i, j, dir, part2) {
						total++

						if total >= leaveSeatThreshold {
							positionsToEmpty[i] = append(positionsToEmpty[i], j)
							break
						}
					}
				}
			}
		}
	}

	var sb strings.Builder

	for i := 0; i < len(seats); i++ {
		if len(positionsToOccupy[i]) > 0 || len(positionsToEmpty[i]) > 0 {
			seats[i] = day11RebuildRow(seats[i], positionsToOccupy[i], positionsToEmpty[i], &sb)
			occupiedSeatDelta += len(positionsToOccupy[i]) - len(positionsToEmpty[i])
		}
	}

	return occupiedSeatDelta
}

func day11CheckOccupied(seats []string, i, j int, dir [2]int, sightMode bool) bool {
	for {
		i += dir[0]
		j += dir[1]

		if !day11CheckIndex(seats, i, j) {
			return false
		}

		switch seats[i][j] {
		case '#':
			return true
		case 'L':
			return false
		}

		if !sightMode {
			return false
		}
	}
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

func day11RebuildRow(row string, positionsToOccupy, positionsToEmpty []int, sb *strings.Builder) string {
	sb.Reset()

	sort.Ints(positionsToOccupy)
	sort.Ints(positionsToEmpty)

	occupyI := 0
	emptyI := 0

	for i := 0; i < len(row); i++ {
		if occupyI < len(positionsToOccupy) && positionsToOccupy[occupyI] == i {
			sb.WriteRune('#')
			occupyI++
		} else if emptyI < len(positionsToEmpty) && positionsToEmpty[emptyI] == i {
			sb.WriteRune('L')
			emptyI++
		} else {
			sb.WriteByte(row[i])
		}
	}

	return sb.String()
}
