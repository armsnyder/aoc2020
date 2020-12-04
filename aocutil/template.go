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

import "io"

var _ = declareDay(%d, func(part2 bool, inputReader io.Reader) interface{} {
	panic("no solution")
})
`, day); err != nil {
		panic(err)
	}

	if _, err := fmt.Fprintf(test, `package main

import (
	"testing"
)

func Test_day%02[1]d(t *testing.T) {
	runDayTests(t, %[1]d, []dayTest{
		{
			input: %[2]s%[2]s,
			want:  nil,
		},
		{
			part2: true,
			input: %[2]s%[2]s,
			want:  nil,
		},
	})
}
`, day, "`"); err != nil {
		panic(err)
	}
}
