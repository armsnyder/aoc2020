package main

import (
	"strconv"
	"strings"
)

var _ = declareDay(2, day02)

func day02(part2 bool, inputUntyped interface{}) interface{} {
	input := inputUntyped.([]string)
	total := 0

	for _, line := range input {
		min, max, letter, password := day02Parse(line)

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
	}

	return total
}

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
