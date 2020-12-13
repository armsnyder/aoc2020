package main

import (
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(13, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day13Part2(inputReader)
	}
	return day13Part1(inputReader)
})

func day13Part1(inputReader io.Reader) interface{} {
	input := aocutil.ReadAllStrings(inputReader)
	earliestTime, _ := strconv.Atoi(input[0])
	notes := strings.Split(input[1], ",")

	bestWaitTime := math.MaxInt32
	bestBusID := 0

	for _, note := range notes {
		if note == "x" {
			continue
		}

		busID, _ := strconv.Atoi(note)
		waitTime := busID - earliestTime%busID

		if waitTime < bestWaitTime {
			bestWaitTime = waitTime
			bestBusID = busID
		}
	}

	return bestWaitTime * bestBusID
}

func day13Part2(inputReader io.Reader) interface{} {
	notes := strings.Split(aocutil.ReadAllStrings(inputReader)[1], ",")

	var expressions []day13ModularExpression

	for i, note := range notes {
		if note == "x" {
			continue
		}

		var expression day13ModularExpression
		expression.modulus, _ = strconv.ParseInt(note, 10, 64)
		expression.remainder = day13PositiveModulo(-int64(i), expression.modulus)

		expressions = append(expressions, expression)
	}

	return day13SolveSystemOfCongruences(expressions)
}

func day13PositiveModulo(a, m int64) int64 {
	return ((a % m) + m) % m
}

type day13ModularExpression struct {
	remainder int64
	modulus   int64
}

// day13SolveSystemOfCongruences uses the Chinese Remainder Theorem to find the smallest value that
// is congruent to each expression's remainder using that expression's modulus. All moduli must be
// co-prime.
func day13SolveSystemOfCongruences(expressions []day13ModularExpression) int64 {
	var moduliProduct int64 = 1

	for _, expression := range expressions {
		moduliProduct *= expression.modulus
	}

	var result int64

	for _, expression := range expressions {
		zeroValue := moduliProduct / expression.modulus
		inverse := day13ModularMultiplicativeInverse(zeroValue, expression.modulus)
		result += zeroValue * expression.remainder * inverse
		result %= moduliProduct
	}

	return result
}

// day13ModularMultiplicativeInverse uses the Extended Euclidean Algorithm to find a modular
// multiplicative inverse. a and m must be co-prime.
func day13ModularMultiplicativeInverse(a, m int64) int64 {
	m0 := m
	y := int64(0)
	x := int64(1)

	if m == 1 {
		return 0
	}

	for a > 1 {
		q := a / m
		t := m
		m = a % m
		a = t
		t = y
		y = x - q*y
		x = t
	}

	if x < 0 {
		x += m0
	}

	return x
}
