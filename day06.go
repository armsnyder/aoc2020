package main

import (
	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(6, func(part2 bool, rawInput []byte) interface{} {
	total := 0
	passengerGroups := aocutil.StringGroups(rawInput)
	for _, passengers := range passengerGroups {
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
	}
	return total
})
