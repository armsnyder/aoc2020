package main

import (
	"io"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(6, func(part2 bool, inputReader io.Reader) interface{} {
	total := 0
	aocutil.VisitStringGroups(inputReader, func(passengers []string) {
		allYesAnswers := make(map[rune]int)
		for _, passengerAnswers := range passengers {
			for _, ch := range passengerAnswers {
				allYesAnswers[ch]++
			}
		}
		for _, answerCount := range allYesAnswers {
			if !part2 || answerCount == len(passengers) {
				total++
			}
		}
	})
	return total
})
