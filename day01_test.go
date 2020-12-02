package main

import (
	"reflect"
	"testing"
)

func Test_day01(t *testing.T) {
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
			args: args{inputUntyped: []int{1721, 979, 366, 299, 675, 1456}},
			want: 514579,
		},
		{
			args: args{part2: true, inputUntyped: []int{1721, 979, 366, 299, 675, 1456}},
			want: 241861950,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day01(tt.args.part2, tt.args.inputUntyped); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("day01(%v, ...) = %v, want %v", tt.args.part2, got, tt.want)
			}
		})
	}
}
