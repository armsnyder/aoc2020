package main

import (
	"bufio"
	"io"
)

var _ = declareDay(11, func(part2 bool, inputReader io.Reader) interface{} {
	seats := day11Read2DByteArray(inputReader)

	seatChangeBuffer := make([][]byte, len(seats))
	for i, row := range seats {
		seatChangeBuffer[i] = make([]byte, len(row))
	}

	for day11DoRound(seats, seatChangeBuffer, part2) {
	}

	return day11CountOccupied(seats)
})

func day11Read2DByteArray(inputReader io.Reader) (result [][]byte) {
	scanner := bufio.NewScanner(inputReader)
	scanner.Buffer(make([]byte, 93), bufio.MaxScanTokenSize)
	for scanner.Scan() {
		if len(scanner.Bytes()) > 0 {
			row := make([]byte, len(scanner.Bytes()))
			copy(row, scanner.Bytes())
			result = append(result, row)
		}
	}
	return
}

func day11CountOccupied(seats [][]byte) (total int) {
	for _, row := range seats {
		for _, ch := range row {
			if ch == '#' {
				total++
			}
		}
	}

	return total
}

func day11DoRound(seats, seatChangeBuffer [][]byte, part2 bool) (changed bool) {
	leaveSeatThreshold := 4
	if part2 {
		leaveSeatThreshold = 5
	}

	for i := 0; i < len(seats); i++ {
	outer:
		for j := 0; j < len(seats[0]); j++ {
			switch seats[i][j] {
			case 'L':
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						if di == 0 && dj == 0 {
							continue
						}
						if day11CheckOccupied(seats, i, j, di, dj, part2) {
							seatChangeBuffer[i][j] = 0
							continue outer
						}
					}
				}

				seatChangeBuffer[i][j] = '#'

			case '#':
				total := 0

				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						if di == 0 && dj == 0 {
							continue
						}
						if day11CheckOccupied(seats, i, j, di, dj, part2) {
							total++

							if total >= leaveSeatThreshold {
								seatChangeBuffer[i][j] = 'L'
								continue outer
							}
						}
					}
				}

				seatChangeBuffer[i][j] = 0
			}
		}
	}

	for i, row := range seatChangeBuffer {
		for j, ch := range row {
			switch ch {
			case 'L':
				seats[i][j] = 'L'
				changed = true
			case '#':
				seats[i][j] = '#'
				changed = true
			}
		}
	}

	return changed
}

func day11CheckOccupied(seats [][]byte, i, j, di, dj int, sightMode bool) bool {
	for {
		i += di
		j += dj

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

func day11CheckIndex(seats [][]byte, i, j int) bool {
	if i < 0 || i >= len(seats) {
		return false
	}

	if j < 0 || j >= len(seats[i]) {
		return false
	}

	return true
}
