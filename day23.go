package main

import (
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(23, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day23Part2(inputReader)
	}
	return day23Part1(inputReader)
})

func day23Part1(inputReader io.Reader) interface{} {
	game := day23Parse(inputReader, false)

	game.play(100)

	var sb strings.Builder
	for cur := game.lookup[1].next; cur.value != 1; cur = cur.next {
		sb.WriteString(strconv.Itoa(cur.value))
	}
	return sb.String()
}

func day23Part2(inputReader io.Reader) interface{} {
	game := day23Parse(inputReader, true)

	game.play(10_000_000)

	return game.lookup[1].next.value * game.lookup[1].next.next.value
}

func day23Parse(inputReader io.Reader, part2 bool) *day23Game {
	game := &day23Game{
		lookup: make(map[int]*day23Cup),
		max:    math.MinInt32,
		min:    math.MaxInt32,
	}

	input := aocutil.ReadAllStrings(inputReader)[0]
	for _, ch := range input {
		game.addNewCup(int(ch - '0'))
	}

	if part2 {
		for i := game.max + 1; i <= 1_000_000; i++ {
			game.addNewCup(i)
		}
	}

	return game
}

// day23Game contains the game state.
type day23Game struct {
	// head points to the current cup in the doubly-lined list.
	head *day23Cup

	// lookup can look up any cup in the game by its value.
	lookup map[int]*day23Cup

	// max is the maximum cup value in the game.
	max int

	// min is the minimum cup value in the game.
	min int
}

// play plays the provided number of rounds of the game.
func (g *day23Game) play(rounds int) {
	var removed [3]*day23Cup

	for i := 0; i < rounds; i++ {
		g.head.removeNextN(removed[:])
		dest := g.findNextLowestCup(removed[:])
		dest.append(removed[:]...)
		g.head = g.head.next
	}
}

// addNewCup is used when building the game to register a new cup and add it before the head.
func (g *day23Game) addNewCup(value int) {
	cup := &day23Cup{value: value}

	g.lookup[value] = cup

	if g.head == nil {
		g.head = cup
		cup.next = cup
		cup.prev = cup
	} else {
		g.head.prev.append(cup)
	}

	if value > g.max {
		g.max = value
	}

	if value < g.min {
		g.min = value
	}
}

// findNextLowestCup finds the cup with the next lowest value than the head. It will not choose a
// cup if it is contained in the provided ignoreList.
func (g *day23Game) findNextLowestCup(ignoreList []*day23Cup) *day23Cup {
outer:
	for i := g.head.value - 1; ; i-- {
		if i < g.min {
			i = g.max
		}

		for _, ignoredCup := range ignoreList {
			if i == ignoredCup.value {
				continue outer
			}
		}

		return g.lookup[i]
	}
}

// day23Cup is a node in a doubly linked list.
type day23Cup struct {
	next  *day23Cup
	prev  *day23Cup
	value int
}

// append appends one or more cups after this cup.
func (c *day23Cup) append(cups ...*day23Cup) {
	first := cups[0]
	last := cups[len(cups)-1]

	c.next, c.next.prev, last.next, first.prev = first, last, c.next, c
}

// removeNextN removes a number of cups after this cup, equal to the length of the provided slice,
// and populates the slice with the removed cups.
func (c *day23Cup) removeNextN(dest []*day23Cup) {
	cur := c.next
	for i := 0; i < len(dest); i++ {
		dest[i] = cur
		cur = cur.next
	}

	last := dest[len(dest)-1]

	c.next = last.next
	last.prev = c

	dest[0].prev = nil
	last.next = nil
}
