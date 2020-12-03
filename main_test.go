package main

import (
	"reflect"
	"testing"
)

type dayTest struct {
	name  string
	part2 bool
	input interface{}
	want  interface{}
}

func runDayTests(t *testing.T, day int, tests []dayTest) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := days[day](tt.part2, tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("day%02d(%v, ...) = %v, want %v", day, tt.part2, got, tt.want)
			}
		})
	}
}
