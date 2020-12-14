package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(14, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day14Part2(inputReader)
	}
	return day14Part1(inputReader)
})

func day14Part1(inputReader io.Reader) interface{} {
	var ones, zeros int64

	visitMaskFn := func(mask string) {
		ones, zeros = 0, 0
		for i, ch := range mask {
			switch ch {
			case '0':
				zeros |= 1 << (35 - i)
			case '1':
				ones |= 1 << (35 - i)
			}
		}
	}

	memory := make(map[int64]int64)

	visitWriteFn := func(addr, value int64) {
		memory[addr] = (value | ones) &^ zeros
	}

	day14VisitInput(inputReader, visitMaskFn, visitWriteFn)

	return day14SumMemory(memory)
}

func day14Part2(inputReader io.Reader) interface{} {
	var ones int64
	var floatingIndices []int

	visitMaskFn := func(mask string) {
		ones = 0
		floatingIndices = nil
		for i, ch := range mask {
			switch ch {
			case '1':
				ones |= 1 << (35 - i)
			case 'X':
				floatingIndices = append(floatingIndices, 35-i)
			}
		}
	}

	memory := make(map[int64]int64)

	visitWriteFn := func(baseAddr, value int64) {
		day14ResolveFloatingAddresses(floatingIndices, ones, 0, baseAddr, func(addr int64) {
			memory[addr] = value
		})
	}

	day14VisitInput(inputReader, visitMaskFn, visitWriteFn)

	return day14SumMemory(memory)
}

func day14VisitInput(inputReader io.Reader, visitMaskFn func(mask string), visitWriteFn func(addr, value int64)) {
	aocutil.VisitStrings(inputReader, func(v string) {
		parts := strings.SplitN(v, " = ", 2)
		switch parts[0][1] {
		case 'a': // "mask"
			visitMaskFn(parts[1])
		case 'e': // "mem[x]"
			addr, _ := strconv.ParseInt(parts[0][4:len(parts[0])-1], 10, 64)
			value, _ := strconv.ParseInt(parts[1], 10, 64)
			visitWriteFn(addr, value)
		}
	})
}

func day14SumMemory(memory map[int64]int64) (sum int64) {
	for _, v := range memory {
		sum += v
	}
	return sum
}

func day14ResolveFloatingAddresses(floatingIndices []int, ones, zeros, baseAddr int64, visitFn func(addr int64)) {
	if len(floatingIndices) == 0 {
		addr := (baseAddr | ones) &^ zeros
		visitFn(addr)
		return
	}

	floatIndex := floatingIndices[0]
	nextFloatingIndices := floatingIndices[1:]

	day14ResolveFloatingAddresses(nextFloatingIndices, ones|(1<<floatIndex), zeros, baseAddr, visitFn)
	day14ResolveFloatingAddresses(nextFloatingIndices, ones, zeros|(1<<floatIndex), baseAddr, visitFn)
}
