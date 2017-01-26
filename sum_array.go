package main

import (
	"fmt"
	"os"
	"strconv"
)

// O(n)
func findSumBF(arr []int, sum int) bool {
	i, j := 0, len(arr)-1
	for i <= j {
		t := arr[i] + arr[j]
		if t == sum {
			return true
		} else if t > sum {
			j--
		} else {
			i++
		}
	}

	return false
}

func findSumBST(arr []int, sum int) bool {
	i := 0
	for i < len(arr) {
		t := sum - arr[i]
		if findSumBST_Util(arr[i+1:], t) {
			return true
		}
		i++
	}
	return false
}

// nlg(n)
func findSumBST_Util(arr []int, sum int) bool {
	l := len(arr)
	m := l / 2
	if l == 1 {
		return sum == arr[0]
	} else if l%2 != 0 {
		m = m + 1
	}

	if arr[m] > sum {
		return findSumBST_Util(arr[:m], sum)
	} else if arr[m] < sum {
		return findSumBST_Util(arr[m:], sum)
	} else {
		return true
	}
}

func findSumHash(arr []int, sum int) bool {
	l := len(arr)
	i := 0
	hmap := [1000]int{}

	for i < l {
		t := sum - arr[i]
		if hmap[t] == 1 {
			return true
		}
		hmap[arr[i]] = 1
		i++
	}

	return false
}

// Sub Array with given sum
// o(n)
func findSumSubArray(arr []int, sum int) (int, int) {
	curr := arr[0]
	i, l := 1, len(arr)
	start := 0

	for i < l {
		curr = curr + arr[i]
		if curr > sum {
			for curr > sum && start < i-1 {
				curr = curr - arr[start]
				start++
			}
		}
		if curr == sum {
			return start, i
		}
		i++
	}
	return -1, -1
}

func main() {
	args := os.Args
	sum, _ := strconv.Atoi(args[1])
	arr := []int{2, 4, 6, 9, 10, 12, 16, 20}
	fmt.Println(findSumBF(arr, sum))
	fmt.Println(findSumBST(arr, sum))
	fmt.Println(findSumHash(arr, sum))

	s, e := findSumSubArray(arr, 25)
	fmt.Println(s, e)
}
