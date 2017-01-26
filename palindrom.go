package main

import (
	"fmt"
	"math"

	"strconv"
)

func checkIfPalindromString(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			break
		}
	}
	return i >= j
}

func checkIfPalindromInt(n int) bool {
	res := 0
	tn := n
	for tn != 0 {
		res = res*10 + (tn - tn/10*10)
		tn = tn / 10
	}
	return n == res
}

func nextPalindromBF(n int) int {
	// 2147483647
	for n < math.MaxInt32 {
		if checkIfPalindromInt(n) {
			return n
		}
		n++
	}
	return 0
}

func nextPalindrom(ns string) int {
	nsl := len(ns)

	if nsl == 1 {
		res, _ := strconv.Atoi(ns)
		return res + 1
	}

	if checkIfPalindromString(ns) {
		if nsl%2 == 0 {
			i := nsl / 2
			j := nsl/2 + 1
			ts1, _ := strconv.Atoi(string(ns[i]))
			ts2, _ := strconv.Atoi(string(ns[j]))
			ts1 = ts1 + 1
			ts2 = ts2 + 1

			str := fmt.Sprintf("%v%d%d%v", ns[:i], ts1, ts2, ns[i+2:])
			res, _ := strconv.Atoi(str)
			return res
		} else {
			t := nsl/2 + 1
			ts, _ := strconv.Atoi(string(ns[t]))
			ts = ts + 1
			str := fmt.Sprintf("%v%d%v", ns[:t-1], ts, ns[t:])
			res, _ := strconv.Atoi(str)
			return res
		}
	} else {
		res, _ := strconv.Atoi(incrementRoutine(ns, nsl))
		return res
	}
}

func incrementRoutine(ns string, len int) string {
	i := 0
	j := len - 1

	for i < j {
		ts1, _ := strconv.Atoi(string(ns[i]))
		ts2, _ := strconv.Atoi(string(ns[j]))
		// fmt.Printf("i: %v, ts1: %v, j: %v, ts2: %v\n", i, ts1, j, ts2)
		if ts1 == ts2 {
			i++
			j--
		} else if ts1 > ts2 {
			ns = fmt.Sprintf("%v%v%v", ns[0:j], string(ns[i]), ns[j+1:])
			i++
			j--
		} else if ts1 < ts2 {
			for ts1 != ts2 {
				ns = incrementAtIndex(ns, i)
				ts1, _ = strconv.Atoi(string(ns[i]))
				ts2, _ = strconv.Atoi(string(ns[j]))
			}
			i++
			j--
		}
	}

	return ns
}

func incrementAtIndex(ns string, k int) string {
	mul := 1
	for k > 0 {
		mul = mul * 10
		k--
	}
	val, _ := strconv.Atoi(ns)

	val = val + mul
	return strconv.Itoa(val)
}

// func main() {
// 	args := os.Args
// 	arg, _ := strconv.Atoi(args[1])
// 	fmt.Println("Expected: ", nextPalindromBF(arg+1))
// 	fmt.Println("Observed: ", nextPalindrom(args[1]))
// }
