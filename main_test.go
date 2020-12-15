package main

import (
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/armsnyder/aoc2020/aocutil"
)

type dayTest struct {
	name  string
	part2 bool
	input string
	want  interface{}
}

func runDayTests(t *testing.T, day int, tests []dayTest) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := days[day](tt.part2, strings.NewReader(tt.input))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("day%02d(%v, ...) = %v, want %v", day, tt.part2, got, tt.want)
			}
		})
	}
}

func BenchmarkBaselineIO(b *testing.B) {
	day := 1
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		input := aocutil.GetInput(day)
		io.Copy(ioutil.Discard, input)
		input.Close()
	}
}
