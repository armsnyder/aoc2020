package aocutil

import (
	"fmt"
	"os"
)

func GenerateStub(day int) {
	code, err := os.Create(fmt.Sprintf("day%02d.go", day))
	if err != nil {
		panic(err)
	}
	defer code.Close()

	test, err := os.Create(fmt.Sprintf("day%02d_test.go", day))
	if err != nil {
		panic(err)
	}
	defer code.Close()

	if _, err := fmt.Fprintf(code, `package main

import (
	"io"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(%[1]d, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day%02[1]dPart2(inputReader)
	}
	return day%02[1]dPart1(inputReader)
})

func day%02[1]dPart1(inputReader io.Reader) interface{} {
	aocutil.VisitStrings(inputReader, func(v string) {})

	panic("no solution")
}

func day%02[1]dPart2(inputReader io.Reader) interface{} {
	aocutil.VisitStrings(inputReader, func(v string) {})

	panic("no solution")
}
`, day); err != nil {
		panic(err)
	}

	if _, err := fmt.Fprintf(test, `package main

import (
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

func TestDay%02[1]dPart1(t *testing.T) {
	runDayTests(t, %[1]d, []dayTest{
		{
			input: %[2]s%[2]s,
			want:  nil,
		},
	})
}

func TestDay%02[1]dPart2(t *testing.T) {
	runDayTests(t, %[1]d, []dayTest{
		{
			part2: true,
			input: %[2]s%[2]s,
			want:  nil,
		},
	})
}

func BenchmarkDay%02[1]dPart1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[%[1]d](false, aocutil.GetInput(%[1]d))
	}
}

func BenchmarkDay%02[1]dPart2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		days[%[1]d](true, aocutil.GetInput(%[1]d))
	}
}
`, day, "`"); err != nil {
		panic(err)
	}
}
