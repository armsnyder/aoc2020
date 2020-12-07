package aocutil

import (
	"bytes"
	"strconv"
)

func Ints(rawInput []byte) []int {
	var result []int
	for _, s := range Strings(rawInput) {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}
	return result
}

func Strings(rawInput []byte) []string {
	var result []string
	for _, line := range Lines(rawInput) {
		if len(line) > 0 {
			result = append(result, string(line))
		}
	}
	return result
}

func StringGroups(rawInput []byte) [][]string {
	var result [][]string
	var group []string

	for _, line := range Lines(rawInput) {
		if len(line) > 0 {
			group = append(group, string(line))
		} else if len(group) > 0 {
			result = append(result, group)
			group = nil
		}
	}

	if len(group) > 0 {
		result = append(result, group)
	}

	return result
}

func Lines(rawInput []byte) [][]byte {
	return bytes.Split(rawInput, []byte{'\n'})
}
