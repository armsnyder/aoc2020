package main

import (
	"io"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(17, func(part2 bool, inputReader io.Reader) interface{} {
	initialState := aocutil.Read2DByteArray(inputReader)
	activeCubes := make(map[day17Coord]bool)

	for y, row := range initialState {
		for x, cube := range row {
			if cube == '#' {
				activeCubes[day17Coord{x: x, y: y}] = true
			}
		}
	}

	for i := 1; i <= 6; i++ {
		day17DoCycle(activeCubes, part2)
	}

	return len(activeCubes)
})

func day17DoCycle(activeCubes map[day17Coord]bool, part2 bool) {
	inactiveCubes := make(map[day17Coord]bool)

	for coord := range activeCubes {
		activeNeighbors := 0

		day17VisitNeighbors(coord, part2, func(neighborCoord day17Coord) {
			if activeCubes[neighborCoord] {
				activeNeighbors++
			} else {
				inactiveCubes[neighborCoord] = true
			}
		})

		if activeNeighbors != 2 && activeNeighbors != 3 {
			defer delete(activeCubes, coord)
		}
	}

	for coord := range inactiveCubes {
		activeNeighbors := 0

		day17VisitNeighbors(coord, part2, func(neighborCoord day17Coord) {
			if activeCubes[neighborCoord] {
				activeNeighbors++
			}
		})

		if activeNeighbors == 3 {
			defer func(coord day17Coord) {
				activeCubes[coord] = true
			}(coord)
		}
	}
}

type day17Coord struct{ x, y, z, w int }

func day17VisitNeighbors(coord day17Coord, part2 bool, visitFn func(neighborCoord day17Coord)) {
	minW, maxW := 0, 0
	if part2 {
		minW, maxW = -1, 1
	}

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := minW; dw <= maxW; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}

					visitFn(day17Coord{coord.x + dx, coord.y + dy, coord.z + dz, coord.w + dw})
				}
			}
		}
	}
}
