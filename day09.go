package main

import (
	"io"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(9, func(part2 bool, inputReader io.Reader) interface{} {
	preamble := 25

	if part2 {
		return day09Part2(inputReader, preamble)
	}

	return day09Part1(inputReader, preamble)
})

func day09Part1(inputReader io.Reader, preamble int) int {
	values := aocutil.ReadAllInts(inputReader)
	return day09FindInvalid(values, preamble)
}

func day09Part2(inputReader io.Reader, preamble int) int {
	values := aocutil.ReadAllInts(inputReader)
	invalid := day09FindInvalid(values, preamble)

outer:
	for i := 0; i < len(values); i++ {
		sum := values[i]
		smallest := sum
		largest := sum

		for j := i + 1; j < len(values); j++ {
			sum += values[j]

			if values[j] < smallest {
				smallest = values[j]
			} else if values[j] > largest {
				largest = values[j]
			}

			if sum == invalid {
				return smallest + largest
			} else if sum > invalid {
				continue outer
			}
		}
	}

	panic("no solution")
}

func day09FindInvalid(values []int, preamble int) int {
	// Naive solution had the best performance at the puzzle scale.
outer:
	for i := preamble; i < len(values); i++ {
		for j := i - preamble; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if values[i] == values[j]+values[k] {
					continue outer
				}
			}
		}

		return values[i]
	}

	panic("no solution")
}
