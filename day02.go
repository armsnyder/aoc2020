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
	var sb strings.Builder
	var i int
	for i < len(line) {
		if line[i] == '-' {
			i++
			break
		}
		sb.WriteByte(line[i])
		i++
	}
	var err error
	min, err = strconv.Atoi(sb.String())
	if err != nil {
		panic(err)
	}
	sb.Reset()
	for i < len(line) {
		if line[i] == ' ' {
			i++
			break
		}
		sb.WriteByte(line[i])
		i++
	}
	max, err = strconv.Atoi(sb.String())
	if err != nil {
		panic(err)
	}
	letter = string(line[i])
	password = line[i+3:]
	return
}
