package main

import (
	. "github.com/armsnyder/aoc2020/common"
)

func main() {
	AssertEqual(514579, day01(false, 1721, 979, 366, 299, 675, 1456))
	AssertEqual(241861950, day01(true, 1721, 979, 366, 299, 675, 1456))

	var puzzleInput []int
	partB := Bootstrap(1, &puzzleInput)

	Submit(1, partB, day01(partB, puzzleInput...))
}

func day01(partB bool, ints ...int) int {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			if partB {
				for k := j + 1; k < len(ints); k++ {
					if ints[i]+ints[j]+ints[k] == 2020 {
						return ints[i] * ints[j] * ints[k]
					}
				}
			} else {
				if ints[i]+ints[j] == 2020 {
					return ints[i] * ints[j]
				}
			}
		}
	}

	panic("no solution")
}
