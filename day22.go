package main

import (
	"io"
	"strconv"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(22, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day22Part2(inputReader)
	}
	return day22Part1(inputReader)
})

func day22Part1(inputReader io.Reader) interface{} {
	player1Deck, player2Deck := day22Parse(inputReader)

	for {
		winner, loser := player1Deck, player2Deck
		if player2Deck.head.value > player1Deck.head.value {
			winner, loser = player2Deck, player1Deck
		}

		if !day22TurnOver(winner, loser) {
			return winner.score()
		}
	}
}

func day22Part2(inputReader io.Reader) interface{} {
	player1Deck, player2Deck := day22Parse(inputReader)

	if day22PlayRecursive(player1Deck, player2Deck) {
		return player1Deck.score()
	}

	return player2Deck.score()
}

func day22PlayRecursive(player1Deck, player2Deck *day22Deck) (winnerIsPlayer1 bool) {
	seen := make(map[[2][5]int64]bool)

	for {
		hash := [2][5]int64{player1Deck.hash(), player2Deck.hash()}
		if seen[hash] {
			return true
		}
		seen[hash] = true

		if player1Deck.cardCount > player1Deck.head.value && player2Deck.cardCount > player2Deck.head.value {
			winnerIsPlayer1 = day22PlayRecursive(player1Deck.copyForRecursion(), player2Deck.copyForRecursion())
		} else {
			winnerIsPlayer1 = player1Deck.head.value > player2Deck.head.value
		}

		winner, loser := player1Deck, player2Deck
		if !winnerIsPlayer1 {
			winner, loser = player2Deck, player1Deck
		}

		if !day22TurnOver(winner, loser) {
			return winnerIsPlayer1
		}
	}
}

func day22TurnOver(winner, loser *day22Deck) (canContinue bool) {
	winner.rotate()
	lostCard := loser.pop()
	winner.append(lostCard)
	return loser.cardCount > 0
}

func day22Parse(inputReader io.Reader) (player1Deck, player2Deck *day22Deck) {
	player1Deck = &day22Deck{}
	player2Deck = &day22Deck{}

	aocutil.VisitStringGroups(inputReader, func(ss []string) {
		playerNumber, err := strconv.Atoi(ss[0][7:8])
		if err != nil {
			panic(err)
		}

		deck := player1Deck
		if playerNumber == 2 {
			deck = player2Deck
		}

		for _, s := range ss[1:] {
			cardValue, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			deck.append(&day22Card{value: cardValue})
		}
	})

	return player1Deck, player2Deck
}

type day22Card struct {
	value int
	next  *day22Card
}

type day22Deck struct {
	head      *day22Card
	tail      *day22Card
	cardCount int
}

func (d *day22Deck) append(c *day22Card) {
	if d.tail == nil {
		d.tail = c
		d.head = c
	}
	c.next, d.tail, d.tail.next = d.head, c, c
	d.cardCount++
}

func (d *day22Deck) pop() *day22Card {
	popped := d.head
	d.head, d.tail.next, d.head.next = d.head.next, d.head.next, nil
	d.cardCount--
	return popped
}

func (d *day22Deck) rotate() {
	d.head, d.tail = d.head.next, d.head
}

func (d *day22Deck) score() int {
	total := 0
	cur := d.head

	for multiplier := d.cardCount; multiplier > 0; multiplier-- {
		total += cur.value * multiplier
		cur = cur.next
	}

	return total
}

// hash returns a unique value representing the order and cards in the deck.
// It assumes the deck size is <= 50 and the maximum card value is <= 64.
func (d *day22Deck) hash() [5]int64 {
	var hash [5]int64
	cur := d.head

	for i := 0; i < d.cardCount; i++ {
		hash[i/10] |= int64(cur.value) << ((i % 10) * 6)
		cur = cur.next
	}

	return hash
}

// copyForRecursion returns a copy of the next n cards in the deck, where n is the value of the head
// card.
func (d *day22Deck) copyForRecursion() *day22Deck {
	result := &day22Deck{}
	cur := d.head.next

	for i := 0; i < d.head.value; i++ {
		result.append(&day22Card{value: cur.value})
		cur = cur.next
	}

	return result
}
