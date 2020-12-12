package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(12, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day12Part2(inputReader)
	}
	return day12Part1(inputReader)
})

func day12Part1(inputReader io.Reader) int {
	shipX, shipY := 0, 0
	shipRotation := 0

	rotationToCardinal := map[int]uint8{0: 'E', 90: 'N', 180: 'W', 270: 'S'}

	day12VisitDirections(inputReader, func(action uint8, value int) {
		if action == 'F' {
			action = rotationToCardinal[shipRotation]
		}

		switch action {
		case 'N':
			shipY += value
		case 'S':
			shipY -= value
		case 'E':
			shipX += value
		case 'W':
			shipX -= value
		case 'L':
			shipRotation += value
			shipRotation %= 360
		case 'R':
			shipRotation += 360 - value
			shipRotation %= 360
		}
	})

	return day12ManhattanDistance(shipX, shipY)
}

func day12Part2(inputReader io.Reader) int {
	shipX, shipY := 0, 0
	waypointX, waypointY := 10, 1

	day12VisitDirections(inputReader, func(action uint8, value int) {
		switch action {
		case 'N':
			waypointY += value
		case 'S':
			waypointY -= value
		case 'E':
			waypointX += value
		case 'W':
			waypointX -= value
		case 'L':
			waypointX, waypointY = day12Rotate(waypointX, waypointY, value)
		case 'R':
			waypointX, waypointY = day12Rotate(waypointX, waypointY, 360-value)
		case 'F':
			shipX += waypointX * value
			shipY += waypointY * value
		}
	})

	return day12ManhattanDistance(shipX, shipY)
}

func day12VisitDirections(inputReader io.Reader, visitFn func(action uint8, value int)) {
	aocutil.VisitStrings(inputReader, func(v string) {
		action := v[0]
		value, _ := strconv.Atoi(v[1:])
		visitFn(action, value)
	})
}

func day12Rotate(x, y, deg int) (x2, y2 int) {
	switch deg {
	case 90:
		return -y, x
	case 180:
		return -x, -y
	case 270:
		return y, -x
	default:
		panic(fmt.Errorf("unhandled angle %d", deg))
	}
}

func day12ManhattanDistance(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}
