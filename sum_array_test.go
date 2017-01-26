package main

import "testing"

func Test_findSumInArrayBF(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"findSumInArrayBF",
			args{[]int{2, 4, 6, 9, 10, 12, 16, 20}, 25},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumInArrayBF(tt.args.arr, tt.args.sum); got != tt.want {
				t.Errorf("findSumInArrayBF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSumInArrayBST(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"findSumInArrayBST",
			args{[]int{2, 4, 6, 9, 10, 12, 16, 20}, 25},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumInArrayBST(tt.args.arr, tt.args.sum); got != tt.want {
				t.Errorf("findSumInArrayBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSumInArrayHash(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"findSumInArrayHash",
			args{[]int{2, 4, 6, 9, 10, 12, 16, 20}, 25},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumInArrayHash(tt.args.arr, tt.args.sum); got != tt.want {
				t.Errorf("findSumInArrayHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSubArraySumInArray(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"findSubArraySumInArray",
			args{[]int{2, 4, 6, 9, 10, 12, 16, 20}, 25},
			"2-4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubArraySumInArray(tt.args.arr, tt.args.sum); got != tt.want {
				t.Errorf("findSubArraySumInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
