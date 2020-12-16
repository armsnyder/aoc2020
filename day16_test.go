package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDay16Part1(t *testing.T) {
	runDayTests(t, 16, []dayTest{
		{
			input: `
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`,
			want: 71,
		},
	})
}

func TestDay16Part2(t *testing.T) {
	inputReader := strings.NewReader(`
class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
`)

	sortedFields, yourTicket := day16MatchFieldsOnYourTicket(inputReader)

	wantSortedFields := []string{"row", "class", "seat"}
	wantYourTicket := day16Ticket{11, 12, 13}

	if !reflect.DeepEqual(sortedFields, wantSortedFields) {
		t.Errorf("sortedFields: got %v; wanted %v", sortedFields, wantSortedFields)
	}

	if !reflect.DeepEqual(yourTicket, wantYourTicket) {
		t.Errorf("yourTicket: got %v; wanted %v", yourTicket, wantYourTicket)
	}
}
