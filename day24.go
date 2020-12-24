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
	state := day24InitialState(inputReader)
	return day24CountBlackTiles(state)
}

func day24CountBlackTiles(state map[day24Coord]bool) int {
	result := 0
	for _, v := range state {
		if v {
			result++
		}
	}
	return result
}

func day24InitialState(inputReader io.Reader) map[day24Coord]bool {
	state := make(map[day24Coord]bool)
	aocutil.VisitStrings(inputReader, func(v string) {
		var pos day24Coord
		day24VisitDirections(v, pos.add)
		if _, ok := state[pos]; ok {
			delete(state, pos)
		} else {
			state[pos] = true
		}
	})
	return state
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

func day24Part2(inputReader io.Reader) interface{} {
	state := day24InitialState(inputReader)
	for i := 0; i < 100; i++ {
		day24Process(state)
	}
	return day24CountBlackTiles(state)
}

func day24Process(blackTiles map[day24Coord]bool) {
	whiteTiles := make(map[day24Coord]bool)

	for blackTile := range blackTiles {
		blackNeighbors := 0

		for _, neighbor := range blackTile.neighbors() {
			if blackTiles[neighbor] {
				blackNeighbors++
			} else {
				whiteTiles[neighbor] = true
			}
		}

		if blackNeighbors == 0 || blackNeighbors > 2 {
			defer delete(blackTiles, blackTile)
		}
	}

	for whiteTile := range whiteTiles {
		blackNeighbors := 0

		for _, neighbor := range whiteTile.neighbors() {
			if blackTiles[neighbor] {
				blackNeighbors++
			}
		}

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

func (c *day24Coord) add(c2 day24Coord) {
	c.x += c2.x
	c.y += c2.y
}

func (c day24Coord) neighbors() [6]day24Coord {
	return [6]day24Coord{
		{x: c.x + 1, y: c.y - 1},
		{x: c.x, y: c.y - 1},
		{x: c.x, y: c.y + 1},
		{x: c.x - 1, y: c.y + 1},
		{x: c.x + 1, y: c.y},
		{x: c.x - 1, y: c.y},
	}
}
