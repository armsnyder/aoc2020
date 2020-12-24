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

	aocutil.VisitStrings(inputReader, func(v string) {
		var pos day24Coord

		day24VisitDirections(v, func(dir day24Coord) {
			pos = pos.plus(dir)
		})

		if _, ok := blackTiles[pos]; ok {
			delete(blackTiles, pos)
		} else {
			blackTiles[pos] = true
		}
	})

	return blackTiles
}

func day24VisitDirections(v string, f func(dir day24Coord)) {
	for i := 0; i < len(v); i++ {
		switch v[i] {
		case 'n':
			switch v[i+1] {
			case 'e':
				f(day24Coord{x: 1, y: -1})
			case 'w':
				f(day24Coord{y: -1})
			}
			i++
		case 's':
			switch v[i+1] {
			case 'e':
				f(day24Coord{y: 1})
			case 'w':
				f(day24Coord{x: -1, y: 1})
			}
			i++
		case 'e':
			f(day24Coord{x: 1})
		case 'w':
			f(day24Coord{x: -1})
		}
	}
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
