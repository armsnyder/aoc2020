package main

import (
	"io"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(1, func(partB bool, inputReader io.Reader) interface{} {
	input := aocutil.ReadAllInts(inputReader)

	var values [2020]bool
	for _, value := range input {
		if value < 2020 {
			values[value] = true
		}
	}

	for i, vi := range input {
		if partB {
			for j, vj := range input {
				if j <= i {
					continue
				}
				if vi+vj < 2020 && values[2020-vi-vj] {
					return vi * vj * (2020 - vi - vj)
				}
			}
		} else if values[2020-vi] {
			return vi * (2020 - vi)
		}
	}

	panic("no solution")
})
