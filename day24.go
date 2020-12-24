package main

import (
	"io"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(24, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day24Part2(inputReader)
	}
	return day24Part1(inputReader)
})

func day24Part1(inputReader io.Reader) interface{} {
	return len(day24InitialBlackTiles(inputReader))
}

func day24Part2(inputReader io.Reader) interface{} {
	blackTiles := day24InitialBlackTiles(inputReader)
	for i := 0; i < 100; i++ {
		day24DoArt(blackTiles)
	}
	return len(blackTiles)
}

func day24InitialBlackTiles(inputReader io.Reader) map[day24Coord]bool {
	blackTiles := make(map[day24Coord]bool)

	aocutil.VisitStrings(inputReader, func(directions string) {
		tile := day24FollowDirectionsToTile(directions)

		if _, ok := blackTiles[tile]; ok {
			delete(blackTiles, tile)
		} else {
			blackTiles[tile] = true
		}
	})

	return blackTiles
}

func day24FollowDirectionsToTile(directions string) day24Coord {
	var result day24Coord

	for i := 0; i < len(directions); i++ {
		switch directions[i] {
		case 'n':
			switch directions[i+1] {
			case 'e':
				result.x++
				result.y--
			case 'w':
				result.y--
			}
			i++
		case 's':
			switch directions[i+1] {
			case 'e':
				result.y++
			case 'w':
				result.x--
				result.y++
			}
			i++
		case 'e':
			result.x++
		case 'w':
			result.x--
		}
	}

	return result
}

func day24DoArt(blackTiles map[day24Coord]bool) {
	whiteTiles := make(map[day24Coord]bool)

	for blackTile := range blackTiles {
		blackNeighbors := 0

		blackTile.visitNeighbors(func(neighbor day24Coord) {
			if blackTiles[neighbor] {
				blackNeighbors++
			} else {
				whiteTiles[neighbor] = true
			}
		})

		if blackNeighbors == 0 || blackNeighbors > 2 {
			defer delete(blackTiles, blackTile)
		}
	}

	for whiteTile := range whiteTiles {
		blackNeighbors := 0

		whiteTile.visitNeighbors(func(neighbor day24Coord) {
			if blackTiles[neighbor] {
				blackNeighbors++
			}
		})

		if blackNeighbors == 2 {
			defer func(whiteTile day24Coord) {
				blackTiles[whiteTile] = true
			}(whiteTile)
		}
	}
}

type day24Coord struct {
	x, y int
}

func (c day24Coord) plus(c2 day24Coord) day24Coord {
	c.x += c2.x
	c.y += c2.y
	return c
}

func (c day24Coord) visitNeighbors(f func(day24Coord)) {
	f(day24Coord{x: c.x + 1, y: c.y - 1})
	f(day24Coord{x: c.x, y: c.y - 1})
	f(day24Coord{x: c.x, y: c.y + 1})
	f(day24Coord{x: c.x - 1, y: c.y + 1})
	f(day24Coord{x: c.x + 1, y: c.y})
	f(day24Coord{x: c.x - 1, y: c.y})
}
