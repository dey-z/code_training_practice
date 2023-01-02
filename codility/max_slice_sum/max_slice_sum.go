package main

import (
	"fmt"
	"os"
)

func main() {
	var A []int
	m := make(map[string][]int)
	m["1"] = []int{3, 2, -6, 4, 0}
	m["2"] = []int{1, 1, 1}
	m["3"] = []int{-2, -2, -2}
	m["4"] = []int{-10}
	args := os.Args[1:]
	A = m[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// algorithm for max slice
	s := 0
	max_slice := -2147483648
	for _, a := range A {
		s += a
		max_slice = max(max_slice, s)
		// -ve check means a new slice needs to begin(which will stall till a +ve is found)
		if s < 0 {
			s = 0
		}
	}
	return max_slice
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
