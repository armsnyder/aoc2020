package main

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(4, func(part2 bool, inputReader io.Reader) interface{} {
	fieldValidators := day04PassportFieldValidators()
	validPassports := 0

	day04VisitPassports(inputReader, func(passport map[string]string) {
		for field, check := range fieldValidators {
			if part2 {
				if !check(passport[field]) {
					return
				}
			} else if _, ok := passport[field]; !ok {
				return
			}
		}

		validPassports++
	})

	return validPassports
})

func day04PassportFieldValidators() map[string]func(string) bool {
	rangeRule := func(low, hi int) func(string) bool {
		return func(s string) bool {
			i, err := strconv.Atoi(s)
			if err != nil {
				return false
			}
			return i >= low && i <= hi
		}
	}

	hgtPattern := regexp.MustCompile(`^(\d+)(in|cm)$`)
	hclPattern := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	pidPattern := regexp.MustCompile(`^[0-9]{9}$`)
	validEyeColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	return map[string]func(string) bool{
		"byr": rangeRule(1920, 2002),
		"iyr": rangeRule(2010, 2020),
		"eyr": rangeRule(2020, 2030),
		"hgt": func(s string) bool {
			match := hgtPattern.FindStringSubmatch(s)
			if len(match) == 0 {
				return false
			}
			if match[2] == "cm" {
				return rangeRule(150, 193)(match[1])
			}
			return rangeRule(59, 76)(match[1])
		},
		"hcl": func(s string) bool {
			return hclPattern.MatchString(s)
		},
		"ecl": func(s string) bool {
			return validEyeColors[s]
		},
		"pid": func(s string) bool {
			return pidPattern.MatchString(s)
		},
	}
}

func day04VisitPassports(inputReader io.Reader, visitFn func(map[string]string)) {
	passport := make(map[string]string)

	flush := func() {
		if len(passport) > 0 {
			visitFn(passport)
			passport = make(map[string]string)
		}
	}

	aocutil.VisitLines(inputReader, func(s string) {
		if len(s) == 0 {
			flush()
			return
		}

		for _, keyPair := range strings.Fields(s) {
			split := strings.Split(keyPair, ":")
			passport[split[0]] = split[1]
		}
	})

	flush()
}
