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

// day13Part2 assumes that all bus IDs are co-prime.
func day13Part2(inputReader io.Reader) interface{} {
	notes := strings.Split(aocutil.ReadAllStrings(inputReader)[1], ",")

	var workingResult int64 = 0
	var workingModuliProduct int64 = 1

	for i, note := range notes {
		if note == "x" {
			continue
		}

		// Find a new workingResult that satisfies all previous modular congruencies, in addition to
		// the new congruency given by this bus ID. This is done by solving the modular equation for
		// k and adding k * workingModuliProduct to the workingResult.
		//     k * workingModuliProduct + workingResult + i ≡ 0 (mod busID)

		busID, _ := strconv.ParseInt(note, 10, 64)
		congruentTo := day13PositiveModulo(-workingResult-int64(i), busID)
		inverse := day13ModularMultiplicativeInverse(workingModuliProduct, busID)
		k := (inverse * congruentTo) % busID

		workingResult += workingModuliProduct * k
		workingModuliProduct *= busID
	}

	return workingResult
}

func day13PositiveModulo(a, m int64) int64 {
	return ((a % m) + m) % m
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
