package main

import "fmt"

// O(n)
func findSumInArrayBF(arr []int, sum int) bool {
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

func findSumInArrayBST(arr []int, sum int) bool {
	i := 0
	for i < len(arr) {
		t := sum - arr[i]
		if findSumInArrayBST_Util(arr[i+1:], t) {
			return true
		}
		i++
	}
	return false
}

// nlg(n)
func findSumInArrayBST_Util(arr []int, sum int) bool {
	l := len(arr)
	m := l / 2
	if l == 1 {
		return sum == arr[0]
	} else if l%2 != 0 {
		m = m + 1
	}

	if arr[m] > sum {
		return findSumInArrayBST_Util(arr[:m], sum)
	} else if arr[m] < sum {
		return findSumInArrayBST_Util(arr[m:], sum)
	} else {
		return true
	}
}

func findSumInArrayHash(arr []int, sum int) bool {
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
func findSubArraySumInArray(arr []int, sum int) string {
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
			return fmt.Sprintf("%v-%v", start, i)
		}
		i++
	}
	return fmt.Sprintf("%v-%v", -1, -1)
}

// func main() {
// 	args := os.Args
// 	sum, _ := strconv.Atoi(args[1])
// 	arr := []int{2, 4, 6, 9, 10, 12, 16, 20}
// 	fmt.Println(findSumInArrayBF(arr, sum))
// 	fmt.Println(findSumInArrayBST(arr, sum))
// 	fmt.Println(findSumInArrayHash(arr, sum))

// 	fmt.Println(findSubArraySumInArray(arr, 25))
// }
