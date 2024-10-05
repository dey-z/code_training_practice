package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/dey-z/code_training_practice/codility/binary_search/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A, _ = testdata.T[args[0]]
	X, _ := strconv.Atoi(args[1])
	fmt.Println(Solution(A, X))
}

func Solution(A []int, X int) int {
	// algorithm to find if X is present in array
	res := -1
	N := len(A)
	sort.Ints(A)
	beg := 0
	end := N - 1
	for beg <= end {
		mid := (beg + end) / 2
		if A[mid] > X {
			end = mid - 1
		} else if A[mid] == X {
			return 1
		} else {
			beg = mid + 1
		}
	}
	return res
}
