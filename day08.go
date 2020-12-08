package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(8, func(part2 bool, inputReader io.Reader) interface{} {
	program := day08NewCorruptedProgram(day08Parse(inputReader))

	if !part2 {
		cpu := day08CPU{program: program}
		cpu.runUntilHaltOrLoop()
		return cpu.acc
	}

	for i := 0; i < program.len(); i++ {
		instruction := program.get(i)
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
		cpu.program = program.replace(i, day08Instruction{op: flipped, arg: instruction.arg})
		if cpu.runUntilHaltOrLoop() {
			return cpu.acc
		}
	}

	panic("no solution")
})

type day08Instruction struct {
	op  string
	arg int
}

type day08CorruptedProgram struct {
	corruptedInstructions []day08Instruction
	corruptedIndex        int
	replacement           day08Instruction
}

func day08NewCorruptedProgram(instructions []day08Instruction) day08CorruptedProgram {
	return day08CorruptedProgram{
		corruptedInstructions: instructions,
		corruptedIndex:        -1,
	}
}

func (d day08CorruptedProgram) get(i int) day08Instruction {
	if i == d.corruptedIndex {
		return d.replacement
	}
	return d.corruptedInstructions[i]
}

func (d day08CorruptedProgram) len() int {
	return len(d.corruptedInstructions)
}

func (d day08CorruptedProgram) replace(i int, instruction day08Instruction) day08CorruptedProgram {
	return day08CorruptedProgram{
		corruptedInstructions: d.corruptedInstructions,
		corruptedIndex:        i,
		replacement:           instruction,
	}
}

type day08CPU struct {
	ip      int
	acc     int
	program day08CorruptedProgram
}

func (d *day08CPU) runUntilHaltOrLoop() (halts bool) {
	programLen := d.program.len()
	seen := make([]bool, programLen)
	for {
		seen[d.ip] = true
		d.step()
		if d.ip >= programLen {
			return true
		}
		if seen[d.ip] {
			return false
		}
	}
}

func (d *day08CPU) step() {
	instruction := d.program.get(d.ip)
	switch instruction.op {
	case "nop":
		d.ip++
	case "acc":
		d.acc += instruction.arg
		d.ip++
	case "jmp":
		d.ip += instruction.arg
	}
}

func day08Parse(inputReader io.Reader) []day08Instruction {
	var instructions []day08Instruction
	aocutil.VisitStrings(inputReader, func(v string) {
		fields := strings.Fields(v)
		arg, _ := strconv.Atoi(fields[1])
		instructions = append(instructions, day08Instruction{op: fields[0], arg: arg})
	})
	return instructions
}
