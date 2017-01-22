package main

import (
	"reflect"
	"testing"
)

func Test_combineArray(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"combineArray",
			args{
				[]int{1, 2, 3},
				[]int{4, 5},
			},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"combineArray",
			args{
				[]int{1, 2, 2},
				[]int{2, 5},
			},
			[]int{1, 2, 2, 2, 5},
		},
		{
			"combineArray",
			args{
				[]int{1, 9, 10},
				[]int{4, 5},
			},
			[]int{1, 4, 5, 9, 10},
		},
		{
			"combineArray",
			args{
				[]int{1},
				[]int{2, 5},
			},
			[]int{1, 2, 5},
		},
		{
			"combineArray",
			args{
				[]int{1, 6, 10},
				[]int{2, 5, 7, 11},
			},
			[]int{1, 2, 5, 6, 7, 10, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combineArray(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combineArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
