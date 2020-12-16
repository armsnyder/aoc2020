package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(16, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day16Part2(inputReader)
	}
	return day16Part1(inputReader)
})

func day16Part1(inputReader io.Reader) interface{} {
	_, rules, _, nearbyTickets := day16Parse(inputReader)
	errorRate := 0
	for _, ticket := range nearbyTickets {
		if value := day16ValidateTicket(ticket, rules); value > -1 {
			errorRate += value
		}
	}
	return errorRate
}

func day16Part2(inputReader io.Reader) interface{} {
	sortedFields, yourTicket := day16MatchFieldsOnYourTicket(inputReader)
	result := 1
	for i, field := range sortedFields {
		if strings.HasPrefix(field, "departure") {
			result *= yourTicket[i]
		}
	}
	return result
}

func day16MatchFieldsOnYourTicket(inputReader io.Reader) ([]string, day16Ticket) {
	fields, rules, yourTicket, nearbyTickets := day16Parse(inputReader)
	validTickets := day16OnlyValidTickets(rules, yourTicket, nearbyTickets)
	sortedFields := day16SortFields(fields, rules, validTickets)
	return sortedFields, yourTicket
}

// day16ValidateTicket returns -1 if the ticket can possibly be valid given the rules.
// Otherwise it returns the value in the ticket for which no rule can apply.
func day16ValidateTicket(ticket day16Ticket, rules []day16Rule) int {
outer:
	for _, value := range ticket {
		for _, rule := range rules {
			if rule.check(value) {
				continue outer
			}
		}
		return value
	}
	return -1
}

func day16OnlyValidTickets(rules []day16Rule, yourTicket day16Ticket, nearbyTickets []day16Ticket) []day16Ticket {
	validTickets := []day16Ticket{yourTicket}
	for _, ticket := range nearbyTickets {
		if value := day16ValidateTicket(ticket, rules); value == -1 {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func day16SortFields(fields []string, rules []day16Rule, validTickets []day16Ticket) []string {
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
		solvedField := queue[0]
		queue = append(queue[1:], day16ApplyConstraint(indexFieldCandidates, solvedField)...)
	}

	sortedFields := make([]string, len(fields))
	for i, candidates := range indexFieldCandidates {
		sortedFields[i] = fields[candidates[0]]
	}

	return sortedFields
}

func day16CheckRuleIndex(rule day16Rule, tickets []day16Ticket, index int) bool {
	for _, ticket := range tickets {
		if !rule.check(ticket[index]) {
			return false
		}
	}

	return true
}

func day16ApplyConstraint(indexFieldCandidates [][]int, solvedField int) (newlySolvedFields []int) {
	for i := range indexFieldCandidates {
		if len(indexFieldCandidates[i]) == 1 {
			continue
		}

		for j, field := range indexFieldCandidates[i] {
			if field == solvedField {
				indexFieldCandidates[i] = append(indexFieldCandidates[i][:j], indexFieldCandidates[i][j+1:]...)
				break
			}
		}

		if len(indexFieldCandidates[i]) == 1 {
			newlySolvedFields = append(newlySolvedFields, indexFieldCandidates[i][0])
		}
	}

	return newlySolvedFields
}

type day16Rule [4]int

func (r day16Rule) check(v int) bool {
	return v >= r[0] && v <= r[1] || v >= r[2] && v <= r[3]
}

type day16Ticket []int

func day16Parse(inputReader io.Reader) (fields []string, rules []day16Rule, yourTicket day16Ticket, nearbyTickets []day16Ticket) {
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

func day16ParseTicket(v string) day16Ticket {
	split := strings.Split(v, ",")
	ticket := make(day16Ticket, len(split))
	for i, s := range split {
		ticket[i], _ = strconv.Atoi(s)
	}
	return ticket
}
