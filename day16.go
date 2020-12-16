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
	_, rules, _, nearbyTickets := day16Parse(inputReader)

	errorRate := 0

	for _, ticket := range nearbyTickets {
		day16ValidateTicket(ticket, rules, func(value int) {
			errorRate += value
		})
	}

	return errorRate
}

func day16Part2(inputReader io.Reader) ([]string, []int) {
	fields, rules, yourTicket, nearbyTickets := day16Parse(inputReader)
	validTickets := [][]int{yourTicket}

	for _, ticket := range nearbyTickets {
		if day16ValidateTicket(ticket, rules) {
			validTickets = append(validTickets, ticket)
		}
	}

	indexFieldCandidates := make([][]int, len(fields))

	for i := range fields {
		for j, rule := range rules {
			if day16CheckRuleIndex(rule, validTickets, i) {
				indexFieldCandidates[i] = append(indexFieldCandidates[i], j)
			}
		}
	}

	var queue []int

	for _, candidates := range indexFieldCandidates {
		if len(candidates) == 1 {
			queue = append(queue, candidates[0])
		}
	}

	for len(queue) > 0 {
		fieldToCheck := queue[0]
		queue = queue[1:]

		for i := range indexFieldCandidates {
			if len(indexFieldCandidates[i]) == 1 {
				continue
			}

			for j, field := range indexFieldCandidates[i] {
				if field == fieldToCheck {
					indexFieldCandidates[i] = append(indexFieldCandidates[i][:j], indexFieldCandidates[i][j+1:]...)
					break
				}
			}

			if len(indexFieldCandidates[i]) == 1 {
				queue = append(queue, indexFieldCandidates[i][0])
			}
		}
	}

	sortedFields := make([]string, len(fields))
	for i, candidates := range indexFieldCandidates {
		sortedFields[i] = fields[candidates[0]]
	}

	return sortedFields, yourTicket
}

func day16ValidateTicket(ticket []int, rules []day16Rule, onInvalidValue ...func(int)) bool {
	valid := true

outer:
	for _, value := range ticket {
		for _, rule := range rules {
			if rule.check(value) {
				continue outer
			}
		}

		if len(onInvalidValue) == 0 {
			return false
		}

		valid = false
		onInvalidValue[0](value)
	}

	return valid
}

func day16CheckRuleIndex(rule day16Rule, tickets [][]int, index int) bool {
	for _, ticket := range tickets {
		if !rule.check(ticket[index]) {
			return false
		}
	}

	return true
}

type day16Rule [4]int

func (r day16Rule) check(v int) bool {
	return v >= r[0] && v <= r[1] || v >= r[2] && v <= r[3]
}

func day16Parse(inputReader io.Reader) (fields []string, rules []day16Rule, yourTicket []int, nearbyTickets [][]int) {
	section := 0

	aocutil.VisitStrings(inputReader, func(v string) {
		if v == "your ticket:" {
			section = 1
			return
		}

		if v == "nearby tickets:" {
			section = 2
			return
		}

		switch section {
		case 0:
			split1 := strings.SplitN(v, ": ", 2)
			split2 := strings.SplitN(split1[1], " or ", 2)
			split3 := strings.SplitN(split2[0], "-", 2)
			split4 := strings.SplitN(split2[1], "-", 2)

			a, _ := strconv.Atoi(split3[0])
			b, _ := strconv.Atoi(split3[1])
			c, _ := strconv.Atoi(split4[0])
			d, _ := strconv.Atoi(split4[1])

			fields = append(fields, split1[0])
			rules = append(rules, day16Rule{a, b, c, d})

		case 1:
			yourTicket = day16ParseTicket(v)

		case 2:
			nearbyTickets = append(nearbyTickets, day16ParseTicket(v))
		}
	})

	return fields, rules, yourTicket, nearbyTickets
}

func day16ParseTicket(v string) []int {
	split := strings.Split(v, ",")
	ticket := make([]int, len(split))
	for i, s := range split {
		ticket[i], _ = strconv.Atoi(s)
	}
	return ticket
}
