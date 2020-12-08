package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(8, func(part2 bool, inputReader io.Reader) interface{} {
	program := day08Parse(inputReader)

	if part2 {
		for i, instruction := range program {
			var flipped string
			switch instruction.op {
			case "nop":
				flipped = "jmp"
			case "jmp":
				flipped = "nop"
			default:
				continue
			}
			var cpu day08CPU
			cpu.loadProgram(program)
			cpu.program[i].op = flipped
			if cpu.runUntilHaltOrLoop() {
				return cpu.acc
			}
		}
		panic("no solution")
	}

	cpu := day08CPU{program: program}
	cpu.runUntilHaltOrLoop()
	return cpu.acc
})

type day08Instruction struct {
	op  string
	arg int
}

type day08Program []day08Instruction

type day08CPU struct {
	ip      int
	acc     int
	program day08Program
}

func (d *day08CPU) loadProgram(program day08Program) {
	d.program = make(day08Program, len(program))
	copy(d.program, program)
}

func (d *day08CPU) runUntilHaltOrLoop() (halts bool) {
	seen := make([]bool, len(d.program))
	for {
		seen[d.ip] = true
		d.step()
		if d.ip >= len(d.program) {
			return true
		}
		if seen[d.ip] {
			return false
		}
	}
}

func (d *day08CPU) step() {
	switch d.program[d.ip].op {
	case "nop":
		d.ip++
	case "acc":
		d.acc += d.program[d.ip].arg
		d.ip++
	case "jmp":
		d.ip += d.program[d.ip].arg
	}
}

func day08Parse(inputReader io.Reader) day08Program {
	var program day08Program
	aocutil.VisitStrings(inputReader, func(v string) {
		fields := strings.Fields(v)
		arg, _ := strconv.Atoi(fields[1])
		program = append(program, day08Instruction{op: fields[0], arg: arg})
	})
	return program
}
