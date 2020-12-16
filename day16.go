package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(16, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		sortedRuleNames, yourTicket := day16Part2(inputReader)
		result := 1

		for i, ruleName := range sortedRuleNames {
			if strings.HasPrefix(ruleName, "departure") {
				result *= yourTicket[i]
			}
		}

		return result
	}

	return day16Part1(inputReader)
})

func day16Part1(inputReader io.Reader) interface{} {
	rules, _, nearbyTickets := day16Parse(inputReader)

	errorRate := 0

	for _, ticket := range nearbyTickets {
	outer:
		for _, v := range ticket {
			for _, rule := range rules {
				if rule(v) {
					continue outer
				}
			}
			errorRate += v
		}
	}

	return errorRate
}

func day16Part2(inputReader io.Reader) ([]string, []int) {
	rules, yourTicket, nearbyTickets := day16Parse(inputReader)

	var validTickets [][]int

nextTicket:
	for _, ticket := range nearbyTickets {
	nextField:
		for _, v := range ticket {
			for _, rule := range rules {
				if rule(v) {
					continue nextField
				}
			}
			continue nextTicket
		}
		validTickets = append(validTickets, ticket)
	}

	sortedRuleNames := day16Helper(make([]string, len(rules)), rules, validTickets)
	if sortedRuleNames == nil {
		panic("no solution")
	}

	return sortedRuleNames, yourTicket
}

func day16CheckRuleIndex(rule func(int) bool, tickets [][]int, index int) bool {
	for _, ticket := range tickets {
		if !rule(ticket[index]) {
			return false
		}
	}
	return true
}

func day16Helper(assumptions []string, rules map[string]func(int) bool, tickets [][]int) []string {
	nextUnsolvedIndex := -1
	for i := 0; i < len(rules); i++ {
		if assumptions[i] == "" {
			nextUnsolvedIndex = i
			break
		}
	}

	if nextUnsolvedIndex == -1 {
		return assumptions
	}

nextRule:
	for ruleName, ruleFn := range rules {
		for _, ass := range assumptions {
			if ass == ruleName {
				continue nextRule
			}
		}

		if day16CheckRuleIndex(ruleFn, tickets, nextUnsolvedIndex) {
			nextAssumptions := make([]string, len(assumptions))
			copy(nextAssumptions, assumptions)
			nextAssumptions[nextUnsolvedIndex] = ruleName
			solution := day16Helper(nextAssumptions, rules, tickets)
			if solution != nil {
				return solution
			}
		}
	}

	return nil
}

func day16Parse(inputReader io.Reader) (rules map[string]func(int) bool, yourTicket []int, nearbyTickets [][]int) {
	rules = make(map[string]func(int) bool)
	status := 0

	aocutil.VisitStrings(inputReader, func(v string) {
		if v == "your ticket:" {
			status = 1
			return
		}

		if v == "nearby tickets:" {
			status = 2
			return
		}

		switch status {
		case 0:
			split1 := strings.SplitN(v, ": ", 2)
			split2 := strings.SplitN(split1[1], " or ", 2)
			split3 := strings.SplitN(split2[0], "-", 2)
			split4 := strings.SplitN(split2[1], "-", 2)
			a, _ := strconv.Atoi(split3[0])
			b, _ := strconv.Atoi(split3[1])
			c, _ := strconv.Atoi(split4[0])
			d, _ := strconv.Atoi(split4[1])
			rules[split1[0]] = func(i int) bool {
				return (i >= a && i <= b) || (i >= c && i <= d)
			}
		case 1:
			split := strings.Split(v, ",")
			yourTicket = make([]int, len(split))
			for i, s := range split {
				yourTicket[i], _ = strconv.Atoi(s)
			}
		case 2:
			split := strings.Split(v, ",")
			ticket := make([]int, len(split))
			for i, s := range split {
				ticket[i], _ = strconv.Atoi(s)
			}
			nearbyTickets = append(nearbyTickets, ticket)
		}
	})

	return rules, yourTicket, nearbyTickets
}
