package main

import (
	"fmt"
	"os"

	"github.com/dey-z/code_training_practice/codility/max_slice_sum/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A = testdata.T[args[0]]
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
