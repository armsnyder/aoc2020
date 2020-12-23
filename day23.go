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
	game := day23Parse(inputReader)

	var removed [3]int
	var chainStart *day23Cup
	var chainEnd *day23Cup
	for move := 0; move < 100; move++ {
		toRemove := game.cur.next
		for i := 0; i < 3; i++ {
			if i == 0 {
				chainStart = toRemove
			} else if i == 2 {
				chainEnd = toRemove
			}
			removed[i] = toRemove.value
			toRemove = toRemove.next
		}
		var destination *day23Cup
	countdown:
		for i := game.cur.value - 1; i >= game.min; i-- {
			if cup, ok := game.lookup[i]; ok {
				for _, v := range removed {
					if v == i {
						continue countdown
					}
				}
				destination = cup
				break
			}
		}
		if destination == nil {
		countdown2:
			for i := game.max; ; i-- {
				if cup, ok := game.lookup[i]; ok {
					for _, v := range removed {
						if v == i {
							continue countdown2
						}
					}
					destination = cup
					break
				}
			}
		}
		oldChainEndNext := chainEnd.next
		oldDestNext := destination.next
		destination.next = chainStart
		chainStart.prev = destination
		oldDestNext.prev = chainEnd
		chainEnd.next = oldDestNext
		game.cur.next = oldChainEndNext
		oldChainEndNext.prev = game.cur
		game.cur = oldChainEndNext
	}

	var sb strings.Builder

	cur := game.lookup[1].next
	for i := 1; i < game.len; i++ {
		sb.WriteString(strconv.Itoa(cur.value))
		cur = cur.next
	}

	return sb.String()
}

func day23Part2(inputReader io.Reader) interface{} {
	game := day23Parse2(inputReader)

	var removed [3]int
	var chainStart *day23Cup
	var chainEnd *day23Cup
	for move := 0; move < 10_000_000; move++ {
		toRemove := game.cur.next
		for i := 0; i < 3; i++ {
			if i == 0 {
				chainStart = toRemove
			} else if i == 2 {
				chainEnd = toRemove
			}
			removed[i] = toRemove.value
			toRemove = toRemove.next
		}
		var destination *day23Cup
	countdown:
		for i := game.cur.value - 1; i >= game.min; i-- {
			if cup, ok := game.lookup[i]; ok {
				for _, v := range removed {
					if v == i {
						continue countdown
					}
				}
				destination = cup
				break
			}
		}
		if destination == nil {
		countdown2:
			for i := game.max; ; i-- {
				if cup, ok := game.lookup[i]; ok {
					for _, v := range removed {
						if v == i {
							continue countdown2
						}
					}
					destination = cup
					break
				}
			}
		}
		oldChainEndNext := chainEnd.next
		oldDestNext := destination.next
		destination.next = chainStart
		chainStart.prev = destination
		oldDestNext.prev = chainEnd
		chainEnd.next = oldDestNext
		game.cur.next = oldChainEndNext
		oldChainEndNext.prev = game.cur
		game.cur = oldChainEndNext
	}

	return game.lookup[1].next.value * game.lookup[1].next.next.value
}

type day23Cup struct {
	next  *day23Cup
	prev  *day23Cup
	value int
}

type day23Game struct {
	cur    *day23Cup
	lookup map[int]*day23Cup
	max    int
	min    int
	len    int
}

func (g *day23Game) add(value int) {
	g.len++
	cup := &day23Cup{value: value}
	g.lookup[value] = cup
	if g.cur == nil {
		g.cur = cup
		cup.next = cup
		cup.prev = cup
	} else {
		prevPrev := g.cur.prev
		prevPrev.next = cup
		g.cur.prev = cup
		cup.next = g.cur
		cup.prev = prevPrev
	}
	if value > g.max {
		g.max = value
	}
	if value < g.min {
		g.min = value
	}
}

func day23Parse(inputReader io.Reader) *day23Game {
	game := &day23Game{
		lookup: make(map[int]*day23Cup),
		max:    math.MinInt32,
		min:    math.MaxInt32,
	}
	input := aocutil.ReadAllStrings(inputReader)[0]
	for _, ch := range input {
		value, err := strconv.Atoi(string(ch))
		if err != nil {
			panic(err)
		}
		game.add(value)
	}
	return game
}

func day23Parse2(inputReader io.Reader) *day23Game {
	game := &day23Game{
		lookup: make(map[int]*day23Cup),
		max:    math.MinInt32,
		min:    math.MaxInt32,
	}
	input := aocutil.ReadAllStrings(inputReader)[0]
	for _, ch := range input {
		value, err := strconv.Atoi(string(ch))
		if err != nil {
			panic(err)
		}
		game.add(value)
	}
	for i := game.max + 1; i <= 1_000_000; i++ {
		game.add(i)
	}
	return game
}
