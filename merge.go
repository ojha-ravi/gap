package main

import "fmt"

func main() {
	a := []int{1, 2, 2}
	b := []int{2, 5}
	fmt.Println(combineArray(a, b))
}

func combineArray(a []int, b []int) []int {
	var res []int

	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else if a[i] > b[j] {
			res = append(res, b[j])
			j++
		}

		if i == len(a) {
			for j < len(b) {
				res = append(res, b[j])
				j++
			}
		} else if j == len(b) {
			for i < len(a) {
				res = append(res, a[i])
				i++
			}
		}
	}
	return res
}
