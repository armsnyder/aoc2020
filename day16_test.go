package main

import (
	"io"
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

//func TestDay16Part2(t *testing.T) {
//	got, _ := day16Part2(strings.NewReader(`
//class: 0-1 or 4-19
//row: 0-5 or 8-19
//seat: 0-13 or 16-19
//
//your ticket:
//11,12,13
//
//nearby tickets:
//3,9,18
//15,1,5
//5,14,9
//`))
//	want := []string{"row", "class", "seat"}
//	if !reflect.DeepEqual(want, got) {
//		t.Errorf("got=%v, want=%v", got, want)
//	}
//}

func Test_day16Part2(t *testing.T) {
	type args struct {
		inputReader io.Reader
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []int
	}{
		{
			args: args{inputReader: strings.NewReader(`
class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
`)},
			want:  []string{"row", "class", "seat"},
			want1: []int{11, 12, 13},
		},
		{
			args: args{inputReader: strings.NewReader(`
class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19
aaa: 14-15 or 18-20

your ticket:
11,12,13,5

nearby tickets:
3,9,18,15
15,1,5,18
5,14,9,19
`)},
			want:  []string{"row", "class", "seat", "aaa"},
			want1: []int{11, 12, 13, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := day16Part2(tt.args.inputReader)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("day16Part2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("day16Part2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
