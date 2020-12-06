package aocutil

import (
	"bufio"
	"io"
	"strconv"
)

func ReadAllInts(inputReader io.Reader) []int {
	var result []int
	VisitInts(inputReader, func(v int) {
		result = append(result, v)
	})
	return result
}

func VisitInts(inputReader io.Reader, visitFn func(v int)) {
	VisitLines(inputReader, func(v string) {
		if len(v) > 0 {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			visitFn(n)
		}
	})
}

func ReadAllStrings(inputReader io.Reader) []string {
	var result []string
	VisitStrings(inputReader, func(v string) {
		result = append(result, v)
	})
	return result
}

func VisitStrings(inputReader io.Reader, visitFn func(v string)) {
	VisitLines(inputReader, func(v string) {
		if len(v) > 0 {
			visitFn(v)
		}
	})
}

func VisitStringGroups(inputReader io.Reader, visitFn func(v []string)) {
	var group []string
	VisitLines(inputReader, func(v string) {
		if len(v) > 0 {
			group = append(group, v)
		} else if len(group) > 0 {
			visitFn(group)
			group = nil
		}
	})
	if len(group) > 0 {
		visitFn(group)
	}
}

func VisitLines(inputReader io.Reader, visitFn func(v string)) {
	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		visitFn(scanner.Text())
	}
}
