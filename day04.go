package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(4, func(part2 bool, rawInput []byte) interface{} {
	fieldValidators := day04PassportFieldValidators()
	validPassports := 0
	passports := day04Passports(rawInput)

outer:
	for _, passport := range passports {
		for field, check := range fieldValidators {
			if part2 {
				if !check(passport[field]) {
					continue outer
				}
			} else if _, ok := passport[field]; !ok {
				continue outer
			}
		}

		validPassports++
	}

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

func day04Passports(rawInput []byte) []map[string]string {
	groups := aocutil.StringGroups(rawInput)
	var passports []map[string]string
	for _, group := range groups {
		passport := make(map[string]string)
		for _, line := range group {
			for _, keyPair := range strings.Fields(line) {
				split := strings.Split(keyPair, ":")
				passport[split[0]] = split[1]
			}
		}
		passports = append(passports, passport)
	}
	return passports
}
