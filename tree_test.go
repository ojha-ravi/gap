package main

import (
	"reflect"
	"testing"
)

var a = []int{4, 1, 3, 2, 5, 6}

type TestData struct {
	inp  []int
	in   []int
	pre  []int
	post []int
}

var testData = []TestData{
	{
		[]int{4, 1, 3, 2, 5, 6},
		[]int{1, 2, 3, 4, 5, 6},
		[]int{4, 1, 3, 2, 5, 6},
		[]int{2, 3, 1, 6, 5, 4},
	},
	{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{1, 2, 3, 4, 5, 6},
		[]int{1, 2, 3, 4, 5, 6},
		[]int{6, 5, 4, 3, 2, 1},
	},
}

func Test_inOrder(t *testing.T) {
	for _, td := range testData {
		tree := createBST(td.inp)

		var res []int
		inOrder(tree, &res)
		if !reflect.DeepEqual(res, td.in) {
			t.Errorf("inOrder(tree, arr) = %v, want %v", res, td.in)
		}

		res = []int{}
		preOrder(tree, &res)
		if !reflect.DeepEqual(res, td.pre) {
			t.Errorf("preOrder(tree, arr) = %v, want %v", res, td.pre)
		}

		res = []int{}
		postOrder(tree, &res)
		if !reflect.DeepEqual(res, td.post) {
			t.Errorf("postOrder(tree, arr) = %v, want %v", res, td.post)
		}
	}
}
