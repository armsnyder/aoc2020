package main

import (
	"regexp"
	"strconv"
	"strings"
)

var _ = declareDay(2, day02)

func day02(part2 bool, inputUntyped interface{}) interface{} {
	input := inputUntyped.([]string)
	pattern := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	total := 0

	for _, val := range input {
		match := pattern.FindStringSubmatch(val)
		min, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		letter := match[3]
		password := match[4]

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
