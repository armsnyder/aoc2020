package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(2, func(part2 bool, inputReader io.Reader) interface{} {
	total := 0

	aocutil.VisitStrings(inputReader, func(s string) {
		min, max, letter, password := day02Parse(s)

		if part2 {
			if (password[min-1] == letter[0]) != (password[max-1] == letter[0]) {
				total++
			}
		} else {
			count := strings.Count(password, letter)
			if count >= min && count <= max {
				total++
			}
		}
	})

	return total
})

func day02Parse(line string) (min, max int, letter string, password string) {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		switch r {
		case '-', ' ', ':':
			return true
		default:
			return false
		}
	})
	min, _ = strconv.Atoi(fields[0])
	max, _ = strconv.Atoi(fields[1])
	letter = fields[2]
	password = fields[3]
	return
}
