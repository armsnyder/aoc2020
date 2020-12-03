package main

var _ = declareDay(3, func(part2 bool, inputUntyped interface{}) interface{} {
	input := inputUntyped.([]string)

	slopes := [][2]int{{3, 1}}
	if part2 {
		slopes = append(slopes, [][2]int{{1, 1}, {5, 1}, {7, 1}, {1, 2}}...)
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
