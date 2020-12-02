package main

import (
	"reflect"
	"testing"
)

func Test_day02(t *testing.T) {
	type args struct {
		part2        bool
		inputUntyped interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			args: args{inputUntyped: []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}},
			want: 2,
		},
		{
			args: args{part2: true, inputUntyped: []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day02(tt.args.part2, tt.args.inputUntyped); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("day02(%v, ...) = %v, want %v", tt.args.part2, got, tt.want)
			}
		})
	}
}
